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

/** @title This contract is used to handle payments for TEMPORAL */
contract Payments {
    using SafeMath for uint256;    
    bytes constant private PREFIX = "\x19Ethereum Signed Message:\n32";
    // this is the address we use to sign payments
    address constant public SIGNER = 0x7E4A2359c745A982a54653128085eAC69E446DE1;
    address constant public TOKENADDRESS = 0x185ae6A87BBB02097923e859D742747Bb979Ae9a;
    address constant public HOTWALLET = 0x7E4A2359c745A982a54653128085eAC69E446DE1;
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

    /** @dev Used to submit a payment for TEMPORAL uploads
        * @param _h This is the message hash that has been signed
        * @param _v This is pulled from the signature
        * @param _r This is pulled from the signature
        * @param _s This is pulled from the signature
        * @param _paymentNumber This is the current payments number (how many payments the user has submitted)
        * @param _paymentMethod This is the payment method (RTC, ETH) being used
        * @param _chargeAmountInWei This is how much the user is to be charged
        * @param _prefixed This indicates whether or not the signature was generated using ERC191 standards
     */
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

    /** @dev This is a helper function used to verify whether or not the provided arguments can reconstruct the message hash
        * @param _h This is the message hash which is signed, and will be reconstructed
        * @param _paymentNumber This is the number of payment
        * @param _paymentMethod This is the payment method (RTC, ETH) being used
        * @param _chargeAmountInWei This is the amount the user is to be charged
        * @param _prefixed This indicates whether the message was signed according to ERC191
     */
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

    /** @dev This is a helper function which can be used to verify the signer of a message
        * @param _h This is the message hash that is signed
        * @param _v This is pulled from the signature
        * @param _r This is pulled from the signature
        * @param _s This is pulled from the signature
        * @param _paymentNumber This is the payment number of this particular payment
        * @param _paymentMethod This is the payment method (RTC, ETH) being used
        * @param _chargeAmountInWei This is the amount hte user is to be charged
        * @param _prefixed This indicates whether or not the message was signed using ERC191
     */
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

    /** @dev This is a helper function used to generate a non ERC191 signed message hash
        * @param _paymentNumber This is the payment number of this payment
        * @param _chargeAmountInWei This is the amount the user is to be charged
        * @param _paymentMethod This is the payment method (RTC, ETH) being used
     */
    function generatePreimage(
        uint256 _paymentNumber,
        uint256 _chargeAmountInWei,
        uint8   _paymentMethod)
        internal
        view
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(msg.sender, _paymentNumber, _paymentMethod, _chargeAmountInWei));
    }

    /** @dev This is a helper function that prepends the ERC191 signed message prefix
        * @param _preimage This is the reconstructed message hash before being prepened with the ERC191 prefix
     */
    function generatePrefixedPreimage(bytes32 _preimage) internal pure returns (bytes32)  {
        return keccak256(abi.encodePacked(PREFIX, _preimage));
    }
}