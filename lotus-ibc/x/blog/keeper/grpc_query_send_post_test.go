package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/papillon6814/lotus-ibc/x/blog/types"
)

func TestSendPostQuery(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestSendPost(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSendPostRequest
		response *types.QueryGetSendPostResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetSendPostRequest{},
			response: &types.QueryGetSendPostResponse{SendPost: &item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SendPost(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
