package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func CmdCreateKv() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-kv [index] [value] [device-name]",
		Short: "Create a new kv",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argValue := args[1]
			argDeviceName := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateKv(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argValue,
				argDeviceName,
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

func CmdUpdateKv() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-kv [index] [value] [device-name]",
		Short: "Update a kv",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argValue := args[1]
			argDeviceName := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateKv(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argValue,
				argDeviceName,
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

func CmdDeleteKv() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-kv [index] [device-name]",
		Short: "Delete a kv",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]
			argDeviceName := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteKv(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argDeviceName,
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
