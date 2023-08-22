package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKv = "create_kv"
	TypeMsgUpdateKv = "update_kv"
	TypeMsgDeleteKv = "delete_kv"
)

var _ sdk.Msg = &MsgCreateKv{}

func NewMsgCreateKv(
	creator string,
	index string,
	value string,

) *MsgCreateKv {
	return &MsgCreateKv{
		Creator: creator,
		Index:   index,
		Value:   value,
	}
}

func (msg *MsgCreateKv) Route() string {
	return RouterKey
}

func (msg *MsgCreateKv) Type() string {
	return TypeMsgCreateKv
}

func (msg *MsgCreateKv) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKv) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKv) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateKv{}

func NewMsgUpdateKv(
	creator string,
	index string,
	value string,

) *MsgUpdateKv {
	return &MsgUpdateKv{
		Creator: creator,
		Index:   index,
		Value:   value,
	}
}

func (msg *MsgUpdateKv) Route() string {
	return RouterKey
}

func (msg *MsgUpdateKv) Type() string {
	return TypeMsgUpdateKv
}

func (msg *MsgUpdateKv) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateKv) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateKv) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteKv{}

func NewMsgDeleteKv(
	creator string,
	index string,

) *MsgDeleteKv {
	return &MsgDeleteKv{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteKv) Route() string {
	return RouterKey
}

func (msg *MsgDeleteKv) Type() string {
	return TypeMsgDeleteKv
}

func (msg *MsgDeleteKv) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteKv) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteKv) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
