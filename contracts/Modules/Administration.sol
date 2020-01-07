pragma solidity 0.4.24;
pragma experimental "v0.5.0";

import "../Math/SafeMath.sol";

contract Administration {

    using SafeMath for uint256;

    address public owner;
    address public admin;

    event AdminSet(address _admin);
    event OwnershipTransferred(address _previousOwner, address _newOwner);


    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier onlyAdmin() {
        require(msg.sender == owner || msg.sender == admin);
        _;
    }

    modifier nonZeroAddress(address _addr) {
        require(_addr != address(0), "must be non zero address");
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
        nonZeroAddress(_newAdmin)
        returns (bool)
    {
        require(_newAdmin != admin);
        admin = _newAdmin;
        emit AdminSet(_newAdmin);
        return true;
    }

    function transferOwnership(
        address _newOwner
    )
        public
        onlyOwner
        nonZeroAddress(_newOwner)
        returns (bool)
    {
        owner = _newOwner;
        emit OwnershipTransferred(msg.sender, _newOwner);
        return true;
    }

}
