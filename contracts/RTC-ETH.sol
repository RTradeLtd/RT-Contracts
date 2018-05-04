pragma solidity 0.4.23;

import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";
import "./Interfaces/RTCoinInterface.sol";
import "./Interfaces/ERC20Interface.sol";

contract RTCETH is Administration {
    using SafeMath for uint256;

    address public hotWallet;
    uint256 public ethUSD;
    uint256 public weiPerRtc;
    bool   public locked;

    RTCoinInterface public rtI;

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
        locked = true;
        rtI = RTCoinInterface(address(0xb4ed44372bbc71dad64956373214c667b717e805));
    }

    function () public payable {
        require(buyRtc());
    }

    function setRtInterfae(
        address _rtcAddress)
        public
        onlyAdmin
        isLocked
        returns (bool)
    {
        rtI = RTCoinInterface(_rtcAddress);
        // event place holder
        return true;
    }


    function updateEthPrice(
        uint256 _ethUSD)
        public
        onlyAdmin
        returns (bool)
    {
        ethUSD = _ethUSD;
        uint256 oneEth = 1 ether;
        uint256 oneUsdOfEth = oneEth.div(ethUSD);
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
        require(rtI.transfer(msg.sender, rtI.balanceOf(address(this))));
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
        require(rtI.transfer(msg.sender, rtcPurchased));
        return true;
    }
}