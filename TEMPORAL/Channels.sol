pragma solidity 0.4.24;

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
    RTCoinInterface constant public RTI = RTCoinInterface(TOKENADDRESS);

    // STATE CHANGEABLE
    address public admin;

    enum ChannelState { nil, opened, closed }

    struct ChannelStruct {
        address opener;
        address recipient;      // this is the recipient of the channel, the person allowed to withdraw funds
        uint256 openedAt;       // this is when the channel was opened
        uint256 balance;        // this is the current balance of the channel
        uint256 lastWithdrawal; // this is the time at which a withdrawal was last made
        uint256 numberOfPayments; // this is usedto track the total number of payments made in this channel
        ChannelState state;     // this is the current state of the channel
    }

    mapping (address => ChannelStruct) public channels;

    event ChannelOpened(address _opener, address _recipient, uint256 _amount, uint256 _durationInSeconds);

    modifier onlyRecipient() {
        require(msg.sender == RECIPIENT);
        _;
    }

    modifier noExistingChannel() {
        require(channels[msg.sender].state == ChannelState.nil);
        _;
    }

    modifier openedChannel(address _channelOpener) {
        require(channels[_channelOpener].state == ChannelState.opened);
        _;
    }

    modifier inactiveChannel() {
        uint256 lastWithdrawal = channels[msg.sender].lastWithdrawal;
        uint256 closureDate = lastWithdrawal.add(30 days);
        require(now >= closureDate);
        _;
    }

    modifier channelOpener() {
        require(msg.sender == channels[msg.sender].opener);
        _;
    }

    constructor() public {
        // prevent deployments if the token address hasnt been setup
        require(TOKENADDRESS != address(0), "token address not set");
        // prevent deployments if recipient is not set
        require(RECIPIENT != address(0), "recipient not set");
        // prevent deployments if the admin is the recipient
        require(msg.sender != RECIPIENT, "admin can't be the recipient");
        admin = msg.sender;
    }

    function openChannel(
        uint256 _channelBalance,
        uint256 _durationInSeconds)
        public
        noExistingChannel
        returns (bool)
    {
        ChannelStruct memory cs = ChannelStruct({
            opener: msg.sender,
            recipient: RECIPIENT,
            openedAt: now,
            balance: _channelBalance,
            lastWithdrawal: now,
            numberOfPayments: 0,
            state: ChannelState.opened
        });
        channels[msg.sender] = cs;
        emit ChannelOpened(msg.sender, RECIPIENT, _channelBalance, _durationInSeconds);
        require(RTI.transferFrom(msg.sender, address(this), _channelBalance), "transfer from failed, most likely needs approval");
        return true;
    }

    // allows channel opener to close the chanel and withdraw their funds after the inactive period
    function closeChannel() public channelOpener inactiveChannel returns (bool) {
        uint256 balance = channels[msg.sender].balance;
        channels[msg.sender].balance = 0;
        channels[msg.sender].state = ChannelState.closed;
        require(RTI.transfer(msg.sender, balance), "failed to close channel");
        return true;
    }

    function withdrawFromChannel(
        address _channelOpener,
        bytes32 _h,
        uint8   _v,
        bytes32 _r,
        bytes32 _s,
        uint256 _amount,
        uint256 _paymentNumber,
        bool    _prefixed)
        public
        onlyRecipient
        openedChannel(_channelOpener)
        returns (bool)
    {
        require(_paymentNumber > channels[_channelOpener].numberOfPayments, "invalid payment number, less than what is in storage");
        bytes32 image;
        if (_prefixed) {
            image = generatePreimage(_channelOpener, _amount, _paymentNumber);
            image = generatePrefixedPreimage(image);
        } else {
            image = generatePreimage(_channelOpener, _amount, _paymentNumber);
        }
        require(image == _h, "constructed preimage does not match hash");
        address signer = ecrecover(_h, _v, _r, _s);
        require(signer == _channelOpener, "recovered signer is not channel opener");
        require(channels[_channelOpener].balance >= _amount, "not enough channel balance");
        channels[_channelOpener].balance = channels[_channelOpener].balance.sub(_amount);
        channels[_channelOpener].numberOfPayments = _paymentNumber;
        channels[_channelOpener].lastWithdrawal = now;
        require(RTI.transfer(msg.sender, _amount), "failed to withdraw funds");
        return true;
    }

    // allows channel opener to replenish the channel balance
    // do note that this will "reset" the last withdrawal date
    function depositFundsIntoChannel(
        uint256 _amount)
        public
        openedChannel(msg.sender)
        returns (bool)
    {
        channels[msg.sender].balance = channels[msg.sender].balance.add(_amount);
        channels[msg.sender].lastWithdrawal = now;
        require(RTI.transferFrom(msg.sender, address(this), _amount), "transfer from failed, most likely needs approval");
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
}