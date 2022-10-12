package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

var _ = strconv.Itoa(0)

func CmdCreateTran() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-tran [vendor] [commission]",
		Short: "Broadcast message createTran",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVendor := args[0]
			argCommission := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTran(
				clientCtx.GetFromAddress().String(),
				argVendor,
				argCommission,
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
