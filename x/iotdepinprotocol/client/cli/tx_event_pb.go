package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func CmdCreateEventPb() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-event-pb [index] [device-name] [payload]",
		Short: "Create a new eventPb",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argDeviceName := args[1]
			argPayload := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateEventPb(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argDeviceName,
				argPayload,
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

func CmdUpdateEventPb() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-event-pb [index] [device-name] [payload]",
		Short: "Update a eventPb",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argDeviceName := args[1]
			argPayload := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateEventPb(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argDeviceName,
				argPayload,
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

func CmdDeleteEventPb() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-event-pb [index] [device-name]",
		Short: "Delete a eventPb",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]
			argDeviceName := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteEventPb(
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
