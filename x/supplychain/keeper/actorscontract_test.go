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

func createNActorscontract(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Actorscontract {
	items := make([]types.Actorscontract, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetActorscontract(ctx, items[i])
	}
	return items
}

func TestActorscontractGet(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	items := createNActorscontract(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActorscontract(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestActorscontractRemove(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	items := createNActorscontract(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActorscontract(ctx,
			item.Index,
		)
		_, found := keeper.GetActorscontract(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestActorscontractGetAll(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	items := createNActorscontract(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActorscontract(ctx)),
	)
}
