<p align="center">
  <img src="./votum-logo.png" width="300">
</p>

# Get Started

```
$ git clone https://github.com/EG-easy/votumchain.git
cd votumchain
export GO111MODULE=on
make install
```

**Initialize configuration files and genesis file**
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
$ votumd add-genesis-account $(votumcli keys show jack -a) 1000votum,100000000stake
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


## Voting

### send Proposal 
```
$ votumcli tx gov submit-proposal --proposal="proposal/proposal.json" --from jack
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

```
# Work from the votumchain Repo
$ cd $GOPATH/src/github.com/EG-easy/votumchain

# Build the linux binary in ./build
$ make build-linux

# Build votumchain/votumdnode image
$ make build-docker-votumdnode
```

### Run Your Testnet

To start a 4 node testnet run:

```
make localnet-start
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
make build-linux localnet-start
```


### Keys & Accounts

To interact with `votumcli` and start querying state or creating txs, you use the
`votumcli` directory of any given node as your `home`, for example:

```shell
votumcli keys list --home ./build/node0/votumcli
```


## License
Licensed under the [Apache v2 License](LICENSE).
