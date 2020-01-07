pragma solidity 0.4.24;
pragma experimental "v0.5.0";

import "../Math/SafeMath.sol";
import "../Modules/Administration.sol";
import "../Interfaces/RTCoinInterface.sol";

/*
This contract is used to facilitate payments between frequent TEMPORAL users (ie, API users). It allows us to continue the same payment model as irregular users,
however we don't have to commit a transaction to the blockchain for each payment.  Whenever the user wishes to upload something through our system, they will 
generate valid signature data. This data is given to us, and then validated. If validation is successful, then the content is injected into our system. RTRade
can then utilize these signatures to redeem our RTC/ETH whenever we wish. 

By doing this, we allow users smart contract validated, per-upload payments in a gas efficient manner.
*/

contract Channels is Administration {
    using SafeMath for uint256;

    // CONSTANTS
    bytes constant private PREFIX = "\x19Ethereum Signed Message:\n32";
    address constant public RECIPIENT = address(0);
    address constant public TOKENADDRESS = address(0);
    address constant public HOTWALLET = address(0);
    uint256 constant private BLOCKACTIVITYLIMIT = 172800; // 172800 blocks
    uint256 constant private TIMEACTIVITYLIMIT = 2592000; // 2592000 seconds
    RTCoinInterface constant public RTI = RTCoinInterface(TOKENADDRESS);

    // STATE CHANGEABLE
    address public admin;

    enum ChannelState { nil, opened, closed }
    enum PaymentMethod { RTC, ETH }

    struct ChannelStruct {
        address owner;          // this is the owner/opener of the channel
        address recipient;      // this is the recipient of the channel, the person allowed to withdraw funds
        uint256 balance;        // this is the current balance of the channel
        uint256 paymentNumber; // this is usedto track the total number of payments made in this channel
        uint256 lastBlockActivity; // this records the last block at which an action was performed on this channel
        uint256 lastTimeActivity; // this records the last time at whichan action was performed
        ChannelState state;     // this is the current state of the channel
        PaymentMethod method;   // this is the payment method
    }

    mapping (address => ChannelStruct) public channels;
    mapping (address => uint256) public ethBalances;
    mapping (address => uint256) public rtcBalances;

    event ChannelOpened(address _opener, address _recipient, uint256 _amount);

    modifier validPaymentMethod(uint8 _method) {
        require(PaymentMethod(_method) == PaymentMethod.RTC || PaymentMethod(_method) == PaymentMethod.ETH, "payment method must be ETH or RTC");
        _;
    }

    modifier validPaymentNumber(address _channelOwner, uint256 _number) {
        require(_number > channels[_channelOwner].paymentNumber);
        _;
    }

    modifier isActiveChannel(address _channelOwner) {
        require(channels[_channelOwner].state == ChannelState.opened);
        _;
    }

    modifier isInactiveChannel() {
        require(channels[msg.sender].state == ChannelState.opened, "channel must be opened");
        require(checkForInactiveChannel(msg.sender), "can't execute before the block count and time count have passed");
        _;
    }

    // make sure the sender has neither an opened, nor closed channel
    modifier noActiveChannel() {
        require(channels[msg.sender].state == ChannelState.nil, "sender can't have an active channel");
        _;
    }

    modifier hasValidBalance(address _channelOwner, uint256 _chargeAmountInWei) {
        require(_chargeAmountInWei <= channels[_channelOwner].balance);
        _;
    }

    modifier onlyRecipient() {
        require(msg.sender == RECIPIENT, "sender must be the recipient");
        _;
    }

    modifier onlyChannelOwner() {
        require(msg.sender == channels[msg.sender].owner, "sender must be channel owner");
        _;
    }

    function () payable external {

    }

    constructor() public {
        require(TOKENADDRESS != address(0), "token address can't be unset");
        require(HOTWALLET != address(0), "hot wallet can't be unset");
        require(RECIPIENT != address(0), "recipient can't be unset");
    }

    function openChannel(
        address _recipient,
        uint256 _channelBalance,
        uint8   _paymentMethod)
        public
        payable
        noActiveChannel
        validPaymentMethod(_paymentMethod)
        returns (bool)
    {
        if (PaymentMethod(_paymentMethod) == PaymentMethod.ETH) {
            require(msg.value == _channelBalance, "msg.value must be equal to _channelBalance");
        }
        ChannelStruct memory cs = ChannelStruct({
            owner: msg.sender,
            recipient: _recipient,
            balance: _channelBalance,
            paymentNumber: 0,
            lastBlockActivity: block.number,
            lastTimeActivity: now,
            state: ChannelState.opened,
            method: PaymentMethod(_paymentMethod)
        });
        channels[msg.sender] = cs;
        if (PaymentMethod(_paymentMethod) == PaymentMethod.RTC) {
            require(RTI.transferFrom(msg.sender, address(this), _channelBalance), "failed to execute transferFrom, likely needs approval");
        }
        return true;
    }

    function redeemPayment(
        address _channelOwner,
        uint256 _paymentNumber,
        uint256 _chargeAmountInWei,
        uint8   _paymentMethod)
        public 
        payable 
        onlyRecipient 
        validPaymentNumber(_channelOwner, _paymentNumber)
        validPaymentMethod(_paymentMethod)
        isActiveChannel(_channelOwner)
        hasValidBalance(_channelOwner, _chargeAmountInWei)
        returns (bool) 
    {
        return true;
    }

    function updateChannelBalance(uint256 _amount) public payable onlyChannelOwner returns (bool) {
        if (channels[msg.sender].method == PaymentMethod.ETH) {
            require(msg.value == _amount, "msg.value must be equal to _amount");
            channels[msg.sender].lastBlockActivity = block.number;
            channels[msg.sender].lastTimeActivity = now;
            channels[msg.sender].balance = channels[msg.sender].balance.add(_amount);
            ethBalances[msg.sender] = ethBalances[msg.sender].add(_amount);
            return true;
        } else if (channels[msg.sender].method == PaymentMethod.RTC) {
            channels[msg.sender].lastBlockActivity = block.number;
            channels[msg.sender].lastTimeActivity = now;
            channels[msg.sender].balance = channels[msg.sender].balance.add(_amount);
            require(RTI.transferFrom(msg.sender, address(this), _amount), "failed to execute transferFrom, likely needs approval");
            return true;
        }
        return false;
    }

    function closeChannel() public onlyChannelOwner isInactiveChannel returns (bool) {
        uint256 balance = channels[msg.sender].balance;
        channels[msg.sender].balance = 0;
        channels[msg.sender].state = ChannelState.closed;
        require(ethBalances[msg.sender] == balance, "eth balances must be equal to channel balance");
        ethBalances[msg.sender] = 0;
        if (channels[msg.sender].method == PaymentMethod.ETH) {
            msg.sender.transfer(balance);
            return true;
        } else {
            require(RTI.transfer(msg.sender, balance), "transfer failed");
            return true;
        }
        return false;
    }

    function generatePreimage(
        address _channelOpener,
        uint256 _chargeAmountInWei,
        uint256 _paymentNumber)
        internal
        view
        returns (bytes32 preimage)
    {
        return keccak256(abi.encodePacked(_channelOpener, msg.sender, _paymentNumber, _chargeAmountInWei));
    }

    function generatePrefixedPreimage(bytes32 _preimage) internal pure returns (bytes32)  {
        return keccak256(abi.encodePacked(PREFIX, _preimage));
    }

    function checkForInactiveChannel(address _channelOwner) public view returns (bool) {
        uint256 inactiveBlockNumber = channels[_channelOwner].lastBlockActivity.add(BLOCKACTIVITYLIMIT);
        require(block.number <= inactiveBlockNumber, "172800 blocks must've passed since last action");
        uint256 inactivityDate = channels[_channelOwner].lastTimeActivity.add(TIMEACTIVITYLIMIT);
        require(now <= inactivityDate, "2592000 seconds must've passed since last action");
        return true;
    }
}