pragma solidity 0.4.24;

import "../Math/SafeMath.sol";
import "../Interfaces/ERC20Interface.sol";

contract TemporalBalances {

    using SafeMath for uint256;

    address constant public TOKENADDRESS = address(0);
    ERC20Interface constant public ercI = ERC20Interface(TOKENADDRESS);

    enum BalanceState { nil, active, locked, closed }
    
    struct Balance {
        uint256 ethBalance;
        uint256 rtcBalance;
        uint256 lastActivity;
        uint256 renewalDate;
        BalanceState state;
    }

    mapping (address => Balance) public users;

    modifier activeUser(address _user) {
        require(users[_user].state == BalanceState.active, "user must be active");
        _;
    }

    modifier nonRegisteredUser(address _user) {
        require(users[_user].state == BalanceState.nil, "user must not have an active account");
        _;
    }

    modifier beforeRenewal(address _user) {
        require(now < users[_user].renewalDate, "time must be before renewal date");
        _;
    }

    constructor() {
        require(TOKENADDRESS != address(0), "token address cant be empty");
    }

    function registerUser() public nonRegisteredUser(msg.sender) returns (bool) {
        users[msg.sender] = Balance({
            ethBalance: 0,
            rtcBalance: 0,
            lastActivity: now,
            renewalDate: now.add(30 days),
            state: BalanceState.active
        });
        return true;
    }

    function depositRTC(uint256 _numRTC) public activeUser(msg.sender) beforeRenewal(msg.sender) returns (bool) {
        users[msg.sender].rtcBalance = users[msg.sender].rtcBalance.add(_numRTC);
        users[msg.sender].lastActivity = now;
        require(ercI.transferFrom(msg.sender, address(this), _numRTC), "transferFrom failed likely needs approval");
        return true;
    }

    function depositETH() public payable activeUser(msg.sender) beforeRenewal(msg.sender) returns (bool) {
        users[msg.sender].ethBalance = users[msg.sender].ethBalance.add(msg.value);
        users[msg.sender].lastActivity = now;
        return true;
    }
}