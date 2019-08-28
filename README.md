<p align="center">
  <img src="./votum-logo.png" width="300">
</p>

[![version](https://img.shields.io/github/v/tag/EG-easy/votumchain)](https://github.com/EG-easy/votumchain/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/EG-easy/votumchain)](https://goreportcard.com/report/github.com/EG-easy/votumchain)
[![GolangCI](https://golangci.com/badges/github.com/EG-easy/votumchain.svg)](https://golangci.com/r/github.com/EG-easy/votumchain)

VotumChain is a blockchain application using cosmos-sdk that anyone can use when making important decisions and when the results need to be published publicly.

**WARNING**: VotumChain is under VERY ACTIVE DEVELOPMENT and should be treated as pre-alpha software. This means it is not meant to be run in production, its APIs are subject to change without warning and should not be relied upon, and it should not be used to hold any value.


# What is it

- Cost reduction when making decisions by the majority, such as voting in elections and general meetings.
- Prevent falsification of voting results.
- It is possible to collate voting results.
- Present the ideal way of community governance decision-making.
- It is finally good that 1741 municipalities nationwide have validators.

# Implementation
## Completed
- basic blockchain functions
- issue my token
- vote using staking tokens

## Current Work
- vote using my tokens
- allow users to decide election params 

## Next Steps
- add new type of voting params 
- CI/CD setting
- release testnet
- integate with blockchain explore

# Quick Start
Go 1.12.1+ is required for the Cosmos SDK.

## Install votumcli and votumd

```
$ mkdir -p $GOPATH/src/github.com/EG-easy
$ cd $GOPATH/src/github.com/EG-easy
$ go clone github.com/EG-easy/votumchain.git
$ cd votumchain && git checkout master
$ export GO111MODULE=on
$ make install
```

Try `votumcli version` and `votumd version` to verify everything is OK!

## **Initialize configuration files and genesis file**

```
$ votumd init eguegu --chain-id testchain
```

**Copy the `Address` output here and save it for later use**
```
$ votumcli keys add jack
```

**Copy the `Address` output here and save it for later use**


**Add account with coins to the genesis file**

```
$ votumd add-genesis-account $(votumcli keys show jack -a) 100000000votum,100000000stake
```


**Configure your CLI to eliminate need for chain-id flag**

```
$ votumcli config chain-id testchain
$ votumcli config output json
$ votumcli config indent true
$ votumcli config trust-node true
```

```
$ votumd gentx --name jack
$ votumd collect-gentxs
$ votumd validate-genesis
```

Now let's start!  
```
$ votumd start
```

## Voting

### send Proposal 
Firstly, you need to broadcast a proposal to the network.
You can modify the title, description, deposit of proposal as you like in `proposal/proposal.json`.

```
$ votumcli tx votum submit-proposal --proposal="proposal/proposal.json" --from jack
```

### deposit 
You need to deposit votum token to start voting period.
The default time of voting period is only 120 seconds.
You can change the parameter in `votum/genesis.go`.

```
$ votumcli tx votum deposit 1 1000000votum --from jack 
```

### vote 
Those who can already stake token can vote the proposal as its stake amount.

```
$ votumcli tx votum vote 1 yes --from jack
```

### Check Result
Check proposal status and final result with this command.

```
$ votumcli query votum proposal 1
```

Check deposit status about the certain address.
```
$ votumcli query votum deposit $(votumcli keys show -a jack)
```

Check vote status of all votes.
```
$ votumcli query votum votes 1
```

## Use docker

```
$ docker build -t votum .
```

```
$ docker run --rm -it votum sh
```

```
$ docker exec -it votum sh
```

## 4 validators local private net

```bash
# Work from the votumchain Repo
$ cd $GOPATH/src/github.com/EG-easy/votumchain

# Build the linux binary in ./build
$ make build-linux

# Build votumchain/votumdnode image
$ make build-docker
```

### Run Your Testnet

To start a 4 node testnet run:

```
$ make localnet-start
```

This command creates a 4-node network using the votumdnode image.
The ports for each node are found in this table:

| Node ID | P2P Port | RPC Port |
| --------|-------|------|
| `votumnode0` | `26656` | `26657` |
| `votumnode1` | `26659` | `26660` |
| `votumnode2` | `26661` | `26662` |
| `votumnode3` | `26663` | `26664` |

To update the binary, just rebuild it and restart the nodes:

```
$ make build-linux localnet-start
```


### Keys & Accounts

To interact with `votumcli` and start querying state or creating txs, you use the
`votumcli` directory of any given node as your `home`, for example:

```shell
$ votumcli keys list --home ./build/node0/votumcli
```


## License
Licensed under the [Apache v2 License](LICENSE).
