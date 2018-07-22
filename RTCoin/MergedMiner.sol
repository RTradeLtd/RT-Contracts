pragma solidity 0.4.24;
pragma experimental "v0.5.0";

import "../Math/SafeMath.sol";
import "../Interfaces/RTCoinInterface.sol";

/// @title Merged Miner Validator allows people who mine mainnet Ethereum blocks to also mint RTC
/// @author Postables, RTrade Technologies Ltd
/// @notice Version 1, future versions will require a non-interactive block submissinon method
/// @dev We able V5 for safety features, see https://solidity.readthedocs.io/en/v0.4.24/security-considerations.html#take-warnings-seriously
contract MergedMinerValidator {

    using SafeMath for uint256;
    
    // 0.5
    uint256 constant public SUBMISSIONREWARD = 500000000000000000;
    // 0.3
    uint256 constant public BLOCKREWARD = 300000000000000000;
    address constant public TOKENADDRESS = 0xB8fe3B2C83014566733B766a27d94CB9AC167Dc6;
    RTCoinInterface constant public RTI = RTCoinInterface(TOKENADDRESS);
    address public admin = address(0);

    enum BlockStateEnum { nil, submitted, claimed }

    struct Blocks {
        uint256 number;
        address coinbase;
        BlockStateEnum state;
    }

    mapping (uint256 => Blocks) public blocks;

    event BlockInformationSubmitted(address indexed _coinbase, uint256 indexed _blockNumber, address _submitter);
    event MergedMinedRewardClaimed(address indexed _claimer, uint256 indexed _blockNumber, uint256 _blockReward);

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

    modifier onlyAdmin() {
        require(msg.sender == admin);
        _;
    }
    constructor() public {
        require(TOKENADDRESS != address(0), "token address not set");
        admin = msg.sender;
        Blocks memory b = Blocks({
            number: block.number,
            coinbase: block.coinbase,
            state: BlockStateEnum.submitted
        });
        blocks[block.number] = b;
        // we use address(0) and don't mint any tokens, since "we are submitting the information" 
        emit BlockInformationSubmitted(block.coinbase, block.number, address(0));
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
        // lets not do a storage lookup so we can avoid SSLOAD gas usage
        emit BlockInformationSubmitted(block.coinbase, block.number, msg.sender);
        require(RTI.mint(msg.sender, SUBMISSIONREWARD), "failed to transfer reward to block submitter");
        return true;
    }

    /** @notice Used by a miner to claim their merged mined RTC
        * @param _blockNumber The block number of the block that the person mined
     */
    function claimReward(uint256 _blockNumber) 
        internal
        isCoinbase(_blockNumber) 
        unclaimed(_blockNumber) 
        submittedBlock(_blockNumber)
        returns (uint256) 
    {
        // mark the reward as claimed
        blocks[_blockNumber].state = BlockStateEnum.claimed;
        return BLOCKREWARD;
    }

    /** @notice Used by a miner to bulk claim their merged mined RTC
        * @dev To prevent expensive looping, we throttle to 20 withdrawals at once
        * @param _blockNumbers Contains the block numbers for which they want to claim
     */
    function bulkClaimReward(uint256[] _blockNumbers) public returns (bool) {
        require(_blockNumbers.length <= 20, "can only claim up to 20 rewards at once");
        uint256 totalMint;
        // here we grab the reward number, so that we don't have to retrieve it from storage every loop, and instead, only retrieve once
        // that means we only have to pay the SSLOAD gas once, and instead us MLOAD each iteration
        uint256 reward = BLOCKREWARD;
        for (uint256 i = 0; i < _blockNumbers.length; i++) {
            // update their total amount minted
            totalMint = totalMint.add(claimReward(_blockNumbers[i]));
            // emit an event
            emit MergedMinedRewardClaimed(
                blocks[_blockNumbers[i]].coinbase,
                blocks[_blockNumbers[i]].number,
                reward
            );
        }
        require(RTI.mint(msg.sender, totalMint), "unable to mint tokens");
        return true;
    }

    /** @notice Used to destroy the contract
     */
    function goodNightSweetPrince() public onlyAdmin returns (bool) {
        selfdestruct(msg.sender);
        return true;
    }
}