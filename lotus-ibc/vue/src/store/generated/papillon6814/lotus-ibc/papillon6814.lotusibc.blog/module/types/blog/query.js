/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { TimedoutPost } from '../blog/timedout_post';
import { SendPost } from '../blog/send_post';
import { Post } from '../blog/post';
export const protobufPackage = 'papillon6814.lotusibc.blog';
const baseQueryGetTimedoutPostRequest = {};
export const QueryGetTimedoutPostRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetTimedoutPostRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseQueryGetTimedoutPostRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryGetTimedoutPostRequest };
        return message;
    }
};
const baseQueryGetTimedoutPostResponse = {};
export const QueryGetTimedoutPostResponse = {
    encode(message, writer = Writer.create()) {
        if (message.TimedoutPost !== undefined) {
            TimedoutPost.encode(message.TimedoutPost, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetTimedoutPostResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.TimedoutPost = TimedoutPost.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetTimedoutPostResponse };
        if (object.TimedoutPost !== undefined && object.TimedoutPost !== null) {
            message.TimedoutPost = TimedoutPost.fromJSON(object.TimedoutPost);
        }
        else {
            message.TimedoutPost = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.TimedoutPost !== undefined && (obj.TimedoutPost = message.TimedoutPost ? TimedoutPost.toJSON(message.TimedoutPost) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetTimedoutPostResponse };
        if (object.TimedoutPost !== undefined && object.TimedoutPost !== null) {
            message.TimedoutPost = TimedoutPost.fromPartial(object.TimedoutPost);
        }
        else {
            message.TimedoutPost = undefined;
        }
        return message;
    }
};
const baseQueryGetSendPostRequest = {};
export const QueryGetSendPostRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetSendPostRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseQueryGetSendPostRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryGetSendPostRequest };
        return message;
    }
};
const baseQueryGetSendPostResponse = {};
export const QueryGetSendPostResponse = {
    encode(message, writer = Writer.create()) {
        if (message.SendPost !== undefined) {
            SendPost.encode(message.SendPost, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetSendPostResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.SendPost = SendPost.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetSendPostResponse };
        if (object.SendPost !== undefined && object.SendPost !== null) {
            message.SendPost = SendPost.fromJSON(object.SendPost);
        }
        else {
            message.SendPost = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.SendPost !== undefined && (obj.SendPost = message.SendPost ? SendPost.toJSON(message.SendPost) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetSendPostResponse };
        if (object.SendPost !== undefined && object.SendPost !== null) {
            message.SendPost = SendPost.fromPartial(object.SendPost);
        }
        else {
            message.SendPost = undefined;
        }
        return message;
    }
};
const baseQueryGetPostRequest = {};
export const QueryGetPostRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPostRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseQueryGetPostRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryGetPostRequest };
        return message;
    }
};
const baseQueryGetPostResponse = {};
export const QueryGetPostResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Post !== undefined) {
            Post.encode(message.Post, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPostResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Post = Post.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetPostResponse };
        if (object.Post !== undefined && object.Post !== null) {
            message.Post = Post.fromJSON(object.Post);
        }
        else {
            message.Post = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Post !== undefined && (obj.Post = message.Post ? Post.toJSON(message.Post) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPostResponse };
        if (object.Post !== undefined && object.Post !== null) {
            message.Post = Post.fromPartial(object.Post);
        }
        else {
            message.Post = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    TimedoutPost(request) {
        const data = QueryGetTimedoutPostRequest.encode(request).finish();
        const promise = this.rpc.request('papillon6814.lotusibc.blog.Query', 'TimedoutPost', data);
        return promise.then((data) => QueryGetTimedoutPostResponse.decode(new Reader(data)));
    }
    SendPost(request) {
        const data = QueryGetSendPostRequest.encode(request).finish();
        const promise = this.rpc.request('papillon6814.lotusibc.blog.Query', 'SendPost', data);
        return promise.then((data) => QueryGetSendPostResponse.decode(new Reader(data)));
    }
    Post(request) {
        const data = QueryGetPostRequest.encode(request).finish();
        const promise = this.rpc.request('papillon6814.lotusibc.blog.Query', 'Post', data);
        return promise.then((data) => QueryGetPostResponse.decode(new Reader(data)));
    }
}
