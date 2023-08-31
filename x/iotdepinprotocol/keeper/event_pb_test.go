package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stc-community/iot-depin-protocol/testutil/keeper"
	"github.com/stc-community/iot-depin-protocol/testutil/nullify"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/keeper"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
	"github.com/stretchr/testify/require"
)

func createNEventPb(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EventPb {
	items := make([]types.EventPb, n)
	for i := range items {
		items[i].Id = keeper.AppendEventPb(ctx, items[i])
	}
	return items
}

func TestEventPbGet(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNEventPb(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEventPb(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEventPbRemove(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNEventPb(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEventPb(ctx, item.Id)
		_, found := keeper.GetEventPb(ctx, item.Id)
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

func TestEventPbCount(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNEventPb(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEventPbCount(ctx))
}
