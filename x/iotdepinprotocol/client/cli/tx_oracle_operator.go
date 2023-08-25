package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

var _ = strconv.Itoa(0)

func CmdOracleOperator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracle-operator [req-id] [operator-type] [data]",
		Short: "Broadcast message OracleOperator",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReqId := args[0]
			argOperatorType := args[1]
			argData := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOracleOperator(
				clientCtx.GetFromAddress().String(),
				argReqId,
				argOperatorType,
				argData,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
