pragma solidity 0.4.24;

import "../Math/SafeMath.sol";
import "../Interfaces/RTCoinInterface.sol";

/*
This contract is used to handle vesting of RTC tokens
*/

contract Vesting {
    using SafeMath for uint256;

    address constant public TOKENADDRESS = address(0);
    RTCoinInterface constant public RTI = RTCoinInterface(TOKENADDRESS);

    enum VestState {nil, vesting, vested}

    struct Vest {
        uint256 totalVest;
        uint256[] releaseDates;
        uint256[] releaseAmounts;
        VestState state;
        mapping (uint256 => bool) claimed;
    }

    mapping (address => Vest) public vests;

    modifier validIndex(uint256 _vestIndex) {
        require(_vestIndex < vests[msg.sender].releaseDates.length);
        _;
    }

    modifier pastClaimDate(uint256 _vestIndex) {
        require(now >= vests[msg.sender].releaseDates[_vestIndex]);
        _;
    }

    modifier unclaimedVest(uint256 _vestIndex) {
        require(!vests[msg.sender].claimed[_vestIndex]);
        _;
    }

    modifier activeVester() {
        require(vests[msg.sender].state == VestState.vesting);
        _;
    }

    modifier nonActiveVester(address _vester) {
        require(vests[_vester].state == VestState.nil);
        _;
    }

    constructor() public {
        // prevent deployments if not properly setup
        require(TOKENADDRESS != address(0), "token address not set");
    }
    function addVest(
        address _vester,
        uint256 _totalAmountToVest,
        uint256[] _releaseDates, // unix time stamp format `time.Now().Unix()` in golang
        uint256[] _releaseAmounts)
        public
        nonActiveVester(_vester)
        returns (bool)
    {
        require(_releaseDates.length == _releaseAmounts.length, "array lengths are not equal");
        uint256 total;
        for (uint256 i = 0; i < _releaseAmounts.length; i++) {
            total = total.add(_releaseAmounts[i]);
            require(now < _releaseDates[i], "invalid release date must be in the future");
        }
        require(total == _totalAmountToVest, "invalid total amount to vest");
        Vest memory v = Vest({
            totalVest: _totalAmountToVest,
            releaseDates: _releaseDates,
            releaseAmounts: _releaseAmounts,
            state: VestState.vesting
        });
        vests[_vester] = v;
        require(RTI.transferFrom(msg.sender, address(this), _totalAmountToVest), "transfer from failed, most likely needs approval");
        return true;
    }

    function withdrawVestedTokens(
        uint256 _vestIndex)
        public
        activeVester
        validIndex(_vestIndex)
        unclaimedVest(_vestIndex)
        pastClaimDate(_vestIndex)
        returns (bool)
    {
        // mark this particular vest as claimed
        vests[msg.sender].claimed[_vestIndex] = true;
        // if this is the last vest, make sure all others have been claimed and then mark as vested
        if (_vestIndex == vests[msg.sender].releaseAmounts.length.sub(1)) {
            bool check;
            for (uint256 i = 0; i < vests[msg.sender].releaseAmounts.length; i++) {
                // if we detect that even one vest hasn't been claimed, set check to false and break out of loop
                if (!vests[msg.sender].claimed[i]) {
                    // this will preventsituations where the first vest may not be claimed but later ones have been
                    // which would result in a "split brain" type scenario, in which the code thinks all vests have been claimed
                    // but they actually haven't
                    check = false;
                    // break out of the loop
                    break;
                }
                check = true;
            }
            // if check is true, this means that all claims have been made and we can mark the vest as having been complete
            if (check) {
                vests[msg.sender].state = VestState.vested;
            }
        }
        uint256 amount = vests[msg.sender].releaseAmounts[_vestIndex];
        require(RTI.transfer(msg.sender, amount), "failed to transfer");
        return true;
    }
}