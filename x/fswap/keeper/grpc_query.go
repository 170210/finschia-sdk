package keeper

import (
	"context"

	sdk "github.com/Finschia/finschia-sdk/types"
	"github.com/Finschia/finschia-sdk/x/fswap/types"
)

var _ types.QueryServer = Keeper{}

// Swapped implements types.QueryServer.
func (k Keeper) Swapped(ctx context.Context, req *types.QuerySwappedRequest) (*types.QuerySwappedResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return &types.QuerySwappedResponse{Swapped: k.GetSwapped(sdkCtx)}, nil
}

// TotalNewCurrencySwapLimit implements types.QueryServer.
func (k Keeper) TotalNewCurrencySwapLimit(ctx context.Context, req *types.QueryTotalSwappableAmountRequest) (*types.QueryTotalSwappableAmountResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return &types.QueryTotalSwappableAmountResponse{SwappableNewCoinAmount: k.GetTotalSupply(sdkCtx)}, nil
}
