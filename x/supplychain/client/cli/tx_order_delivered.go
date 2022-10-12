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

func CmdOrderDelivered() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "order-delivered [deal-id] [contract-id]",
		Short: "Broadcast message orderDelivered",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argDealId := args[0]
             argContractId := args[1]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOrderDelivered(
				clientCtx.GetFromAddress().String(),
				argDealId,
				argContractId,
				
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