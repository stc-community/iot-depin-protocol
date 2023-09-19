/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stccommunity.iotdepinprotocol.iotdepinprotocol";

export interface MsgCreateDevice {
  creator: string;
  deviceName: string;
  address: string;
  value: string;
}

export interface MsgCreateDeviceResponse {
}

export interface MsgUpdateDevice {
  creator: string;
  deviceName: string;
  address: string;
  value: string;
}

export interface MsgUpdateDeviceResponse {
}

export interface MsgDeleteDevice {
  creator: string;
  deviceName: string;
}

export interface MsgDeleteDeviceResponse {
}

export interface MsgCreateKv {
  creator: string;
  index: string;
  value: string;
  deviceName: string;
}

export interface MsgCreateKvResponse {
}

export interface MsgUpdateKv {
  creator: string;
  index: string;
  value: string;
  deviceName: string;
}

export interface MsgUpdateKvResponse {
}

export interface MsgDeleteKv {
  creator: string;
  index: string;
  deviceName: string;
}

export interface MsgDeleteKvResponse {
}

export interface MsgCreateEventPb {
  creator: string;
  index: string;
  deviceName: string;
  payload: string;
  createdAt: string;
  updatedAt: string;
}

export interface MsgCreateEventPbResponse {
}

export interface MsgUpdateEventPb {
  creator: string;
  index: string;
  deviceName: string;
  payload: string;
  createdAt: string;
  updatedAt: string;
}

export interface MsgUpdateEventPbResponse {
}

export interface MsgDeleteEventPb {
  creator: string;
  index: string;
  deviceName: string;
}

export interface MsgDeleteEventPbResponse {
}

export interface MsgCreateDeviceRegistry {
  creator: string;
  mid: string;
  metaData: string;
}

export interface MsgCreateDeviceRegistryResponse {
}

export interface MsgUpdateDeviceRegistry {
  creator: string;
  mid: string;
  metaData: string;
}

export interface MsgUpdateDeviceRegistryResponse {
}

export interface MsgDeleteDeviceRegistry {
  creator: string;
  mid: string;
}

export interface MsgDeleteDeviceRegistryResponse {
}

function createBaseMsgCreateDevice(): MsgCreateDevice {
  return { creator: "", deviceName: "", address: "", value: "" };
}

