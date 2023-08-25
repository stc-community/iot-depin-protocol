package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateEventPb = "create_event_pb"
	TypeMsgUpdateEventPb = "update_event_pb"
	TypeMsgDeleteEventPb = "delete_event_pb"
)

var _ sdk.Msg = &MsgCreateEventPb{}

func NewMsgCreateEventPb(
	creator string,
	pubId string,
	topic string,
	pubType string,
	payload string,
	pubTime uint64,

) *MsgCreateEventPb {
	return &MsgCreateEventPb{
		Creator: creator,
		PubId:   pubId,
		Topic:   topic,
		PubType: pubType,
		Payload: payload,
		PubTime: pubTime,
	}
}

func (msg *MsgCreateEventPb) Route() string {
	return RouterKey
}

func (msg *MsgCreateEventPb) Type() string {
	return TypeMsgCreateEventPb
}

func (msg *MsgCreateEventPb) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEventPb) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEventPb) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateEventPb{}

func NewMsgUpdateEventPb(
	creator string,
	pubId string,
	topic string,
	pubType string,
	payload string,
	pubTime uint64,

) *MsgUpdateEventPb {
	return &MsgUpdateEventPb{
		Creator: creator,
		PubId:   pubId,
		Topic:   topic,
		PubType: pubType,
		Payload: payload,
		PubTime: pubTime,
	}
}

func (msg *MsgUpdateEventPb) Route() string {
	return RouterKey
}

func (msg *MsgUpdateEventPb) Type() string {
	return TypeMsgUpdateEventPb
}

func (msg *MsgUpdateEventPb) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateEventPb) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateEventPb) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteEventPb{}

func NewMsgDeleteEventPb(
	creator string,
	pubId string,

) *MsgDeleteEventPb {
	return &MsgDeleteEventPb{
		Creator: creator,
		PubId:   pubId,
	}
}
func (msg *MsgDeleteEventPb) Route() string {
	return RouterKey
}

func (msg *MsgDeleteEventPb) Type() string {
	return TypeMsgDeleteEventPb
}

func (msg *MsgDeleteEventPb) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteEventPb) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteEventPb) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
