/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { TimedoutPost } from '../blog/timedout_post'
import { SendPost } from '../blog/send_post'
import { Post } from '../blog/post'

export const protobufPackage = 'papillon6814.lotusibc.blog'

/** this line is used by starport scaffolding # 3 */
export interface QueryGetTimedoutPostRequest {}

export interface QueryGetTimedoutPostResponse {
  TimedoutPost: TimedoutPost | undefined
}

export interface QueryGetSendPostRequest {}

export interface QueryGetSendPostResponse {
  SendPost: SendPost | undefined
}

export interface QueryGetPostRequest {}

export interface QueryGetPostResponse {
  Post: Post | undefined
}

const baseQueryGetTimedoutPostRequest: object = {}

export const QueryGetTimedoutPostRequest = {
  encode(_: QueryGetTimedoutPostRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTimedoutPostRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetTimedoutPostRequest } as QueryGetTimedoutPostRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): QueryGetTimedoutPostRequest {
    const message = { ...baseQueryGetTimedoutPostRequest } as QueryGetTimedoutPostRequest
    return message
  },

  toJSON(_: QueryGetTimedoutPostRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetTimedoutPostRequest>): QueryGetTimedoutPostRequest {
    const message = { ...baseQueryGetTimedoutPostRequest } as QueryGetTimedoutPostRequest
    return message
  }
}

const baseQueryGetTimedoutPostResponse: object = {}

