var web3;
var contract;
var rtcContract;
var stakes;
const abi = [{"constant":false,"inputs":[],"name":"disableNewStakes","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"allowNewStakes","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_stakeNumber","type":"uint256"}],"name":"withdrawInitialStake","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"TOKENADDRESS","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"uint256"}],"name":"stakes","outputs":[{"name":"initialStake","type":"uint256"},{"name":"blockLocked","type":"uint256"},{"name":"blockUnlocked","type":"uint256"},{"name":"releaseDate","type":"uint256"},{"name":"totalCoinsToMint","type":"uint256"},{"name":"coinsMinted","type":"uint256"},{"name":"rewardPerBlock","type":"uint256"},{"name":"lastBlockWithdrawn","type":"uint256"},{"name":"state","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"RTI","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"newStakesAllowed","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_tokenAddress","type":"address"},{"name":"_recipient","type":"address"},{"name":"_amount","type":"uint256"}],"name":"transferForeignToken","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_stakeNumber","type":"uint256"}],"name":"mint","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"canMint","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_numRTC","type":"uint256"}],"name":"depositStake","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"numberOfStakes","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"activeStakes","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"internalRTCBalances","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"admin","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"VERSION","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_admin","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[],"name":"StakesDisabled","type":"event"},{"anonymous":false,"inputs":[],"name":"StakesEnabled","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_staker","type":"address"},{"indexed":true,"name":"_stakeNum","type":"uint256"},{"indexed":false,"name":"_coinsToMint","type":"uint256"},{"indexed":false,"name":"_releaseDate","type":"uint256"},{"indexed":false,"name":"_releaseBlock","type":"uint256"}],"name":"StakeDeposited","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_staker","type":"address"},{"indexed":true,"name":"_stakeNum","type":"uint256"},{"indexed":false,"name":"_reward","type":"uint256"}],"name":"StakeRewardWithdrawn","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_staker","type":"address"},{"indexed":true,"name":"_stakeNumber","type":"uint256"},{"indexed":false,"name":"_amount","type":"uint256"}],"name":"InitialStakeWithdrawn","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_sender","type":"address"},{"indexed":true,"name":"_recipient","type":"address"},{"indexed":false,"name":"_amount","type":"uint256"}],"name":"ForeignTokenTransfer","type":"event"}];
const contractAddress = "0xD6e33C11CFF866162787b7198030aaC101A61F29";
const rtcAbi = [{"constant":false,"inputs":[],"name":"freezeTransfers","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_owner","type":"address"},{"name":"_recipient","type":"address"},{"name":"_amount","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"stakeContractAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balances","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"stake","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_recipient","type":"address"},{"name":"_amount","type":"uint256"}],"name":"mint","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_contractAddress","type":"address"}],"name":"setStakeContract","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"allowed","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_subtractedValue","type":"uint256"}],"name":"decreaseApproval","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_newAdmin","type":"address"}],"name":"setAdmin","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_holder","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"transferOutEth","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_tokenAddress","type":"address"},{"name":"_recipient","type":"address"},{"name":"_amount","type":"uint256"}],"name":"transferForeignToken","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_recipient","type":"address"},{"name":"_amount","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_mergedMinerValidator","type":"address"}],"name":"setMergedMinerValidator","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"thawTransfers","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_addedValue","type":"uint256"}],"name":"increaseApproval","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_contractAddress","type":"address"}],"name":"setFailOverStakeContract","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"INITIALSUPPLY","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"stakeFailOverRestrictionLifted","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"transfersFrozen","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"mergedMinerValidatorAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_newOwner","type":"address"}],"name":"transferOwnership","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"minters","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"admin","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"VERSION","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_sender","type":"address"},{"indexed":true,"name":"_recipient","type":"address"},{"indexed":false,"name":"_amount","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_owner","type":"address"},{"indexed":true,"name":"_spender","type":"address"},{"indexed":false,"name":"_amount","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_transfersFrozen","type":"bool"}],"name":"TransfersFrozen","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_transfersThawed","type":"bool"}],"name":"TransfersThawed","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_sender","type":"address"},{"indexed":true,"name":"_recipient","type":"address"},{"indexed":false,"name":"_amount","type":"uint256"}],"name":"ForeignTokenTransfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_recipient","type":"address"},{"indexed":false,"name":"_amount","type":"uint256"}],"name":"EthTransferOut","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_contractAddress","type":"address"}],"name":"MergedMinerValidatorSet","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_contractAddress","type":"address"}],"name":"StakeContractSet","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_contractAddress","type":"address"}],"name":"FailOverStakeContractSet","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_stakeContract","type":"address"},{"indexed":true,"name":"_recipient","type":"address"},{"indexed":false,"name":"_mintAmount","type":"uint256"}],"name":"CoinsMinted","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_admin","type":"address"}],"name":"AdminSet","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_previousOwner","type":"address"},{"indexed":false,"name":"_newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"}];
var rtcAddress = "0xecc043b92834c1ebDE65F2181B59597a6588D616";

window.addEventListener('load', function ()
{
    if (typeof web3 !== 'undefined')
    {
        //use current provider if available from metamask
        web3 = new Web3(web3.currentProvider);
        //add function for waiting on transaction success
        //console.log(web3);
        startApp();
    }
    else
    {
        //show error if no metamask account available
    }
});

//add event listener to the deposit stake form button
$("#depositStakeForm").submit(function(event)
{
    // cancels the form submission
    event.preventDefault();
    submitDepositStakeForm(contract);
});

