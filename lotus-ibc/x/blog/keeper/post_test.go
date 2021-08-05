package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/papillon6814/lotus-ibc/x/blog/types"
)

func createTestPost(keeper *Keeper, ctx sdk.Context) types.Post {
	item := types.Post{
		Creator: "any",
	}
	keeper.SetPost(ctx, item)
	return item
}

func TestPostGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	item := createTestPost(keeper, ctx)
	rst, found := keeper.GetPost(ctx)
	assert.True(t, found)
	assert.Equal(t, item, rst)
}
func TestPostRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	createTestPost(keeper, ctx)
	keeper.RemovePost(ctx)
	_, found := keeper.GetPost(ctx)
	assert.False(t, found)
}
