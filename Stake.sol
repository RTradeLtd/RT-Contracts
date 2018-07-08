pragma solidity 0.4.24;

import "./Interfaces/RTCoinInterface.sol";
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";

contract Stake is Administration {
    using SafeMath for uint256;

    RTCoinInterface constant public RTI = RTCoinInterface(address(0));
    uint256 constant public MULTIPLIER = 10000000000000000;
    uint256 constant public BLOCKHOLDPERIOD = 2103840;
    uint256 constant public BLOCKSEC = 13;

    enum StakeStateEnum { nil, pending, registered, finished }

    struct StakeStruct {
        bytes32 stakeID;
        uint256 stakeAmount;
        uint256 blockLocked;
        uint256 blockUnlocked;
        uint256 releaseDate;
        uint256 coinsMinted;
        uint256 rewardPerBlock;
        uint256 lastBlockWithdrawn;
    }

    mapping (address => mapping (uint256 => bytes32)) public stakeNumToIDMap;
    mapping (address => mapping (bytes32 => uint256)) public stakeIDToNumMap;
    mapping (address => mapping (uint256 => StakeStruct)) public stakes;
    mapping (address => uint256) public numberOfStakes;
    mapping (address => uint256) public internalRTCBalances;

    modifier isWholeNumber(uint256 _number) {
        assert(_number % 1 == 0);
        _;
    }
    
    modifier validInitialStakeRelease(uint256 _stakeNum, address _staker) {
        require(now >= stakes[_staker][_stakeNum].releaseDate && block.number >= stakes[_staker][_stakeNum].blockUnlocked);
        _;
    }

    constructor () {}

    function depositStake(uint256 _numRTC)
        external
        isWholeNumber(_numRTC)
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
            stakeAmount: _numRTC,
            blockLocked: blockLocked,
            blockUnlocked: blockReleased,
            releaseDate: releaseDate,
            coinsMinted: totalCoinsMinted,
            rewardPerBlock: rewardPerBlock,
            lastBlockWithdrawn: block.number
        });
        stakes[msg.sender][stakeCount] = ss;
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