package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/papillon6814/lotus-ibc/x/blog/keeper"
	"github.com/papillon6814/lotus-ibc/x/blog/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set if defined
	if genState.TimedoutPost != nil {
		k.SetTimedoutPost(ctx, *genState.TimedoutPost)
	}

	// Set if defined
	if genState.SendPost != nil {
		k.SetSendPost(ctx, *genState.SendPost)
	}

	// Set if defined
	if genState.Post != nil {
		k.SetPost(ctx, *genState.Post)
	}

	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all timedoutPost
	timedoutPost, found := k.GetTimedoutPost(ctx)
	if found {
		genesis.TimedoutPost = &timedoutPost
	}

	// Get all sendPost
	sendPost, found := k.GetSendPost(ctx)
	if found {
		genesis.SendPost = &sendPost
	}

	// Get all post
	post, found := k.GetPost(ctx)
	if found {
		genesis.Post = &post
	}

	genesis.PortId = k.GetPort(ctx)

	return genesis
}
