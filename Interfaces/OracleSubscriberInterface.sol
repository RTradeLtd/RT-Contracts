pragma solidity 0.4.24;
pragma experimental "v0.5.0";

interface OracleSubscriberInterface {
    function updateRtcPrice(uint256 _rtcUSD) external returns (bool);
    function updateEthPrice(uint256 _ethUSD) external returns (bool);
}