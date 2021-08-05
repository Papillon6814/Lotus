import { Reader, Writer } from 'protobufjs/minimal';
import { TimedoutPost } from '../blog/timedout_post';
import { SendPost } from '../blog/send_post';
import { Post } from '../blog/post';
export declare const protobufPackage = "papillon6814.lotusibc.blog";
/** this line is used by starport scaffolding # 3 */
export interface QueryGetTimedoutPostRequest {
}
export interface QueryGetTimedoutPostResponse {
    TimedoutPost: TimedoutPost | undefined;
}
export interface QueryGetSendPostRequest {
}
export interface QueryGetSendPostResponse {
    SendPost: SendPost | undefined;
}
export interface QueryGetPostRequest {
}
export interface QueryGetPostResponse {
    Post: Post | undefined;
}
export declare const QueryGetTimedoutPostRequest: {
    encode(_: QueryGetTimedoutPostRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTimedoutPostRequest;
    fromJSON(_: any): QueryGetTimedoutPostRequest;
    toJSON(_: QueryGetTimedoutPostRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetTimedoutPostRequest>): QueryGetTimedoutPostRequest;
};
export declare const QueryGetTimedoutPostResponse: {
    encode(message: QueryGetTimedoutPostResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTimedoutPostResponse;
    fromJSON(object: any): QueryGetTimedoutPostResponse;
    toJSON(message: QueryGetTimedoutPostResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetTimedoutPostResponse>): QueryGetTimedoutPostResponse;
};
export declare const QueryGetSendPostRequest: {
    encode(_: QueryGetSendPostRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSendPostRequest;
    fromJSON(_: any): QueryGetSendPostRequest;
    toJSON(_: QueryGetSendPostRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetSendPostRequest>): QueryGetSendPostRequest;
};
export declare const QueryGetSendPostResponse: {
    encode(message: QueryGetSendPostResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSendPostResponse;
    fromJSON(object: any): QueryGetSendPostResponse;
    toJSON(message: QueryGetSendPostResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSendPostResponse>): QueryGetSendPostResponse;
};
export declare const QueryGetPostRequest: {
    encode(_: QueryGetPostRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPostRequest;
    fromJSON(_: any): QueryGetPostRequest;
    toJSON(_: QueryGetPostRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetPostRequest>): QueryGetPostRequest;
};
export declare const QueryGetPostResponse: {
    encode(message: QueryGetPostResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPostResponse;
    fromJSON(object: any): QueryGetPostResponse;
    toJSON(message: QueryGetPostResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetPostResponse>): QueryGetPostResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a timedoutPost by index. */
    TimedoutPost(request: QueryGetTimedoutPostRequest): Promise<QueryGetTimedoutPostResponse>;
    /** Queries a sendPost by index. */
    SendPost(request: QueryGetSendPostRequest): Promise<QueryGetSendPostResponse>;
    /** Queries a post by index. */
    Post(request: QueryGetPostRequest): Promise<QueryGetPostResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    TimedoutPost(request: QueryGetTimedoutPostRequest): Promise<QueryGetTimedoutPostResponse>;
    SendPost(request: QueryGetSendPostRequest): Promise<QueryGetSendPostResponse>;
    Post(request: QueryGetPostRequest): Promise<QueryGetPostResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
