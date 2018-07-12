pragma solidity 0.4.24;

import "../Math/SafeMath.sol";
import "../Modules/Administration.sol";
/*
This contract is used to facilitate payments between frequent TEMPORAL users (ie, API users). It allows us to continue the same payment model as irregular users,
however we don't have to commit a transaction to the blockchain for each payment.  Whenever the user wishes to upload something through our system, they will 
generate valid signature data. This data is given to us, and then validated. If validation is successful, then the content is injected into our system. RTRade
can then utilize these signatures to redeem our RTC/ETH whenever we wish. 

By doing this, we allow users smart contract validated, per-upload payments in a gas efficient manner.
*/