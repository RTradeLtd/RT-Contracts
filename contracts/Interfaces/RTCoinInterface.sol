pragma solidity 0.4.23;

/**
RTCoin Interface
*/

interface RTCoinInterface {
    
    function freezeTransfers() external returns (bool);

    function thawTransfers() external returns (bool);

    function transfer(address _recipient, uint256 _amount) external returns (bool transferred);

    function transferFrom(address _owner, address _recipient, uint256 _amount) external returns (bool transferredFrom);

    function approve(address _spender, uint256 _amount) external returns (bool approved);
    /**GETTERS */

    function totalSupply() external view returns (uint256);

    function balanceOf(address _holder) external view returns (uint256);

    function allowance(address _owner, address _spender) external view returns (uint256);
}
