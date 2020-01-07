# RT-Contracts

In this repository you'll find all of the smart contracts that we use at RTrade for our various services. Along with them, you'll find a detailed summary of our smart contract infrastructure for RTC in this readme.

For code audit reports check the [Documents Directory](https://github.com/RTradeLtd/RT-Contracts/tree/master/Documents/audits)

For our official launch/deployment announcment see this [blog post](https://steemit.com/cryptocurrency/@rtrade/rtrade-technologies-ltd-officially-announces-the-mainnet-launch-of-rtcoin)

## Deployed Contracts

### Mainnet

* [MultiSignatureWallet](https://etherscan.io/address/0x41fb0e5bd1dfe3b61e9a09ebd4105c2e35b0bcbd)
* [RTCoin](https://etherscan.io/token/0xecc043b92834c1ebde65f2181b59597a6588d616)
* [Stake](https://etherscan.io/address/0xD6e33C11CFF866162787b7198030aaC101A61F29#code)
* [RTCETH](https://etherscan.io/address/0x40e68e3F58b9C1928954BEe5dEcC09A45aA531f8#code)
* [Vesting](https://etherscan.io/address/0x211D8B3EB985626B3363D3AeDd9B071113660330#code)
* [Merged Miner](https://etherscan.io/address/0x19Cdf52Ce778ef01BAc6A87615dC292A0BC14d5C#code)
* [Payments](https://etherscan.io/address/0x00c8162c86a977edc10c641636bdff1c59327983#code)

### Rinkeby

* [RTCoin](https://rinkeby.etherscan.io/address/0x675b45856257ceef650100c7ca1b2e8c6ff42e7c#code)
* [Stake](https://rinkeby.etherscan.io/address/0x0aa0277cecb6ec0872d55b7bdc9f8a9dc160022a#code)
* [RTCETH](https://rinkeby.etherscan.io/address/0xC2a87A17b28316469152e0c2892D20095369bd97#code)
* [Vesting](https://rinkeby.etherscan.io/address/0x8B8784c083b90AA0D699014d750748B3B2c1aD64#code)
* [Merged Miner](https://rinkeby.etherscan.io/address/0x64d577fd98B000871ae99Fb1d9C1C32c3018dC10#code)
* [Payments](https://rinkeby.etherscan.io/address/0x0b1906fd5509edbac3db619553fbacb64ebad132#code)

## RTCoin

---
RTCoin (RTC) is an "mmPOS" (merged-mining Proof Of Stake) ERC20 compliant utility token that gives the user access to RTrade's services. The supply can only ever be increased, and not burned. There are two ways to generate RTC, either by staking or through merged mining with the Ethereum blockchain.

By default, token transfers are frozen for RTC, so they need to manually be enabled after deployment.

As of this commit, all tests use the previous RTC contract which didn't have the additional stake failover ability. This will be changed shortly.

### Supply Distribution And Coin Information [more info](https://www.rtradetechnologies.com/en/coin/)

* Initial/Total Supply: 61.6Million
* Max Supply: uncapped
* POS Coin Generation: 10%
* Merged Mining Coin Generation: Theoretical max of 3.15% (see later section for explanation of why it is a theoretical max)

| Allocation Categories | Allocation Percentage |
|-----------------------|-----------------------|
| Advisors/Partnership Incentives | 5%          |
| Founders/Early Contributors | 10%             |
| Company | 15%                                 |
| Community Purchase | 70%                      |

### Exchanges

RTrade will not engage in public discussion about potential exchange partnerships. Repeated asks in our community channels will result in a permanent ban as it will be considered spamming.

## Proof Of Stake

---
By utilizing the `Stake.sol` smart contract, users are able to stake, at a minimum, 1RTC for a period of 2103840 blocks, generating 10% (note, this may be subject to change before release) of the initial stake as newly minted RTC tokens over the lockup time (2103840 blocks).

The staking system features per-block coin generation, allowing the user to mint coins every single block directly to their Ethereum address. After a period of 2103840 blocks, and after 31557600 (we reach this figure by taking an avg 15 second block time, multiplied by the lockup blocks) seconds have passed, the initial stake can be withdrawn to the users wallet.

### Objectives - Proof Of Stake

* [x] - Allow per block coin minting
* [x] - Allow per block coin mint withdrawal
* [x] - Allow staking for a period of 2103840 blocks
* [x] - Allow coin mints to continue even if initial stake has been withdrawn
* [x] - Do not allow initial stake withdrawal until 2103840 blocks, and after 31557600 seconds have passed since initial stake was deposited

### Setup - Proof Of Stake

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

### Gas Usage - Proof Of Stake

| function | gas |
|----------|------|
| deposit  stake | 278321 |
| withdraw initial stake | 28796 |
| mint | 71574|

### Statistics - Proof Of Stake

Our PoS system allows for for a total supply increase of 10%

### Tests - Proof Of Stake

The proof of stake tests are configured to use a block hold period of 5 blocks to allow for easy end-to-end tests.
To run these tests, simply update the `RTCoin/stake/main.go` file to include your eth account, and IPC path for a blockchain or testnet.

## Merged Mining

---
Currently in development, a Merged Mining contract will allow anyone who mines a block on the Ethereum mainnet, to submit the block headers from the block which they mined to our Merged Mining contract, and be awarded freshly minted RTC! Currently the merged mining contract requires that each block, the block hash, and the coinbase (miner) are stored in a smart contract, allowing the miner to claim their minted tokens whenever.

The ability to submit block hash and coinbase information is incentivized and can also mint RTC. The first person to submit the blockhash and coinbase information for a given block will receive a small amount of RTC, directly minted to their address. We incentivize storing this information as the user has to pay for the gas costs to invoke the transaction.

There is an experimental command line client that can be used to "mine" the `submitBlock` function. Please see the `cli/miner` directory.

### Objectives - Merged Mining

* [x] Incentivize Block Information (block number, coinbase) Submission
* [x] Reward ETH block miners who have had their block information for the mined block submitted

### Setup - Merged Mining

Deployment
1) Set RTCoin Interface on merged mining contract
2) Set Merged Mining contract address on RTCoin contract

### Gas Usage - Merged Mining

| function | gas |
| ---------|------|
| submit block | 92428 |
| bulk reward claim (10 rewards) | 119239 |

### Statistics - Merged Mining

The total supply increase based off a starting RTC supply of 61.6M is 3.15% if all rewards are claimed. This roughly equates to 242750.7375USD/year and 1942005.9RTC/year however in reality this will likely be much, much less as it is highly unlikely that all blocks will be submitted and that all blocks will have their rewards claimed.

### Block Rewards

With an average 13 second block time, if every block had their rewards claimed an average 728252.1 RTC would be minted a year with a reward of 0.3RTC

Formula to reach this is:
`(seconds per year/block time seconds) * (blocks mined reward)`

### Block Information Submission

With an average 13 second block time, if all blocks have their information submitted, an average 1213753.8RTC would be minted a year with a reward of 0.5RTC per block information submitted

Formula to reach this is:
`(seconds per year/block time seconds) * (block submission reward)`

### Tests - Merged Mining

No special configuration is needed to use the merged mining test located at `RTCoin/miner/main.go`, simply ensure that the blockchain network you're testing it on is actually PoW ETHHASH and not a simulated, or mock testing network. Compatability with truffle/ganache is undetermined as I do not use ganache in my development workflow.

### Limitations - Merged Mining

Block number, and coinbase will only be stored when the transaction is mined. If you want your transaction to include the same information from the block right after you submit your transaction you will need to increase your gas price appropriately in order to ensure the transaction is mined in time. Otherwise, the information for the block at which at which your TX is included will be what is stored in the contract.

The current system isn't as efficient in terms of performance, usability, and gas requirements as I would like it to be. Future versions will avoid the block information submission thereby removing storage costs, and instead only focus on block header validation to determine the validity, and reward the miner. Although this will remove the incentivize block information submission, it will be incredibly more efficient, and usable since it won't require any setup phase for block miners to get their reawrds. They will simply have to submit the required header information. This system is still being designed and should not be expected to be released on the mainnet for quite sometime.

## TEMPORAL (WIP)

---
TEMPORAL is RTrade's first major product, allowing users and enterprise businesses to gain the full benefits of distributed and decentralized storage protocols without having to run any infrastructure themselves. RTrade targets enterprise use case by creating an easy to consume API. For end-users there will be an easy to use web interface, allowing non technical people to use distributed and decentralized storage protocols, just like you were usign a regular website.

### TEMPORAL Links

* [Code Repository](https://github.com/RTradeLtd/Temporal)
* [General Information](https://www.rtradetechnologies.com/en/temporal/)

## MultiSig

---
Within `MultiSig/MultiSigWalletModded.sol` you can find a modified version of Gnosis' MultiSig wallet. It includes a few modifications, such as compiler version 0.4.24, as well as implementing the recommendations pointed out in the OpenZeppelin audit. Additionally, formatting was applied so that the code adheres to standard solidity format and passes linter checks.

### Thanks

---
Thanks to Figs999 for the EventStorage.sol contract which is serving as a basis for block header parsing.
    > [EventStorage.sol](https://github.com/figs999/Ethereum/blob/master/EventStorage.sol)
