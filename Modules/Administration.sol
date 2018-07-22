pragma solidity 0.4.24;
pragma experimental "v0.5.0";

contract Administration {

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
        emit AdminSet(_newAdmin);
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
        emit OwnershipTransferred(msg.sender, _newOwner);
        return true;
    }

}
