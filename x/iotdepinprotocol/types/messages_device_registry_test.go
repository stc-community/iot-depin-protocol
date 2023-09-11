package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stc-community/iot-depin-protocol/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateDeviceRegistry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDeviceRegistry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDeviceRegistry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateDeviceRegistry{
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

func TestMsgUpdateDeviceRegistry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateDeviceRegistry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateDeviceRegistry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateDeviceRegistry{
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

func TestMsgDeleteDeviceRegistry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteDeviceRegistry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteDeviceRegistry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteDeviceRegistry{
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
