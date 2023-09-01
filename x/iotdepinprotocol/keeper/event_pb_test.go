package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stc-community/iot-depin-protocol/testutil/keeper"
	"github.com/stc-community/iot-depin-protocol/testutil/nullify"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/keeper"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNEventPb(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EventPb {
	items := make([]types.EventPb, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEventPb(ctx, items[i])
	}
	return items
}

func TestEventPbGet(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNEventPb(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEventPb(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEventPbRemove(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNEventPb(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEventPb(ctx,
			item.Index,
		)
		_, found := keeper.GetEventPb(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEventPbGetAll(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNEventPb(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEventPb(ctx)),
	)
}
