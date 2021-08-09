// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdatePoll } from "./types/voter/tx";
import { MsgCreateVote } from "./types/voter/tx";
import { MsgDeletePoll } from "./types/voter/tx";
import { MsgDeleteVote } from "./types/voter/tx";
import { MsgCreatePoll } from "./types/voter/tx";
import { MsgUpdateVote } from "./types/voter/tx";
const types = [
    ["/papillon6814.voter.voter.MsgUpdatePoll", MsgUpdatePoll],
    ["/papillon6814.voter.voter.MsgCreateVote", MsgCreateVote],
    ["/papillon6814.voter.voter.MsgDeletePoll", MsgDeletePoll],
    ["/papillon6814.voter.voter.MsgDeleteVote", MsgDeleteVote],
    ["/papillon6814.voter.voter.MsgCreatePoll", MsgCreatePoll],
    ["/papillon6814.voter.voter.MsgUpdateVote", MsgUpdateVote],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgUpdatePoll: (data) => ({ typeUrl: "/papillon6814.voter.voter.MsgUpdatePoll", value: data }),
        msgCreateVote: (data) => ({ typeUrl: "/papillon6814.voter.voter.MsgCreateVote", value: data }),
        msgDeletePoll: (data) => ({ typeUrl: "/papillon6814.voter.voter.MsgDeletePoll", value: data }),
        msgDeleteVote: (data) => ({ typeUrl: "/papillon6814.voter.voter.MsgDeleteVote", value: data }),
        msgCreatePoll: (data) => ({ typeUrl: "/papillon6814.voter.voter.MsgCreatePoll", value: data }),
        msgUpdateVote: (data) => ({ typeUrl: "/papillon6814.voter.voter.MsgUpdateVote", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
