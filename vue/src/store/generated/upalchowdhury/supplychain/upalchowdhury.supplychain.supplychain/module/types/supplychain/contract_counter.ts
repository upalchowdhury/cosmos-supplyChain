/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "upalchowdhury.supplychain.supplychain";

export interface ContractCounter {
  dealId: string;
  idValue: number;
}

const baseContractCounter: object = { dealId: "", idValue: 0 };

export const ContractCounter = {
  encode(message: ContractCounter, writer: Writer = Writer.create()): Writer {
    if (message.dealId !== "") {
      writer.uint32(10).string(message.dealId);
    }
    if (message.idValue !== 0) {
      writer.uint32(16).uint64(message.idValue);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ContractCounter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseContractCounter } as ContractCounter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.dealId = reader.string();
          break;
        case 2:
          message.idValue = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ContractCounter {
    const message = { ...baseContractCounter } as ContractCounter;
    if (object.dealId !== undefined && object.dealId !== null) {
      message.dealId = String(object.dealId);
    } else {
      message.dealId = "";
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = Number(object.idValue);
    } else {
      message.idValue = 0;
    }
    return message;
  },

  toJSON(message: ContractCounter): unknown {
    const obj: any = {};
    message.dealId !== undefined && (obj.dealId = message.dealId);
    message.idValue !== undefined && (obj.idValue = message.idValue);
    return obj;
  },

  fromPartial(object: DeepPartial<ContractCounter>): ContractCounter {
    const message = { ...baseContractCounter } as ContractCounter;
    if (object.dealId !== undefined && object.dealId !== null) {
      message.dealId = object.dealId;
    } else {
      message.dealId = "";
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue;
    } else {
      message.idValue = 0;
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
