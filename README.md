<p align="center">
  <img src="./votum-logo.png" width="300">
</p>

[![version](https://img.shields.io/github/v/tag/EG-easy/votumchain)](https://github.com/EG-easy/votumchain/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/EG-easy/votumchain)](https://goreportcard.com/report/github.com/EG-easy/votumchain)
[![GolangCI](https://golangci.com/badges/github.com/EG-easy/votumchain.svg)](https://golangci.com/r/github.com/EG-easy/votumchain)

VotumChain is a blockchain application using cosmos-sdk that anyone can use when making important decisions and when the results need to be published publicly, presenting the ideal way of community governance decision-making.VotumChain can reduce cost when making decisions by the majority, such as voting in elections and general meetings, and prevent falsification of voting results. It is possible to refer every voting results.

**WARNING**: VotumChain is under VERY ACTIVE DEVELOPMENT and should be treated as pre-alpha software. This means it is not meant to be run in production, its APIs are subject to change without warning and should not be relied upon, and it should not be used to hold any value.

# Implementation
## Completed
- [x] basic blockchain functions
- [x] issue my token
- [x] vote using staking tokens

## Current Work
- [ ] vote using my tokens
- [ ] allow users to decide election params 

## Next Steps
- [ ] add new type of voting params 
- [ ] CI/CD setting
- [ ] release testnet
- [ ] integate with blockchain explore

# Quick Start
Go 1.12.1+ is required for the Cosmos SDK.

Firstly, watch [Demo](https://drive.google.com/file/d/157b35fApxevHui0C2ABp00ugYwjwofba/view?usp=sharing)!

## Install votumcli and votumd

```bash
$ mkdir -p $GOPATH/src/github.com/EG-easy
$ cd $GOPATH/src/github.com/EG-easy
$ git clone github.com/EG-easy/votumchain.git
$ cd votumchain && git checkout master
$ export GO111MODULE=on
$ make install
```

Try `votumcli version` and `votumd version` to verify everything is OK!

## **Initialize configuration files and genesis file**

Just use shell scripts bellow.
```
$ sh scripts/start.sh
```

Or you can follow the command.

```bash
$ votumd init eguegu --chain-id testchain
```

**Copy the `Address` output here and save it for later use**

```bash
$ votumcli keys add jack
```

**Add account with coins to the genesis file**

```bash
$ votumd add-genesis-account $(votumcli keys show jack -a) 100000000votum,100000000stake
``` 

**Configure your CLI to eliminate need for chain-id flag**

```bash
$ votumcli config chain-id testchain
$ votumcli config output json
$ votumcli config indent true
$ votumcli config trust-node true
```

```bash
$ votumd gentx --name jack
$ votumd collect-gentxs
$ votumd validate-genesis
```

**Now let's start!**
```bash
$ votumd start
```

## Voting

### send Proposal 
Firstly, you need to broadcast a proposal to the network.
You can modify the title, description, deposit of proposal as you like in `proposal/proposal.json`.

```bash
$ votumcli tx votum submit-proposal --proposal="proposal/proposal.json" --from jack
```

### deposit 
You need to deposit votum token to start voting period.
The default time of voting period is only 120 seconds.
You can change the parameter in `votum/genesis.go`.

```bash
$ votumcli tx votum deposit 1 1000000votum --from jack 
```

### vote 
Those who can already stake token can vote the proposal as its stake amount.

```bash
$ votumcli tx votum vote 1 yes --from jack
```

### Check Result
Check proposal status and final result with this command.

```bash
$ votumcli query votum proposal 1
```

Check deposit status about the certain address.
```bash
$ votumcli query votum deposit $(votumcli keys show -a jack)
```

Check vote status of all votes.
```bash
$ votumcli query votum votes 1
```

## Use docker

```bash
$ docker build -t votum .
```

```bash
$ docker run --rm -it votum sh
```

```bash
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

```bash
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

```bash
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
