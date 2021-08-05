export declare type BlogMsgSendIbcPostResponse = object;
export interface BlogPost {
    creator?: string;
    title?: string;
    content?: string;
}
export interface BlogQueryGetPostResponse {
    Post?: BlogPost;
}
export interface BlogQueryGetSendPostResponse {
    SendPost?: BlogSendPost;
}
export interface BlogQueryGetTimedoutPostResponse {
    TimedoutPost?: BlogTimedoutPost;
}
export interface BlogSendPost {
    creator?: string;
    postID?: string;
    title?: string;
    chain?: string;
}
export interface BlogTimedoutPost {
    creator?: string;
    title?: string;
    chain?: string;
}
export interface ProtobufAny {
    typeUrl?: string;
    /** @format byte */
    value?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title blog/genesis.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryPost
     * @summary Queries a post by index.
     * @request GET:/papillon6814/lotusibc/blog/post
     */
    queryPost: (params?: RequestParams) => Promise<HttpResponse<BlogQueryGetPostResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QuerySendPost
     * @summary Queries a sendPost by index.
     * @request GET:/papillon6814/lotusibc/blog/sendPost
     */
    querySendPost: (params?: RequestParams) => Promise<HttpResponse<BlogQueryGetSendPostResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryTimedoutPost
     * @summary Queries a timedoutPost by index.
     * @request GET:/papillon6814/lotusibc/blog/timedoutPost
     */
    queryTimedoutPost: (params?: RequestParams) => Promise<HttpResponse<BlogQueryGetTimedoutPostResponse, RpcStatus>>;
}
export {};
