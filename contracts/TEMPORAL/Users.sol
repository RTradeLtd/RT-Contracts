pragma solidity 0.4.24;

import "../Modules/UsersAdministration.sol";
import "../Modules/Utils.sol";
import "../Math/SafeMath.sol";

interface ERC20I {
    function transferFrom(address _owner, address _recipient, uint256 _amount) external returns (bool);
    function transfer(address _recipient, uint256 _amount) external returns (bool);
}


/**
    Current Limitations:
        Locked balance can't be unlocked, it can't only be spent (this will change)
*/

contract Users is UsersAdministration, Utils {

    using SafeMath for uint256;

    ERC20I public rtcI = ERC20I(address(0));
    address public hotWallet;
    address public paymentProcessorAddress;
    address[] public uploaders;

    enum UserStateEnum { nil, registered, disabled }

    struct UserStruct {
        address uploaderAddress;
        uint256 availableEthBalance;
        uint256 availableRtcBalance;
        UserStateEnum state;
    }


    mapping (address => UserStruct) public users;

    event PaymentProcessorSet(address _paymentProcessor);

    event HotWalletSet(address _wallet);

    event RTCInterfaceSet(address _tokenContractAddress);

    event UserRegistered(address _uploader);

    event RtcDeposited(address indexed _uploader, uint256 _amount);
    event RtcPaymentWithdrawnForUploader(address indexed _uploader, uint256 _amount, bytes32 _hashedCID);

    event EthDeposited(address indexed _uploader, uint256 _amount);
    event EthPaymentWithdrawnForUpload(address indexed _uploader, uint256 _amount, bytes32 _hashedCID);

    modifier nonRegisteredUser(address _uploaderAddress) {
        require(users[_uploaderAddress].state == UserStateEnum.nil);
        _;
    }

    modifier onlyPaymentProcessor(address _sender) {
        require(_sender == paymentProcessorAddress);
        _;
    }

    modifier isRegistered(address _uploaderAddress) {
        require(users[_uploaderAddress].state == UserStateEnum.registered);
        _;
    }

    modifier validAvailableEthBalance(address _uploaderAddress, uint256 _amount) {
        require(users[_uploaderAddress].availableEthBalance >= _amount);
        _;
    }

    modifier validAvailableRtcBalance(address _uploaderAddress, uint256 _amount) {
        require(users[_uploaderAddress].availableRtcBalance >= _amount);
        _;
    }

    function paymentProcessorWithdrawRtcForUploader(
        address _uploaderAddress,
        uint256 _amount,
        bytes32 _hashedCID)
        public
        isRegistered(_uploaderAddress)
        greaterThanZeroU(_amount)
        onlyPaymentProcessor(msg.sender)
        validAvailableRtcBalance(_uploaderAddress, _amount)
        returns (bool)
    {
        uint256 remaining = users[_uploaderAddress].availableRtcBalance.sub(_amount);
        users[_uploaderAddress].availableRtcBalance = remaining;
        emit RtcPaymentWithdrawnForUploader(_uploaderAddress, _amount, _hashedCID);
        require(rtcI.transfer(hotWallet, _amount));
        return true;
    }

    function paymentProcessorWithdrawEthForUploader(
        address _uploaderAddress,
        uint256 _amount,
        bytes32 _hashedCID)
        public
        isRegistered(_uploaderAddress)
        greaterThanZeroU(_amount)
        onlyPaymentProcessor(msg.sender)
        validAvailableEthBalance(_uploaderAddress, _amount)
        returns (bool)
    {
        uint256 remaining = users[_uploaderAddress].availableEthBalance.sub(_amount);
        users[_uploaderAddress].availableEthBalance = remaining;
        emit EthPaymentWithdrawnForUpload(_uploaderAddress, _amount, _hashedCID);
        hotWallet.transfer(_amount);
        return true;
    }
    
    function registerUser()
        public
        nonRegisteredUser(msg.sender)
        returns (bool)
    {
        users[msg.sender].state = UserStateEnum.registered;
        users[msg.sender].uploaderAddress = msg.sender;
        uploaders.push(msg.sender);
        emit UserRegistered(msg.sender);
        return true;
    }


    function depositEther()
        public
        payable
        isRegistered(msg.sender)
        greaterThanZeroU(msg.value)
        returns (bool)
    {
        users[msg.sender].availableEthBalance = users[msg.sender].availableEthBalance.add(msg.value);
        emit EthDeposited(msg.sender, msg.value);
        return true;
    }


    function depositRtc(
        uint256 _amount)
        public
        isRegistered(msg.sender)
        greaterThanZeroU(_amount)
        returns (bool)
    {
        users[msg.sender].availableRtcBalance = users[msg.sender].availableRtcBalance.add(_amount);
        emit RtcDeposited(msg.sender, _amount);
        require(rtcI.transferFrom(msg.sender, address(this), _amount));
        return true;
    }

    function setRTCInterface(
        address _tokenContractAddress)
        public
        onlyAdmin(msg.sender)
        nonZeroAddress(_tokenContractAddress)
        returns (bool)
    {
        rtcI = ERC20I(_tokenContractAddress);
        emit RTCInterfaceSet(_tokenContractAddress);
        return true;
    }

    function setHotWallet(
        address _hotWalletAddress)
        public
        onlyAdmin(msg.sender)
        nonZeroAddress(_hotWalletAddress)
        returns (bool)
    {
        hotWallet = _hotWalletAddress;
        emit HotWalletSet(_hotWalletAddress);
        return true;
    }

    function setPaymentProcessorAddress(
        address _paymentProcessorAddress)
        public
        onlyAdmin(msg.sender)
        nonZeroAddress(_paymentProcessorAddress)
        returns (bool)
    {
        paymentProcessorAddress = _paymentProcessorAddress;
        emit PaymentProcessorSet(_paymentProcessorAddress);
        return true;
    }
}
