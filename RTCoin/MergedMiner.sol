pragma solidity 0.4.24;

import "../Math/SafeMath.sol";
import "../Interfaces/RTCoinInterface.sol";

contract MinerValidator {

    using SafeMath for uint256;

    uint256 constant public SUBMISSIONREWARD = 1;
    uint256 constant public BLOCKREWARD = 1;
    uint256 constant public MINWITHDRAWAL = 1;
    address constant public TOKENADDRESS = 0xB8fe3B2C83014566733B766a27d94CB9AC167Dc6;
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

    struct BlockMiners {
        mapping (uint256 => bool) claims;
        mapping (uint256 => uint256) indexes;
        uint256[] blocks;
    }

    mapping (uint256 => Blocks) public blocks;
    mapping (address => BlockMiners) private miners;

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
        miners[block.coinbase].blocks.push(block.number);
        miners[block.coinbase].claims[block.number] = false;
        miners[block.coinbase].indexes[block.number] = miners[block.coinbase].blocks.length;
        blocks[block.number] = b;
        require(RTI.transfer(msg.sender, SUBMISSIONREWARD), "failed to transfer reward to block submitter");
        return true;
    }

    function claimReward(uint256 _blockNumber) public isCoinbase(_blockNumber) unclaimed(_blockNumber) set(_blockNumber) returns (bool) {
        blocks[_blockNumber].claimed = true;
        miners[msg.sender].claims[block.number] = true;
        uint256 index = miners[msg.sender].indexes[_blockNumber];
        // free up some space
        delete miners[msg.sender].blocks[index];
        // set the claim status
        miners[msg.sender].claims[_blockNumber] = true;
        require(RTI.transfer(msg.sender, BLOCKREWARD), "failed to transfer block reward");
        return true;
    }

    function bulkClaimReward(uint256[] _blocks) public returns (bool) {
        uint256 cumulativeReward;
        for (uint256 i = 0; i < _blocks.length; i++) {
            uint256 num = _blocks[i];
            // if this block has been claimed, skip processing
            if (blocks[num].claimed) {
                continue;
            }
            // if tx submitter isn't coinbase, skip processing
            if (blocks[num].coinbase != msg.sender) {
                continue;
            }
            // also check here that the block wasn't claimed
            if (miners[msg.sender].claims[num]) {
                continue;
            }
            // if the element is empty this has been claimed
            if (miners[msg.sender].indexes[num] == 0) {
                continue;
            }
            uint256 index = miners[msg.sender].indexes[num];
            // free up some space
            delete miners[msg.sender].blocks[index];
            // set claimed to true
            blocks[num].claimed = true;
            // record the users claims
            miners[msg.sender].claims[num] = true;
            // update the cumulative reward
            cumulativeReward = cumulativeReward.add(BLOCKREWARD);
        }
        require(cumulativeReward > 0, "no rewards recorded");
        require(RTI.transfer(msg.sender, cumulativeReward), "failed to send tokens to claimer");
    }

    function getBlocksForMiner(address _miner) public view returns (uint256[]) {
        return miners[_miner].blocks;
    }

}