export const QueryGetTimedoutPostResponse = {
  encode(message: QueryGetTimedoutPostResponse, writer: Writer = Writer.create()): Writer {
    if (message.TimedoutPost !== undefined) {
      TimedoutPost.encode(message.TimedoutPost, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTimedoutPostResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetTimedoutPostResponse } as QueryGetTimedoutPostResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.TimedoutPost = TimedoutPost.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetTimedoutPostResponse {
    const message = { ...baseQueryGetTimedoutPostResponse } as QueryGetTimedoutPostResponse
    if (object.TimedoutPost !== undefined && object.TimedoutPost !== null) {
      message.TimedoutPost = TimedoutPost.fromJSON(object.TimedoutPost)
    } else {
      message.TimedoutPost = undefined
    }
    return message
  },

  toJSON(message: QueryGetTimedoutPostResponse): unknown {
    const obj: any = {}
    message.TimedoutPost !== undefined && (obj.TimedoutPost = message.TimedoutPost ? TimedoutPost.toJSON(message.TimedoutPost) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetTimedoutPostResponse>): QueryGetTimedoutPostResponse {
    const message = { ...baseQueryGetTimedoutPostResponse } as QueryGetTimedoutPostResponse
    if (object.TimedoutPost !== undefined && object.TimedoutPost !== null) {
      message.TimedoutPost = TimedoutPost.fromPartial(object.TimedoutPost)
    } else {
      message.TimedoutPost = undefined
    }
    return message
  }
}

const baseQueryGetSendPostRequest: object = {}

export const QueryGetSendPostRequest = {
  encode(_: QueryGetSendPostRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetSendPostRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetSendPostRequest } as QueryGetSendPostRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): QueryGetSendPostRequest {
    const message = { ...baseQueryGetSendPostRequest } as QueryGetSendPostRequest
    return message
  },

  toJSON(_: QueryGetSendPostRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetSendPostRequest>): QueryGetSendPostRequest {
    const message = { ...baseQueryGetSendPostRequest } as QueryGetSendPostRequest
    return message
  }
}

const baseQueryGetSendPostResponse: object = {}

export const QueryGetSendPostResponse = {
  encode(message: QueryGetSendPostResponse, writer: Writer = Writer.create()): Writer {
    if (message.SendPost !== undefined) {
      SendPost.encode(message.SendPost, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetSendPostResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetSendPostResponse } as QueryGetSendPostResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.SendPost = SendPost.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetSendPostResponse {
    const message = { ...baseQueryGetSendPostResponse } as QueryGetSendPostResponse
    if (object.SendPost !== undefined && object.SendPost !== null) {
      message.SendPost = SendPost.fromJSON(object.SendPost)
    } else {
      message.SendPost = undefined
    }
    return message
  },

  toJSON(message: QueryGetSendPostResponse): unknown {
    const obj: any = {}
    message.SendPost !== undefined && (obj.SendPost = message.SendPost ? SendPost.toJSON(message.SendPost) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetSendPostResponse>): QueryGetSendPostResponse {
    const message = { ...baseQueryGetSendPostResponse } as QueryGetSendPostResponse
    if (object.SendPost !== undefined && object.SendPost !== null) {
      message.SendPost = SendPost.fromPartial(object.SendPost)
    } else {
      message.SendPost = undefined
    }
    return message
  }
}

const baseQueryGetPostRequest: object = {}

export const QueryGetPostRequest = {
  encode(_: QueryGetPostRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPostRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPostRequest } as QueryGetPostRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): QueryGetPostRequest {
    const message = { ...baseQueryGetPostRequest } as QueryGetPostRequest
    return message
  },

  toJSON(_: QueryGetPostRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetPostRequest>): QueryGetPostRequest {
    const message = { ...baseQueryGetPostRequest } as QueryGetPostRequest
    return message
  }
}

const baseQueryGetPostResponse: object = {}

export const QueryGetPostResponse = {
  encode(message: QueryGetPostResponse, writer: Writer = Writer.create()): Writer {
    if (message.Post !== undefined) {
      Post.encode(message.Post, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPostResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPostResponse } as QueryGetPostResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.Post = Post.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPostResponse {
    const message = { ...baseQueryGetPostResponse } as QueryGetPostResponse
    if (object.Post !== undefined && object.Post !== null) {
      message.Post = Post.fromJSON(object.Post)
    } else {
      message.Post = undefined
    }
    return message
  },

  toJSON(message: QueryGetPostResponse): unknown {
    const obj: any = {}
    message.Post !== undefined && (obj.Post = message.Post ? Post.toJSON(message.Post) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPostResponse>): QueryGetPostResponse {
    const message = { ...baseQueryGetPostResponse } as QueryGetPostResponse
    if (object.Post !== undefined && object.Post !== null) {
      message.Post = Post.fromPartial(object.Post)
    } else {
      message.Post = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a timedoutPost by index. */
  TimedoutPost(request: QueryGetTimedoutPostRequest): Promise<QueryGetTimedoutPostResponse>
  /** Queries a sendPost by index. */
  SendPost(request: QueryGetSendPostRequest): Promise<QueryGetSendPostResponse>
  /** Queries a post by index. */
  Post(request: QueryGetPostRequest): Promise<QueryGetPostResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  TimedoutPost(request: QueryGetTimedoutPostRequest): Promise<QueryGetTimedoutPostResponse> {
    const data = QueryGetTimedoutPostRequest.encode(request).finish()
    const promise = this.rpc.request('papillon6814.lotusibc.blog.Query', 'TimedoutPost', data)
    return promise.then((data) => QueryGetTimedoutPostResponse.decode(new Reader(data)))
  }

  SendPost(request: QueryGetSendPostRequest): Promise<QueryGetSendPostResponse> {
    const data = QueryGetSendPostRequest.encode(request).finish()
    const promise = this.rpc.request('papillon6814.lotusibc.blog.Query', 'SendPost', data)
    return promise.then((data) => QueryGetSendPostResponse.decode(new Reader(data)))
  }

  Post(request: QueryGetPostRequest): Promise<QueryGetPostResponse> {
    const data = QueryGetPostRequest.encode(request).finish()
    const promise = this.rpc.request('papillon6814.lotusibc.blog.Query', 'Post', data)
    return promise.then((data) => QueryGetPostResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
