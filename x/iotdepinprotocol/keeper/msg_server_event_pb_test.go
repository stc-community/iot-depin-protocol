package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func TestEventPbMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateEventPb(ctx, &types.MsgCreateEventPb{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
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
			desc:    "Completed",
			request: &types.MsgUpdateEventPb{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateEventPb{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateEventPb{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateEventPb(ctx, &types.MsgCreateEventPb{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateEventPb(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
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
			desc:    "Completed",
			request: &types.MsgDeleteEventPb{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteEventPb{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteEventPb{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateEventPb(ctx, &types.MsgCreateEventPb{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteEventPb(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