function startApp()
{
    //check if metamask is avaialble
    if(web3.eth.coinbase === null)
    {
        window.alert("Please login to Metamask");
    }
    
    //get user address and display it
    var address = web3.eth.coinbase;
    web3.eth.defaultAccount = web3.eth.coinbase;
    $("#currentAccount").html(address);
    //get rt coin contract
    rtcContract = web3.eth.contract(rtcAbi).at(rtcAddress);
    //get the staking contract
    contract = web3.eth.contract(abi).at(contractAddress);

    //get user eth balance and display it
    web3.eth.getBalance(address, function (error, balance)
    {
        if (!error)
        {
            $("#ethBalance").html(balance.toNumber() / 1000000000000000000 + " ether");
            contract.numberOfStakes(address, function(error, result)
            {
                if(!error)
                {
                    console.log("Num stakes: " + result);
                    stakes = new Array(result.toNumber());
                    getStakerDetails(address);
                }
                else
                {
                    console.log();
                }
            });
        }
        else
        {
            console.error(error);
        }
    });
    
    //get RTC balance and display it
    getRTCBalance(address);
}

function getRTCBalance(address)
{
    rtcContract.balanceOf(address, function(error, result)
    {
        if(!error)
        {
            //console.log(result);
            $("#rtcBalance").html(result.toNumber() / 1000000000000000000 + " RTC");
        }
        else
        {
            console.log(error);
        }
    });
}

function getStakerDetails(address)
{
    //get all stakes for user
    if(stakes.length > 0)
    {
        getStakes(address, 0);
    }
    else
    {
        $("#totalRtcStaked").html("not currently staking");
    }
    
    //displayStakerDetails(stakes[0]);
}

//recursively retreives all stakes
function getStakes(address, id)
{
    contract.stakes(address, id, function(error, result)
    {
        if(!error)
        {
            console.log(result);
            stakes[id] = result;
            addStakeListItem(id);
            if(id === stakes.length-1)
            {
                displayTotals();
            }
            else if(id < stakes.length-1)
            {
                getStakes(address, id+1);
            }
        }
        else
        {
            console.log();
        }
    });
}

function addStakeListItem(id)
{
    var num = id;
    var listItem = $('<li><a href="#">' + num + '</a></li>');
    $("#stakeList").append(listItem);
    
    listItem.click(function(e)
    {
        e.preventDefault();
        displayStakeDetails(stakes[id], id);
		showMintButton(id);
    });
}

function showMintButton(id)
{
	var mintButton = $('<button type="submit" class="btn btn-primary has-spinner">Mint</button>');
	mintButton.click(function(e)
	{
		e.preventDefault();
		mintStake(id);
	});
	$("#mintButton").html(mintButton);
}

function displayStakeDetails(stake, id)
{
    var rtcStaked = stake[0].toNumber() / 1000000000000000000;
    if(rtcStaked > 0)
    {
        //$("#stakerAddress").html(staker[0]);
        $("#rtcStaked").html(rtcStaked);
        $("#rtcMinted").html(stake[5].toNumber() / 1000000000000000000);
        $("#releaseDate").html(new Date(stake[3].toNumber()*1000).toString().split("(")[0]);
        $("#id").html(id);
        $("#state").html(stake[8].toNumber());
    }
    else
    {
        $("#stakerAddress").html(stake[0]);
        $("#rtcStaked").html("not currently staking");
    }
}

function displayTotals()
{
    var totalRtc = 0;
	var totalMinted = 0;
    
    for(var i = 0; i < stakes.length; i++)
    {
        totalRtc += stakes[i][0].toNumber() / 1000000000000000000;
		totalMinted += stakes[i][5].toNumber() / 1000000000000000000;
    }
    
    $("#totalRtcStaked").html(totalRtc);
	$("#totalMinted").html(totalMinted);
}

function displayRewardDetails(reward)
{
    $("#rtcAwarded").html(reward[1].toNumber() / 1000000000000000000);
}

function submitDepositStakeForm(contract)
{
    //get inputs from from
    var rtcToStake = $("#rtcToStake").val() * 1000000000000000000;
    console.log(rtcToStake);
    
    //approve the stake transfer
    rtcContract.approve(contractAddress, rtcToStake, function(error, result)
    {
        if(error)
        {
            console.log(error);
        }
        else
        {
            console.log("waiting for approval confirmation...  (" + result + ")");
            $("#spinnerText").html(" waiting for approval confirmation...");
            $("#spinnerDiv").show();
            getTransactionReceiptMined(result).then(function (receipt)
            {
                console.log(receipt);
                
                //submit inputs to the contract
                console.log("Staking " + rtcToStake + " rtc");
                contract.depositStake(rtcToStake, function(error, result)
                {
                    if(error)
                    {
                        console.log(error);
                    }
                    else
                    {
                        console.log("waiting for stake confirmation...  (" + result + ")");
                        $("#spinnerText").html(" waiting for stake confirmation...");
                        getTransactionReceiptMined(result).then(function (receipt)
                        {
                            console.log("stake completed");
                            $("#spinnerDiv").html("Stake completed. Refresh the page to see staking details.");
                        });
                    }
                });
            });
        }
    });
}

function mintStake(id)
{
	contract.mint(id, function(error, result)
	{
		if(error)
		{
			console.log(error);
		}
		else
		{
			console.log("Minting Stake ID " + id + "...  (" + result + ")");
			$("#spinnerText").html(" minting in progress...");
			getTransactionReceiptMined(result).then(function (receipt)
			{
				console.log("Mint complete");
				$("#spinnerDiv").html("Your freshly minted RTC have been deposited! Refresh to see updated balance.");
			});
		}
	});
}