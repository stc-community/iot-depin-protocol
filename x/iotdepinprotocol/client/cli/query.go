package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group iotdepinprotocol queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListDevice())
	cmd.AddCommand(CmdShowDevice())
	cmd.AddCommand(CmdListKv())
	cmd.AddCommand(CmdShowKv())
	cmd.AddCommand(CmdListEventPb())
	cmd.AddCommand(CmdShowEventPb())
	cmd.AddCommand(CmdListDeviceRegistry())
	cmd.AddCommand(CmdShowDeviceRegistry())
	// this line is used by starport scaffolding # 1

	return cmd
}