export const MsgCreateDevice = {
  encode(message: MsgCreateDevice, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.deviceName !== "") {
      writer.uint32(18).string(message.deviceName);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.value !== "") {
      writer.uint32(34).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDevice {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDevice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.deviceName = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDevice {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
      address: isSet(object.address) ? String(object.address) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: MsgCreateDevice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.deviceName !== undefined && (obj.deviceName = message.deviceName);
    message.address !== undefined && (obj.address = message.address);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDevice>, I>>(object: I): MsgCreateDevice {
    const message = createBaseMsgCreateDevice();
    message.creator = object.creator ?? "";
    message.deviceName = object.deviceName ?? "";
    message.address = object.address ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseMsgCreateDeviceResponse(): MsgCreateDeviceResponse {
  return {};
}

export const MsgCreateDeviceResponse = {
  encode(_: MsgCreateDeviceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDeviceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDeviceResponse();
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

  fromJSON(_: any): MsgCreateDeviceResponse {
    return {};
  },

  toJSON(_: MsgCreateDeviceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDeviceResponse>, I>>(_: I): MsgCreateDeviceResponse {
    const message = createBaseMsgCreateDeviceResponse();
    return message;
  },
};

function createBaseMsgUpdateDevice(): MsgUpdateDevice {
  return { creator: "", deviceName: "", address: "", value: "" };
}

export const MsgUpdateDevice = {
  encode(message: MsgUpdateDevice, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.deviceName !== "") {
      writer.uint32(18).string(message.deviceName);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    if (message.value !== "") {
      writer.uint32(34).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDevice {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDevice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.deviceName = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        case 4:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateDevice {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
      address: isSet(object.address) ? String(object.address) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: MsgUpdateDevice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.deviceName !== undefined && (obj.deviceName = message.deviceName);
    message.address !== undefined && (obj.address = message.address);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDevice>, I>>(object: I): MsgUpdateDevice {
    const message = createBaseMsgUpdateDevice();
    message.creator = object.creator ?? "";
    message.deviceName = object.deviceName ?? "";
    message.address = object.address ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseMsgUpdateDeviceResponse(): MsgUpdateDeviceResponse {
  return {};
}

export const MsgUpdateDeviceResponse = {
  encode(_: MsgUpdateDeviceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDeviceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDeviceResponse();
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

  fromJSON(_: any): MsgUpdateDeviceResponse {
    return {};
  },

  toJSON(_: MsgUpdateDeviceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDeviceResponse>, I>>(_: I): MsgUpdateDeviceResponse {
    const message = createBaseMsgUpdateDeviceResponse();
    return message;
  },
};

function createBaseMsgDeleteDevice(): MsgDeleteDevice {
  return { creator: "", deviceName: "" };
}

export const MsgDeleteDevice = {
  encode(message: MsgDeleteDevice, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.deviceName !== "") {
      writer.uint32(18).string(message.deviceName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDevice {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDevice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.deviceName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteDevice {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
    };
  },

  toJSON(message: MsgDeleteDevice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.deviceName !== undefined && (obj.deviceName = message.deviceName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDevice>, I>>(object: I): MsgDeleteDevice {
    const message = createBaseMsgDeleteDevice();
    message.creator = object.creator ?? "";
    message.deviceName = object.deviceName ?? "";
    return message;
  },
};

function createBaseMsgDeleteDeviceResponse(): MsgDeleteDeviceResponse {
  return {};
}

export const MsgDeleteDeviceResponse = {
  encode(_: MsgDeleteDeviceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDeviceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDeviceResponse();
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

  fromJSON(_: any): MsgDeleteDeviceResponse {
    return {};
  },

  toJSON(_: MsgDeleteDeviceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDeviceResponse>, I>>(_: I): MsgDeleteDeviceResponse {
    const message = createBaseMsgDeleteDeviceResponse();
    return message;
  },
};

function createBaseMsgCreateKv(): MsgCreateKv {
  return { creator: "", index: "", value: "", deviceName: "" };
}

export const MsgCreateKv = {
  encode(message: MsgCreateKv, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    if (message.deviceName !== "") {
      writer.uint32(34).string(message.deviceName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateKv {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateKv();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        case 4:
          message.deviceName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateKv {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      value: isSet(object.value) ? String(object.value) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
    };
  },

  toJSON(message: MsgCreateKv): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.value !== undefined && (obj.value = message.value);
    message.deviceName !== undefined && (obj.deviceName = message.deviceName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateKv>, I>>(object: I): MsgCreateKv {
    const message = createBaseMsgCreateKv();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.value = object.value ?? "";
    message.deviceName = object.deviceName ?? "";
    return message;
  },
};

function createBaseMsgCreateKvResponse(): MsgCreateKvResponse {
  return {};
}

export const MsgCreateKvResponse = {
  encode(_: MsgCreateKvResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateKvResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateKvResponse();
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

  fromJSON(_: any): MsgCreateKvResponse {
    return {};
  },

  toJSON(_: MsgCreateKvResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateKvResponse>, I>>(_: I): MsgCreateKvResponse {
    const message = createBaseMsgCreateKvResponse();
    return message;
  },
};

function createBaseMsgUpdateKv(): MsgUpdateKv {
  return { creator: "", index: "", value: "", deviceName: "" };
}

export const MsgUpdateKv = {
  encode(message: MsgUpdateKv, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    if (message.deviceName !== "") {
      writer.uint32(34).string(message.deviceName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateKv {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateKv();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        case 4:
          message.deviceName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateKv {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      value: isSet(object.value) ? String(object.value) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
    };
  },

  toJSON(message: MsgUpdateKv): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.value !== undefined && (obj.value = message.value);
    message.deviceName !== undefined && (obj.deviceName = message.deviceName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateKv>, I>>(object: I): MsgUpdateKv {
    const message = createBaseMsgUpdateKv();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.value = object.value ?? "";
    message.deviceName = object.deviceName ?? "";
    return message;
  },
};

function createBaseMsgUpdateKvResponse(): MsgUpdateKvResponse {
  return {};
}

export const MsgUpdateKvResponse = {
  encode(_: MsgUpdateKvResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateKvResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateKvResponse();
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

  fromJSON(_: any): MsgUpdateKvResponse {
    return {};
  },

  toJSON(_: MsgUpdateKvResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateKvResponse>, I>>(_: I): MsgUpdateKvResponse {
    const message = createBaseMsgUpdateKvResponse();
    return message;
  },
};

function createBaseMsgDeleteKv(): MsgDeleteKv {
  return { creator: "", index: "", deviceName: "" };
}

export const MsgDeleteKv = {
  encode(message: MsgDeleteKv, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.deviceName !== "") {
      writer.uint32(26).string(message.deviceName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteKv {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteKv();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.deviceName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteKv {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
    };
  },

  toJSON(message: MsgDeleteKv): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.deviceName !== undefined && (obj.deviceName = message.deviceName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteKv>, I>>(object: I): MsgDeleteKv {
    const message = createBaseMsgDeleteKv();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.deviceName = object.deviceName ?? "";
    return message;
  },
};

function createBaseMsgDeleteKvResponse(): MsgDeleteKvResponse {
  return {};
}

export const MsgDeleteKvResponse = {
  encode(_: MsgDeleteKvResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteKvResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteKvResponse();
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

  fromJSON(_: any): MsgDeleteKvResponse {
    return {};
  },

  toJSON(_: MsgDeleteKvResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteKvResponse>, I>>(_: I): MsgDeleteKvResponse {
    const message = createBaseMsgDeleteKvResponse();
    return message;
  },
};

function createBaseMsgCreateEventPb(): MsgCreateEventPb {
  return { creator: "", index: "", deviceName: "", payload: "", createdAt: "", updatedAt: "" };
}

export const MsgCreateEventPb = {
  encode(message: MsgCreateEventPb, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.deviceName !== "") {
      writer.uint32(26).string(message.deviceName);
    }
    if (message.payload !== "") {
      writer.uint32(34).string(message.payload);
    }
    if (message.createdAt !== "") {
      writer.uint32(42).string(message.createdAt);
    }
    if (message.updatedAt !== "") {
      writer.uint32(50).string(message.updatedAt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateEventPb {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateEventPb();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.deviceName = reader.string();
          break;
        case 4:
          message.payload = reader.string();
          break;
        case 5:
          message.createdAt = reader.string();
          break;
        case 6:
          message.updatedAt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateEventPb {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
      payload: isSet(object.payload) ? String(object.payload) : "",
      createdAt: isSet(object.createdAt) ? String(object.createdAt) : "",
      updatedAt: isSet(object.updatedAt) ? String(object.updatedAt) : "",
    };
  },

  toJSON(message: MsgCreateEventPb): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.deviceName !== undefined && (obj.deviceName = message.deviceName);
    message.payload !== undefined && (obj.payload = message.payload);
    message.createdAt !== undefined && (obj.createdAt = message.createdAt);
    message.updatedAt !== undefined && (obj.updatedAt = message.updatedAt);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateEventPb>, I>>(object: I): MsgCreateEventPb {
    const message = createBaseMsgCreateEventPb();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.deviceName = object.deviceName ?? "";
    message.payload = object.payload ?? "";
    message.createdAt = object.createdAt ?? "";
    message.updatedAt = object.updatedAt ?? "";
    return message;
  },
};

function createBaseMsgCreateEventPbResponse(): MsgCreateEventPbResponse {
  return {};
}

export const MsgCreateEventPbResponse = {
  encode(_: MsgCreateEventPbResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateEventPbResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateEventPbResponse();
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

  fromJSON(_: any): MsgCreateEventPbResponse {
    return {};
  },

  toJSON(_: MsgCreateEventPbResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateEventPbResponse>, I>>(_: I): MsgCreateEventPbResponse {
    const message = createBaseMsgCreateEventPbResponse();
    return message;
  },
};

function createBaseMsgUpdateEventPb(): MsgUpdateEventPb {
  return { creator: "", index: "", deviceName: "", payload: "", createdAt: "", updatedAt: "" };
}

export const MsgUpdateEventPb = {
  encode(message: MsgUpdateEventPb, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.deviceName !== "") {
      writer.uint32(26).string(message.deviceName);
    }
    if (message.payload !== "") {
      writer.uint32(34).string(message.payload);
    }
    if (message.createdAt !== "") {
      writer.uint32(42).string(message.createdAt);
    }
    if (message.updatedAt !== "") {
      writer.uint32(50).string(message.updatedAt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateEventPb {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateEventPb();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.deviceName = reader.string();
          break;
        case 4:
          message.payload = reader.string();
          break;
        case 5:
          message.createdAt = reader.string();
          break;
        case 6:
          message.updatedAt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateEventPb {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
      payload: isSet(object.payload) ? String(object.payload) : "",
      createdAt: isSet(object.createdAt) ? String(object.createdAt) : "",
      updatedAt: isSet(object.updatedAt) ? String(object.updatedAt) : "",
    };
  },

  toJSON(message: MsgUpdateEventPb): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.deviceName !== undefined && (obj.deviceName = message.deviceName);
    message.payload !== undefined && (obj.payload = message.payload);
    message.createdAt !== undefined && (obj.createdAt = message.createdAt);
    message.updatedAt !== undefined && (obj.updatedAt = message.updatedAt);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateEventPb>, I>>(object: I): MsgUpdateEventPb {
    const message = createBaseMsgUpdateEventPb();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.deviceName = object.deviceName ?? "";
    message.payload = object.payload ?? "";
    message.createdAt = object.createdAt ?? "";
    message.updatedAt = object.updatedAt ?? "";
    return message;
  },
};

function createBaseMsgUpdateEventPbResponse(): MsgUpdateEventPbResponse {
  return {};
}

export const MsgUpdateEventPbResponse = {
  encode(_: MsgUpdateEventPbResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateEventPbResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateEventPbResponse();
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

  fromJSON(_: any): MsgUpdateEventPbResponse {
    return {};
  },

  toJSON(_: MsgUpdateEventPbResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateEventPbResponse>, I>>(_: I): MsgUpdateEventPbResponse {
    const message = createBaseMsgUpdateEventPbResponse();
    return message;
  },
};

function createBaseMsgDeleteEventPb(): MsgDeleteEventPb {
  return { creator: "", index: "", deviceName: "" };
}

export const MsgDeleteEventPb = {
  encode(message: MsgDeleteEventPb, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== "") {
      writer.uint32(18).string(message.index);
    }
    if (message.deviceName !== "") {
      writer.uint32(26).string(message.deviceName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteEventPb {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteEventPb();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = reader.string();
          break;
        case 3:
          message.deviceName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteEventPb {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? String(object.index) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
    };
  },

  toJSON(message: MsgDeleteEventPb): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = message.index);
    message.deviceName !== undefined && (obj.deviceName = message.deviceName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteEventPb>, I>>(object: I): MsgDeleteEventPb {
    const message = createBaseMsgDeleteEventPb();
    message.creator = object.creator ?? "";
    message.index = object.index ?? "";
    message.deviceName = object.deviceName ?? "";
    return message;
  },
};

function createBaseMsgDeleteEventPbResponse(): MsgDeleteEventPbResponse {
  return {};
}

export const MsgDeleteEventPbResponse = {
  encode(_: MsgDeleteEventPbResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteEventPbResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteEventPbResponse();
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

  fromJSON(_: any): MsgDeleteEventPbResponse {
    return {};
  },

  toJSON(_: MsgDeleteEventPbResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteEventPbResponse>, I>>(_: I): MsgDeleteEventPbResponse {
    const message = createBaseMsgDeleteEventPbResponse();
    return message;
  },
};

function createBaseMsgCreateDeviceRegistry(): MsgCreateDeviceRegistry {
  return { creator: "", mid: "", metaData: "" };
}

export const MsgCreateDeviceRegistry = {
  encode(message: MsgCreateDeviceRegistry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.mid !== "") {
      writer.uint32(18).string(message.mid);
    }
    if (message.metaData !== "") {
      writer.uint32(26).string(message.metaData);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDeviceRegistry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDeviceRegistry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.mid = reader.string();
          break;
        case 3:
          message.metaData = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDeviceRegistry {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      mid: isSet(object.mid) ? String(object.mid) : "",
      metaData: isSet(object.metaData) ? String(object.metaData) : "",
    };
  },

  toJSON(message: MsgCreateDeviceRegistry): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.mid !== undefined && (obj.mid = message.mid);
    message.metaData !== undefined && (obj.metaData = message.metaData);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDeviceRegistry>, I>>(object: I): MsgCreateDeviceRegistry {
    const message = createBaseMsgCreateDeviceRegistry();
    message.creator = object.creator ?? "";
    message.mid = object.mid ?? "";
    message.metaData = object.metaData ?? "";
    return message;
  },
};

function createBaseMsgCreateDeviceRegistryResponse(): MsgCreateDeviceRegistryResponse {
  return {};
}

export const MsgCreateDeviceRegistryResponse = {
  encode(_: MsgCreateDeviceRegistryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateDeviceRegistryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateDeviceRegistryResponse();
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

  fromJSON(_: any): MsgCreateDeviceRegistryResponse {
    return {};
  },

  toJSON(_: MsgCreateDeviceRegistryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateDeviceRegistryResponse>, I>>(_: I): MsgCreateDeviceRegistryResponse {
    const message = createBaseMsgCreateDeviceRegistryResponse();
    return message;
  },
};

function createBaseMsgUpdateDeviceRegistry(): MsgUpdateDeviceRegistry {
  return { creator: "", mid: "", metaData: "" };
}

export const MsgUpdateDeviceRegistry = {
  encode(message: MsgUpdateDeviceRegistry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.mid !== "") {
      writer.uint32(18).string(message.mid);
    }
    if (message.metaData !== "") {
      writer.uint32(26).string(message.metaData);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDeviceRegistry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDeviceRegistry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.mid = reader.string();
          break;
        case 3:
          message.metaData = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateDeviceRegistry {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      mid: isSet(object.mid) ? String(object.mid) : "",
      metaData: isSet(object.metaData) ? String(object.metaData) : "",
    };
  },

  toJSON(message: MsgUpdateDeviceRegistry): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.mid !== undefined && (obj.mid = message.mid);
    message.metaData !== undefined && (obj.metaData = message.metaData);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDeviceRegistry>, I>>(object: I): MsgUpdateDeviceRegistry {
    const message = createBaseMsgUpdateDeviceRegistry();
    message.creator = object.creator ?? "";
    message.mid = object.mid ?? "";
    message.metaData = object.metaData ?? "";
    return message;
  },
};

function createBaseMsgUpdateDeviceRegistryResponse(): MsgUpdateDeviceRegistryResponse {
  return {};
}

export const MsgUpdateDeviceRegistryResponse = {
  encode(_: MsgUpdateDeviceRegistryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDeviceRegistryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDeviceRegistryResponse();
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

  fromJSON(_: any): MsgUpdateDeviceRegistryResponse {
    return {};
  },

  toJSON(_: MsgUpdateDeviceRegistryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDeviceRegistryResponse>, I>>(_: I): MsgUpdateDeviceRegistryResponse {
    const message = createBaseMsgUpdateDeviceRegistryResponse();
    return message;
  },
};

function createBaseMsgDeleteDeviceRegistry(): MsgDeleteDeviceRegistry {
  return { creator: "", mid: "" };
}

export const MsgDeleteDeviceRegistry = {
  encode(message: MsgDeleteDeviceRegistry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.mid !== "") {
      writer.uint32(18).string(message.mid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDeviceRegistry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDeviceRegistry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.mid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteDeviceRegistry {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      mid: isSet(object.mid) ? String(object.mid) : "",
    };
  },

  toJSON(message: MsgDeleteDeviceRegistry): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.mid !== undefined && (obj.mid = message.mid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDeviceRegistry>, I>>(object: I): MsgDeleteDeviceRegistry {
    const message = createBaseMsgDeleteDeviceRegistry();
    message.creator = object.creator ?? "";
    message.mid = object.mid ?? "";
    return message;
  },
};

function createBaseMsgDeleteDeviceRegistryResponse(): MsgDeleteDeviceRegistryResponse {
  return {};
}

export const MsgDeleteDeviceRegistryResponse = {
  encode(_: MsgDeleteDeviceRegistryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteDeviceRegistryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteDeviceRegistryResponse();
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

  fromJSON(_: any): MsgDeleteDeviceRegistryResponse {
    return {};
  },

  toJSON(_: MsgDeleteDeviceRegistryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteDeviceRegistryResponse>, I>>(_: I): MsgDeleteDeviceRegistryResponse {
    const message = createBaseMsgDeleteDeviceRegistryResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateDevice(request: MsgCreateDevice): Promise<MsgCreateDeviceResponse>;
  UpdateDevice(request: MsgUpdateDevice): Promise<MsgUpdateDeviceResponse>;
  DeleteDevice(request: MsgDeleteDevice): Promise<MsgDeleteDeviceResponse>;
  CreateKv(request: MsgCreateKv): Promise<MsgCreateKvResponse>;
  UpdateKv(request: MsgUpdateKv): Promise<MsgUpdateKvResponse>;
  DeleteKv(request: MsgDeleteKv): Promise<MsgDeleteKvResponse>;
  CreateEventPb(request: MsgCreateEventPb): Promise<MsgCreateEventPbResponse>;
  UpdateEventPb(request: MsgUpdateEventPb): Promise<MsgUpdateEventPbResponse>;
  DeleteEventPb(request: MsgDeleteEventPb): Promise<MsgDeleteEventPbResponse>;
  CreateDeviceRegistry(request: MsgCreateDeviceRegistry): Promise<MsgCreateDeviceRegistryResponse>;
  UpdateDeviceRegistry(request: MsgUpdateDeviceRegistry): Promise<MsgUpdateDeviceRegistryResponse>;
  DeleteDeviceRegistry(request: MsgDeleteDeviceRegistry): Promise<MsgDeleteDeviceRegistryResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateDevice = this.CreateDevice.bind(this);
    this.UpdateDevice = this.UpdateDevice.bind(this);
    this.DeleteDevice = this.DeleteDevice.bind(this);
    this.CreateKv = this.CreateKv.bind(this);
    this.UpdateKv = this.UpdateKv.bind(this);
    this.DeleteKv = this.DeleteKv.bind(this);
    this.CreateEventPb = this.CreateEventPb.bind(this);
    this.UpdateEventPb = this.UpdateEventPb.bind(this);
    this.DeleteEventPb = this.DeleteEventPb.bind(this);
    this.CreateDeviceRegistry = this.CreateDeviceRegistry.bind(this);
    this.UpdateDeviceRegistry = this.UpdateDeviceRegistry.bind(this);
    this.DeleteDeviceRegistry = this.DeleteDeviceRegistry.bind(this);
  }
  CreateDevice(request: MsgCreateDevice): Promise<MsgCreateDeviceResponse> {
    const data = MsgCreateDevice.encode(request).finish();
    const promise = this.rpc.request("stccommunity.iotdepinprotocol.iotdepinprotocol.Msg", "CreateDevice", data);
    return promise.then((data) => MsgCreateDeviceResponse.decode(new _m0.Reader(data)));
  }

  UpdateDevice(request: MsgUpdateDevice): Promise<MsgUpdateDeviceResponse> {
    const data = MsgUpdateDevice.encode(request).finish();
    const promise = this.rpc.request("stccommunity.iotdepinprotocol.iotdepinprotocol.Msg", "UpdateDevice", data);
    return promise.then((data) => MsgUpdateDeviceResponse.decode(new _m0.Reader(data)));
  }

  DeleteDevice(request: MsgDeleteDevice): Promise<MsgDeleteDeviceResponse> {
    const data = MsgDeleteDevice.encode(request).finish();
    const promise = this.rpc.request("stccommunity.iotdepinprotocol.iotdepinprotocol.Msg", "DeleteDevice", data);
    return promise.then((data) => MsgDeleteDeviceResponse.decode(new _m0.Reader(data)));
  }

  CreateKv(request: MsgCreateKv): Promise<MsgCreateKvResponse> {
    const data = MsgCreateKv.encode(request).finish();
    const promise = this.rpc.request("stccommunity.iotdepinprotocol.iotdepinprotocol.Msg", "CreateKv", data);
    return promise.then((data) => MsgCreateKvResponse.decode(new _m0.Reader(data)));
  }

  UpdateKv(request: MsgUpdateKv): Promise<MsgUpdateKvResponse> {
    const data = MsgUpdateKv.encode(request).finish();
    const promise = this.rpc.request("stccommunity.iotdepinprotocol.iotdepinprotocol.Msg", "UpdateKv", data);
    return promise.then((data) => MsgUpdateKvResponse.decode(new _m0.Reader(data)));
  }

  DeleteKv(request: MsgDeleteKv): Promise<MsgDeleteKvResponse> {
    const data = MsgDeleteKv.encode(request).finish();
    const promise = this.rpc.request("stccommunity.iotdepinprotocol.iotdepinprotocol.Msg", "DeleteKv", data);
    return promise.then((data) => MsgDeleteKvResponse.decode(new _m0.Reader(data)));
  }

  CreateEventPb(request: MsgCreateEventPb): Promise<MsgCreateEventPbResponse> {
    const data = MsgCreateEventPb.encode(request).finish();
    const promise = this.rpc.request("stccommunity.iotdepinprotocol.iotdepinprotocol.Msg", "CreateEventPb", data);
    return promise.then((data) => MsgCreateEventPbResponse.decode(new _m0.Reader(data)));
  }

  UpdateEventPb(request: MsgUpdateEventPb): Promise<MsgUpdateEventPbResponse> {
    const data = MsgUpdateEventPb.encode(request).finish();
    const promise = this.rpc.request("stccommunity.iotdepinprotocol.iotdepinprotocol.Msg", "UpdateEventPb", data);
    return promise.then((data) => MsgUpdateEventPbResponse.decode(new _m0.Reader(data)));
  }

  DeleteEventPb(request: MsgDeleteEventPb): Promise<MsgDeleteEventPbResponse> {
    const data = MsgDeleteEventPb.encode(request).finish();
    const promise = this.rpc.request("stccommunity.iotdepinprotocol.iotdepinprotocol.Msg", "DeleteEventPb", data);
    return promise.then((data) => MsgDeleteEventPbResponse.decode(new _m0.Reader(data)));
  }

  CreateDeviceRegistry(request: MsgCreateDeviceRegistry): Promise<MsgCreateDeviceRegistryResponse> {
    const data = MsgCreateDeviceRegistry.encode(request).finish();
    const promise = this.rpc.request(
      "stccommunity.iotdepinprotocol.iotdepinprotocol.Msg",
      "CreateDeviceRegistry",
      data,
    );
    return promise.then((data) => MsgCreateDeviceRegistryResponse.decode(new _m0.Reader(data)));
  }

  UpdateDeviceRegistry(request: MsgUpdateDeviceRegistry): Promise<MsgUpdateDeviceRegistryResponse> {
    const data = MsgUpdateDeviceRegistry.encode(request).finish();
    const promise = this.rpc.request(
      "stccommunity.iotdepinprotocol.iotdepinprotocol.Msg",
      "UpdateDeviceRegistry",
      data,
    );
    return promise.then((data) => MsgUpdateDeviceRegistryResponse.decode(new _m0.Reader(data)));
  }

  DeleteDeviceRegistry(request: MsgDeleteDeviceRegistry): Promise<MsgDeleteDeviceRegistryResponse> {
    const data = MsgDeleteDeviceRegistry.encode(request).finish();
    const promise = this.rpc.request(
      "stccommunity.iotdepinprotocol.iotdepinprotocol.Msg",
      "DeleteDeviceRegistry",
      data,
    );
    return promise.then((data) => MsgDeleteDeviceRegistryResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
