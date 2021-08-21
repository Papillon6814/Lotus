package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/papillon6814/scavenge/x/scavenge/types"
	"github.com/tendermint/tendermint/crypto"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SubmitScavenge(goCtx context.Context, msg *types.MsgSubmitScavenge) (*types.MsgSubmitScavengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var scavenge = types.Scavenge{
		Index: msg.SolutionHash,
		Creator: msg.Creator,
		Description: msg.Description,
		SolutionHash: msg.SolutionHash,
		Reward: msg.Reward,
	}

	// SolutionHashをkeyとしてstoreからscavengeを取り出す
	_, isFound := k.GetScavenge(ctx, scavenge.SolutionHash)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Scavenge with that solution hash already exists")
	}

	// Scavenge module accountのアドレスを取得
	// 実際のモジュールが所有しているアカウントへの参照
	// このアカウントはScavengeが解決されるまで報奨金を保持するために使用される。
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	// MessageのCreator addressをstringからsdk.AccAddress
	scavenger, err := sdk,AccAddressFromBech32(scavenge.Creator)
	if err != nil {
		panic(err)
	}

	// tokenをstringからsdk.Coinsに変換する
	reward, err := sdk.ParseCoinsNormalized(scavenge.Reward)
	if err != nil {
		panic(err)
	}

	// scavenge creatorからmodule accountへtokenを送信する
	sdkError := k.bankKeeper.SendCoins(ctx, scavenger, moduleAcct, reward)
	if sdkError != nil {
		return nil, sdkError
	}

	// scavengeをstoreに書き込む
	k.SetScavenge(ctx, scavenge)

	return &types.MsgSubmitScavengeResponse{}, nil
}
