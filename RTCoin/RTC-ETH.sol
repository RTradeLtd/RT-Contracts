pragma solidity 0.4.24;
pragma experimental "v0.5.0";

import "../Modules/Administration.sol";
import "../Math/SafeMath.sol";
import "../Interfaces/RTCoinInterface.sol";
import "../Interfaces/ERC20Interface.sol";

/// @title RTCETH allows the sale of RTC for ETH with an updatable ETH price
/// @author Postables, RTrade Technologies Ltd
/// @dev We able V5 for safety features, see https://solidity.readthedocs.io/en/v0.4.24/security-considerations.html#take-warnings-seriously
contract RTCETH is Administration {
    using SafeMath for uint256;

    // we mark as constant private to save gas
    address constant private TOKENADDRESS = address(0);
    RTCoinInterface constant public RTI = RTCoinInterface(TOKENADDRESS);
    string constant public VERSION = "production";

    address public hotWallet;
    uint256 public ethUSD;
    uint256 public weiPerRtc;
    bool   public locked;

    event EthUsdPriceUpdated(uint256 _ethUSD);
    event EthPerRtcUpdated(uint256 _ethPerRtc);
    event RtcPurchased(uint256 _rtcPurchased);

    modifier notLocked() {
        require(!locked);
        _;
    }

    modifier isLocked() {
        require(locked);
        _;
    }

    function lockSales()
        public
        onlyAdmin
        notLocked
        returns (bool)
    {
        locked = true;
        // place holder
        return true;
    }

    function unlockSales()
        public
        onlyAdmin
        isLocked
        returns (bool)
    {
        locked = false;
        // place holder
        return true;
    }

    constructor() public {
        // prevent deployment if the token address isnt set
        require(TOKENADDRESS != address(0));
        locked = true;
    }

    function () external payable {
        require(buyRtc());
    }

    function updateEthPrice(
        uint256 _ethUSD)
        public
        onlyAdmin
        returns (bool)
    {
        ethUSD = _ethUSD;
        uint256 oneEth = 1 ether;
        // here we calculate how many ETH 1 USD is worth
        uint256 oneUsdOfEth = oneEth.div(ethUSD);
        // for the duration of this contract, RTC will be at a fixed price of 0.125USD, which divides into 1 8 times
        weiPerRtc = oneUsdOfEth.div(8);
        emit EthUsdPriceUpdated(ethUSD);
        emit EthPerRtcUpdated(weiPerRtc);
        return true;
    }

    function setHotWallet(
        address _hotWalletAddress)
        public
        onlyAdmin
        isLocked
        returns (bool)
    {
        hotWallet = _hotWalletAddress;
        return true;
    }

    function withdrawRemainingRtc()
        public
        onlyAdmin
        isLocked
        returns (bool)
    {
        require(RTI.transfer(msg.sender, RTI.balanceOf(address(this))));
        return true;
    }

    function buyRtc()
        public
        payable
        notLocked
        returns (bool)
    {
        require(hotWallet != address(0));
        require(msg.value > 0);
        uint256 rtcPurchased = (msg.value.div(weiPerRtc)).mul(1 ether);
        emit RtcPurchased(rtcPurchased);
        hotWallet.transfer(msg.value);
        emit RtcPurchased(rtcPurchased);
        require(RTI.transfer(msg.sender, rtcPurchased));
        return true;
    }
}