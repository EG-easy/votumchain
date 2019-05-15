# cosmosのモジュールについて
cosmos-sdkにはいくつか便利なモジュールが準備されていて、それらをうまく使うことで、拡張性がある実装をすることができる。
https://github.com/cosmos/cosmos-sdk/blob/master/docs/spec/README.md

## auth

トランザクションとアカウントのの基本となる型を決めるパッケージで、トランザクションのvalidationの仕組みなどが組み込まれている

### **gasとfeeについて**

gasには、無制限にfull nodeのstateを増やしていくのを抑制する目的(anti-spam)とvalidatorに対する経済的なインセンティブを与えるという目的がある。stateのread/write、署名の検証などのトランザクションごとにgasは発生し、トランザクションのサイズによって変化する。トランザクションがmempoolに格納される時に、validatorによってそのトランザクションがfeeを上乗せしているかcheckする仕組みになっている。

### **stateについて**

accounyは次のinterfaceを実装している。
ただし、accountの状態をstoreにwriteする時は、account keeperが必要になる。
```golang
type Account interface {
  GetAddress() AccAddress
  SetAddress(AccAddress)

  GetPubKey() PubKey
  SetPubKey(PubKey)

  GetAccountNumber() uint64
  SetAccountNumber(uint64)

  GetSequence() uint64
  SetSequence(uint64)

  GetCoins() Coins
  SetCoins(Coins)
}

```

base account
```golang
type BaseAccount struct {
  Address       AccAddress
  Coins         Coins
  PubKey        PubKey
  AccountNumber uint64
  Sequence      uint64
}
```

### **Messageについて**

TODO

ante handlerは、mempoolに送られる際のトランザクションをcheckをする際に呼ばれる。

**Typesについて**

|項目|説明|
|---|---|
|StdFee|fee amountとgas limitを規定|
|StdSignature|public keyと署名を規定|
|StdTx|StdFeeとStdSignatureを使って、sdk.Txを実装したもの|

### **Keeperについて**

account keeperは、全てのaccountの各フィールドに対してwrite/read権限を持っている
```golang
type AccountKeeper interface {
  // Return a new account with the next account number and the specified address. Does not save the new account to the store.
  NewAccountWithAddress(AccAddress) Account

  // Return a new account with the next account number. Does not save the new account to the store.
  NewAccount(Account) Account

  // Retrieve an account from the store
  GetAccount(AccAddress) Account

  // Set an account in the store
  SetAccount(Account)

  // Remove an account from the store
  RemoveAccount(Account)

  // Iterate over all accounts, calling the provided function. Stop iteraiton when it returns false.
  IterateAccounts(func(Account) (bool))

  // Fetch the public key of an account at a specified address
  GetPubKey(AccAddress) PubKey

  // Fetch the sequence of an account at a specified address
  GetSequence(AccAddress) uint64

  // Fetch the next account number, and increment the internal counter
  GetNextAccountNumber() uint64
}
```

### **Vestingについて**

TODO
validatorにdelegate/undelegateするために必要となるaccountの形式。
genesis時にbalance Xと vesting time Tを決める。


## bank
- トークンの送受信
- delegate/undelegate

### Keepers
3つのKeeperのinterfaceが準備されている。
それぞれのinterfaceのpermissionが違うので、用途によって使い分ける必要がある。

- 共通のtype
複数のトークンを送受信する時に必要になるstruct

input
```golang
type Input struct {
  Address AccAddress
  Coins   Coins
}
```
output
```golang
type Output struct {
  Address AccAddress
  Coins   Coins
}
```

- BaseKeeperについて

全ての権限が付与されているアカウント（バランスの変更/mint/burn）
```golang
type BaseKeeper interface {
  SetCoins(addr AccAddress, amt Coins)
  SubtractCoins(addr AccAddress, amt Coins)
  AddCoins(addr AccAddress, amt Coins)
  InputOutputCoins(inputs []Input, outputs []Output)
}
```

