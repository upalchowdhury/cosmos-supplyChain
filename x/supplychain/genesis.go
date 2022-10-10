package supplychain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/upalchowdhury/supplychain/x/supplychain/keeper"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.ConCounter != nil {
		k.SetConCounter(ctx, *genState.ConCounter)
	}
	// Set all the actorscontract
	for _, elem := range genState.ActorscontractList {
		k.SetActorscontract(ctx, elem)
	}
	// Set if defined
	if genState.ContractCounter != nil {
		k.SetContractCounter(ctx, *genState.ContractCounter)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all conCounter
	conCounter, found := k.GetConCounter(ctx)
	if found {
		genesis.ConCounter = &conCounter
	}
	genesis.ActorscontractList = k.GetAllActorscontract(ctx)
	// Get all contractCounter
	contractCounter, found := k.GetContractCounter(ctx)
	if found {
		genesis.ContractCounter = &contractCounter
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
