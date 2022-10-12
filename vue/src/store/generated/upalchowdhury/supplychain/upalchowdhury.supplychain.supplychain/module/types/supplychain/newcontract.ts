/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "upalchowdhury.supplychain.supplychain";

export interface Newcontract {
  index: string;
  dealId: string;
  contractId: string;
  consumer: string;
  desc: string;
  ownereta: number;
  vendoreta: number;
  status: string;
  fees: string;
  expiry: string;
  shippingdelay: number;
  starttime: string;
  deliverydelay: number;
}

const baseNewcontract: object = {
  index: "",
  dealId: "",
  contractId: "",
  consumer: "",
  desc: "",
  ownereta: 0,
  vendoreta: 0,
  status: "",
  fees: "",
  expiry: "",
  shippingdelay: 0,
  starttime: "",
  deliverydelay: 0,
};

export const Newcontract = {
  encode(message: Newcontract, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.dealId !== "") {
      writer.uint32(18).string(message.dealId);
    }
    if (message.contractId !== "") {
      writer.uint32(26).string(message.contractId);
    }
    if (message.consumer !== "") {
      writer.uint32(34).string(message.consumer);
    }
    if (message.desc !== "") {
      writer.uint32(42).string(message.desc);
    }
    if (message.ownereta !== 0) {
      writer.uint32(48).uint64(message.ownereta);
    }
    if (message.vendoreta !== 0) {
      writer.uint32(56).uint64(message.vendoreta);
    }
    if (message.status !== "") {
      writer.uint32(66).string(message.status);
    }
    if (message.fees !== "") {
      writer.uint32(74).string(message.fees);
    }
    if (message.expiry !== "") {
      writer.uint32(82).string(message.expiry);
    }
    if (message.shippingdelay !== 0) {
      writer.uint32(88).uint64(message.shippingdelay);
    }
    if (message.starttime !== "") {
      writer.uint32(98).string(message.starttime);
    }
    if (message.deliverydelay !== 0) {
      writer.uint32(104).uint64(message.deliverydelay);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Newcontract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNewcontract } as Newcontract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.dealId = reader.string();
          break;
        case 3:
          message.contractId = reader.string();
          break;
        case 4:
          message.consumer = reader.string();
          break;
        case 5:
          message.desc = reader.string();
          break;
        case 6:
          message.ownereta = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.vendoreta = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.status = reader.string();
          break;
        case 9:
          message.fees = reader.string();
          break;
        case 10:
          message.expiry = reader.string();
          break;
        case 11:
          message.shippingdelay = longToNumber(reader.uint64() as Long);
          break;
        case 12:
          message.starttime = reader.string();
          break;
        case 13:
          message.deliverydelay = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Newcontract {
    const message = { ...baseNewcontract } as Newcontract;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.dealId !== undefined && object.dealId !== null) {
      message.dealId = String(object.dealId);
    } else {
      message.dealId = "";
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = String(object.contractId);
    } else {
      message.contractId = "";
    }
    if (object.consumer !== undefined && object.consumer !== null) {
      message.consumer = String(object.consumer);
    } else {
      message.consumer = "";
    }
    if (object.desc !== undefined && object.desc !== null) {
      message.desc = String(object.desc);
    } else {
      message.desc = "";
    }
    if (object.ownereta !== undefined && object.ownereta !== null) {
      message.ownereta = Number(object.ownereta);
    } else {
      message.ownereta = 0;
    }
    if (object.vendoreta !== undefined && object.vendoreta !== null) {
      message.vendoreta = Number(object.vendoreta);
    } else {
      message.vendoreta = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.fees !== undefined && object.fees !== null) {
      message.fees = String(object.fees);
    } else {
      message.fees = "";
    }
    if (object.expiry !== undefined && object.expiry !== null) {
      message.expiry = String(object.expiry);
    } else {
      message.expiry = "";
    }
    if (object.shippingdelay !== undefined && object.shippingdelay !== null) {
      message.shippingdelay = Number(object.shippingdelay);
    } else {
      message.shippingdelay = 0;
    }
    if (object.starttime !== undefined && object.starttime !== null) {
      message.starttime = String(object.starttime);
    } else {
      message.starttime = "";
    }
    if (object.deliverydelay !== undefined && object.deliverydelay !== null) {
      message.deliverydelay = Number(object.deliverydelay);
    } else {
      message.deliverydelay = 0;
    }
    return message;
  },

  toJSON(message: Newcontract): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.dealId !== undefined && (obj.dealId = message.dealId);
    message.contractId !== undefined && (obj.contractId = message.contractId);
    message.consumer !== undefined && (obj.consumer = message.consumer);
    message.desc !== undefined && (obj.desc = message.desc);
    message.ownereta !== undefined && (obj.ownereta = message.ownereta);
    message.vendoreta !== undefined && (obj.vendoreta = message.vendoreta);
    message.status !== undefined && (obj.status = message.status);
    message.fees !== undefined && (obj.fees = message.fees);
    message.expiry !== undefined && (obj.expiry = message.expiry);
    message.shippingdelay !== undefined &&
      (obj.shippingdelay = message.shippingdelay);
    message.starttime !== undefined && (obj.starttime = message.starttime);
    message.deliverydelay !== undefined &&
      (obj.deliverydelay = message.deliverydelay);
    return obj;
  },

  fromPartial(object: DeepPartial<Newcontract>): Newcontract {
    const message = { ...baseNewcontract } as Newcontract;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.dealId !== undefined && object.dealId !== null) {
      message.dealId = object.dealId;
    } else {
      message.dealId = "";
    }
    if (object.contractId !== undefined && object.contractId !== null) {
      message.contractId = object.contractId;
    } else {
      message.contractId = "";
    }
    if (object.consumer !== undefined && object.consumer !== null) {
      message.consumer = object.consumer;
    } else {
      message.consumer = "";
    }
    if (object.desc !== undefined && object.desc !== null) {
      message.desc = object.desc;
    } else {
      message.desc = "";
    }
    if (object.ownereta !== undefined && object.ownereta !== null) {
      message.ownereta = object.ownereta;
    } else {
      message.ownereta = 0;
    }
    if (object.vendoreta !== undefined && object.vendoreta !== null) {
      message.vendoreta = object.vendoreta;
    } else {
      message.vendoreta = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.fees !== undefined && object.fees !== null) {
      message.fees = object.fees;
    } else {
      message.fees = "";
    }
    if (object.expiry !== undefined && object.expiry !== null) {
      message.expiry = object.expiry;
    } else {
      message.expiry = "";
    }
    if (object.shippingdelay !== undefined && object.shippingdelay !== null) {
      message.shippingdelay = object.shippingdelay;
    } else {
      message.shippingdelay = 0;
    }
    if (object.starttime !== undefined && object.starttime !== null) {
      message.starttime = object.starttime;
    } else {
      message.starttime = "";
    }
    if (object.deliverydelay !== undefined && object.deliverydelay !== null) {
      message.deliverydelay = object.deliverydelay;
    } else {
      message.deliverydelay = 0;
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
