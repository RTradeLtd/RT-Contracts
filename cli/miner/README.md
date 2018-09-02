# Miner

This is a simple comand line program designed to "mine" on the Merged Miner Validator smart contract. By this, we mean that it will repeatedly submit transactions to the smart contract submitting information. It is currently a WIP, and bugs should be expected. As such, please use at your own risk. There is no implied warranty for this piece of software, and any loss of funds that occur as such is the responsibility of the person running this program. If you disagree with any of this, then don't run this program. As of this commit, "Mining" is the only free way to earn RTC.

Currently only keystore files, like those generated via MEW, or Mist are supported.

## "Automated" Setup

1) Run the `setup.sh` script, enter details as needed

### Manual Setup

1) Update `config.json` with an appropriate endpoint to connect to the ethereum blockchain with
2) Update `config.json` with the absolute path to your keystore key file
3) Update `config.json` with the password to unlock your keystore key file
4) Update `config.json` with the address of the merged miner contract that you wish to use