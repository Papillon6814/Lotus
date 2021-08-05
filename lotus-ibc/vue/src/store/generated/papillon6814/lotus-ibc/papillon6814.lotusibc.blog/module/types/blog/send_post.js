/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'papillon6814.lotusibc.blog';
const baseSendPost = { creator: '', postID: '', title: '', chain: '' };
export const SendPost = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.postID !== '') {
            writer.uint32(18).string(message.postID);
        }
        if (message.title !== '') {
            writer.uint32(26).string(message.title);
        }
        if (message.chain !== '') {
            writer.uint32(34).string(message.chain);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSendPost };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.postID = reader.string();
                    break;
                case 3:
                    message.title = reader.string();
                    break;
                case 4:
                    message.chain = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseSendPost };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.postID !== undefined && object.postID !== null) {
            message.postID = String(object.postID);
        }
        else {
            message.postID = '';
        }
        if (object.title !== undefined && object.title !== null) {
            message.title = String(object.title);
        }
        else {
            message.title = '';
        }
        if (object.chain !== undefined && object.chain !== null) {
            message.chain = String(object.chain);
        }
        else {
            message.chain = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.postID !== undefined && (obj.postID = message.postID);
        message.title !== undefined && (obj.title = message.title);
        message.chain !== undefined && (obj.chain = message.chain);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSendPost };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.postID !== undefined && object.postID !== null) {
            message.postID = object.postID;
        }
        else {
            message.postID = '';
        }
        if (object.title !== undefined && object.title !== null) {
            message.title = object.title;
        }
        else {
            message.title = '';
        }
        if (object.chain !== undefined && object.chain !== null) {
            message.chain = object.chain;
        }
        else {
            message.chain = '';
        }
        return message;
    }
};
