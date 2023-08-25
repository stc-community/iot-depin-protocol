package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func CmdListDevice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-device",
		Short: "list all device",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			from, _ := cmd.Flags().GetString(flags.FlagFrom)

			params := &types.QueryAllDeviceRequest{
				Pagination: pageReq,
				Creator:    from,
			}

			res, err := queryClient.DeviceAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String(flags.FlagFrom, "", "For query address")
	return cmd
}

func CmdShowDevice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-device [address]",
		Short: "shows a device",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]

			from, _ := cmd.Flags().GetString(flags.FlagFrom)

			params := &types.QueryGetDeviceRequest{
				Address: argAddress,
				Creator: from,
			}

			res, err := queryClient.Device(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String(flags.FlagFrom, "", "For query address")
	return cmd
}
