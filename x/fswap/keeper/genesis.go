package keeper

import (
	sdk "github.com/Finschia/finschia-sdk/types"
	"github.com/Finschia/finschia-sdk/x/fswap/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, bk types.BankKeeper, genState types.GenesisState) error {
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		return err
	}
	// todo cony
	totalOldCoinsSupply := bk.GetSupply(ctx, "cony").Amount
	// todo check & modify
	a := int64(genState.Params.SwapRate)
	b := int(genState.Params.SwapRateDecimals)
	totalNewCoinsSupply := totalOldCoinsSupply.Mul(sdk.NewIntWithDecimal(a, b))
	totalNewCoins := sdk.NewCoin(genState.Params.NewCoinDenom, totalNewCoinsSupply)
	if err := k.SetTotalSupply(ctx, totalNewCoins); err != nil {
		return err
	}
	return nil
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
