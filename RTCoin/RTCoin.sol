pragma solidity 0.4.24;

import "../Modules/Administration.sol";
import "../Math/SafeMath.sol";
import "../Interfaces/ERC20Interface.sol";
import "../Interfaces/StakeInterface.sol";

contract RTCoin is Administration {

    using SafeMath for uint256;

    uint256 constant public INITIALSUPPLY = 61000000000000000000000000;

    StakeInterface public stake;
    address public  stakeContractAddress;
    address public  mergedMinerValidatorAddress;
    string  public  name;
    string  public  symbol;
    uint256 public  totalSupply;
    uint8   public  decimals;
    bool    public  transfersFrozen;

    mapping (address => uint256) public balances;
    mapping (address => mapping (address => uint256)) public allowed;

    event Transfer(address indexed _sender, address indexed _recipient, uint256 _amount);
    event Approval(address indexed _owner, address indexed _spender, uint256 _amount);
    event TransfersFrozen(bool indexed _transfersFrozen);
    event TransfersThawed(bool indexed _transfersThawed);
    event ForeignTokenTransfer(address indexed _sender, address indexed _recipient, uint256 _amount);
    event EthTransferOut(address indexed _recipient, uint256 _amount);
    event MergedMinerValidatorSet(address _contractAddress);
    event StakeContractSet(address _contractAddress);
    event CoinsMinted(address indexed _stakeContract, address indexed _recipient, uint256 _mintAmount);

    modifier transfersNotFrozen() {
        require(!transfersFrozen);
        _;
    }

    modifier transfersAreFrozen() {
        require(transfersFrozen);
        _;
    }

    modifier onlyMinters() {
        require(msg.sender == stakeContractAddress || msg.sender == mergedMinerValidatorAddress);
        _;
    }

    constructor() public {
        name = "RTCoin";
        symbol = "RTC";
        decimals = 18;
        // 88il in wei
        totalSupply = INITIALSUPPLY;
        balances[msg.sender] = totalSupply;
        emit Transfer(address(0), msg.sender, totalSupply);
    }

    function transfer(
        address _recipient,
        uint256 _amount
    )
        public
        transfersNotFrozen
        returns (bool transferred)
    {
        // check that the sender has a valid balance
        require(balances[msg.sender] >= _amount);
        balances[msg.sender] = balances[msg.sender].sub(_amount);
        balances[_recipient] = balances[_recipient].add(_amount);
        emit Transfer(msg.sender, _recipient, _amount);
        return true;
    }

    function transferFrom(
        address _owner,
        address _recipient,
        uint256 _amount
    )
        public
        transfersNotFrozen
        returns (bool transferredFrom)
    {
        // ensure owner has a valid balance
        require(balances[_owner] >= _amount);
        // ensure that the spender has a valid allowance
        require(allowed[_owner][msg.sender] >= _amount);
        require(allowed[_owner][msg.sender].sub(_amount) >= 0);
        // reduce the allowance
        allowed[_owner][msg.sender] = allowed[_owner][msg.sender].sub(_amount);
        // reduce balance of owner
        balances[_owner] = balances[_owner].sub(_amount);
        // increase balance of recipient
        balances[_recipient] = balances[_recipient].add(_amount);
        emit Transfer(_owner, _recipient, _amount);
        return true;
    }

    function approve(
        address _spender,
        uint256 _amount
    )
        public
        returns (bool approved)
    {
        require(_spender != address(0x0));
        require(_amount > 0);
        require(allowed[msg.sender][_spender].add(_amount) > allowed[msg.sender][_spender]);
        allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_amount);
        emit Approval(msg.sender, _spender, _amount);
        return true;
    }

    // NON STANDARD FUNCTIONS //

    function setMergedMinerValidator(address _mergedMinerValidator) external onlyAdmin returns (bool) {
        mergedMinerValidatorAddress = _mergedMinerValidator;
        emit MergedMinerValidatorSet(_mergedMinerValidator);
        return true;
    }


    function setStakeContract(address _contractAddress) external onlyAdmin returns (bool) {
        // this prevents us from changing contracts while there are active stakes going on
        if (stakeContractAddress != address(0)) {
            require(stake.activeStakes() == 0);
        }
        stakeContractAddress = _contractAddress;
        stake = StakeInterface(_contractAddress);
        emit StakeContractSet(_contractAddress);
        return true;
    }


    // This needs to be locked down so only the staking contract can invoke this function
    function mint(
        address _recipient,
        uint256 _amount)
        public
        onlyMinters
        returns (bool)
    {
        balances[_recipient] = balances[_recipient].add(_amount);
        totalSupply = totalSupply.add(_amount);
        emit Transfer(address(0), _recipient, _amount);
        emit CoinsMinted(msg.sender, _recipient, _amount);
        return true;
    }

    function transferForeignToken(
        address _tokenAddress,
        address _recipient,
        uint256 _amount)
        public
        onlyAdmin
        returns (bool)
    {
        // prevent sending of RTC token
        require(_tokenAddress != address(this));
        ERC20Interface eI = ERC20Interface(_tokenAddress);
        require(eI.balanceOf(address(this)) >= _amount);
        require(eI.transfer(_recipient, _amount));
        emit ForeignTokenTransfer(msg.sender, _recipient, _amount);
    }
    
    //This will only ever have to be called in cases where a contract suicides and ether is forcably sent
    // or if someone sends ehter to the address the contract will reside at before it is deployed
    function transferOutEth()
        public
        onlyAdmin
        returns (bool)
    {
        uint256 balance = address(this).balance;
        msg.sender.transfer(address(this).balance);
        emit EthTransferOut(msg.sender, balance);
    }

    function freezeTransfers()
        public
        onlyAdmin
        returns (bool)
    {
        transfersFrozen = true;
        emit TransfersFrozen(true);
        return true;
    }

    function thawTransfers()
        public
        onlyAdmin
        returns (bool)
    {
        transfersFrozen = false;
        emit TransfersThawed(true);
        return true;
    }


    /**GETTERS */

    function totalSupply()
        public
        view
        returns (uint256)
    {
        return totalSupply;
    }

    function balanceOf(
        address _holder
    )
        public
        view
        returns (uint256)
    {
        return balances[_holder];
    }

    function allowance(
        address _owner,
        address _spender
    )
        public
        view
        returns (uint256)
    {
        return allowed[_owner][_spender];
    }

}
