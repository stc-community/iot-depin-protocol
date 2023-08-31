package iotdepinprotocol_test

import (
	"testing"

	keepertest "github.com/stc-community/iot-depin-protocol/testutil/keeper"
	"github.com/stc-community/iot-depin-protocol/testutil/nullify"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		KvList: []types.Kv{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		DeviceList: []types.Device{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		EventPbList: []types.EventPb{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		EventPbCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IotdepinprotocolKeeper(t)
	iotdepinprotocol.InitGenesis(ctx, *k, genesisState)
	got := iotdepinprotocol.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.KvList, got.KvList)
	require.ElementsMatch(t, genesisState.DeviceList, got.DeviceList)
	require.ElementsMatch(t, genesisState.EventPbList, got.EventPbList)
	require.Equal(t, genesisState.EventPbCount, got.EventPbCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
