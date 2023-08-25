package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stc-community/iot-depin-protocol/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateEventPb_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateEventPb
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateEventPb{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateEventPb{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateEventPb_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateEventPb
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateEventPb{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateEventPb{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteEventPb_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteEventPb
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteEventPb{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteEventPb{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
