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

func createTestContractCounter(keeper *keeper.Keeper, ctx sdk.Context) types.ContractCounter {
	item := types.ContractCounter{}
	keeper.SetContractCounter(ctx, item)
	return item
}

func TestContractCounterGet(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	item := createTestContractCounter(keeper, ctx)
	rst, found := keeper.GetContractCounter(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestContractCounterRemove(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	createTestContractCounter(keeper, ctx)
	keeper.RemoveContractCounter(ctx)
	_, found := keeper.GetContractCounter(ctx)
	require.False(t, found)
}
