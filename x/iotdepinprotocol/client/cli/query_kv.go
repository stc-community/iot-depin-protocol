package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func CmdListKv() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-kv",
		Short: "list all kv",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			from, _ := cmd.Flags().GetString(flags.FlagFrom)

			params := &types.QueryAllKvRequest{
				Pagination: pageReq,
				Creator:    from,
			}
			res, err := queryClient.KvAll(context.Background(), params)
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

func CmdShowKv() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-kv [index]",
		Short: "shows a kv",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]
			from, _ := cmd.Flags().GetString(flags.FlagFrom)

			params := &types.QueryGetKvRequest{
				Index:   argIndex,
				Creator: from,
			}

			res, err := queryClient.Kv(context.Background(), params)
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
