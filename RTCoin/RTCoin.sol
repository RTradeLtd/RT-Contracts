pragma solidity 0.4.24;
pragma experimental "v0.5.0";

import "../Modules/Administration.sol";
import "../Math/SafeMath.sol";
import "../Interfaces/ERC20Interface.sol";
import "../Interfaces/StakeInterface.sol";

/// @title RTC Token Contract
/// @author Postables, RTrade Technologies Ltd
/// @dev We able V5 for safety features, see https://solidity.readthedocs.io/en/v0.4.24/security-considerations.html#take-warnings-seriously
contract RTCoin is Administration {

    using SafeMath for uint256;

    // this is the initial supply of tokens, 61.6 Million
    uint256 constant public INITIALSUPPLY = 61600000000000000000000000;

    // this is the interface that allows interaction with the staking contract
    StakeInterface public stake = StakeInterface(0);
    // this is the address of the staking contract
    address public  stakeContractAddress = address(0);
    // This is the address of the merged mining contract, not yet developed
    address public  mergedMinerValidatorAddress = address(0);
    string  public  name = "RTCoin";
    string  public  symbol = "RTC";
    uint256 public  totalSupply = INITIALSUPPLY;
    uint8   public  decimals = 18;
    // allows transfers to be frozen, but enable them by default
    bool    public  transfersFrozen = true;

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

    // makes sure that only the stake contract, or merged miner validator contract can mint coins
    modifier onlyMinters() {
        if (mergedMinerValidatorAddress != address(0)) {
            require(msg.sender == stakeContractAddress || msg.sender == mergedMinerValidatorAddress, "sender is neither stake nor validator contract");
            _;
        } else {
            require(msg.sender == stakeContractAddress, "sender is not stake contract");
            _;
        }
    }

    constructor() public {
        balances[msg.sender] = totalSupply;
        emit Transfer(address(0), msg.sender, totalSupply);
    }

    /** @notice Used to transfer tokens
        * @param _recipient This is the recipient of the transfer
        * @param _amount This is the amount of tokens to send
     */
    function transfer(
        address _recipient,
        uint256 _amount
    )
        public
        transfersNotFrozen
        returns (bool)
    {
        // check that the sender has a valid balance
        require(balances[msg.sender] >= _amount, "sender does not have enough tokens");
        balances[msg.sender] = balances[msg.sender].sub(_amount);
        balances[_recipient] = balances[_recipient].add(_amount);
        emit Transfer(msg.sender, _recipient, _amount);
        return true;
    }

    /** @notice Used to transfer tokens on behalf of someone else
        * @param _recipient This is the recipient of the transfer
        * @param _amount This is the amount of tokens to send
     */
    function transferFrom(
        address _owner,
        address _recipient,
        uint256 _amount
    )
        public
        transfersNotFrozen
        returns (bool)
    {
        // ensure owner has a valid balance
        require(balances[_owner] >= _amount, "owner does not have enough tokens");
        // ensure that the spender has a valid allowance
        require(allowed[_owner][msg.sender] >= _amount, "sender does not have enough allowance");
        // reduce the allowance
        allowed[_owner][msg.sender] = allowed[_owner][msg.sender].sub(_amount);
        // reduce balance of owner
        balances[_owner] = balances[_owner].sub(_amount);
        // increase balance of recipient
        balances[_recipient] = balances[_recipient].add(_amount);
        emit Transfer(_owner, _recipient, _amount);
        return true;
    }

    /** @notice This is used to approve someone to send tokens on your behalf
        * @param _spender This is the person who can spend on your behalf
        * @param _amount This is the amount of tokens that they can spend
     */
    function approve(
        address _spender,
        uint256 _amount
    )
        public
        returns (bool)
    {
        require(_amount > 0, "amount must be greater than 0");
        allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_amount);
        emit Approval(msg.sender, _spender, _amount);
        return true;
    }

    // NON STANDARD FUNCTIONS //

    /** @notice This is used to set the merged miner validator contract
        * @param _mergedMinerValidator this is the address of the mergedmining contract
     */
    function setMergedMinerValidator(address _mergedMinerValidator) external onlyAdmin returns (bool) {
        mergedMinerValidatorAddress = _mergedMinerValidator;
        emit MergedMinerValidatorSet(_mergedMinerValidator);
        return true;
    }

    /** @notice This is used to set the staking contract
        * @param _contractAddress this is the address of the staking contract
    */
    function setStakeContract(address _contractAddress) external onlyAdmin returns (bool) {
        // this prevents us from changing contracts while there are active stakes going on
        if (stakeContractAddress != address(0)) {
            require(stake.activeStakes() == 0, "staking contract already configured, to change it must have 0 active stakes");
        }
        stakeContractAddress = _contractAddress;
        stake = StakeInterface(_contractAddress);
        emit StakeContractSet(_contractAddress);
        return true;
    }


    /** @notice This is used to mint new tokens
        * @dev Can only be executed by the staking, and merged miner validator contracts
        * @param _recipient This is the person who will received the mint tokens
        * @param _amount This is the amount of tokens that they will receive and which will be generated
     */
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
        // don't allow us to transfer RTC tokens
        require(_tokenAddress != address(this), "token address can't be this contract");
        ERC20Interface eI = ERC20Interface(_tokenAddress);
        require(eI.balanceOf(address(this)) >= _amount, "attempting to send more tokens than current balance");
        require(eI.transfer(_recipient, _amount), "token transfer failed");
        emit ForeignTokenTransfer(msg.sender, _recipient, _amount);
        return true;
    }
    
    /** @notice Transfers eth that is stuck in this contract
        * ETH can be sent to the address this contract resides at before the contract is deployed
        * A contract can be suicided, forcefully sending ether to this contract
     */
    function transferOutEth()
        public
        onlyAdmin
        returns (bool)
    {
        uint256 balance = address(this).balance;
        msg.sender.transfer(address(this).balance);
        emit EthTransferOut(msg.sender, balance);
        return true;
    }

    /** @notice Used to freeze token transfers
     */
    function freezeTransfers()
        public
        onlyAdmin
        returns (bool)
    {
        transfersFrozen = true;
        emit TransfersFrozen(true);
        return true;
    }

    /** @notice Used to thaw token transfers
     */
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

    /** @notice Used to get the total supply
     */
    function totalSupply()
        public
        view
        returns (uint256)
    {
        return totalSupply;
    }

    /** @notice Used to get the balance of a holder
        * @param _holder The address of the token holder
     */
    function balanceOf(
        address _holder
    )
        public
        view
        returns (uint256)
    {
        return balances[_holder];
    }

    /** @notice Used to get the allowance of someone
        * @param _owner The address of the token owner
        * @param _spender The address of thhe person allowed to spend funds on behalf of the owner
     */
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
