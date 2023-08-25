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

func TestEventPbMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.IotdepinprotocolKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateEventPb{Creator: creator,
			PubId: strconv.Itoa(i),
		}
		_, err := srv.CreateEventPb(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetEventPb(ctx,
			expected.PubId,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestEventPbMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateEventPb
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateEventPb{Creator: creator,
				PubId: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateEventPb{Creator: "B",
				PubId: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateEventPb{Creator: creator,
				PubId: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IotdepinprotocolKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateEventPb{Creator: creator,
				PubId: strconv.Itoa(0),
			}
			_, err := srv.CreateEventPb(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateEventPb(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetEventPb(ctx,
					expected.PubId,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestEventPbMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteEventPb
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteEventPb{Creator: creator,
				PubId: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteEventPb{Creator: "B",
				PubId: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteEventPb{Creator: creator,
				PubId: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IotdepinprotocolKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateEventPb(wctx, &types.MsgCreateEventPb{Creator: creator,
				PubId: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteEventPb(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetEventPb(ctx,
					tc.request.PubId,
				)
				require.False(t, found)
			}
		})
	}
}
