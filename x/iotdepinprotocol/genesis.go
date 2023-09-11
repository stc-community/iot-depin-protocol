package iotdepinprotocol

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/keeper"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the device
	for _, elem := range genState.DeviceList {
		k.SetDevice(ctx, elem)
	}
	// Set all the kv
	for _, elem := range genState.KvList {
		k.SetKv(ctx, elem)
	}
	// Set all the eventPb
	for _, elem := range genState.EventPbList {
		k.SetEventPb(ctx, elem)
	}
	// Set all the deviceRegistry
	for _, elem := range genState.DeviceRegistryList {
		k.SetDeviceRegistry(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DeviceList = k.GetAllDevice(ctx)
	genesis.KvList = k.GetAllKv(ctx)
	genesis.EventPbList = k.GetAllEventPb(ctx)
	genesis.DeviceRegistryList = k.GetAllDeviceRegistry(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
