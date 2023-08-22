package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stc-community/iot-depin-protocol/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateDevice_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDevice
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDevice{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateDevice{
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

func TestMsgUpdateDevice_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateDevice
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateDevice{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateDevice{
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

func TestMsgDeleteDevice_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteDevice
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteDevice{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteDevice{
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
