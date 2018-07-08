pragma solidity 0.4.24;

import "./Interfaces/RTCoinInterface.sol";
import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";

contract Stake is Administration {
    using SafeMath for uint256;

    uint256 constant public MULTIPLIER = 10000000000000000;
    uint256 constant public BLOCKHOLDPERIOD = 2103840;

    enum StakeStateEnum { nil, pending, registered, finished }

    struct StakeStruct {
        bytes32 stakeID;
        uint256 stakeAmount;
        uint256 blockLocked;
        uint256 blockUnlocked;
    }

    mapping (address => mapping (uint256 => bytes32)) public stakeNumToIDMap;
    mapping (address => mapping (bytes32 => uint256)) public stakeIDToNumMap;
    mapping (address => mapping (uint256 => StakeStruct)) public stakes;
    mapping (address => uint256) public numberOfStakes;

    constructor () {}
}