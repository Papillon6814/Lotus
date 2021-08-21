package cli

import (
	"github.com/spf13/cobra"
	"strconv"
	"crypto/sha256"
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/papillon6814/scavenge/x/scavenge/types"
)

var _ = strconv.Itoa(0)

func CmdSubmitScavenge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-scavenge [solution] [description] [reward]",
		Short: "Broadcast message submit-scavenge",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			solutionHash := sha256.Sum256([]byte(args[0]))
			solutionHashString := hex.EncodeToString(solutionHash[:])
			argsDescription := string(args[1])
			argsReward := string(args[2])

			// submit a new message
			msg := types.NewMsgSubmitScavenge(clientCtx.GetFromAddress().String(), string(solutionHashString), string(argsDescription), string(argsReward))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// トランザクション付きメッセージをブロックチェーンに向かってブロードキャストする
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
