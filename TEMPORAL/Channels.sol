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
        uint256 numberOfPayments; // this is usedto track the total number of payments made in this channel
        uint256 lastBlockActivity; // this records the last block at which an action was performed on this channel
        uint256 lastTimeActivity; // this records the last time at whichan action was performed
        ChannelState state;     // this is the current state of the channel
        PaymentMethod method;   // this is the payment method
    }

    mapping (address => ChannelStruct) public channels;

    event ChannelOpened(address _opener, address _recipient, uint256 _amount);

    modifier inactiveChannel() {
        require(checkForInactiveChannel(msg.sender));
        _;
    }

    modifier validPaymentMethod(uint8 _method) {
        require(PaymentMethod(_method) == PaymentMethod.RTC || PaymentMethod(_method) == PaymentMethod.ETH);
        _;
    }

    // make sure the sender has neither an opened, nor closed channel
    modifier noActiveChannel() {
        require(channels[msg.sender].state == ChannelState.nil);
        _;
    }

    modifier onlyRecipient() {
        require(msg.sender == RECIPIENT);
        _;
    }

    modifier onlyChannelOwner() {
        require(msg.sender == channels[msg.sender].owner);
        _;
    }

    function () payable external {
        
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
            require(msg.value == _channelBalance);
        }
        ChannelStruct memory cs = ChannelStruct({
            owner: msg.sender,
            recipient: _recipient,
            balance: _channelBalance,
            numberOfPayments: 0,
            lastBlockActivity: block.number,
            lastTimeActivity: now,
            state: ChannelState.opened,
            method: PaymentMethod(_paymentMethod)
        });
        channels[msg.sender] = cs;
        if (PaymentMethod(_paymentMethod) == PaymentMethod.RTC) {
            require(RTI.transferFrom(msg.sender, address(this), _channelBalance));
        }
        return true;
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