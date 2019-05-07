## client command

### status
Query remote node for status
``` json
$ votumchaincli status
{
  "node_info": {
    "protocol_version": {
      "p2p": "7",
      "block": "10",
      "app": "0"
    },
    "id": "790aaf8c23e1a71cc47f20a423a9dd44ceb622c0",
    "listen_addr": "tcp://0.0.0.0:26656",
    "network": "testchain",
    "version": "0.30.1",
    "channels": "4020212223303800",
    "moniker": "eguegu.local",
    "other": {
      "tx_index": "on",
      "rpc_address": "tcp://0.0.0.0:26657"
    }
  },
  "sync_info": {
    "latest_block_hash": "50822848E9D1DAF0D907D871404EC1C93B5833B45BF9DDB24FFD8E3BFCFBD2DC",
    "latest_app_hash": "2FB85BB1FA7D51E4EB706C11A82B686AD06B012A627A3C824E272CDADD1E5E0D",
    "latest_block_height": "236",
    "latest_block_time": "2019-05-07T03:07:24.246668Z",
    "catching_up": false
  },
  "validator_info": {
    "address": "CD0338521A1AA2948D700921474F6CE1967F33A6",
    "pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "AtA1u+LM9fKJlbtp/82O6HfP1Z3pk5w5+00Ks/0sNZY="
    },
    "voting_power": "10"
  }
}
```

### config
Create or query a Gaia CLI configuration file
```json
votumchaincli config
chain-id = "testchain"
indent = true
output = "json"
trust-node = true
```

### query
Querying subcommands

- tendermint-validator-set
Get the full tendermint validator set at given height
```json
$ votumchaincli query tendermint-validator-set
{
 "block_height": "658",
 "validators": [
  {
   "address": "cosmosvalcons1e5pns5s6r23ffrtspys5wnmvuxt87vax0gkwq6",
   "pub_key": "cosmosvalconspub1zcjduepqqtgrtwlzen6l9zv4hd5llnvwapmul4vaaxfecw0mf59t8lfvxktqsadurq",
   "proposer_priority": "0",
   "voting_power": "10"
  }
 ]
}
```

- block
Get verified data for a the block at given height

