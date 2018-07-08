pragma solidity 0.4.24;

import "./Interfaces/RTCoinInterface.sol";
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";

contract Stake is Administration {
    using SafeMath for uint256;

    RTCoinInterface constant public RTI = RTCoinInterface(address(0));
    uint256 constant public MULTIPLIER = 10000000000000000;
    // we use an average blocks per year of 2,103,840 assuming an average block time of 15 seconds.
    // the only thing effected by this, is when they can withdraw their initial stake. 
    // To do so, 2103840 blocks must've passed. The current block time must also be equal,or greater to
    // the "release date" which is calculated based off the time the initial stake is deposited added to
    // the number of seconds per block (15) multiplied by the block hold period of 2103840 blocks.
    // all other stake reward creditation and withdrawal is ultimately controlled by real-time block generation speeds.
    uint256 constant public BLOCKHOLDPERIOD = 2103840;
    uint256 constant public BLOCKSEC = 15;

    enum StakeStateEnum { nil, staking, staked }

    struct StakeStruct {
        bytes32 stakeID;
        uint256 initialStake;
        uint256 blockLocked;
        uint256 blockUnlocked;
        uint256 releaseDate;
        uint256 coinsMinted;
        uint256 coinsWithdrawn;
        uint256 rewardPerBlock;
        uint256 lastBlockWithdrawn;
        StakeStateEnum    state;
    }

    mapping (address => mapping (uint256 => bytes32)) public stakeNumToIDMap;
    mapping (address => mapping (bytes32 => uint256)) public stakeIDToNumMap;
    mapping (address => mapping (uint256 => StakeStruct)) public stakes;
    mapping (address => uint256) public numberOfStakes;
    mapping (address => uint256) public internalRTCBalances;

    modifier validInitialStakeRelease(uint256 _stakeNum) {
        require(stakes[msg.sender][_stakeNum].state == StakeStateEnum.staking);
        require(now >= stakes[msg.sender][_stakeNum].releaseDate && block.number >= stakes[msg.sender][_stakeNum].blockUnlocked);
        require(internalRTCBalances[msg.sender] >= stakes[msg.sender][_stakeNum].initialStake);
        _;
    }

    modifier validRewardWithdrawal(uint256 _stakeNumber) {
        // allow people to withdraw their rewards even if the staking period is over
        require(stakes[msg.sender][_stakeNumber].state == StakeStateEnum.staking || stakes[msg.sender][_stakeNumber].state == StakeStateEnum.staked);
        require(stakes[msg.sender][_stakeNumber].coinsWithdrawn < stakes[msg.sender][_stakeNumber].coinsMinted);
        uint256 currentBlock = block.number;
        uint256 lastBlockWithdrawn = stakes[msg.sender][_stakeNumber].lastBlockWithdrawn;
        require(currentBlock > lastBlockWithdrawn);
        _;
    }

    constructor () {}

    function withdrawReward(uint256 _stakeNumber) external validRewardWithdrawal(_stakeNumber) returns (bool) {
        uint256 currentBlock = block.number;
        if (currentBlock >= stakes[msg.sender][_stakeNumber].blockUnlocked) {
            currentBlock = stakes[msg.sender][_stakeNumber].blockUnlocked;
        }
        uint256 lastBlockWithdrawn = stakes[msg.sender][_stakeNumber].lastBlockWithdrawn;
        uint256 blocksToReward = currentBlock.sub(lastBlockWithdrawn);
        uint256 reward = blocksToReward.mul(stakes[msg.sender][_stakeNumber].rewardPerBlock);
        reward = reward.div(1 ether);
        stakes[msg.sender][_stakeNumber].coinsWithdrawn = stakes[msg.sender][_stakeNumber].coinsWithdrawn.add(reward);
        stakes[msg.sender][_stakeNumber].lastBlockWithdrawn = block.number;
        require(RTI.mint(msg.sender, reward));
    }

    function withdrawInitialStake(uint256 _stakeNumber) external validInitialStakeRelease(_stakeNumber) {
        uint256 initialStake = stakes[msg.sender][_stakeNumber].initialStake;
        stakes[msg.sender][_stakeNumber].state = StakeStateEnum.staked;
        require(RTI.transfer(msg.sender, initialStake));
    }

    function depositStake(uint256 _numRTC)
        external
        returns (bool)
    {
        uint256 stakeCount = getStakeCount(msg.sender);
        (uint256 blockLocked, 
        uint256 blockReleased, 
        uint256 releaseDate, 
        uint256 totalCoinsMinted,
        uint256 rewardPerBlock) = calculateStake(_numRTC);
        StakeStruct memory ss = StakeStruct({
            stakeID: keccak256(blockLocked, blockReleased, releaseDate, totalCoinsMinted, stakeCount),
            initialStake: _numRTC,
            blockLocked: blockLocked,
            blockUnlocked: blockReleased,
            releaseDate: releaseDate,
            coinsMinted: totalCoinsMinted,
            coinsWithdrawn: 0,
            rewardPerBlock: rewardPerBlock,
            lastBlockWithdrawn: block.number,
            state: StakeStateEnum.staking
        });
        stakes[msg.sender][stakeCount] = ss;
        stakeNumToIDMap[msg.sender][stakeCount] = keccak256(blockLocked, blockReleased, releaseDate, totalCoinsMinted, stakeCount);
        stakeIDToNumMap[msg.sender][keccak256(blockLocked, blockReleased, releaseDate, totalCoinsMinted, stakeCount)] = stakeCount;
        numberOfStakes[msg.sender] = numberOfStakes[msg.sender].add(1);
        internalRTCBalances[msg.sender] = internalRTCBalances[msg.sender].add(_numRTC);
        require(RTI.transferFrom(msg.sender, address(this), _numRTC));
        // event place holder
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
        blockReleased = blockLocked.mul(BLOCKHOLDPERIOD);
        releaseDate = now.add(BLOCKHOLDPERIOD.mul(1 seconds));
        totalCoinsMinted = _numRTC.mul(MULTIPLIER);
        totalCoinsMinted = totalCoinsMinted.div(1 ether);
        rewardPerBlock = totalCoinsMinted.div(BLOCKHOLDPERIOD);
    }

    function calculateTotalCoinsMinted(uint256 _numRTC) internal view returns (uint256 totalCoinsMinted) {
        totalCoinsMinted = _numRTC.mul(MULTIPLIER);
        totalCoinsMinted = totalCoinsMinted.div(1 ether);
    }

    function getStakeCount(address _staker) internal view returns (uint256) {
        return numberOfStakes[_staker];
    }


}