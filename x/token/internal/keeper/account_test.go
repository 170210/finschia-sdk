package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/line/link/x/token/internal/types"
	"github.com/stretchr/testify/require"
)

func verifyAccountFunc(t *testing.T, expected types.Account, actual types.Account) {
	require.Equal(t, expected.GetSymbol(), actual.GetSymbol())
	require.Equal(t, expected.GetAddress(), actual.GetAddress())
	require.Equal(t, expected.GetBalance(), actual.GetBalance())
}

func TestKeeper_SetAccount(t *testing.T) {
	ctx := cacheKeeper()
	t.Log("Set Account")
	expected := types.NewBaseAccountWithAddress(defaultSymbol, addr1)
	{
		require.NoError(t, keeper.SetAccount(ctx, expected))
	}
	t.Log("Compare Account")
	{
		store := ctx.KVStore(keeper.storeKey)
		bz := store.Get(types.AccountKey(expected.GetSymbol(), addr1))
		actual := keeper.mustDecodeAccount(bz)
		verifyAccountFunc(t, expected, actual)
	}
}

func TestKeeper_GetAccount(t *testing.T) {
	ctx := cacheKeeper()
	t.Log("Prepare Account")
	expected := types.NewBaseAccountWithAddress(defaultSymbol, addr1)
	{
		store := ctx.KVStore(keeper.storeKey)
		store.Set(types.AccountKey(expected.GetSymbol(), addr1), keeper.cdc.MustMarshalBinaryBare(expected))
	}
	t.Log("Get Account")
	{
		actual, err := keeper.GetAccount(ctx, defaultSymbol, addr1)
		require.NoError(t, err)
		verifyAccountFunc(t, expected, actual)
	}
}

func TestKeeper_UpdateAccount(t *testing.T) {
	ctx := cacheKeeper()
	t.Log("Set Account")
	acc := types.NewBaseAccountWithAddress(defaultSymbol, addr1)
	{
		require.NoError(t, keeper.SetAccount(ctx, acc))
	}
	t.Log("Update Account")
	var expected types.Account
	expected = types.NewBaseAccountWithAddress(defaultSymbol, addr1)
	expected = expected.SetBalance(sdk.OneInt())
	{
		require.NoError(t, keeper.UpdateAccount(ctx, expected))
	}
	t.Log("Compare Account")
	{
		store := ctx.KVStore(keeper.storeKey)
		bz := store.Get(types.AccountKey(acc.GetSymbol(), addr1))
		actual := keeper.mustDecodeAccount(bz)
		verifyAccountFunc(t, expected, actual)
	}
}

func TestKeeper_GetOrNewAccount(t *testing.T) {
	ctx := cacheKeeper()
	t.Log("Prepare Account")
	expected := types.NewBaseAccountWithAddress(defaultSymbol, addr1)
	{
		store := ctx.KVStore(keeper.storeKey)
		store.Set(types.AccountKey(expected.GetSymbol(), addr1), keeper.cdc.MustMarshalBinaryBare(expected))
	}
	t.Log("Get Account addr1")
	{
		actual, err := keeper.GetOrNewAccount(ctx, defaultSymbol, addr1)
		require.NoError(t, err)
		verifyAccountFunc(t, expected, actual)
	}

	expected = types.NewBaseAccountWithAddress(defaultSymbol, addr2)
	t.Log("Get Account addr2")
	{
		actual, err := keeper.GetOrNewAccount(ctx, defaultSymbol, addr2)
		require.NoError(t, err)
		verifyAccountFunc(t, expected, actual)
	}
}
