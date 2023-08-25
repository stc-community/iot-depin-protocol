package iotdepinprotocol

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stc-community/iot-depin-protocol/testutil/sample"
	iotdepinprotocolsimulation "github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/simulation"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = iotdepinprotocolsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateKv = "op_weight_msg_kv"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKv int = 100

	opWeightMsgUpdateKv = "op_weight_msg_kv"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKv int = 100

	opWeightMsgDeleteKv = "op_weight_msg_kv"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKv int = 100

	opWeightMsgCreateDevice = "op_weight_msg_device"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDevice int = 100

	opWeightMsgUpdateDevice = "op_weight_msg_device"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDevice int = 100

	opWeightMsgDeleteDevice = "op_weight_msg_device"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDevice int = 100

	opWeightMsgOracleOperator = "op_weight_msg_oracle_operator"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOracleOperator int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	iotdepinprotocolGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		KvList: []types.Kv{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		DeviceList: []types.Device{
			{
				Creator: sample.AccAddress(),
				Address: "0",
			},
			{
				Creator: sample.AccAddress(),
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&iotdepinprotocolGenesis)
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

	var weightMsgCreateKv int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKv, &weightMsgCreateKv, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKv = defaultWeightMsgCreateKv
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKv,
		iotdepinprotocolsimulation.SimulateMsgCreateKv(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateKv int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateKv, &weightMsgUpdateKv, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateKv = defaultWeightMsgUpdateKv
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateKv,
		iotdepinprotocolsimulation.SimulateMsgUpdateKv(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteKv int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteKv, &weightMsgDeleteKv, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteKv = defaultWeightMsgDeleteKv
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteKv,
		iotdepinprotocolsimulation.SimulateMsgDeleteKv(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateDevice int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDevice, &weightMsgCreateDevice, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDevice = defaultWeightMsgCreateDevice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDevice,
		iotdepinprotocolsimulation.SimulateMsgCreateDevice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDevice int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDevice, &weightMsgUpdateDevice, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDevice = defaultWeightMsgUpdateDevice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDevice,
		iotdepinprotocolsimulation.SimulateMsgUpdateDevice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDevice int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDevice, &weightMsgDeleteDevice, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDevice = defaultWeightMsgDeleteDevice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDevice,
		iotdepinprotocolsimulation.SimulateMsgDeleteDevice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgOracleOperator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOracleOperator, &weightMsgOracleOperator, nil,
		func(_ *rand.Rand) {
			weightMsgOracleOperator = defaultWeightMsgOracleOperator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOracleOperator,
		iotdepinprotocolsimulation.SimulateMsgOracleOperator(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}