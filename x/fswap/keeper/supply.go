package keeper

import (
	sdk "github.com/Finschia/finschia-sdk/types"
	"github.com/Finschia/finschia-sdk/x/fswap/types"
)

func (k Keeper) GetTotalSupply(ctx sdk.Context) sdk.Coin {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte{types.TotalSupplyKey})
	var totalSupply sdk.Coin
	if bz == nil {
		//todo
		panic(types.ErrParamsNotFound)
	}
	k.cdc.MustUnmarshal(bz, &totalSupply)
	return totalSupply
}

func (k Keeper) SetTotalSupply(ctx sdk.Context, totalSupply sdk.Coin) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&totalSupply)
	if err != nil {
		return err
	}
	store.Set([]byte{types.TotalSupplyKey}, bz)
	return nil
}
