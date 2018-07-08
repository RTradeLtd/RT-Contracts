# RT-Contracts
Collection of RTrade's Smart Contracts


# RTCoin

RTC can be considered a "Merged Mining Proof-Of-Stake" coin, in that the total supply can only be increased by staking RTC in the Stake contract. Coin generation is controlled via block number increase, which is where the "merged mining" aspect comes from. We believe this to be a unique, and exciting way to stake RTC.  The staking system is completely hands off, except for a total of three steps:
* 1) Depositing stake
* 2) Withdrawing stake reward
* 3) Withdrawing inital stake after stake period is over

Your stake reward can be withdrawn whenever you please, so long as at least one block has passed since the last time you withdrew your reward.  Initiating a withdrawal process, kicks off the supply increase on the RTC token contract.