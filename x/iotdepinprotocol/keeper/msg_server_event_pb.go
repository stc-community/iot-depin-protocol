package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func (k msgServer) CreateEventPb(goCtx context.Context, msg *types.MsgCreateEventPb) (*types.MsgCreateEventPbResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetEventPb(
		ctx,
		msg.PubId,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var eventPb = types.EventPb{
		Creator: msg.Creator,
		PubId:   msg.PubId,
		Topic:   msg.Topic,
		PubType: msg.PubType,
		Payload: msg.Payload,
		PubTime: msg.PubTime,
	}
	// Publish 事件
	ctx.EventManager().EmitTypedEvent(&eventPb)

	k.SetEventPb(
		ctx,
		eventPb,
	)
	return &types.MsgCreateEventPbResponse{}, nil
}

func (k msgServer) UpdateEventPb(goCtx context.Context, msg *types.MsgUpdateEventPb) (*types.MsgUpdateEventPbResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetEventPb(
		ctx,
		msg.PubId,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var eventPb = types.EventPb{
		Creator: msg.Creator,
		PubId:   msg.PubId,
		Topic:   msg.Topic,
		PubType: msg.PubType,
		Payload: msg.Payload,
		PubTime: msg.PubTime,
	}

	k.SetEventPb(ctx, eventPb)

	return &types.MsgUpdateEventPbResponse{}, nil
}

func (k msgServer) DeleteEventPb(goCtx context.Context, msg *types.MsgDeleteEventPb) (*types.MsgDeleteEventPbResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetEventPb(
		ctx,
		msg.PubId,
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
		msg.PubId,
	)

	return &types.MsgDeleteEventPbResponse{}, nil
}
