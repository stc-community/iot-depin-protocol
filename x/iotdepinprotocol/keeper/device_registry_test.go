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

func createNDeviceRegistry(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DeviceRegistry {
	items := make([]types.DeviceRegistry, n)
	for i := range items {
		items[i].Mid = strconv.Itoa(i)

		keeper.SetDeviceRegistry(ctx, items[i])
	}
	return items
}

func TestDeviceRegistryGet(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNDeviceRegistry(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDeviceRegistry(ctx,
			item.Mid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDeviceRegistryRemove(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNDeviceRegistry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDeviceRegistry(ctx,
			item.Mid,
		)
		_, found := keeper.GetDeviceRegistry(ctx,
			item.Mid,
		)
		require.False(t, found)
	}
}

func TestDeviceRegistryGetAll(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNDeviceRegistry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDeviceRegistry(ctx)),
	)
}
