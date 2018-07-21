pragma solidity 0.4.24;
pragma experimental "v0.5.0";

interface TokenLockupInterface {
	function setRtI(address _rtcAddress) external returns (bool);
	function setRtHotWallet(address _rtHotWallet) external returns (bool);
	function depositStake(uint256 _rtcToStake, uint256 _durationInWeeks) external returns (bool);
	function routeRtcRewards(address[] _stakers, uint256 _rtcPerStaker) external returns (bool);
	function routeEthRewards(address[] _stakers, uint256 _ethPerStaker) external returns (bool);
	function verifyEnabledStake(address _staker, uint256 _id) external view returns (bool);
	function owner() external view returns (address);
	function admin() external view returns (address);
    function updateRtcPrice(uint256 _rtcUSD) external returns (bool);
}