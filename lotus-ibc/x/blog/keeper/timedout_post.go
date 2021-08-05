package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/papillon6814/lotus-ibc/x/blog/types"
)

// SetTimedoutPost set timedoutPost in the store
func (k Keeper) SetTimedoutPost(ctx sdk.Context, timedoutPost types.TimedoutPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutPostKey))
	b := k.cdc.MustMarshalBinaryBare(&timedoutPost)
	store.Set([]byte{0}, b)
}

// GetTimedoutPost returns timedoutPost
func (k Keeper) GetTimedoutPost(ctx sdk.Context) (val types.TimedoutPost, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutPostKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveTimedoutPost removes timedoutPost from the store
func (k Keeper) RemoveTimedoutPost(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutPostKey))
	store.Delete([]byte{0})
}
