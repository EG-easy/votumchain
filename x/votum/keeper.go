package votum

//Keeper は、storageを管理するstructで、getter/setterのメソッドを実装する
// type Keeper struct {
// 	coinKeeper bank.Keeper  //coinの発行にはKeeperより上位権限をもつBaseKeeperが必要
// 	cdc        *codec.Codec //codecはgo-aminoを用いて、byte codeをdecode/encodeしてtendermint側と通信するときに必要となる
// }
//
// //NewKeeper 新しいKeeperを生成する
// func NewKeeper(cdc *codec.Codec, coinKeeper bank.Keeper) Keeper {
// 	return Keeper{
// 		coinKeeper: coinKeeper,
// 		cdc:        cdc,
// 	}
// }

// Gets the entire Whois metadata struct for a name
// func (k Keeper) GetIssuers(ctx sdk.Context, addr sdk.AccAddress) types.Issuers {
// 	coins := k.coinKeeper.GetCoins(ctx, addr)
// 	store := ctx.KVStore(k.storeKey)
// 	if !k.IsNamePresent(ctx, name) {
// 		return NewWhois()
// 	}
// 	bz := store.Get([]byte(name))
// 	var issuers Issuers
// 	k.cdc.MustUnmarshalBinaryBare(bz, &issuers)
// 	return issuers
// }
//
// // Sets the entire Whois metadata struct for a name
// func (k Keeper) SetIssuers(ctx sdk.Context, name string, whois Whois) {
// 	if whois.Owner.Empty() {
// 		return
// 	}
// 	store := ctx.KVStore(k.storeKey)
// 	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))
// }
