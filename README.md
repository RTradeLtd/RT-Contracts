# RT-Contracts
Collection of RTrade's Smart Contracts

*NOTE: golang bindings, and solc compiled binarys are out of date

## RTCoin

RTCoin (RTC) is an "mmPOS" (merged-mining Proof Of Stake) ERC20 compliant utility token that gives the user access to RTrade's services. Initially starting out with 61.6Million tokens, the supply can only ever be increased, and not burned. There are two ways to generate RTC, either by staking or through merged mining with the Ethereum blockchain. For release, PoS will be supported however the merged mining contract will be released at a later date.

### RTC - Proof Of Stake

By utilizing the `Stake.sol` smart contract, users are able to stake, at a minimum, 1RTC for a period of 2103840 blocks, generating 10% (note, this may be subject to change before release) of the initial stake as newly minted RTC tokens over the lockup time (2103840 blocks). The staking system features per-block coin generation, allowing the user to mint coins every single block directly to their Ethereum address. After a period of 2103840 blocks, and after 31557600 (we reach this figure by taking an avg 15 second block time, multiplied by the lockup blocks) seconds have passed, the initial stake can be withdrawn to the users wallet.

### RTC - Merged Mining

Currently in development, a Merged Mining contract will allow anyone who mines a block on the Ethereum mainnet, to submit the block headers from the block which they mined to our Merged Mining contract, and be awarded freshly minted RTC! Currently the merged mining contract requires that each block, the block hash, and the coinbase (miner) are stored in a smart contract, allowing the miner to claim their minted tokens whenever. The ability to submit block hash and coinbase information is incentivized and can also mint RTC. The first person to submit the blockhash and coinbase information for a given block will receive a small amount of RTC, directly minted to their address. We incentivize storing this information as the user has to pay for the gas costs to invoke the transaction.

#### RTC - Merged Mining Statistics

The total supply increase based off a starting RTC supply of 61.6M is 3.15% if all rewards are claimed. This roughly equates to 242750.7375USD/year and 1942005.9RTC/year
NOTE: These figures aren't yet finalized

##### Merged Mined Block Rewards

With an average 13 second block time, if every block had their rewards claimed an average 728252.1 RTC would be minted a year with a reward of 0.3RTC

Formula to reach this is:
`(seconds per year/block time seconds) * (blocks mined reward)`

##### Merged Mined Block Information Submission

With an average 13 second block time, if all blocks have their information submitted, an average 1213753.8RTC would be minted a year with a reward of 0.5RTC per block information submitted

Formula to reach this is:
`(seconds per year/block time seconds) * (block submission reward)`

## Thanks

Thanks to @Figs999 for the EventStorage.sol contract which is serving as a basis for block header parsing. 
    > https://github.com/figs999/Ethereum/blob/master/EventStorage.sol
