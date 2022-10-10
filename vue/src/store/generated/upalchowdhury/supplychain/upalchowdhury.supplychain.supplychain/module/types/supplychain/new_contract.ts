/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "upalchowdhury.supplychain.supplychain";

export interface NewContract {
  dealId: string;
  contractId: string;
  consumer: string;
  desc: string;
  ownerETA: number;
  vendorETA: number;
  status: string;
  fees: number;
  expiry: string;
  shippingDelay: number;
  startTime: string;
  deliveryDelay: number;
}

const baseNewContract: object = {
  dealId: "",
  contractId: "",
  consumer: "",
  desc: "",
  ownerETA: 0,
  vendorETA: 0,
  status: "",
  fees: 0,
  expiry: "",
  shippingDelay: 0,
  startTime: "",
  deliveryDelay: 0,
};

export const NewContract = {
  encode(message: NewContract, writer: Writer = Writer.create()): Writer {
    if (message.dealId !== "") {
      writer.uint32(10).string(message.dealId);
    }
    if (message.contractId !== "") {
      writer.uint32(18).string(message.contractId);
    }
    if (message.consumer !== "") {
      writer.uint32(26).string(message.consumer);
    }
    if (message.desc !== "") {
      writer.uint32(34).string(message.desc);
    }
    if (message.ownerETA !== 0) {
      writer.uint32(40).uint32(message.ownerETA);
    }
    if (message.vendorETA !== 0) {
      writer.uint32(48).uint32(message.vendorETA);
    }
    if (message.status !== "") {
      writer.uint32(58).string(message.status);
    }
    if (message.fees !== 0) {
      writer.uint32(64).uint64(message.fees);
    }
    if (message.expiry !== "") {
      writer.uint32(74).string(message.expiry);
    }
    if (message.shippingDelay !== 0) {
      writer.uint32(80).uint32(message.shippingDelay);
    }
    if (message.startTime !== "") {
      writer.uint32(90).string(message.startTime);
    }
    if (message.deliveryDelay !== 0) {
      writer.uint32(96).uint32(message.deliveryDelay);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NewContract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNewContract } as NewContract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.dealId = reader.string();
          break;
        case 2:
          message.contractId = reader.string();
          break;
        case 3:
          message.consumer = reader.string();
          break;
        case 4:
          message.desc = reader.string();
          break;
        case 5:
          message.ownerETA = reader.uint32();
          break;
        case 6:
          message.vendorETA = reader.uint32();
          break;
        case 7:
          message.status = reader.string();
          break;
        case 8:
          message.fees = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.expiry = reader.string();
          break;
        case 10:
          message.shippingDelay = reader.uint32();
          break;
        case 11:
          message.startTime = reader.string();
          break;
        case 12:
          message.deliveryDelay = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NewContract {
    const message = { ...baseNewContract } as NewContract;
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
    if (object.ownerETA !== undefined && object.ownerETA !== null) {
      message.ownerETA = Number(object.ownerETA);
    } else {
      message.ownerETA = 0;
    }
    if (object.vendorETA !== undefined && object.vendorETA !== null) {
      message.vendorETA = Number(object.vendorETA);
    } else {
      message.vendorETA = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.fees !== undefined && object.fees !== null) {
      message.fees = Number(object.fees);
    } else {
      message.fees = 0;
    }
    if (object.expiry !== undefined && object.expiry !== null) {
      message.expiry = String(object.expiry);
    } else {
      message.expiry = "";
    }
    if (object.shippingDelay !== undefined && object.shippingDelay !== null) {
      message.shippingDelay = Number(object.shippingDelay);
    } else {
      message.shippingDelay = 0;
    }
    if (object.startTime !== undefined && object.startTime !== null) {
      message.startTime = String(object.startTime);
    } else {
      message.startTime = "";
    }
    if (object.deliveryDelay !== undefined && object.deliveryDelay !== null) {
      message.deliveryDelay = Number(object.deliveryDelay);
    } else {
      message.deliveryDelay = 0;
    }
    return message;
  },

  toJSON(message: NewContract): unknown {
    const obj: any = {};
    message.dealId !== undefined && (obj.dealId = message.dealId);
    message.contractId !== undefined && (obj.contractId = message.contractId);
    message.consumer !== undefined && (obj.consumer = message.consumer);
    message.desc !== undefined && (obj.desc = message.desc);
    message.ownerETA !== undefined && (obj.ownerETA = message.ownerETA);
    message.vendorETA !== undefined && (obj.vendorETA = message.vendorETA);
    message.status !== undefined && (obj.status = message.status);
    message.fees !== undefined && (obj.fees = message.fees);
    message.expiry !== undefined && (obj.expiry = message.expiry);
    message.shippingDelay !== undefined &&
      (obj.shippingDelay = message.shippingDelay);
    message.startTime !== undefined && (obj.startTime = message.startTime);
    message.deliveryDelay !== undefined &&
      (obj.deliveryDelay = message.deliveryDelay);
    return obj;
  },

  fromPartial(object: DeepPartial<NewContract>): NewContract {
    const message = { ...baseNewContract } as NewContract;
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
    if (object.ownerETA !== undefined && object.ownerETA !== null) {
      message.ownerETA = object.ownerETA;
    } else {
      message.ownerETA = 0;
    }
    if (object.vendorETA !== undefined && object.vendorETA !== null) {
      message.vendorETA = object.vendorETA;
    } else {
      message.vendorETA = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.fees !== undefined && object.fees !== null) {
      message.fees = object.fees;
    } else {
      message.fees = 0;
    }
    if (object.expiry !== undefined && object.expiry !== null) {
      message.expiry = object.expiry;
    } else {
      message.expiry = "";
    }
    if (object.shippingDelay !== undefined && object.shippingDelay !== null) {
      message.shippingDelay = object.shippingDelay;
    } else {
      message.shippingDelay = 0;
    }
    if (object.startTime !== undefined && object.startTime !== null) {
      message.startTime = object.startTime;
    } else {
      message.startTime = "";
    }
    if (object.deliveryDelay !== undefined && object.deliveryDelay !== null) {
      message.deliveryDelay = object.deliveryDelay;
    } else {
      message.deliveryDelay = 0;
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
