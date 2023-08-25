package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/keeper"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func SimulateMsgOracleOperator(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgOracleOperator{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the OracleOperator simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "OracleOperator simulation not implemented"), nil, nil
	}
}
