package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/papillon6814/lotus-ibc/x/blog/types"
)

func createTestSendPost(keeper *Keeper, ctx sdk.Context) types.SendPost {
	item := types.SendPost{
		Creator: "any",
	}
	keeper.SetSendPost(ctx, item)
	return item
}

func TestSendPostGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	item := createTestSendPost(keeper, ctx)
	rst, found := keeper.GetSendPost(ctx)
	assert.True(t, found)
	assert.Equal(t, item, rst)
}
func TestSendPostRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	createTestSendPost(keeper, ctx)
	keeper.RemoveSendPost(ctx)
	_, found := keeper.GetSendPost(ctx)
	assert.False(t, found)
}
