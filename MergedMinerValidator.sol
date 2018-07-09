pragma solidity 0.4.24;

/*
This contract was originally written by the following author, Chance Santana-Wees
I (Postables, RTrade Technologies LTD have repurposed it for merged mining rewards

Original:
Author: Chance Santana-Wees
Contact Email: figs999@gmail.com
*/

// Version 1
    
import "./Math/SafeMath.sol";
import "./Interfaces/RTCoinInterface.sol";

contract MergedMinerValidator {

    using SafeMath for uint256;

    address constant private TOKENCONTRACT = address(0);
    // set to 1000 for testing purposes
    uint256 constant private MINWITHDRAWAL = 1000000000000000000000;
    // set to 0.1  for testing purposes
    uint256 constant private BLOCKREWARD = 100000000000000000000;
    RTCoinInterface constant private RTI = RTCoinInterface(TOKENCONTRACT);

    struct BlockHeader {
        bytes32 derivedHash;        
        bytes32 parentHash;         
        bytes32 ommersHash;         
        
        bytes32 stateRoot;          
        bytes32 transactionsRoot;   
        bytes32 receiptsRoot;         
        
        bytes32 mixHash;            
        bytes32 extraData;          
        
        address miner;              
        
        bytes8 nonce;               

        uint difficulty;            
        uint32 blockNumber;         
        uint32 gasLimit;            
        uint32 gasUsed;             
        uint32 timeStamp;           
                                    
        bytes logsBloom;
    }
    
    struct Miners {
        uint256 totalMined;
        uint256 currentBalance;
    }

    enum BlockState {nil, stored}

    struct Block {
        bytes32 blockHash;
        BlockState state;
    }
    mapping (address => Miners) public miners;
    mapping (uint256 => Block) public blocks;

    modifier validClaim() {
        require(miners[msg.sender].currentBalance >= MINWITHDRAWAL);
        _;
    }

    function submitBlockValidation(bytes _rlpData) public returns (bool) {
        require(parseBlockHeader(_rlpData));
        miners[msg.sender].totalMined = miners[msg.sender].totalMined.add(BLOCKREWARD);
        miners[msg.sender].currentBalance = miners[msg.sender].currentBalance.add(BLOCKREWARD);
        return true;
    }

    function claimTokens() public validClaim returns (bool) {
        uint256 balance = miners[msg.sender].currentBalance;
        miners[msg.sender].currentBalance = 0;
        require(RTI.mint(msg.sender, balance));
        return true;
    }

    function parseBlockHeader(bytes _rlpData) internal view returns (bool) {
        BlockHeader memory parsedHeader;
        
        parsedHeader.derivedHash = keccak256(_rlpData);
        bytes memory logsBloom = new bytes(256);
        
        assembly {
            calldatacopy(add(parsedHeader,32), 104, 32)                 //parentHash
            calldatacopy(add(parsedHeader,64), 137, 32)                 //ommersHash
            calldatacopy(add(parsedHeader,268), 170, 20)                //miner    
            calldatacopy(add(parsedHeader,96), 191, 32)                 //stateRoot
            calldatacopy(add(parsedHeader,128), 224, 32)                //transactionsRoot
            calldatacopy(add(parsedHeader,160), 257, 32)                //receiptsRoot
            
            calldatacopy(add(logsBloom,32), 292, 256)                   //logsBloom
            
            let _size := sub(and(calldataload(517), 0xFF), 128)
            calldatacopy(add(parsedHeader,sub(352,_size)), 549, _size)  //difficulty
            
            let _idx := add(add(549,_size),1)
            _size := sub(and(calldataload(sub(_idx,32)), 0xFF), 128)
            calldatacopy(add(parsedHeader,sub(384,_size)), _idx, _size) //blockNumber
            
            _idx := add(add(_idx,_size),1)
            _size := sub(and(calldataload(sub(_idx,32)), 0xFF), 128)
            calldatacopy(add(parsedHeader,sub(416,_size)), _idx, _size) //gasLimit
            
            _idx := add(add(_idx,_size),1)
            _size := sub(and(calldataload(sub(_idx,32)), 0xFF), 128)
            calldatacopy(add(parsedHeader,sub(448,_size)), _idx, _size) //gasUsed
            
            _idx := add(add(_idx,_size),1)
            _size := sub(and(calldataload(sub(_idx,32)), 0xFF), 128)
            calldatacopy(add(parsedHeader,sub(480,_size)), _idx, _size) //timeStamp
            
            _idx := add(add(_idx,_size),1)
            _size := sub(and(calldataload(sub(_idx,32)), 0xFF), 128)
            calldatacopy(add(parsedHeader,sub(256,_size)), _idx, _size) //extraData
            
            _idx := add(add(_idx,_size),1)
            calldatacopy(add(parsedHeader,192), _idx, 32)               //mixHash

            _idx := add(_idx,33)
            calldatacopy(add(parsedHeader,288), _idx, 8)                //nonce
        }
        
        require(parsedHeader.derivedHash == blockhash(parsedHeader.blockNumber));
        require(parsedHeader.miner == msg.sender);
        return true;
    }
}