/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "stccommunity.iotdepinprotocol.iotdepinprotocol";

export interface DeviceRegistry {
  mid: string;
  metaData: string;
  creator: string;
}

function createBaseDeviceRegistry(): DeviceRegistry {
  return { mid: "", metaData: "", creator: "" };
}

export const DeviceRegistry = {
  encode(message: DeviceRegistry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.mid !== "") {
      writer.uint32(10).string(message.mid);
    }
    if (message.metaData !== "") {
      writer.uint32(18).string(message.metaData);
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeviceRegistry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeviceRegistry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mid = reader.string();
          break;
        case 2:
          message.metaData = reader.string();
          break;
        case 3:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeviceRegistry {
    return {
      mid: isSet(object.mid) ? String(object.mid) : "",
      metaData: isSet(object.metaData) ? String(object.metaData) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: DeviceRegistry): unknown {
    const obj: any = {};
    message.mid !== undefined && (obj.mid = message.mid);
    message.metaData !== undefined && (obj.metaData = message.metaData);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeviceRegistry>, I>>(object: I): DeviceRegistry {
    const message = createBaseDeviceRegistry();
    message.mid = object.mid ?? "";
    message.metaData = object.metaData ?? "";
    message.creator = object.creator ?? "";
    return message;
  },
};

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
