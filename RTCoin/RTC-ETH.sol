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
    address constant private TOKENADDRESS = 0xecc043b92834c1ebDE65F2181B59597a6588D616;
    RTCoinInterface constant public RTI = RTCoinInterface(TOKENADDRESS);
    string constant public VERSION = "production";

    address public hotWallet;
    uint256 public ethUSD;
    uint256 public weiPerRtc;
    bool   public locked;

    event EthUsdPriceUpdated(uint256 _ethUSD);
    event EthPerRtcUpdated(uint256 _ethPerRtc);
    event RtcPurchased(uint256 _rtcPurchased);
    event ForeignTokenTransfer(address indexed _sender, address indexed _recipient, uint256 _amount);

    modifier notLocked() {
        require(!locked, "sale must not be locked");
        _;
    }

    modifier isLocked() {
        require(locked, "sale must be locked");
        _;
    }

    function lockSales()
        public
        onlyAdmin
        notLocked
        returns (bool)
    {
        locked = true;
        return true;
    }

    function unlockSales()
        public
        onlyAdmin
        isLocked
        returns (bool)
    {
        locked = false;
        return true;
    }

    constructor() public {
        // prevent deployment if the token address isnt set
        require(TOKENADDRESS != address(0), "token address cant be unset");
        locked = true;
    }

    function () external payable {
        require(msg.data.length == 0, "data length must be 0");
        require(buyRtc(), "buying rtc failed");
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
        onlyOwner
        isLocked
        returns (bool)
    {
        hotWallet = _hotWalletAddress;
        return true;
    }

    function withdrawRemainingRtc()
        public
        onlyOwner
        isLocked
        returns (bool)
    {
        require(RTI.transfer(msg.sender, RTI.balanceOf(address(this))), "transfer failed");
        return true;
    }

    function buyRtc()
        public
        payable
        notLocked
        returns (bool)
    {
        require(hotWallet != address(0), "hot wallet cant be unset");
        require(msg.value > 0, "msg value must be greater than zero");
        uint256 rtcPurchased = (msg.value.mul(1 ether)).div(weiPerRtc);
        hotWallet.transfer(msg.value);
        require(RTI.transfer(msg.sender, rtcPurchased), "transfer failed");
        emit RtcPurchased(rtcPurchased);
        return true;
    }

    /** @notice Allow us to transfer tokens that someone might've accidentally sent to this contract
        @param _tokenAddress this is the address of the token contract
        @param _recipient This is the address of the person receiving the tokens
        @param _amount This is the amount of tokens to send
     */
    function transferForeignToken(
        address _tokenAddress,
        address _recipient,
        uint256 _amount)
        public
        onlyAdmin
        returns (bool)
    {
        require(_recipient != address(0), "recipient address can't be empty");
        // don't allow us to transfer RTC tokens stored in this contract
        require(_tokenAddress != TOKENADDRESS, "token can't be RTC");
        ERC20Interface eI = ERC20Interface(_tokenAddress);
        require(eI.transfer(_recipient, _amount), "token transfer failed");
        emit ForeignTokenTransfer(msg.sender, _recipient, _amount);
        return true;
    }
}