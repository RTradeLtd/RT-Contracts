pragma solidity 0.4.24;
pragma experimental "v0.5.0";

interface StakeInterface {
    function activeStakes() external view returns (uint256);
}