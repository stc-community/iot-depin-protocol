package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func CmdCreateEventPb() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-event-pb [pub-id] [topic] [pub-type] [payload] [pub-time]",
		Short: "Create a new EventPb",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexPubId := args[0]

			// Get value arguments
			argTopic := args[1]
			argPubType := args[2]
			argPayload := args[3]
			argPubTime, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateEventPb(
				clientCtx.GetFromAddress().String(),
				indexPubId,
				argTopic,
				argPubType,
				argPayload,
				argPubTime,
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
		Use:   "update-event-pb [pub-id] [topic] [pub-type] [payload] [pub-time]",
		Short: "Update a EventPb",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexPubId := args[0]

			// Get value arguments
			argTopic := args[1]
			argPubType := args[2]
			argPayload := args[3]
			argPubTime, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateEventPb(
				clientCtx.GetFromAddress().String(),
				indexPubId,
				argTopic,
				argPubType,
				argPayload,
				argPubTime,
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
		Use:   "delete-event-pb [pub-id]",
		Short: "Delete a EventPb",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexPubId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteEventPb(
				clientCtx.GetFromAddress().String(),
				indexPubId,
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
