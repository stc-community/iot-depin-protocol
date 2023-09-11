package types_test

import (
	"testing"

	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				DeviceList: []types.Device{
					{
						DeviceName: "0",
					},
					{
						DeviceName: "1",
					},
				},
				KvList: []types.Kv{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				EventPbList: []types.EventPb{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				DeviceRegistryList: []types.DeviceRegistry{
					{
						Mid: "0",
					},
					{
						Mid: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated device",
			genState: &types.GenesisState{
				DeviceList: []types.Device{
					{
						DeviceName: "0",
					},
					{
						DeviceName: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated kv",
			genState: &types.GenesisState{
				KvList: []types.Kv{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated eventPb",
			genState: &types.GenesisState{
				EventPbList: []types.EventPb{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated deviceRegistry",
			genState: &types.GenesisState{
				DeviceRegistryList: []types.DeviceRegistry{
					{
						Mid: "0",
					},
					{
						Mid: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