```json
$ votumchaincli query block 120
{
  "block_meta": {
    "block_id": {
      "hash": "3292A0E704F7CD88B7A896790393753B47364CD96DB523CF4B897FA73442513B",
      "parts": {
        "total": "1",
        "hash": "617201F78A5B401928F62E75528903D7158F59835F18DA1D44814EC315B21ABF"
      }
    },
    "header": {
      "version": {
        "block": "10",
        "app": "0"
      },
      "chain_id": "testchain",
      "height": "120",
      "time": "2019-05-07T02:57:40.416132Z",
      "num_txs": "0",
      "total_txs": "0",
      "last_block_id": {
        "hash": "584254C51CFED45149B7620DFB6F1F145BB4B9C2509471F8C242390B9511D16C",
        "parts": {
          "total": "1",
          "hash": "6F98820E8C7D85BDF8D51F832CB15D8FC4358C2FBD85A4069BC65AEFEE5C1E85"
        }
      },
      "last_commit_hash": "86CC9FAEAF0F34902BC4F724FBA815ACC4E80D12F085F81D8E98081A90303B2A",
      "data_hash": "",
      "validators_hash": "0D97B27AAF0C799C9886DF0AFF3A096A9EF29EB6770CEF31EAFEAE7BA60FABF2",
      "next_validators_hash": "0D97B27AAF0C799C9886DF0AFF3A096A9EF29EB6770CEF31EAFEAE7BA60FABF2",
      "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
      "app_hash": "30C2A16A18B804A6F57DABED47C7A7B2433A73AEEF3D3B153A94B0AC516DB1D6",
      "last_results_hash": "",
      "evidence_hash": "",
      "proposer_address": "CD0338521A1AA2948D700921474F6CE1967F33A6"
    }
  },
  "block": {
    "header": {
      "version": {
        "block": "10",
        "app": "0"
      },
      "chain_id": "testchain",
      "height": "120",
      "time": "2019-05-07T02:57:40.416132Z",
      "num_txs": "0",
      "total_txs": "0",
      "last_block_id": {
        "hash": "584254C51CFED45149B7620DFB6F1F145BB4B9C2509471F8C242390B9511D16C",
        "parts": {
          "total": "1",
          "hash": "6F98820E8C7D85BDF8D51F832CB15D8FC4358C2FBD85A4069BC65AEFEE5C1E85"
        }
      },
      "last_commit_hash": "86CC9FAEAF0F34902BC4F724FBA815ACC4E80D12F085F81D8E98081A90303B2A",
      "data_hash": "", "validators_hash": "0D97B27AAF0C799C9886DF0AFF3A096A9EF29EB6770CEF31EAFEAE7BA60FABF2", "next_validators_hash": "0D97B27AAF0C799C9886DF0AFF3A096A9EF29EB6770CEF31EAFEAE7BA60FABF2", "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F", "app_hash": "30C2A16A18B804A6F57DABED47C7A7B2433A73AEEF3D3B153A94B0AC516DB1D6", "last_results_hash": "", "evidence_hash": "", "proposer_address": "CD0338521A1AA2948D700921474F6CE1967F33A6"
    },
    "data": {
      "txs": null
    },
    "evidence": {
      "evidence": null
    },
    "last_commit": {
      "block_id": {
        "hash": "584254C51CFED45149B7620DFB6F1F145BB4B9C2509471F8C242390B9511D16C",
        "parts": {
          "total": "1",
          "hash": "6F98820E8C7D85BDF8D51F832CB15D8FC4358C2FBD85A4069BC65AEFEE5C1E85"
        }
      },
      "precommits": [
        {
          "type": 2,
          "height": "119",
          "round": "0",
          "block_id": {
            "hash": "584254C51CFED45149B7620DFB6F1F145BB4B9C2509471F8C242390B9511D16C",
            "parts": {
              "total": "1",
              "hash": "6F98820E8C7D85BDF8D51F832CB15D8FC4358C2FBD85A4069BC65AEFEE5C1E85"
            }
          },
          "timestamp": "2019-05-07T02:57:40.416132Z",
          "validator_address": "CD0338521A1AA2948D700921474F6CE1967F33A6",
          "validator_index": "0",
          "signature": "TfscmSDF4fpgkklAkBiOW6z1Bjf1eW4WmAz4H/lqNzBH36gwcA/AzdFSci+9/IEWCFeV3XWCdwtIxnsf03EyAQ=="
        }
      ]
    }
  }
}
```

- tx
Matches this txhash over all committed blocks
```json
$ votumchaincli query tx C3610DABFF0262C04D702A9422C35B43AE69AB78AD9F3926461478A7B7D1C30F
{
 "height": "764",
 "txhash": "C3610DABFF0262C04D702A9422C35B43AE69AB78AD9F3926461478A7B7D1C30F",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "24101",
 "tags": [
  {
   "key": "action",
   "value": "send"
  },
  {
   "key": "sender",
   "value": "cosmos1hj6ds7akwuu3ys5h6200epsenfm7t2dae7rzth"
  },
  {
   "key": "recipient",
   "value": "cosmos1xzzcenw2wfu8nrpev00jyynl2rrpp3tsdlxfdm"
  }
 ],
 "tx": {
  "type": "auth/StdTx",
  "value": {
   "msg": [
    {
     "type": "cosmos-sdk/Send",
     "value": {
      "from_address": "cosmos1hj6ds7akwuu3ys5h6200epsenfm7t2dae7rzth",
      "to_address": "cosmos1xzzcenw2wfu8nrpev00jyynl2rrpp3tsdlxfdm",
      "amount": [
       {
        "denom": "mycoin",
        "amount": "10"
       }
      ]
     }
    }
   ],
   "fee": {
    "amount": null,
    "gas": "200000"
   },
   "signatures": [
    {
     "pub_key": {
      "type": "tendermint/PubKeySecp256k1",
      "value": "Ao7upJFN0vMh097STrZ3vJXtPvixAjty3/0vdeOV6rJY"
     },
     "signature": "8/8E579DaWXTWNAL4v3Z4Ckc/8AYS4SsWUH6B8F6YzYqzQ7SAhZOIdx6onbMrm6K1MG8LZpxA0f49ud6Ao9Mqg=="
    }
   ],
   "memo": ""
  }
 }
}
```

