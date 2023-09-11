package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/stc-community/iot-depin-protocol/testutil/keeper"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/keeper"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDeviceRegistryMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.IotdepinprotocolKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDeviceRegistry{Creator: creator,
			Mid: strconv.Itoa(i),
		}
		_, err := srv.CreateDeviceRegistry(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetDeviceRegistry(ctx,
			expected.Mid,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestDeviceRegistryMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDeviceRegistry
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDeviceRegistry{Creator: creator,
				Mid: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDeviceRegistry{Creator: "B",
				Mid: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateDeviceRegistry{Creator: creator,
				Mid: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IotdepinprotocolKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateDeviceRegistry{Creator: creator,
				Mid: strconv.Itoa(0),
			}
			_, err := srv.CreateDeviceRegistry(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateDeviceRegistry(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetDeviceRegistry(ctx,
					expected.Mid,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestDeviceRegistryMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteDeviceRegistry
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteDeviceRegistry{Creator: creator,
				Mid: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteDeviceRegistry{Creator: "B",
				Mid: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteDeviceRegistry{Creator: creator,
				Mid: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IotdepinprotocolKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateDeviceRegistry(wctx, &types.MsgCreateDeviceRegistry{Creator: creator,
				Mid: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteDeviceRegistry(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetDeviceRegistry(ctx,
					tc.request.Mid,
				)
				require.False(t, found)
			}
		})
	}
}
