package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func (k msgServer) OracleOperator(goCtx context.Context, msg *types.MsgOracleOperator) (*types.MsgOracleOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 事件发布
	ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgOracleOperatorResponse{}, nil
}
