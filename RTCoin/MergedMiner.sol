pragma solidity 0.4.24;

import "../Math/SafeMath.sol";
import "../Interfaces/RTCoinInterface.sol";

contract MinerValidator {

    using SafeMath for uint256;

    uint256 constant public SUBMISSIONREWARD = 1;
    uint256 constant public BLOCKREWARD = 1;
    uint256 constant public MINWITHDRAWAL = 1;
    address constant public TOKENADDRESS = 0x185ae6A87BBB02097923e859D742747Bb979Ae9a;
    RTCoinInterface constant public RTI = RTCoinInterface(TOKENADDRESS);

    struct RewardStruct {
        uint256 totalRewards;
        uint256 balance;
    }

    struct Blocks {
        uint256 number;
        address coinbase;
        bytes32 hash;
        bool set;
        bool claimed;
    }

    mapping (uint256 => Blocks) public blocks;

    modifier set(uint256 _blockNum) {
        require(blocks[_blockNum].set);
        _;

    }
    modifier notSet(uint256 _blockNum) {
        require(!blocks[_blockNum].set);
        _;
    }

    modifier isCoinbase(uint256 _blockNumber) {
        require(msg.sender == blocks[_blockNumber].coinbase);
        _;
    }

    modifier unclaimed(uint256 _blockNumber) {
        require(!blocks[_blockNumber].claimed);
        _;
    }

    constructor() public {
        require(TOKENADDRESS != address(0), "token address not set");
    }

    function submitBlock() public notSet(block.number) returns (bool) {
        Blocks memory b = Blocks({
            number: block.number,
            coinbase: block.coinbase,
            hash: blockhash(block.number),
            set: true,
            claimed: false
        });
        blocks[block.number] = b;
        require(RTI.transfer(msg.sender, SUBMISSIONREWARD), "failed to transfer reward to block submitter");
        return true;
    }

    function claimReward(uint256 _blockNumber) public isCoinbase(_blockNumber) unclaimed(_blockNumber) set(_blockNumber) returns (bool) {
        blocks[_blockNumber].claimed = true;
        require(RTI.transfer(msg.sender, BLOCKREWARD), "failed to transfer block reward");
        return true;
    }

}