/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "upalchowdhury.supplychain.supplychain";

export interface Actorscontract {
  index: string;
  contractId: string;
  owner: string;
  vendor: string;
  commission: number;
}

const baseActorscontract: object = {
  index: "",
  contractId: "",
  owner: "",
  vendor: "",
  commission: 0,
};

export const Actorscontract = {
  encode(message: Actorscontract, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.contractId !== "") {
      writer.uint32(18).string(message.contractId);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    if (message.vendor !== "") {
      writer.uint32(34).string(message.vendor);
    }
    if (message.commission !== 0) {
      writer.uint32(40).uint64(message.commission);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Actorscontract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseActorscontract } as Actorscontract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.contractId = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        case 4:
          message.vendor = reader.string();
          break;
        case 5:
          message.commission = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Actorscontract {
    const message = { ...baseActorscontract } as Actorscontract;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = String(object.contractId);
    } else {
      message.contractId = "";
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

  toJSON(message: Actorscontract): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.contractId !== undefined && (obj.contractId = message.contractId);
    message.owner !== undefined && (obj.owner = message.owner);
    message.vendor !== undefined && (obj.vendor = message.vendor);
    message.commission !== undefined && (obj.commission = message.commission);
    return obj;
  },

  fromPartial(object: DeepPartial<Actorscontract>): Actorscontract {
    const message = { ...baseActorscontract } as Actorscontract;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = object.contractId;
    } else {
      message.contractId = "";
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
