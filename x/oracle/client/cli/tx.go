package cli

import (
	"strconv"

	"github.com/Carina-hackatom/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(NewUpdateStateCmd())

	return cmd
}

func NewUpdateStateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update_state [from_key_or_address] [amount] [decimal] [block_height] [app_hash] [chain_id]",
		Short: "",
		Long:  "",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Set(flags.FlagFrom, args[0]); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			decimal, err := strconv.ParseUint(args[2], 10, 32)
			if err != nil {
				return err
			}

			blockHeight, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return err
			}

			appHash := []byte(args[4])
			chainId := args[5]

			msg := types.NewMsgUpdateChainState(clientCtx.GetFromAddress(), chainId, amount, uint32(decimal), blockHeight, appHash)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
