package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) KvAll(goCtx context.Context, req *types.QueryAllKvRequest) (*types.QueryAllKvResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var kvs []types.Kv
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	kvStore := prefix.NewStore(store, types.KeyPrefix(types.KvKeyPrefix+req.DeviceName))

	pageRes, err := query.Paginate(kvStore, req.Pagination, func(key []byte, value []byte) error {
		var kv types.Kv
		if err := k.cdc.Unmarshal(value, &kv); err != nil {
			return err
		}

		kvs = append(kvs, kv)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKvResponse{Kv: kvs, Pagination: pageRes}, nil
}

func (k Keeper) Kv(goCtx context.Context, req *types.QueryGetKvRequest) (*types.QueryGetKvResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetKv(
		ctx,
		req.Index,
		req.DeviceName,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetKvResponse{Kv: val}, nil
}
