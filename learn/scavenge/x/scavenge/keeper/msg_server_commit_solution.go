package keeper

import (
	"context"

	"github.com/papillon6814/scavenge/x/scavenge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CommitSolution(goCtx context.Context, msg *types.MsgCommitSolution) (*types.MsgCommitSolutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var commit = types.Commit{
		Index: msg.SolutionScavengerHash,
		Creator: msg.Creator,
		SolutionHash: msg.SolutionHash,
		SolutionScavengerHash: msg.SolutionScavengerHash,
	}

	// solution + scavengerのhashをkeyとして使用し、commitをstoreから取得する
	// 見つかったらダメ
	_, isFound := k.GetCommit(ctx, commit.SolutionScavengerHash)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Commit with that hash already exists")
	}

	// Write commit to the store
	k.SetCommit(ctx, commit)

	return &types.MsgCommitSolutionResponse{}, nil
}
