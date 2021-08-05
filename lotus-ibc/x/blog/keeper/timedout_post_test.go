package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/papillon6814/lotus-ibc/x/blog/types"
)

func createTestTimedoutPost(keeper *Keeper, ctx sdk.Context) types.TimedoutPost {
	item := types.TimedoutPost{
		Creator: "any",
	}
	keeper.SetTimedoutPost(ctx, item)
	return item
}

func TestTimedoutPostGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	item := createTestTimedoutPost(keeper, ctx)
	rst, found := keeper.GetTimedoutPost(ctx)
	assert.True(t, found)
	assert.Equal(t, item, rst)
}
func TestTimedoutPostRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	createTestTimedoutPost(keeper, ctx)
	keeper.RemoveTimedoutPost(ctx)
	_, found := keeper.GetTimedoutPost(ctx)
	assert.False(t, found)
}
