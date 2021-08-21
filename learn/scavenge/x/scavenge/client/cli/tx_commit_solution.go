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

func CmdCommitSolution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-solution [solution]",
		Short: "Broadcast message commit-solution",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			solution := args[0]
			solutionHash := sha256.Sum256([]byte(solution))
			solutionHashString := hex.EncodeToString(solutionHash[:])
			// scavengerのアドレスをstriんｇに変換する
			var scavenger = clientCtx.GetFromAddress().String()
			var solutionScavengerHash = sha256.Sum256([]byte(solution + scavenger))

			var solutionScavengerHashString = hex.EncodeToString(solutionScavengerHash[:])

			// Messageを作成する
			msg := types.NewMsgCommitSolution(clientCtx.GetFromAddress().String(), string(solutionHashString), string(solutionScavengerHashString))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// ブロックチェーンに向かってメッセージ付きのトランザクションをブロードキャストする
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
