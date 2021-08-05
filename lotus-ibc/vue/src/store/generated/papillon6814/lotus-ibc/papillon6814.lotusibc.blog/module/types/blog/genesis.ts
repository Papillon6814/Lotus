/* eslint-disable */
import { TimedoutPost } from '../blog/timedout_post'
import { SendPost } from '../blog/send_post'
import { Post } from '../blog/post'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'papillon6814.lotusibc.blog'

/** GenesisState defines the blog module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  timedoutPost: TimedoutPost | undefined
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  sendPost: SendPost | undefined
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  post: Post | undefined
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  portId: string
}

const baseGenesisState: object = { portId: '' }

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.timedoutPost !== undefined) {
      TimedoutPost.encode(message.timedoutPost, writer.uint32(34).fork()).ldelim()
    }
    if (message.sendPost !== undefined) {
      SendPost.encode(message.sendPost, writer.uint32(26).fork()).ldelim()
    }
    if (message.post !== undefined) {
      Post.encode(message.post, writer.uint32(18).fork()).ldelim()
    }
    if (message.portId !== '') {
      writer.uint32(10).string(message.portId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 4:
          message.timedoutPost = TimedoutPost.decode(reader, reader.uint32())
          break
        case 3:
          message.sendPost = SendPost.decode(reader, reader.uint32())
          break
        case 2:
          message.post = Post.decode(reader, reader.uint32())
          break
        case 1:
          message.portId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    if (object.timedoutPost !== undefined && object.timedoutPost !== null) {
      message.timedoutPost = TimedoutPost.fromJSON(object.timedoutPost)
    } else {
      message.timedoutPost = undefined
    }
    if (object.sendPost !== undefined && object.sendPost !== null) {
      message.sendPost = SendPost.fromJSON(object.sendPost)
    } else {
      message.sendPost = undefined
    }
    if (object.post !== undefined && object.post !== null) {
      message.post = Post.fromJSON(object.post)
    } else {
      message.post = undefined
    }
    if (object.portId !== undefined && object.portId !== null) {
      message.portId = String(object.portId)
    } else {
      message.portId = ''
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    message.timedoutPost !== undefined && (obj.timedoutPost = message.timedoutPost ? TimedoutPost.toJSON(message.timedoutPost) : undefined)
    message.sendPost !== undefined && (obj.sendPost = message.sendPost ? SendPost.toJSON(message.sendPost) : undefined)
    message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined)
    message.portId !== undefined && (obj.portId = message.portId)
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    if (object.timedoutPost !== undefined && object.timedoutPost !== null) {
      message.timedoutPost = TimedoutPost.fromPartial(object.timedoutPost)
    } else {
      message.timedoutPost = undefined
    }
    if (object.sendPost !== undefined && object.sendPost !== null) {
      message.sendPost = SendPost.fromPartial(object.sendPost)
    } else {
      message.sendPost = undefined
    }
    if (object.post !== undefined && object.post !== null) {
      message.post = Post.fromPartial(object.post)
    } else {
      message.post = undefined
    }
    if (object.portId !== undefined && object.portId !== null) {
      message.portId = object.portId
    } else {
      message.portId = ''
    }
    return message
  }
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
