package supplychain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/upalchowdhury/supplychain/testutil/sample"
	supplychainsimulation "github.com/upalchowdhury/supplychain/x/supplychain/simulation"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = supplychainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateTran = "op_weight_msg_create_tran"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTran int = 100

	opWeightMsgCreateContract = "op_weight_msg_create_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateContract int = 100

	opWeightMsgCommitContract = "op_weight_msg_commit_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCommitContract int = 100

	opWeightMsgApproveContract = "op_weight_msg_approve_contract"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveContract int = 100

	opWeightMsgShipOrder = "op_weight_msg_ship_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgShipOrder int = 100

	opWeightMsgOrderDelivered = "op_weight_msg_order_delivered"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOrderDelivered int = 100

	opWeightMsgCancelOrder = "op_weight_msg_cancel_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelOrder int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	supplychainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&supplychainGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTran int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTran, &weightMsgCreateTran, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTran = defaultWeightMsgCreateTran
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTran,
		supplychainsimulation.SimulateMsgCreateTran(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateContract, &weightMsgCreateContract, nil,
		func(_ *rand.Rand) {
			weightMsgCreateContract = defaultWeightMsgCreateContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateContract,
		supplychainsimulation.SimulateMsgCreateContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCommitContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCommitContract, &weightMsgCommitContract, nil,
		func(_ *rand.Rand) {
			weightMsgCommitContract = defaultWeightMsgCommitContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCommitContract,
		supplychainsimulation.SimulateMsgCommitContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveContract int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveContract, &weightMsgApproveContract, nil,
		func(_ *rand.Rand) {
			weightMsgApproveContract = defaultWeightMsgApproveContract
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveContract,
		supplychainsimulation.SimulateMsgApproveContract(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgShipOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgShipOrder, &weightMsgShipOrder, nil,
		func(_ *rand.Rand) {
			weightMsgShipOrder = defaultWeightMsgShipOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgShipOrder,
		supplychainsimulation.SimulateMsgShipOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOrderDelivered int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOrderDelivered, &weightMsgOrderDelivered, nil,
		func(_ *rand.Rand) {
			weightMsgOrderDelivered = defaultWeightMsgOrderDelivered
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOrderDelivered,
		supplychainsimulation.SimulateMsgOrderDelivered(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelOrder, &weightMsgCancelOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelOrder = defaultWeightMsgCancelOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelOrder,
		supplychainsimulation.SimulateMsgCancelOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