``setCoins``は、指定されたアドレスのバランスの変更を行う関数
```golang
setCoins(addr AccAddress, amt Coins)
  account = accountKeeper.getAccount(addr)
  if account == nil
    fail with "no account found"
  account.Coins = amt
  accountKeeper.setAccount(account)
```

``subtractCoins``は、指定されたアドレスから指定されたamountを引く関数
```golang
subtractCoins(addr AccAddress, amt Coins)
  oldCoins = getCoins(addr)
  newCoins = oldCoins - amt
  if newCoins < 0
    fail with "cannot end up with negative coins"
  setCoins(addr, newCoins)
```

``addCoins``は、指定されたアドレスから指定されたamountを足す関数
```golang
addCoins(addr AccAddress, amt Coins)
  oldCoins = getCoins(addr)
  newCoins = oldCoins + amt
  setCoins(addr, newCoins)
```

``inputOutputCoins``は、指定されたアドレスからtokenをtransferする関数
```golang
inputOutputCoins(inputs []Input, outputs []Output)
  for input in inputs
    subtractCoins(input.Address, input.Coins)
  for output in outputs
    addCoins(output.Address, output.Coins)
```

- SendKeeper

トークンの送受信が可能
```golang
type SendKeeper interface {
  SendCoins(from AccAddress, to AccAddress, amt Coins)
}
```

`sendCoins`は、トークンをあるアドレスから別のアドレスに送る
```golang
sendCoins(from AccAddress, to AccAddress, amt Coins) {
  subtractCoins(from, amt)
  addCoins(to, amt)
}
```


- ViewKeeper

read-only accessしかないので、アドレスの残高を参照することができる。
```golang
type ViewKeeper interface {
  GetCoins(addr AccAddress) Coins
  HasCoins(addr AccAddress, amt Coins) bool
}

getCoins(addr AccAddress) {
  account = accountKeeper.getAccount(addr)
  if account == nil
    return Coins{}
  return account.Coins
}

hasCoins(addr AccAddress, amt Coins) {
  account = accountKeeper.getAccount(addr)
  coins = getCoins(addr)
  return coins >= amt
}
```

### Messagesについて
```golang
type MsgSend struct {
  Inputs  []Input
  Outputs []Output
}
```

`handleMsgSend`は、`inputOutputCoins`関数を実行する
```golang
handleMsgSend(msg MsgSend) {
  inputSum = 0
  for input in inputs
    inputSum += input.Amount
  outputSum = 0
  for output in outputs
    outputSum += output.Amount
  if inputSum != outputSum:
    fail with "input/output amount mismatch"

  return inputOutputCoins(msg.Inputs, msg.Outputs)
}
```

## governance
提案と投票

現在governance モジュールでサポートされている機能
- proposalの提出

proposerは、一定のdeposit額と共にproposalを提出することができ、一定の投票期間が設けられる
- 投票

参加者は、MinDepositに達したproposalに投票できる
- 投票権の継承と罰則

validatorは、delegattorsが投票しない場合、delegatorsの投票権を継承することができる
- depositの復活

proposerは、提案が受け入れられる、または投票が行われなかった場合は、depositを復活させることができる

native token 保有者は、`TxGovProposal`transactionによって、proposalを提出することができる。一度、提出されると、`proposalID`が振られる。
spamを避けるために、proposalの提出には、一定額のdeposit(MinDeposit)が必要となる。`TxGovDeposit`transactionを発行することによって、proposerに限らず、depositをすることができる。`MinDeposit`に達すると、投票期間となる。`MaxDepositPeriod`の間に、`MinDeposit`に達しない場合は、proposalは棄却され、depositした分は元の持ち主に返される。proposalの種類は、`PlainTextProposal`と`SoftwareUpgradeProposal`がある。

### 投票
native tokenをvalidatorにbondしているuserが、参加者とみなされる。
投票には、`Yes`, `No`, `NoWithVeto`, `Abstain`の4つの種類がある。









## staking
PoSやbonding
## slashing
validatorに罰則を与えれるメカニズム
## distribution
手数料の分配
## inflation
インフレによるトークンの新規作成
## IBC
Inter-Blockchain Communication (IBC) protocol.


