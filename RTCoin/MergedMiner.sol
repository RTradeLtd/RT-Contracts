pragma solidity 0.4.24;
pragma experimental "v0.5.0";

import "../Math/SafeMath.sol";
import "../Interfaces/RTCoinInterface.sol";

/// @title Merged Miner allows people who mine mainnet Ethereum blocks to also mint RTC
/// @author Postables, RTrade Technologies Ltd
/// @notice Version 1, future versions will require a non-interactive block submissinon method
/// @dev We able V5 for safety features, see https://solidity.readthedocs.io/en/v0.4.24/security-considerations.html#take-warnings-seriously
contract MergedMiner {

    using SafeMath for uint256;

    uint256 constant public SUBMISSIONREWARD = 1;
    uint256 constant public BLOCKREWARD = 1;
    uint256 constant public MINWITHDRAWAL = 1;
    address constant public TOKENADDRESS = 0xB8fe3B2C83014566733B766a27d94CB9AC167Dc6;
    RTCoinInterface constant public RTI = RTCoinInterface(TOKENADDRESS);

    address public blockDeployedAtMiner;
    uint256 public blockDeployedAt;

    enum BlockStateEnum { nil, submitted, claimed }

    struct Blocks {
        uint256 number;
        address coinbase;
        BlockStateEnum state;
    }

    mapping (uint256 => Blocks) public blocks;

    modifier submittedBlock(uint256 _blockNum) {
        require(blocks[_blockNum].state == BlockStateEnum.submitted);
        _;

    }

    modifier nonSubmittedBlock(uint256 _blockNum) {
        require(blocks[_blockNum].state == BlockStateEnum.nil);
        _;
    }

    modifier isCoinbase(uint256 _blockNumber) {
        require(msg.sender == blocks[_blockNumber].coinbase);
        _;
    }

    modifier unclaimed(uint256 _blockNumber) {
        require(blocks[_blockNumber].state == BlockStateEnum.submitted);
        _;
    }

    constructor() public {
        require(TOKENADDRESS != address(0), "token address not set");
        blockDeployedAt = block.number;
        blockDeployedAtMiner = block.coinbase;
    }

    /** @notice Used to submit block hash, and block miner information for the current block
        * @dev Future iterations will avoid this process entirely, and use RLP encoded block headers to parse the data.
     */
    function submitBlock() public nonSubmittedBlock(block.number) returns (bool) {
        Blocks memory b = Blocks({
            number: block.number,
            coinbase: block.coinbase,
            state: BlockStateEnum.submitted
        });
        blocks[block.number] = b;
        require(RTI.mint(msg.sender, SUBMISSIONREWARD), "failed to transfer reward to block submitter");
        return true;
    }

    /** @notice Used by a miner to claim their merged mined RTC
        * @param _blockNumber The block number of the block that the person mined
     */
    function claimReward(uint256 _blockNumber) 
        public 
        isCoinbase(_blockNumber) 
        unclaimed(_blockNumber) 
        submittedBlock(_blockNumber)
        returns (bool) 
    {
        // mark the reward as claimed
        blocks[_blockNumber].state = BlockStateEnum.claimed;
        require(RTI.mint(msg.sender, BLOCKREWARD), "failed to transfer block reward");
        return true;
    }

    /** @notice Used by a miner to bulk claim their merged mined RTC
        * @param _blockNumbers Contains the block numbers for which they want to claim
     */
    function bulkClaimReward(uint256[] _blockNumbers) public returns (bool) {
        for (uint256 i = 0; i < _blockNumbers.length; i++) {
            require(claimReward(_blockNumbers[i]));
        }
        return true;
    }
}