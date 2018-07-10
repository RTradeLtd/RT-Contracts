pragma solidity 0.4.24;

interface StakeInterface {
    function activeStakes() external view returns (uint256);
}