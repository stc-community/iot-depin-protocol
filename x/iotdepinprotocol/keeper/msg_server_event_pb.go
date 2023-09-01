package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func (k msgServer) CreateEventPb(goCtx context.Context, msg *types.MsgCreateEventPb) (*types.MsgCreateEventPbResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// 自动生成 Index
	if msg.Index == "" {
		//msg.Index = hex.EncodeToString(tmhash.Sum(ctx.TxBytes()))
		msg.Index = ctx.HeaderHash().String()
	}
	// Check if the value already exists
	_, isFound := k.GetEventPb(
		ctx,
		msg.Index,
		msg.DeviceName,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var eventPb = types.EventPb{
		Creator:    msg.Creator,
		Index:      msg.Index,
		DeviceName: msg.DeviceName,
		Payload:    msg.Payload,
	}

	k.SetEventPb(
		ctx,
		eventPb,
	)
	// publish 事件
	err := ctx.EventManager().EmitTypedEvent(&eventPb)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateEventPbResponse{}, nil
}

func (k msgServer) UpdateEventPb(goCtx context.Context, msg *types.MsgUpdateEventPb) (*types.MsgUpdateEventPbResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetEventPb(
		ctx,
		msg.Index,
		msg.DeviceName,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	//if msg.Creator != valFound.Creator {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	//}

	var eventPb = types.EventPb{
		Creator:    valFound.Creator,
		Index:      msg.Index,
		DeviceName: msg.DeviceName,
		Payload:    msg.Payload,
	}

	k.SetEventPb(ctx, eventPb)

	return &types.MsgUpdateEventPbResponse{}, nil
}

func (k msgServer) DeleteEventPb(goCtx context.Context, msg *types.MsgDeleteEventPb) (*types.MsgDeleteEventPbResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetEventPb(
		ctx,
		msg.Index,
		msg.DeviceName,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveEventPb(
		ctx,
		msg.Index,
		msg.DeviceName,
	)

	return &types.MsgDeleteEventPbResponse{}, nil
}
