package zestchain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/cytruslabs/zestchain/testutil/sample"
	zestchainsimulation "github.com/cytruslabs/zestchain/x/zestchain/simulation"
	"github.com/cytruslabs/zestchain/x/zestchain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = zestchainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAd = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAd int = 100

	opWeightMsgPayView = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPayView int = 100

	opWeightMsgPayClick = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPayClick int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	zestchainGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&zestchainGenesis)
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

	var weightMsgCreateAd int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAd, &weightMsgCreateAd, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAd = defaultWeightMsgCreateAd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAd,
		zestchainsimulation.SimulateMsgCreateAd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPayView int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPayView, &weightMsgPayView, nil,
		func(_ *rand.Rand) {
			weightMsgPayView = defaultWeightMsgPayView
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPayView,
		zestchainsimulation.SimulateMsgPayView(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPayClick int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPayClick, &weightMsgPayClick, nil,
		func(_ *rand.Rand) {
			weightMsgPayClick = defaultWeightMsgPayClick
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPayClick,
		zestchainsimulation.SimulateMsgPayClick(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
