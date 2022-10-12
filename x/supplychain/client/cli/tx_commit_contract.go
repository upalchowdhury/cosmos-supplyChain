package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

var _ = strconv.Itoa(0)

func CmdCommitContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-contract [deal-id] [contract-id] [vendoreta] [contract-status]",
		Short: "Broadcast message commitContract",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argDealId := args[0]
             argContractId := args[1]
             argVendoreta := args[2]
             argContractStatus := args[3]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCommitContract(
				clientCtx.GetFromAddress().String(),
				argDealId,
				argContractId,
				argVendoreta,
				argContractStatus,
				
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