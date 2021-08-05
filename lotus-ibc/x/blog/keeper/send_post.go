package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/papillon6814/lotus-ibc/x/blog/types"
)

// SetSendPost set sendPost in the store
func (k Keeper) SetSendPost(ctx sdk.Context, sendPost types.SendPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendPostKey))
	b := k.cdc.MustMarshalBinaryBare(&sendPost)
	store.Set([]byte{0}, b)
}

// GetSendPost returns sendPost
func (k Keeper) GetSendPost(ctx sdk.Context) (val types.SendPost, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendPostKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveSendPost removes sendPost from the store
func (k Keeper) RemoveSendPost(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SendPostKey))
	store.Delete([]byte{0})
}
