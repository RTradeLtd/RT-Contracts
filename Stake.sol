pragma solidity 0.4.24;

import "./Interfaces/RTCoinInterface.sol";
import "./Interfaces/ERC20Interface.sol";
import "./Math/SafeMath.sol";

contract Stake {
    using SafeMath for uint256;

    // we mark as constant private to reduce gas costs
    
    // Minimum stake of 1RTC
    uint256 constant public MINSTAKE = 1000000000000000000;
    // NOTE ON MULTIPLIER: this is right now set to 1% this may however change before token is released
    uint256 constant public MULTIPLIER = 10000000000000000;
    // we use an average blocks per year of 2,103,840 assuming an average block time of 15 seconds.
    // the only thing effected by this, is when they can withdraw their initial stake. 
    // To do so, 2103840 blocks must've passed. The current block time must also be equal,or greater to
    // the "release date" which is calculated based off the time the initial stake is deposited added to
    // the number of seconds per block (15) multiplied by the block hold period of 2103840 blocks.
    // all other stake reward creditation and withdrawal is ultimately controlled by real-time block generation speeds.
    uint256 constant public BLOCKHOLDPERIOD = 2103840;
    uint256 constant public BLOCKSEC = 15;
    address  public TOKENCONTRACT = 0x3fDe03720917246B73ba532e4650c656D6020578;
    RTCoinInterface   public RTI = RTCoinInterface(TOKENCONTRACT);

    uint256 public activeStakes;
    address public admin;
    bool public newStakesAllowed;

    enum StakeStateEnum { nil, staking, staked }

    struct StakeStruct {
        uint256 initialStake;
        uint256 blockLocked;
        uint256 blockUnlocked;
        uint256 releaseDate;
        uint256 totalCoinsToMint;
        uint256 coinsMinted;
        uint256 rewardPerBlock;
        uint256 lastBlockWithdrawn;
        StakeStateEnum    state;
    }

    event StakesDisabled();
    event StakesEnabled();
    event StakeDeposited(address indexed _staker, uint256 indexed _stakeNum, uint256 _coinsToMint, uint256 _releaseDate, uint256 _releaseBlock);
    event StakeRewardWithdrawn(address indexed _staker, uint256 indexed _stakeNum, uint256 _reward);
    event InitialStakeWithdrawn(address indexed _staker, uint256 indexed _stakeNumber, uint256 _amount);

    mapping (address => mapping (uint256 => StakeStruct)) public stakes;
    mapping (address => uint256) public numberOfStakes;
    mapping (address => uint256) public internalRTCBalances;

    modifier validInitialStakeRelease(uint256 _stakeNum) {
        require(stakes[msg.sender][_stakeNum].state == StakeStateEnum.staking);
        require(now >= stakes[msg.sender][_stakeNum].releaseDate && block.number >= stakes[msg.sender][_stakeNum].blockUnlocked);
        require(internalRTCBalances[msg.sender] >= stakes[msg.sender][_stakeNum].initialStake);
        _;
    }

    modifier validMint(uint256 _stakeNumber) {
        // allow people to withdraw their rewards even if the staking period is over
        require(stakes[msg.sender][_stakeNumber].state == StakeStateEnum.staking || stakes[msg.sender][_stakeNumber].state == StakeStateEnum.staked);
        require(stakes[msg.sender][_stakeNumber].coinsMinted < stakes[msg.sender][_stakeNumber].totalCoinsToMint);
        uint256 currentBlock = block.number;
        uint256 lastBlockWithdrawn = stakes[msg.sender][_stakeNumber].lastBlockWithdrawn;
        require(currentBlock > lastBlockWithdrawn);
        _;
    }

    modifier stakingEnabled(uint256 _numRTC) {
        require(canMint());
        require(newStakesAllowed);
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

    function setRTI(address _contract) public onlyAdmin returns (bool) {
        RTI = RTCoinInterface(_contract);
        TOKENCONTRACT = _contract;
        return true;
    }

    function disableNewStakes() public onlyAdmin returns (bool) {
        newStakesAllowed = false;
        return true;
    }

    function allowNewStakes() public onlyAdmin returns (bool) {
        newStakesAllowed = true;
        require(RTI.stakeContract() == address(this));
        return true;
    }

    function mint(uint256 _stakeNumber) public validMint(_stakeNumber) returns (bool) {
        uint256 mintAmount = calculateMint(_stakeNumber);
        require(stakes[msg.sender][_stakeNumber].coinsMinted.add(mintAmount) <= stakes[msg.sender][_stakeNumber].totalCoinsToMint);
        stakes[msg.sender][_stakeNumber].coinsMinted = stakes[msg.sender][_stakeNumber].coinsMinted.add(mintAmount);
        stakes[msg.sender][_stakeNumber].lastBlockWithdrawn = block.number;
        emit StakeRewardWithdrawn(msg.sender, _stakeNumber, mintAmount);
        require(RTI.mint(msg.sender, mintAmount));
        return true;
    }

    function withdrawInitialStake(uint256 _stakeNumber) public validInitialStakeRelease(_stakeNumber) returns (bool) {
        uint256 initialStake = stakes[msg.sender][_stakeNumber].initialStake;
        stakes[msg.sender][_stakeNumber].state = StakeStateEnum.staked;
        activeStakes = activeStakes.sub(1);
        internalRTCBalances[msg.sender] = internalRTCBalances[msg.sender].sub(initialStake);
        emit InitialStakeWithdrawn(msg.sender, _stakeNumber, initialStake);
        require(RTI.transfer(msg.sender, initialStake));
        return true;
    }

    function depositStake(uint256 _numRTC) public stakingEnabled(_numRTC) returns (bool) {
        uint256 stakeCount = getStakeCount(msg.sender);

        (uint256 blockLocked, 
        uint256 blockReleased, 
        uint256 releaseDate, 
        uint256 totalCoinsMinted,
        uint256 rewardPerBlock) = calculateStake(_numRTC);

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

        stakes[msg.sender][stakeCount] = ss;
        numberOfStakes[msg.sender] = numberOfStakes[msg.sender].add(1);
        internalRTCBalances[msg.sender] = internalRTCBalances[msg.sender].add(_numRTC);
        activeStakes = activeStakes.add(1);
        
        emit StakeDeposited(msg.sender, stakeCount, totalCoinsMinted, releaseDate, blockReleased);

        require(RTI.transferFrom(msg.sender, address(this), _numRTC));

        return true;
    }


    // UTILITY FUNCTIONS //

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
        blockLocked = block.number;
        blockReleased = blockLocked.add(BLOCKHOLDPERIOD);
        releaseDate = now.add(BLOCKHOLDPERIOD.mul(BLOCKSEC));
        totalCoinsMinted = _numRTC.mul(MULTIPLIER);
        totalCoinsMinted = totalCoinsMinted.div(1 ether);
        rewardPerBlock = totalCoinsMinted.div(BLOCKHOLDPERIOD);
    }

    function calculateMint(uint256 _stakeNumber)
        internal
        view
        returns (uint256 reward)
    {
        uint256 currentBlock = calculateCurrentBlock(_stakeNumber);
        uint256 lastBlockWithdrawn = stakes[msg.sender][_stakeNumber].lastBlockWithdrawn;
        uint256 blocksToReward = currentBlock.sub(lastBlockWithdrawn);
        reward = blocksToReward.mul(stakes[msg.sender][_stakeNumber].rewardPerBlock);
    }

    function calculateTotalCoinsMinted(uint256 _numRTC) internal pure returns (uint256 totalCoinsMinted) {
        totalCoinsMinted = _numRTC.mul(MULTIPLIER);
        totalCoinsMinted = totalCoinsMinted.div(1 ether);
    }

    function calculateCurrentBlock(uint256 _stakeNumber) internal view returns (uint256 currentBlock) {
        currentBlock = block.number;
        if (currentBlock >= stakes[msg.sender][_stakeNumber].blockUnlocked) {
            currentBlock = stakes[msg.sender][_stakeNumber].blockUnlocked;
        }
    }
    
    function getStakeCount(address _staker) internal view returns (uint256) {
        return numberOfStakes[_staker];
    }

    // canMint checks to see that this contract can actually mint tokens on RTC
    // this should only ever NOT be true if a serious vulnerability was discovered in this contract and it had to be replaced
    // after it had been deployed.
    function canMint() public view returns (bool) {
        assert(RTI.stakeContract() == address(this));
        return true;
    }
}