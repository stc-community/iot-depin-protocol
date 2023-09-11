package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/stc-community/iot-depin-protocol/testutil/keeper"
	"github.com/stc-community/iot-depin-protocol/testutil/nullify"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDeviceRegistryQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDeviceRegistry(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDeviceRegistryRequest
		response *types.QueryGetDeviceRegistryResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDeviceRegistryRequest{
				Mid: msgs[0].Mid,
			},
			response: &types.QueryGetDeviceRegistryResponse{DeviceRegistry: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDeviceRegistryRequest{
				Mid: msgs[1].Mid,
			},
			response: &types.QueryGetDeviceRegistryResponse{DeviceRegistry: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDeviceRegistryRequest{
				Mid: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.DeviceRegistry(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestDeviceRegistryQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDeviceRegistry(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDeviceRegistryRequest {
		return &types.QueryAllDeviceRegistryRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DeviceRegistryAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.DeviceRegistry), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.DeviceRegistry),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DeviceRegistryAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.DeviceRegistry), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.DeviceRegistry),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.DeviceRegistryAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.DeviceRegistry),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DeviceRegistryAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
