pragma solidity 0.4.24;

import "../Interfaces/RTCoinInterface.sol";
import "../Math/SafeMath.sol";

/*
This contract is used to facilitate payments for file uploads, or content pins through TEMPORAL.
It is inteded to be utilized by infrequent uploaders, who don't need a payment channel, but want smart contract validated payments.
The way it works is that when attempting to upload a file, or pin content to TEMPORAL, the size of the data is calculated, and a price in USD is determined based
on how may months you want the content in our system for. TEMPORAL will generate valid signature data, which when submitted to a smart contract will be validated.
IF validation passes the specified amount of RTC or ETH is taken from your account, and sent to one of our hot wallets. After transaction confirmation, the data
will be injected into our system.
*/

contract Payments {
    using SafeMath for uint256;    
    bytes constant private PREFIX = "\x19Ethereum Signed Message:\n32";
    // this is the address we use to sign payments
    address constant public SIGNER = 0x7E4A2359c745A982a54653128085eAC69E446DE1;
    address constant public TOKENADDRESS = address(0);
    address constant public HOTWALLET = address(0);
    RTCoinInterface constant public RTI = RTCoinInterface(TOKENADDRESS);

    // PaymentState will keep track of the state of a payment, nil means we havent seen th payment before
    enum PaymentState{ nil, paid }
    // How payments can be made, RTC or eth
    enum PaymentMethod{ RTC, ETH }

    struct PaymentStruct {
        uint256 paymentNumber;
        uint256 chargeAmountInWei;
        PaymentMethod method;
        PaymentState state;
    }

    mapping (address => uint256) public numPayments;
    mapping (address => mapping(uint256 => PaymentStruct)) public payments;

    event PaymentMade(address _payer, uint256 _paymentNumber, uint8 _paymentMethod, uint256 _paymentAmount);

    modifier validPayment(uint256 _paymentNumber) {
        require(payments[msg.sender][_paymentNumber].state == PaymentState.nil, "payment already made");
        _;
    }

    function makePayment(
        bytes32 _h,
        uint8   _v,
        bytes32 _r,
        bytes32 _s,
        uint256 _paymentNumber,
        uint8   _paymentMethod,
        uint256 _chargeAmountInWei,
        bool   _prefixed) // this allows us to sign messages on our own, without prefix https://github.com/ethereum/EIPs/issues/191
        public
        payable
        validPayment(_paymentNumber)
        returns (bool)
    {
        require(_paymentMethod == 0 || _paymentMethod == 1, "invalid payment method");
        bytes32 image;
        if (_prefixed) {
            bytes32 preimage = generatePreimage(_paymentNumber, _chargeAmountInWei, _paymentMethod);
            image = generatePrefixedPreimage(preimage);
        } else {
            image = generatePreimage(_paymentNumber, _chargeAmountInWei, _paymentMethod);
        }
        // ensure that the preimages construct properly
        require(image == _h, "reconstructed preimage does not match");
        address signer = ecrecover(_h, _v, _r, _s);
        // ensure that we actually signed this message
        require(signer == SIGNER, "recovered signer does not match");
        PaymentStruct memory ps = PaymentStruct({
            paymentNumber: _paymentNumber,
            chargeAmountInWei: _chargeAmountInWei,
            method: PaymentMethod(_paymentMethod),
            state: PaymentState.paid
        });
        payments[msg.sender][_paymentNumber] = ps;
        numPayments[msg.sender] = numPayments[msg.sender].add(1);
        if (PaymentMethod(_paymentMethod) == PaymentMethod.ETH) {
            require(msg.value == _chargeAmountInWei, "msg.value does not equal charge amount");
            emit PaymentMade(msg.sender, _paymentNumber, _paymentMethod, _chargeAmountInWei);
            HOTWALLET.transfer(msg.value);
            return true;
        }
        emit PaymentMade(msg.sender, _paymentNumber, _paymentMethod, _chargeAmountInWei);
        require(RTI.transferFrom(msg.sender, HOTWALLET, _chargeAmountInWei), "trasferFrom failed, most likely needs approval");
        return true;
    }

    function verifyImages(
        bytes32 _h,
        uint256 _paymentNumber,
        uint8   _paymentMethod,
        uint256 _chargeAmountInWei,
        bool   _prefixed)
        public
        view
        returns (bool)
    {
        require(_paymentMethod == 0 || _paymentMethod == 1, "invalid payment method");
        bytes32 image;
        if (_prefixed) {
            bytes32 preimage = generatePreimage(_paymentNumber, _chargeAmountInWei, _paymentMethod);
            image = generatePrefixedPreimage(preimage);
        } else {
            image = generatePreimage(_paymentNumber, _chargeAmountInWei, _paymentMethod);
        }
        return image == _h;
    }

    function verifySigner(
        bytes32 _h,
        uint8   _v,
        bytes32 _r,
        bytes32 _s,
        uint256 _paymentNumber,
        uint8   _paymentMethod,
        uint256 _chargeAmountInWei,
        bool   _prefixed)
        public
        view
        returns (bool)
    {
        require(_paymentMethod == 0 || _paymentMethod == 1, "invalid payment method");
        bytes32 image;
        if (_prefixed) {
            bytes32 preimage = generatePreimage(_paymentNumber, _chargeAmountInWei, _paymentMethod);
            image = generatePrefixedPreimage(preimage);
        } else {
            image = generatePreimage(_paymentNumber, _chargeAmountInWei, _paymentMethod);
        }
        require(image == _h, "failed to reconstruct preimages");
        return ecrecover(_h, _v, _r, _s) == SIGNER;
    }

    function generatePreimage(
        uint256 _paymentNumber,
        uint256 _chargeAmountInWei,
        uint8   _paymentMethod)
        internal
        view
        returns (bytes32 preimage)
    {
        return keccak256(abi.encodePacked(msg.sender, _paymentNumber, _paymentMethod, _chargeAmountInWei));
    }

    function generatePrefixedPreimage(bytes32 _preimage) internal pure returns (bytes32)  {
        return keccak256(abi.encodePacked(PREFIX, _preimage));
    }
}