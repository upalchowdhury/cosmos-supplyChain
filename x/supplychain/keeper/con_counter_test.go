package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/upalchowdhury/supplychain/testutil/keeper"
	"github.com/upalchowdhury/supplychain/testutil/nullify"
	"github.com/upalchowdhury/supplychain/x/supplychain/keeper"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

func createTestConCounter(keeper *keeper.Keeper, ctx sdk.Context) types.ConCounter {
	item := types.ConCounter{}
	keeper.SetConCounter(ctx, item)
	return item
}

func TestConCounterGet(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	item := createTestConCounter(keeper, ctx)
	rst, found := keeper.GetConCounter(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestConCounterRemove(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	createTestConCounter(keeper, ctx)
	keeper.RemoveConCounter(ctx)
	_, found := keeper.GetConCounter(ctx)
	require.False(t, found)
}
