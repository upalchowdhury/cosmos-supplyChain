package supplychain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/upalchowdhury/supplychain/testutil/keeper"
	"github.com/upalchowdhury/supplychain/testutil/nullify"
	"github.com/upalchowdhury/supplychain/x/supplychain"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ConCounter: &types.ConCounter{
			IdValue: 51,
		},
		ActorscontractList: []types.Actorscontract{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ContractCounter: &types.ContractCounter{
			DealId:  "49",
			IdValue: 42,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SupplychainKeeper(t)
	supplychain.InitGenesis(ctx, *k, genesisState)
	got := supplychain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.ConCounter, got.ConCounter)
	require.ElementsMatch(t, genesisState.ActorscontractList, got.ActorscontractList)
	require.Equal(t, genesisState.ContractCounter, got.ContractCounter)
	// this line is used by starport scaffolding # genesis/test/assert
}
