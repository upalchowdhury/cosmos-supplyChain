package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/upalchowdhury/supplychain/testutil/keeper"
	"github.com/upalchowdhury/supplychain/testutil/nullify"
	"github.com/upalchowdhury/supplychain/x/supplychain/keeper"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNewcontract(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Newcontract {
	items := make([]types.Newcontract, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetNewcontract(ctx, items[i])
	}
	return items
}

func TestNewcontractGet(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	items := createNNewcontract(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNewcontract(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNewcontractRemove(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	items := createNNewcontract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNewcontract(ctx,
			item.Index,
		)
		_, found := keeper.GetNewcontract(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestNewcontractGetAll(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	items := createNNewcontract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNewcontract(ctx)),
	)
}
