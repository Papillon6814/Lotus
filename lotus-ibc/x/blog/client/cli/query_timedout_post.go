package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/papillon6814/lotus-ibc/x/blog/types"
	"github.com/spf13/cobra"
)

func CmdShowTimedoutPost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-timedout-post",
		Short: "shows timedoutPost",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetTimedoutPostRequest{}

			res, err := queryClient.TimedoutPost(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
