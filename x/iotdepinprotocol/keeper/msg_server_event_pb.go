package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func (k msgServer) CreateEventPb(goCtx context.Context, msg *types.MsgCreateEventPb) (*types.MsgCreateEventPbResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var eventPb = types.EventPb{
		Creator: msg.Creator,
		Topic:   msg.Topic,
		Payload: msg.Payload,
	}

	id := k.AppendEventPb(
		ctx,
		eventPb,
	)
	eventPb.Id = id
	err := ctx.EventManager().EmitTypedEvent(&eventPb)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateEventPbResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateEventPb(goCtx context.Context, msg *types.MsgUpdateEventPb) (*types.MsgUpdateEventPbResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var eventPb = types.EventPb{
		Creator: msg.Creator,
		Id:      msg.Id,
		Topic:   msg.Topic,
		Payload: msg.Payload,
	}

	// Checks that the element exists
	val, found := k.GetEventPb(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// 只有设备管理者可以修改
	if msg.Creator != eventPb.Topic {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	// 还原创建者
	eventPb.Creator = val.Creator

	//// Checks if the msg creator is the same as the current owner
	//if msg.Creator != val.Creator {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	//}

	k.SetEventPb(ctx, eventPb)

	return &types.MsgUpdateEventPbResponse{}, nil
}

func (k msgServer) DeleteEventPb(goCtx context.Context, msg *types.MsgDeleteEventPb) (*types.MsgDeleteEventPbResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetEventPb(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveEventPb(ctx, msg.Id)

	return &types.MsgDeleteEventPbResponse{}, nil
}
