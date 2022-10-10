/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "upalchowdhury.supplychain.supplychain";

export interface Parties {
  Id: string;
  owner: string;
  vendor: string;
  commission: number;
}

const baseParties: object = { Id: "", owner: "", vendor: "", commission: 0 };

export const Parties = {
  encode(message: Parties, writer: Writer = Writer.create()): Writer {
    if (message.Id !== "") {
      writer.uint32(10).string(message.Id);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.vendor !== "") {
      writer.uint32(26).string(message.vendor);
    }
    if (message.commission !== 0) {
      writer.uint32(32).uint64(message.commission);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Parties {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParties } as Parties;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Id = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.vendor = reader.string();
          break;
        case 4:
          message.commission = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Parties {
    const message = { ...baseParties } as Parties;
    if (object.Id !== undefined && object.Id !== null) {
      message.Id = String(object.Id);
    } else {
      message.Id = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.vendor !== undefined && object.vendor !== null) {
      message.vendor = String(object.vendor);
    } else {
      message.vendor = "";
    }
    if (object.commission !== undefined && object.commission !== null) {
      message.commission = Number(object.commission);
    } else {
      message.commission = 0;
    }
    return message;
  },

  toJSON(message: Parties): unknown {
    const obj: any = {};
    message.Id !== undefined && (obj.Id = message.Id);
    message.owner !== undefined && (obj.owner = message.owner);
    message.vendor !== undefined && (obj.vendor = message.vendor);
    message.commission !== undefined && (obj.commission = message.commission);
    return obj;
  },

  fromPartial(object: DeepPartial<Parties>): Parties {
    const message = { ...baseParties } as Parties;
    if (object.Id !== undefined && object.Id !== null) {
      message.Id = object.Id;
    } else {
      message.Id = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.vendor !== undefined && object.vendor !== null) {
      message.vendor = object.vendor;
    } else {
      message.vendor = "";
    }
    if (object.commission !== undefined && object.commission !== null) {
      message.commission = object.commission;
    } else {
      message.commission = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
