package votum

// //Keeper は、storageを管理するstructで、getter/setterのメソッドを実装する
// type Keeper struct {
// 	bk  types.BankKeeper
// 	ak  types.AccountKeeper
// 	cdc *codec.Codec //codecはgo-aminoを用いて、byte codeをdecode/encodeしてtendermint側と通信するときに必要となる
// }
//
// //NewKeeper 新しいKeeperを生成する
// func NewKeeper(cdc *codec.Codec, bk bank.BaseKeeper, ak auth.AccountKeeper) Keeper {
// 	return Keeper{
// 		bk:  bk,
// 		ak:  ak,
// 		cdc: cdc,
// 	}
// }
//
// func (k Keeper) IssueToken(ctx sdk.Context, issuer sdk.AccAddress, newCoin sdk.Coin) sdk.Error {
//
// 	// acc := k.bk.GetAccount(ctx, issuer)
// 	// coins := acc.GetCoins()
// 	// newCoins := sdk.NewCoins(newCoin).Add(coins)
// 	coins := k.bk.GetCoins(ctx, issuer)
// 	newCoins := sdk.NewCoins(newCoin).Add(coins)
//
// 	// log.Println(acc)
// 	log.Println(coins)
// 	log.Println(newCoins)
//
// 	if ok := newCoins.IsValid(); !ok {
// 		return sdk.ErrInvalidCoins("Issuing New Coin is failed")
// 	}
//
// 	if _, err := k.bk.AddCoins(ctx, issuer, newCoins); err != nil {
// 		return sdk.ErrInvalidCoins("Issuing New Coin is failed")
// 	}
// 	return nil
// }
//
// // GetCoins returns the coins at the addr.
// func (keeper Keeper) GetCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
// 	acc := keeper.ak.GetAccount(ctx, addr)
// 	if acc == nil {
// 		return sdk.NewCoins()
// 	}
// 	return acc.GetCoins()
// }
//
// // GetAccount implements sdk.AccountKeeper.
// func (keeper Keeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) exported.Account {
// 	store := ctx.KVStore(keeper.ak.key)
// 	bz := store.Get(types.AddressStoreKey(addr))
// 	if bz == nil {
// 		return nil
// 	}
// 	acc := keeper.decodeAccount(bz)
// 	return acc
// }
