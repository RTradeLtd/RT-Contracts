pragma solidity 0.4.24;

interface ERC20Interface {
    function owner() external view returns (address);
    function decimals() external view returns (uint8);
    function transfer(address _to, uint256 _value) external returns (bool);
    function transferFrom(address _from, address _to, uint256 _value) external returns (bool);
    function approve(address _spender, uint256 _amount) external returns (bool);
    function totalSupply() external view returns (uint256);
    function balanceOf(address _owner) external view returns (uint256);
    function allowance(address _owner, address _spender) external view returns (uint256);
}

library SafeMath {

  // We use `pure` bbecause it promises that the value for the function depends ONLY
  // on the function arguments
    function mul(uint256 a, uint256 b) internal pure  returns (uint256) {
        uint256 c = a * b;
        require(a == 0 || c / a == b);
        return c;
    }

    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a / b;
        return c;
    }

    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a);
        return a - b;
    }

    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a);
        return c;
    }
}

contract Administration {

    address public owner;
    address public admin;
    bool    public frozen;

    mapping (address => bool) public moderators;

    event AdminSet(address indexed _admin, bool indexed _adminSet);
    event OwnershipTransferred(address indexed _previousOwner, address indexed _newOwner, bool indexed _ownershipTransferred);

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier onlyAdmin() {
        require(msg.sender == owner || msg.sender == admin);
        _;
    }

    modifier onlyPrivileged() {
        require(msg.sender == owner || msg.sender == admin || moderators[msg.sender] == true);
        _;
    }

    constructor() public {
        owner = msg.sender;
        admin = msg.sender;
    }

    function setAdmin(
        address _newAdmin
    )
        public
        onlyOwner
        returns (bool)
    {
        require(_newAdmin != admin);
        admin = _newAdmin;
        emit AdminSet(_newAdmin, true);
        return true;
    }

    function transferOwnership(
        address _newOwner
    )
        public
        onlyOwner
        returns (bool)
    {
        require(_newOwner != owner);
        owner = _newOwner;
        emit OwnershipTransferred(msg.sender, _newOwner, true);
        return true;
    }

    function owner() external view returns (address) {
        return owner;
    }

    function admin() external view returns (address) {
        return admin;
    }
}

interface StakeInterface {
    function activeStakes() external view returns (uint256);
}

contract RTCoin is Administration {

    using SafeMath for uint256;

    StakeInterface public stake;

    address public  stakeContract;
    string  public  name;
    string  public  symbol;
    uint256 public  totalSupply;
    uint256 public  initialSupply = 61600000000000000000000000;
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

    modifier transfersNotFrozen() {
        require(!transfersFrozen);
        _;
    }

    modifier transfersAreFrozen() {
        require(transfersFrozen);
        _;
    }

    modifier onlyStakeContract() {
        require(msg.sender == stakeContract);
        _;
    }

    constructor() public {
        name = "RTCoin";
        symbol = "RTC";
        decimals = 18;
        // 88il in wei
        totalSupply = initialSupply;
        balances[msg.sender] = totalSupply;
        emit Transfer(address(0), msg.sender, totalSupply);
    }

    function setStakeContract(address _contractAddress)
        external
        onlyAdmin
        returns (bool)
    {
        // this prevents us from changing contracts while there are active stakes going on
        if (stakeContract != address(0)) {
            require(stake.activeStakes() == 0);
        }
        stakeContract = _contractAddress;
        stake = StakeInterface(_contractAddress);
        // event place holder
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


    // This needs to be locked down so only the staking contract can invoke this function
    function mint(
        address _recipient,
        uint256 _amount)
        public
        onlyStakeContract
        returns (bool)
    {
        balances[_recipient] = balances[_recipient].add(_amount);
        totalSupply = totalSupply.add(_amount);
        emit Transfer(address(0), _recipient, _amount);
        // total supply increase place holder
        return true;
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

    function transfer(
        address _recipient,
        uint256 _amount
    )
        public
        transfersNotFrozen
        returns (bool transferred)
    {
        require(transferCheck(msg.sender, _recipient, _amount));
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
        require(transferCheck(_owner, _recipient, _amount));
        require(allowed[_owner][msg.sender] >= _amount);
        require(allowed[_owner][msg.sender].sub(_amount) >= 0);
        allowed[_owner][msg.sender] = allowed[_owner][msg.sender].sub(_amount);
        balances[_owner] = balances[_owner].sub(_amount);
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

    /**INTERNALS */

    function transferCheck(
        address _sender,
        address _recipient,
        uint256 _amount
    )
        internal
        view
        returns (bool valid)
    {
        require(_sender != address(0x0) && _recipient != address(0x0) && _amount > 0);
        require(balances[_sender].sub(_amount) >= 0);
        require(balances[_recipient].add(_amount) > 0);
        require(balances[_recipient].add(_amount) > balances[_recipient]);
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

