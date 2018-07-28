pragma solidity 0.4.24;
pragma experimental "v0.5.0";

import "../Math/SafeMath.sol";

contract Administration {

    using SafeMath for uint256;

    address public owner;
    address public admin;
    OwnerTimeDelayStruct public delay;

    event AdminSet(address _admin);
    event OwnershipTransferred(address _previousOwner, address _newOwner);
    event OwnerTransferDelayStarted(
        address _previousOwner,
        address _newOwner,
        uint256 _activationBlock,
        uint256 _activationTime
    );

    enum DelayEnum { nil, pending, changed }

    struct OwnerTimeDelayStruct {
        address previousOwner;
        address newOwner;
        uint256 delayExpirationBlock;
        uint256 delayExpirationTime;
        DelayEnum state;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier onlyAdmin() {
        require(msg.sender == owner || msg.sender == admin);
        _;
    }

    modifier validOwnerChange(address _newOwner) {
        require(delay.state == DelayEnum.pending);
        require(delay.newOwner == _newOwner, "incorrect new owner");
        require(now >= delay.delayExpirationTime);
        require(block.number >= delay.delayExpirationBlock);
        _;
    }

    modifier noPendingDelay() {
        require(delay.state == DelayEnum.nil || delay.state == DelayEnum.changed);
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

    /** @notice Starts an ownership transfer process
        * @dev Can be executed by admin or owner, but the final step requires owner execution
        * @param _newOwner This is the address of the new owner
     */
    function startOwnerTransferDelay(address _newOwner) public onlyAdmin noPendingDelay returns (bool) {
        OwnerTimeDelayStruct memory os = OwnerTimeDelayStruct({
            previousOwner: msg.sender,
            newOwner: _newOwner,
            delayExpirationBlock: block.number.add(10),
            delayExpirationTime: now.add(uint256(100).mul(1 seconds)),
            state: DelayEnum.pending
        });
        delay = os;
        emit OwnerTransferDelayStarted(os.previousOwner, os.newOwner, os.delayExpirationBlock, os.delayExpirationTime);
        return true;
    }

    function transferOwnership(
        address _newOwner
    )
        public
        onlyOwner
        validOwnerChange(_newOwner)
        returns (bool)
    {
        delay.state = DelayEnum.changed;
        owner = _newOwner;
        emit OwnershipTransferred(msg.sender, _newOwner);
        return true;
    }

}
