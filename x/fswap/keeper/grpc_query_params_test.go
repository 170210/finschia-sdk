// package keeper_test

// import (
// 	"testing"

// 	"github.com/Finschia/finschia-sdk/x/fswap/types"

// 	testkeeper "github.com/Finschia/finschia-sdk/testutil/keeper"
// 	sdk "github.com/Finschia/finschia-sdk/types"
// 	"github.com/stretchr/testify/require"
// )

// func TestParamsQuery(t *testing.T) {
// 	keeper, ctx := testkeeper.FswapKeeper(t)
// 	wctx := sdk.WrapSDKContext(ctx)
// 	params := types.DefaultParams()
// 	keeper.SetParams(ctx, params)

// 	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
// 	require.NoError(t, err)
// 	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
// }
