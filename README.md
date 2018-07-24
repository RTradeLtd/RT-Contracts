# RT-Contracts
Collection of RTrade's Smart Contracts

*NOTE: golang bindings, and solc compiled binarys are out of date

## RTCoin

RTCoin (RTC) is an "mmPOS" (merged-mining Proof Of Stake) ERC20 compliant utility token that gives the user access to RTrade's services. Initially starting out with 61.6Million tokens, the supply can only ever be increased, and not burned. There are two ways to generate RTC, either by staking or through merged mining with the Ethereum blockchain. For release, PoS will be supported however the merged mining contract will be released at a later date.

By default, token transfers are frozen so they need to manually be enabled after deployment

### Proof Of Stake

By utilizing the `Stake.sol` smart contract, users are able to stake, at a minimum, 1RTC for a period of 2103840 blocks, generating 10% (note, this may be subject to change before release) of the initial stake as newly minted RTC tokens over the lockup time (2103840 blocks). The staking system features per-block coin generation, allowing the user to mint coins every single block directly to their Ethereum address. After a period of 2103840 blocks, and after 31557600 (we reach this figure by taking an avg 15 second block time, multiplied by the lockup blocks) seconds have passed, the initial stake can be withdrawn to the users wallet.

#### Objectives - Proof Of Stake

* [x] - Allow per block coin minting
* [x] - Allow per block coin mint withdrawal
* [x] - Allow staking for a period of 2103840 blocks
* [x] - Allow coin mints to continue even if initial stake has been withdrawn
* [x] - Do not allow initial stake withdrawal until 2103840 blocks, and after 31557600 seconds have passed since initial stake was deposited

#### Proof Of Stake Setup

Deployment:
1) Ensure that the RTC token is deployed, along with transfers enable
2) Deploy stake contract
3) Set stake contract address on the RTC token
4) Set RTC token interface and address on the stake contract
5) Allow new stakes on the stake contract

Interaction:
1) Approve the stake contract to spend funds on your behalf
2) Deposit the stake
3) Wait one block and you can start minting tokens
4) After 2103840 blocks and (2103840 * 15 seconds) have passed you can withdraw your initial stake

#### Gas Usage - Proof Of Stake

| function | gas |
|----------|------|
| deposit  stake | 278321 |
| withdraw initial stake | 28796 |
| mint | 71574|

#### Proof Of Stake Tests

The proof of stake tests are configured to use a block hold period of 5 blocks to allow for easy end-to-end tests

### Merged Mining

Currently in development, a Merged Mining contract will allow anyone who mines a block on the Ethereum mainnet, to submit the block headers from the block which they mined to our Merged Mining contract, and be awarded freshly minted RTC! Currently the merged mining contract requires that each block, the block hash, and the coinbase (miner) are stored in a smart contract, allowing the miner to claim their minted tokens whenever. The ability to submit block hash and coinbase information is incentivized and can also mint RTC. The first person to submit the blockhash and coinbase information for a given block will receive a small amount of RTC, directly minted to their address. We incentivize storing this information as the user has to pay for the gas costs to invoke the transaction.

#### Objectives - Merged Mining

* [x] Incentivize Block Information (block number, coinbase) Submission 
* [x] Reward ETH block miners who have had their block information for the mined block submitted

#### Merged Mining Setup

Deployment
1) Set RTCoin Interface on merged mining contract
2) Set Merged Mining contract address on RTCoin contract

##### Gas Usage - Merged Mining

| function | gas |
| ---------|------|
| submit block | 92428 |
| bulk reward claim (10 rewards) | 119239 |

#### Statistics

The total supply increase based off a starting RTC supply of 61.6M is 3.15% if all rewards are claimed. This roughly equates to 242750.7375USD/year and 1942005.9RTC/year

##### Block Rewards

With an average 13 second block time, if every block had their rewards claimed an average 728252.1 RTC would be minted a year with a reward of 0.3RTC

Formula to reach this is:
`(seconds per year/block time seconds) * (blocks mined reward)`

##### Block Information Submission

With an average 13 second block time, if all blocks have their information submitted, an average 1213753.8RTC would be minted a year with a reward of 0.5RTC per block information submitted

Formula to reach this is:
`(seconds per year/block time seconds) * (block submission reward)`



##### Limitations

Block number, and coinbase will only be stored when the transaction is mined. If you want your transaction to include the same information from the block right after you submit your transaction you will need to increase your gas price appropriately in order to ensure the transaction is mined in time. Otherwise, the information for the block at which at which your TX is included will be what is stored in the contract.

### Thanks

Thanks to Figs999 for the EventStorage.sol contract which is serving as a basis for block header parsing.
    > [EventStorage.sol](https://github.com/figs999/Ethereum/blob/master/EventStorage.sol)
