# RTrade Technologies Ltd smart contracts are now live on the mainnet

It's been a long awaited road however we are finally happy to announce that as of now, all initial smart contracts for RTrade Technologies Ltd are now live on the Ethereum mainnet!
You can find a list of all contracts on our [RT-Contracts Repository](https://github.com/RTradeLtd/RT-Contracts). In the spirit of transparency, we have full verified all smart contracts on Etherscan. We are also making public the address of the multisignature wallet used to control all ownership functions.

## Contract List

* [MultiSignatureWallet](https://etherscan.io/address/0x41fb0e5bd1dfe3b61e9a09ebd4105c2e35b0bcbd)
* [RTCoin](https://etherscan.io/token/0xecc043b92834c1ebde65f2181b59597a6588d616)
* [Stake](https://etherscan.io/address/0xD6e33C11CFF866162787b7198030aaC101A61F29#code)
* [RTCETH](https://etherscan.io/address/0x40e68e3F58b9C1928954BEe5dEcC09A45aA531f8#code)
* [Vesting](https://etherscan.io/address/0x211D8B3EB985626B3363D3AeDd9B071113660330#code)
* [Merged Miner](https://etherscan.io/address/0x19Cdf52Ce778ef01BAc6A87615dC292A0BC14d5C#code)

## Notable Transactions

* [Etherscan Token Supply and Tracking](https://etherscan.io/token/0xecc043b92834c1ebde65f2181b59597a6588d616)
* [First Stake Deposit](https://etherscan.io/tx/0x573df9c610e72240b2777364b1a157913a69cf53d0bda53ba7d68afa15c6a5c0)
* [First Stake Minting](https://etherscan.io/tx/0x42641231b78d87ab0f2c408829ce4f94580ade35d1a82961a2bd10d47163fadc)
* [First Merged Mined Block Submission](https://etherscan.io/tx/0x524cdcb7f5f503a44506ab5a35c288022c080556406b61263315f8c247dbeac0)

## General Information

We were originally intending to deploy the TEMPORAL payment contract today, however due to the long deployment process we had today, as well as the fact that the payment contract won't be used for a little bit, we decided it would be best to post-pone that for another week. We have activated the core components of our token, and you can mint tokens from our merged mining contract, or staking contract! Don't have any tokens? You can buy them from our RTCETH smart contract, or try your chance at the merged mining contract and get tokens for just submitting a transaction.

## Instructions 

Staking is quite simple, however here is a quick tutorial:

* 1) You'll want to approve the staking contract to spend funds on your behalf, this can be done via the `approve` function on the RTC token contract.
* 2) After the approve transaction is successfully mined, you'll want to call the `depositStake` function on the Staking smart contract.
* 3) After your deposit transaction has been mined, and at least 1 block has passed, you may claim your stake! To do so, invoke the `mint` function on the Staking smart contract, and you'll receive freshly minted RTC directly to your wallet.
* 4) After 2103840 blocks, and 31557600 seconds have passed sicne your stake deposit, you may withdraw your initial stake via the `withdrawInitialStake` function.

Merged mining is also very simple:

* 1) Once per a block, information for that block may be submitted to the merged mining smart contract via the `submitBlock` function.Keep in mind whatever transaction gets mined first is the address that receives the reward.