- account
Query account balance
```json
$  votumchaincli query account cosmos1xzzcenw2wfu8nrpev00jyynl2rrpp3tsdlxfdm
{
 "type": "auth/Account",
 "value": {
  "address": "cosmos1xzzcenw2wfu8nrpev00jyynl2rrpp3tsdlxfdm",
  "coins": [
   {
    "denom": "jackcoin",
    "amount": "1000"
   },
   {
    "denom": "mycoin",
    "amount": "1010"
   }
  ],
  "public_key": null,
  "account_number": "0",
  "sequence": "0"
 }
}
```

### tx
Transactions subcommands

- send
Create and sign a send tx
```bash
$ votumchaincli tx send cosmos1xzzcenw2wfu8nrpev00jyynl2rrpp3tsdlxfdm 10mycoin --from alice
```
```json
{
 "height": "764",
 "txhash": "C3610DABFF0262C04D702A9422C35B43AE69AB78AD9F3926461478A7B7D1C30F",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "24101",
 "tags": [
  {
   "key": "action",
   "value": "send"
  },
  {
   "key": "sender",
   "value": "cosmos1hj6ds7akwuu3ys5h6200epsenfm7t2dae7rzth"
  },
  {
   "key": "recipient",
   "value": "cosmos1xzzcenw2wfu8nrpev00jyynl2rrpp3tsdlxfdm"
  }
 ]
}
```

- sign and broadcast

Sign transactions generated offline

Broadcast transactions generated offline


### rest-server
Start LCD (light-client daemon), a local REST server


### keys
Add or view local private keys

- add

Add an encrypted private key (either newly generated or recovered), encrypt it, and save to disk
```json
$ votumchaincli keys add test
{
  "name": "test",
  "type": "local",
  "address": "cosmos1xrt2t6p0rjcxthpu66200te9n2cahchauwswzg",
  "pub_key": "cosmospub1addwnpepqdgehftmx7ee0qkfumelx6ztdjyldy7qggxrua65yleq6hmlfjzm7l8uhdn",
  "mnemonic": "fluid flavor system column rack caught grant parent duty cotton ship depth glad argue index exercise note broken pyramid carbon lawn hunt announce energy"
}
```

- list 
List all keys
```json
$ votumchaincli keys list
[
  {
    "name": "alice",
    "type": "local",
    "address": "cosmos1hj6ds7akwuu3ys5h6200epsenfm7t2dae7rzth",
    "pub_key": "cosmospub1addwnpepq28wafy3fhf0xgwnmmfyadnhhj2760hckyprkukll5hhtcu4a2e9sy3trp4"
  },
  {
    "name": "jack",
    "type": "local",
    "address": "cosmos1xzzcenw2wfu8nrpev00jyynl2rrpp3tsdlxfdm",
    "pub_key": "cosmospub1addwnpepqtv8lasvgd687ukdcdavh3cn9qjhxml2kpckm62kjvvmc44p99hxsusa4a6"
  },
  {
    "name": "test",
    "type": "local",
    "address": "cosmos1xrt2t6p0rjcxthpu66200te9n2cahchauwswzg",
    "pub_key": "cosmospub1addwnpepqdgehftmx7ee0qkfumelx6ztdjyldy7qggxrua65yleq6hmlfjzm7l8uhdn"
  }
]
```

- show 
Show key info for the given name
```bash
$ votumchaincli keys show jack -a
cosmos1xzzcenw2wfu8nrpev00jyynl2rrpp3tsdlxfdm
```

- update
Change the password used to protect private key
```bash
$ votumchaincli keys update test
Password successfully updated!
```

- delete
Delete the given key
```bash
$ votumchaincli keys delete test
Key deleted forever (uh oh!)
```
### Flags
```
      --chain-id string   Chain ID of tendermint node
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
  -h, --help              help for votumchaincli
      --home string       directory for config and data (default "$HOME/.votumchaincli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors
```
