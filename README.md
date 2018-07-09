# RT-Contracts
Collection of RTrade's Smart Contracts


## RTCoin

RTC can be considered a "Merged Mining Proof-Of-Stake" coin, in that the total supply can only be increased by staking RTC in the Stake contract.  The staking system is completely hands off, except for a total of three steps:
* 1) Depositing stake
* 2) Withdrawing stake reward
* 3) Withdrawing inital stake after stake period is over

Your stake reward can be withdrawn whenever you please, so long as at least one block has passed since the last time you withdrew your reward.  Initiating a withdrawal process, kicks off the supply increase on the RTC token contract.

## RTCoin - Merged Mining

It is possible for all miners of ethereum, to also be rewarded RTC. In order to do so, the miner of a block must submit sufficient validation information to the Merged Miner Contract. This validation information must be the entire block header, as well as necessary components to parse through, and reconstruct the block header to validate that the person submitting the transaction, is the person who mined the block. There will be a semi-significant mininum withdrawal amount from the merged mining contract.

Development roadmap for Merged Mining:
    1) require simple rlp encoded block header submission (75% complete)
    2) Allow simple, but incentivized, block hash + number submission (not yet started)
    3) require block header validation  (not yet started)


## Thanks

Thanks to @Figs999 for the EventStorage.sol contract which is serving as a basis for block header parsing. 
    > https://github.com/figs999/Ethereum/blob/master/EventStorage.sol