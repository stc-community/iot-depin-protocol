package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EventPbAll(goCtx context.Context, req *types.QueryAllEventPbRequest) (*types.QueryAllEventPbResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var eventPbs []types.EventPb
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	eventPbStore := prefix.NewStore(store, types.KeyPrefix(types.EventPbKeyPrefix))

	// 使用自定义的 Paginate 查询
	pageRes, err := types.Paginate(eventPbStore, req.Pagination, func(key []byte, value []byte) (bool, error) {
		var eventPb types.EventPb
		if err := k.cdc.Unmarshal(value, &eventPb); err != nil {
			return false, err
		}
		// Filter
		if req.PubId != "" {
			if eventPb.PubId != req.PubId {
				return false, nil
			}
		}
		if req.Topic != "" {
			if eventPb.Topic != req.Topic {
				return false, nil
			}
		}
		eventPbs = append(eventPbs, eventPb)
		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEventPbResponse{EventPb: eventPbs, Pagination: pageRes}, nil
}

func (k Keeper) EventPb(goCtx context.Context, req *types.QueryGetEventPbRequest) (*types.QueryGetEventPbResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetEventPb(
		ctx,
		req.PubId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEventPbResponse{EventPb: val}, nil
}
