pragma solidity 0.4.24;

import "../Math/SafeMath.sol";

/*
This contract is used to handle vesting of RTC tokens
*/

contract Vesting {
    using SafeMath for uint256;


    struct Vest {
        uint256 totalVest;
        
    }
}