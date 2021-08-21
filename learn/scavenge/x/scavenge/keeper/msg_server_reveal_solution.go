package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/papillon6814/scavenge/x/scavenge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) RevealSolution(goCtx context.Context, msg *types.MsgRevealSolution) (*types.MsgRevealSolutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// SolutionとScavengerのアドレスを連結して、byteに変換する
	var solutionScavengerBytes = []byte(msg.Solution + msg.Creator)
	// 変換したbyteのhashを見つける
	var solutionScavengerHash = sha256.Sum256(solutionScavengerBytes)
	// hashからstringに変換する
	var solutionScavengerHashString = hex.EncodeToString(solutionScavengerHash[:])

	// commitが見つからなかった場合はエラーを返す
	_, isFound := k.GetCommit(ctx, solutionScavengerHashString)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Commit with that hash does not exist")
	}

	// solutionのhashを見つける
	var solutionHash = sha256.Sum256([]byte(msg.Solution))
	// solutionのhashを文字列に変換する
	var solutionHashString = hex.EncodeToString(solutionHash[:])

	// SolutionHashを使ってscavengeをstoreから取得する
	var scavenge types.Scavenge
	scavenge, isFound = k.GetScavenge(ctx, solutionHashString)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors,ErrInvalidRequest, "Scavenge with that hash does not exist")
	}

	// scavenge変数のScavengerプロパティが有効なアドレスを持っているかを確認する
	// scavengeがすでに解決されている場合はエラーが吐かれる
	_, err := sdk.AccAddressFromBech32(scavenge.Scavenger)
	if err == nil {
		return nil, sdkerrors.Wrap(sdkerrors,ErrInvalidRequest, "Scavenge has already been solved")
	}

	// scavengerのアドレスをscavenge変数に保存する
	scavenge.Scavenger = msg.Creator
	// 正解をscavengeに保存する
	scavenge.Solution = msg.Solution
	// モジュールの口座を取得する
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	scavenger, err := sdk.AccAddressFromBech32(scavenge.Scavenger)
	if err != nil {
		panic(err)
	}

	// tokenをstringからcoinに変換する
	reward, err := sdk.ParseCoinsNormalized(scavenge.Reward)
	if err != nil {
		panic(err)
	}

	// tokenをmoduleのアカウントからscavengerへ送信する
	sdkError := k.bankKeeper.SendCoins(ctx, moduleAcct, scavenger, reward)
	if sdkError != nil {
		return nil, sdkError
	}

	// 更新されたScavengeをstoreに保存する
	k.SetScavenge(ctx, scavenge)

	return &types.MsgRevealSolutionResponse{}, nil
}
