pragma solidity 0.4.24;
pragma experimental "v0.5.0";

interface RTCoinInterface {
    

    /** Functions - ERC20 */
    function transfer(address _recipient, uint256 _amount) external returns (bool);

    function transferFrom(address _owner, address _recipient, uint256 _amount) external returns (bool);

    function approve(address _spender, uint256 _amount) external returns (bool approved);

    /** Getters - ERC20 */
    function totalSupply() external view returns (uint256);

    function balanceOf(address _holder) external view returns (uint256);

    function allowance(address _owner, address _spender) external view returns (uint256);

    /** Getters - Custom */
    function mint(address _recipient, uint256 _amount) external returns (bool);

    function stakeContractAddress() external view returns (address);

    function mergedMinerValidatorAddress() external view returns (address);
    
    /** Functions - Custom */
    function freezeTransfers() external returns (bool);

    function thawTransfers() external returns (bool);
}
