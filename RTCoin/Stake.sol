pragma solidity 0.4.24;

import "../Interfaces/RTCoinInterface.sol";
import "../Interfaces/ERC20Interface.sol";
import "../Math/SafeMath.sol";

/** @title This contract is used to handle staking, and subsequently can increase RTC token supply */
contract Stake {
    using SafeMath for uint256;

    // we mark as constant private to reduce gas costs
    
    // Minimum stake of 1RTC
    uint256 constant private MINSTAKE = 1000000000000000000;
    // NOTE ON MULTIPLIER: this is right now set to 10% this may however change before token is released
    uint256 constant private MULTIPLIER = 100000000000000000;
    // BLOCKHOLDPERIOD is used to determine how many blocks a stake is held for, and how many blocks will mint tokens
    uint256 constant private BLOCKHOLDPERIOD = 2103840;
    // BLOCKSEC uses 15 seconds as an average block time. Ultimately the only thing this "restricts" is the time at which a stake is withdrawn
    // Yes, we use block timestamps which can be influenced to some degree by miners, however since this only determines the time at which an initial stake can be withdrawn at
    // due to the fact that this is also limited by block height, it is an acceptable risk
    uint256 constant private BLOCKSEC = 15;
    // this is the address of the RTC token contract
    address  public TOKENCONTRACT = 0x3fDe03720917246B73ba532e4650c656D6020578;
    // this is the interface used to interact with the RTC Token
    RTCoinInterface   public RTI = RTCoinInterface(TOKENCONTRACT);

    // keeps track of the number of active stakes
    uint256 public activeStakes;
    // keeps track of the admin address. For security purposes this can't be changed once set
    address public admin;
    // keeps track of whether or not new stakes can be made
    bool public newStakesAllowed;

    // tracks the state of a stake
    enum StakeStateEnum { nil, staking, staked }

    struct StakeStruct {
        // how many tokens were initially staked
        uint256 initialStake;
        // the block that the stake was made
        uint256 blockLocked;
        // the block at which the initial stake can be withdrawn
        uint256 blockUnlocked;
        // the time at which the initial stake can be withdrawn
        uint256 releaseDate;
        // the total number of coins to mint
        uint256 totalCoinsToMint;
        // the current number of coins that have been minted
        uint256 coinsMinted;
        // the amount of coins generated per block
        uint256 rewardPerBlock;
        // the block at which a stake was last withdrawn at 
        uint256 lastBlockWithdrawn;
        // the current state of this stake
        StakeStateEnum    state;
    }

    event StakesDisabled();
    event StakesEnabled();
    // we use indexed parameters to allow for more fine grained event filtration
    event StakeDeposited(address indexed _staker, uint256 indexed _stakeNum, uint256 _coinsToMint, uint256 _releaseDate, uint256 _releaseBlock);
    // we use indexed parameters to allow for more fine grained event filtration
    event StakeRewardWithdrawn(address indexed _staker, uint256 indexed _stakeNum, uint256 _reward);
    // we used indexed parameters to allow for more fine grained event filtration
    event InitialStakeWithdrawn(address indexed _staker, uint256 indexed _stakeNumber, uint256 _amount);

    // keeps track of the stakes a user has
    mapping (address => mapping (uint256 => StakeStruct)) public stakes;
    // keeps track of the total number of stakes a user has
    mapping (address => uint256) public numberOfStakes;
    // keeps track of the user's current RTC balance
    mapping (address => uint256) public internalRTCBalances;

    modifier validInitialStakeRelease(uint256 _stakeNum) {
        // make sure that the stake is active
        require(stakes[msg.sender][_stakeNum].state == StakeStateEnum.staking);
        require(now >= stakes[msg.sender][_stakeNum].releaseDate && block.number >= stakes[msg.sender][_stakeNum].blockUnlocked);
        require(internalRTCBalances[msg.sender] >= stakes[msg.sender][_stakeNum].initialStake);
        _;
    }

    modifier validMint(uint256 _stakeNumber) {
        // allow people to withdraw their rewards even if the staking period is over
        require(stakes[msg.sender][_stakeNumber].state == StakeStateEnum.staking || stakes[msg.sender][_stakeNumber].state == StakeStateEnum.staked);
        // make sure that the current coins minted are less than the total coins minted
        require(stakes[msg.sender][_stakeNumber].coinsMinted < stakes[msg.sender][_stakeNumber].totalCoinsToMint);
        uint256 currentBlock = block.number;
        uint256 lastBlockWithdrawn = stakes[msg.sender][_stakeNumber].lastBlockWithdrawn;
        // verify that the current block is one higher than the last block a withdrawal was made
        require(currentBlock > lastBlockWithdrawn);
        _;
    }

    modifier stakingEnabled(uint256 _numRTC) {
        // make sure this contract can mint coins on the RTC token contract
        require(canMint());
        // make sure new stakes are allowed
        require(newStakesAllowed);
        // make sure they are staking at least one RTC
        require(_numRTC >= MINSTAKE);
        _;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin);
        _;
    }

    constructor() public {
        // prevent deployment if the token contract hasn't been set yet
        if (TOKENCONTRACT == address(0)) {
            revert();
        }
        admin = msg.sender;
    }

    /** @dev Used to set the interface for the RTC token
        * Only usable by contract admin
        * @param _contract This is the address of the RTC token contract
     */
    function setRTI(address _contract) public onlyAdmin returns (bool) {
        // update the interface
        RTI = RTCoinInterface(_contract);
        // update the address
        TOKENCONTRACT = _contract;
        return true;
    }

    /** @dev Used to disable new stakes from being made
        * Only usable by contract admin
     */
    function disableNewStakes() public onlyAdmin returns (bool) {
        newStakesAllowed = false;
        return true;
    }

    /** @dev Used to allow new stakes to be made
        * For this to be enabled, the RTC token contract must be configured properly
        * This means that it is configured to use this contract as the staking contract
     */
    function allowNewStakes() public onlyAdmin returns (bool) {
        newStakesAllowed = true;
        require(RTI.stakeContract() == address(this));
        return true;
    }

    /** @dev Used by a staker to claim currently staked coins
        * Can only be executed when at least one block has passed from the last execution
        * @param _stakeNumber This is the particular stake to withdraw from
     */
    function mint(uint256 _stakeNumber) public validMint(_stakeNumber) returns (bool) {
        // determine the amount of coins to be minted in this withdrawal
        uint256 mintAmount = calculateMint(_stakeNumber);
        // make sure that we can't mint more than allowed
        require(stakes[msg.sender][_stakeNumber].coinsMinted.add(mintAmount) <= stakes[msg.sender][_stakeNumber].totalCoinsToMint);
        // update current coins minted
        stakes[msg.sender][_stakeNumber].coinsMinted = stakes[msg.sender][_stakeNumber].coinsMinted.add(mintAmount);
        // update the last block a withdrawal was made at
        stakes[msg.sender][_stakeNumber].lastBlockWithdrawn = block.number;
        // emit an event
        emit StakeRewardWithdrawn(msg.sender, _stakeNumber, mintAmount);
        // mint the tokenz
        require(RTI.mint(msg.sender, mintAmount));
        return true;
    }

    /** @dev Used by a staker to withdraw their initial stake
        * Can only be executed after the specified block number, and unix timestamp has been passed
        * @param _stakeNumber This is the particular stake to withdraw from
     */
    function withdrawInitialStake(uint256 _stakeNumber) public validInitialStakeRelease(_stakeNumber) returns (bool) {
        // get the initial stake amount
        uint256 initialStake = stakes[msg.sender][_stakeNumber].initialStake;
        // de-activate the stake
        stakes[msg.sender][_stakeNumber].state = StakeStateEnum.staked;
        // decrease the total number of stakes
        activeStakes = activeStakes.sub(1);
        // reduce their internal RTC balance
        internalRTCBalances[msg.sender] = internalRTCBalances[msg.sender].sub(initialStake);
        // emit an event
        emit InitialStakeWithdrawn(msg.sender, _stakeNumber, initialStake);
        // transfer the tokenz
        require(RTI.transfer(msg.sender, initialStake));
        return true;
    }

    /** @dev This is used to deposit coins and start staking
        * Staking must be enabled or this function will not execute
        * Must deposit at least one RTC
        * @param _numRTC This is the number of RTC tokens to stake
     */
    function depositStake(uint256 _numRTC) public stakingEnabled(_numRTC) returns (bool) {
        uint256 stakeCount = getStakeCount(msg.sender);

        // calculate the various stake parameters
        (uint256 blockLocked, 
        uint256 blockReleased, 
        uint256 releaseDate, 
        uint256 totalCoinsMinted,
        uint256 rewardPerBlock) = calculateStake(_numRTC);

        // initialize this struct in memory
        StakeStruct memory ss = StakeStruct({
            initialStake: _numRTC,
            blockLocked: blockLocked,
            blockUnlocked: blockReleased,
            releaseDate: releaseDate,
            totalCoinsToMint: totalCoinsMinted,
            coinsMinted: 0,
            rewardPerBlock: rewardPerBlock,
            lastBlockWithdrawn: blockLocked,
            state: StakeStateEnum.staking
        });

        // update the users list of stakes
        stakes[msg.sender][stakeCount] = ss;
        // update the users total stakes
        numberOfStakes[msg.sender] = numberOfStakes[msg.sender].add(1);
        // update their internal RTC balance
        internalRTCBalances[msg.sender] = internalRTCBalances[msg.sender].add(_numRTC);
        // increase the number of active stakes
        activeStakes = activeStakes.add(1);
        // emit an event
        emit StakeDeposited(msg.sender, stakeCount, totalCoinsMinted, releaseDate, blockReleased);
        // transfer tokens
        require(RTI.transferFrom(msg.sender, address(this), _numRTC));
        return true;
    }


    // UTILITY FUNCTIONS //

    /** @dev This is a helper function used to calculate the parameters of a stake
        * Will determine the block that the initial stake can be withdraw at
        * Will determine the time that the initial stake can be withdrawn at
        * Will determine the total number of RTC to be minted throughout hte stake
        * Will determine how many RTC the stakee will be awarded per block
        * @param _numRTC This is the number of RTC to be staked
     */
    function calculateStake(uint256 _numRTC) 
        internal
        view
        returns (
            uint256 blockLocked, 
            uint256 blockReleased, 
            uint256 releaseDate, 
            uint256 totalCoinsMinted,
            uint256 rewardPerBlock
        ) 
    {
        // the block that the stake is being made at
        blockLocked = block.number;
        // the block at which the initial stake will be released
        blockReleased = blockLocked.add(BLOCKHOLDPERIOD);
        // the time at which the initial stake will be released
        // please see comment at top of contract about why we consider it safe to use block times
        // linter warnings are left enabled on purpose
        releaseDate = now.add(BLOCKHOLDPERIOD.mul(BLOCKSEC));
        // total coins that will be minted
        totalCoinsMinted = _numRTC.mul(MULTIPLIER);
        // make sure to scale down
        totalCoinsMinted = totalCoinsMinted.div(1 ether);
        // calculate the coins minted per block
        rewardPerBlock = totalCoinsMinted.div(BLOCKHOLDPERIOD);
    }

    /** @dev This is a helper function used to calculate how many coins will be awarded in a given internal
        * @param _stakeNumber This is the particular stake to calculate from
     */
    function calculateMint(uint256 _stakeNumber)
        internal
        view
        returns (uint256 reward)
    {
        // calculate how many blocks they can claim a stake for
        uint256 currentBlock = calculateCurrentBlock(_stakeNumber);
        //get the last block a withdrawal was made at
        uint256 lastBlockWithdrawn = stakes[msg.sender][_stakeNumber].lastBlockWithdrawn;
        // determine the number of blocks to generate a reward for
        uint256 blocksToReward = currentBlock.sub(lastBlockWithdrawn);
        // calculate the reward
        reward = blocksToReward.mul(stakes[msg.sender][_stakeNumber].rewardPerBlock);
        // get total number of coins to be minted
        uint256 totalToMint = stakes[msg.sender][_stakeNumber].totalCoinsToMint;
        // get current number of coins minted
        uint256 currentCoinsMinted = stakes[msg.sender][_stakeNumber].coinsMinted;
        // get the new numberof total coins to be minted
        uint256 newCoinsMinted = currentCoinsMinted.add(reward);
        // if for some reason more would be generated, prevent that from happening
        if (newCoinsMinted > totalToMint) {
            reward = newCoinsMinted.sub(totalToMint);
        }
    }

    /** @dev This is a helper function used to calculate the total number of tokens to be minted
        * @param _numRTC This is the number of RTC being staked
     */
    function calculateTotalCoinsMinted(uint256 _numRTC) internal pure returns (uint256 totalCoinsMinted) {
        totalCoinsMinted = _numRTC.mul(MULTIPLIER);
        totalCoinsMinted = totalCoinsMinted.div(1 ether);
    }

    /** @dev This is a helper function used to calculate how many blocks to mint coins for
        * @param _stakeNumber This is the stake to be used for calculations
     */
    function calculateCurrentBlock(uint256 _stakeNumber) internal view returns (uint256 currentBlock) {
        currentBlock = block.number;
        // if the current block is greater than the block at which coins can be unlocked at, 
        // prevent them from generating more coins that allowed
        if (currentBlock >= stakes[msg.sender][_stakeNumber].blockUnlocked) {
            currentBlock = stakes[msg.sender][_stakeNumber].blockUnlocked;
        }
    }
    
    /** @dev This is a helper function used to get the total number of stakes a 
        * @param _staker This is the address of the stakee
     */
    function getStakeCount(address _staker) internal view returns (uint256) {
        return numberOfStakes[_staker];
    }

    /** @dev This is a helper function that checks whether or not this contract can mint tokens
        * This should only ever be false under extreme circumstances such as a potential vulnerability
     */
    function canMint() public view returns (bool) {
        assert(RTI.stakeContract() == address(this));
        return true;
    }
}