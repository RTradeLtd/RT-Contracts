pragma solidity 0.4.24;

import "./Math/SafeMath.sol";
import "./Modules/Administration.sol";
import "./Interfaces/OracleSubscriberInterface.sol";

/*
    Used to facilitate oracle style updates of our smart contracts without having to rely on third-party products
    Currently only one update is supported, ETH-USD prices
*/

contract Oracle is Administration {
    using SafeMath for uint256;

    struct AuthorizedContractStruct {
        address contractAddress;
        uint256 updateFrequencyInHours;
        uint256 nextUpdate;
        bool    enabled;
        bytes4[] enabledFunctions;
        mapping (bytes4 => bool) validFunctions;
    }

    mapping (address => AuthorizedContractStruct) public contracts;

    event AuthorizedContractAdded(address _contractAddress);

    modifier authorizedContract(address _contractAddress) {
        require(contracts[_contractAddress].enabled);
        _;
    }

    modifier authorizedFunctionCall(address _contractAddress, bytes4 _functionCall) {
        require(contracts[_contractAddress].validFunctions[_functionCall]);
        _;
    }

    function addAuthorizedContract(
        address _contractAddress,
        uint256 _updateFrequencyInHours,
        bytes4[] _enabledFunctions)
        public
        onlyAdmin
        returns (bool)
    {
        AuthorizedContractStruct memory a;
        a.contractAddress = _contractAddress;
        a.updateFrequencyInHours = _updateFrequencyInHours;
        a.nextUpdate = now.add(_updateFrequencyInHours.mul(1 hours));
        a.enabledFunctions = _enabledFunctions;
        contracts[_contractAddress] = a;
        for (uint256 i = 0; i < _enabledFunctions.length; i++) {
            contracts[_contractAddress].validFunctions[_enabledFunctions[i]] = true;
        }
        emit AuthorizedContractAdded(_contractAddress);
        return true;
    }



    function updateRtcPrice(
        address _destinationContract,
        uint256 _rtcUSD)
        public
        onlyAdmin
        authorizedContract(_destinationContract)
        authorizedFunctionCall(_destinationContract, bytes4(keccak256("updateRtcPrice(uint256)")))
        returns (bool)
    {
        require(OracleSubscriberInterface(_destinationContract).updateRtcPrice(_rtcUSD));
        return true;
    }

    function updateEthPrice(
        address _destinationContract,
        uint256 _ethUSD)
        public
        onlyAdmin
        authorizedContract(_destinationContract)
        authorizedFunctionCall(_destinationContract, bytes4(keccak256("updateEthPrice(uint256)")))
        returns (bool)
    {
        require(OracleSubscriberInterface(_destinationContract).updateEthPrice(_ethUSD));
        return true;
    }
}