package keeper

import (
	sdk "github.com/Finschia/finschia-sdk/types"
	"github.com/Finschia/finschia-sdk/x/fswap/types"
)

const (
	DefaultOldCoins string = "cony"
)

var (
	DefaultSwapRate = sdk.NewDecWithPrec(148079656, 6)
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, bk types.BankKeeper, genState types.GenesisState) error {
	if err := k.SetParams(ctx, genState.Params); err != nil {
		return err
	}
	if err := k.SetSwapped(ctx, genState.Swapped); err != nil {
		return err
	}
	totalOldCoinsSupply := bk.GetSupply(ctx, DefaultOldCoins).Amount
	totalNewCoinsSupply := DefaultSwapRate.MulInt(totalOldCoinsSupply)
	totalNewCoins := sdk.NewDecCoinFromDec(genState.Params.NewCoinDenom, totalNewCoinsSupply)
	if err := k.SetTotalSupply(ctx, totalNewCoins); err != nil {
		return err
	}
	return nil
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:  k.GetParams(ctx),
		Swapped: k.GetSwapped(ctx),
	}
}
