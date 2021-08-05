/* eslint-disable */
import { TimedoutPost } from '../blog/timedout_post';
import { SendPost } from '../blog/send_post';
import { Post } from '../blog/post';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'papillon6814.lotusibc.blog';
const baseGenesisState = { portId: '' };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.timedoutPost !== undefined) {
            TimedoutPost.encode(message.timedoutPost, writer.uint32(34).fork()).ldelim();
        }
        if (message.sendPost !== undefined) {
            SendPost.encode(message.sendPost, writer.uint32(26).fork()).ldelim();
        }
        if (message.post !== undefined) {
            Post.encode(message.post, writer.uint32(18).fork()).ldelim();
        }
        if (message.portId !== '') {
            writer.uint32(10).string(message.portId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 4:
                    message.timedoutPost = TimedoutPost.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.sendPost = SendPost.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.post = Post.decode(reader, reader.uint32());
                    break;
                case 1:
                    message.portId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        if (object.timedoutPost !== undefined && object.timedoutPost !== null) {
            message.timedoutPost = TimedoutPost.fromJSON(object.timedoutPost);
        }
        else {
            message.timedoutPost = undefined;
        }
        if (object.sendPost !== undefined && object.sendPost !== null) {
            message.sendPost = SendPost.fromJSON(object.sendPost);
        }
        else {
            message.sendPost = undefined;
        }
        if (object.post !== undefined && object.post !== null) {
            message.post = Post.fromJSON(object.post);
        }
        else {
            message.post = undefined;
        }
        if (object.portId !== undefined && object.portId !== null) {
            message.portId = String(object.portId);
        }
        else {
            message.portId = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.timedoutPost !== undefined && (obj.timedoutPost = message.timedoutPost ? TimedoutPost.toJSON(message.timedoutPost) : undefined);
        message.sendPost !== undefined && (obj.sendPost = message.sendPost ? SendPost.toJSON(message.sendPost) : undefined);
        message.post !== undefined && (obj.post = message.post ? Post.toJSON(message.post) : undefined);
        message.portId !== undefined && (obj.portId = message.portId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        if (object.timedoutPost !== undefined && object.timedoutPost !== null) {
            message.timedoutPost = TimedoutPost.fromPartial(object.timedoutPost);
        }
        else {
            message.timedoutPost = undefined;
        }
        if (object.sendPost !== undefined && object.sendPost !== null) {
            message.sendPost = SendPost.fromPartial(object.sendPost);
        }
        else {
            message.sendPost = undefined;
        }
        if (object.post !== undefined && object.post !== null) {
            message.post = Post.fromPartial(object.post);
        }
        else {
            message.post = undefined;
        }
        if (object.portId !== undefined && object.portId !== null) {
            message.portId = object.portId;
        }
        else {
            message.portId = '';
        }
        return message;
    }
};
