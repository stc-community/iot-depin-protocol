package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/stc-community/iot-depin-protocol/testutil/keeper"
	"github.com/stc-community/iot-depin-protocol/testutil/nullify"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func TestEventPbQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEventPb(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetEventPbRequest
		response *types.QueryGetEventPbResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetEventPbRequest{Id: msgs[0].Id},
			response: &types.QueryGetEventPbResponse{EventPb: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetEventPbRequest{Id: msgs[1].Id},
			response: &types.QueryGetEventPbResponse{EventPb: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetEventPbRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.EventPb(wctx, tc.request)
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

func TestEventPbQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEventPb(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllEventPbRequest {
		return &types.QueryAllEventPbRequest{
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
			resp, err := keeper.EventPbAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.EventPb), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.EventPb),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EventPbAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.EventPb), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.EventPb),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.EventPbAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.EventPb),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.EventPbAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
