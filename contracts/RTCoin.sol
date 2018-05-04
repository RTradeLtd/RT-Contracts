pragma solidity 0.4.23;

/**
    [AUTHOR] - Postables
    Property of RTrade Technology Ltd only. If you wish to use this code you must seek approval from the company first.
    This is the prototype version of RTCoin. It will not be the final version used by the RTCoin network however
    it will be directly exchangeable on a 1:1 basis via smart contract automatically at the time RTCoin is released.
*/

import "./Modules/Administration.sol";
import "./Math/SafeMath.sol";

contract RTCoin is Administration {

    using SafeMath for uint256;

    string public name;
    string  public  symbol;
    uint256 public  totalSupply;
    uint8   public  decimals;
    bool    public  transfersFrozen;

    enum AclState { nil, pending, active, disabled }
    struct AclStruct {
        address holder;
        address[] validRecipientsArray;
        address[] validSourcesArray;
        mapping (address => bool) validRecipient;
        mapping (address => bool) validSources;
        AclState state;
    }

    mapping (address => uint256) public balances;
    mapping (address => mapping (address => uint256)) public allowed;

    event Transfer(address indexed _sender, address indexed _recipient, uint256 _amount);
    event Approval(address indexed _owner, address indexed _spender, uint256 _amount);
    event TransfersFrozen(bool indexed _transfersFrozen);
    event TransfersThawed(bool indexed _transfersThawed);

    modifier transfersNotFrozen() {
        require(!transfersFrozen);
        _;
    }

    modifier transfersAreFrozen() {
        require(transfersFrozen);
        _;
    }

    constructor() public {
        name = "RTCoin";
        symbol = "RTC";
        decimals = 18;
        // 88il in wei
        totalSupply = 88888888000000000000000000;
        balances[msg.sender] = totalSupply;
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
