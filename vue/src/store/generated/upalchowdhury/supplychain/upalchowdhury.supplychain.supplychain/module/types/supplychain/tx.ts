/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "upalchowdhury.supplychain.supplychain";

export interface MsgCreateTran {
  creator: string;
  vendor: string;
  commission: number;
}

export interface MsgCreateTranResponse {
  idValue: string;
}

export interface MsgCreateContract {
  creator: string;
  dealId: string;
  consumer: string;
  desc: string;
  ownereta: string;
  expiry: string;
  fees: number;
}

export interface MsgCreateContractResponse {
  idValue: string;
}

const baseMsgCreateTran: object = { creator: "", vendor: "", commission: 0 };

export const MsgCreateTran = {
  encode(message: MsgCreateTran, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.vendor !== "") {
      writer.uint32(18).string(message.vendor);
    }
    if (message.commission !== 0) {
      writer.uint32(24).uint64(message.commission);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTran {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTran } as MsgCreateTran;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.vendor = reader.string();
          break;
        case 3:
          message.commission = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTran {
    const message = { ...baseMsgCreateTran } as MsgCreateTran;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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

  toJSON(message: MsgCreateTran): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.vendor !== undefined && (obj.vendor = message.vendor);
    message.commission !== undefined && (obj.commission = message.commission);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateTran>): MsgCreateTran {
    const message = { ...baseMsgCreateTran } as MsgCreateTran;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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

const baseMsgCreateTranResponse: object = { idValue: "" };

export const MsgCreateTranResponse = {
  encode(
    message: MsgCreateTranResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.idValue !== "") {
      writer.uint32(10).string(message.idValue);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTranResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTranResponse } as MsgCreateTranResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.idValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTranResponse {
    const message = { ...baseMsgCreateTranResponse } as MsgCreateTranResponse;
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue);
    } else {
      message.idValue = "";
    }
    return message;
  },

  toJSON(message: MsgCreateTranResponse): unknown {
    const obj: any = {};
    message.idValue !== undefined && (obj.idValue = message.idValue);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateTranResponse>
  ): MsgCreateTranResponse {
    const message = { ...baseMsgCreateTranResponse } as MsgCreateTranResponse;
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue;
    } else {
      message.idValue = "";
    }
    return message;
  },
};

const baseMsgCreateContract: object = {
  creator: "",
  dealId: "",
  consumer: "",
  desc: "",
  ownereta: "",
  expiry: "",
  fees: 0,
};

export const MsgCreateContract = {
  encode(message: MsgCreateContract, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.dealId !== "") {
      writer.uint32(18).string(message.dealId);
    }
    if (message.consumer !== "") {
      writer.uint32(26).string(message.consumer);
    }
    if (message.desc !== "") {
      writer.uint32(34).string(message.desc);
    }
    if (message.ownereta !== "") {
      writer.uint32(42).string(message.ownereta);
    }
    if (message.expiry !== "") {
      writer.uint32(50).string(message.expiry);
    }
    if (message.fees !== 0) {
      writer.uint32(56).uint64(message.fees);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateContract {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateContract } as MsgCreateContract;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.dealId = reader.string();
          break;
        case 3:
          message.consumer = reader.string();
          break;
        case 4:
          message.desc = reader.string();
          break;
        case 5:
          message.ownereta = reader.string();
          break;
        case 6:
          message.expiry = reader.string();
          break;
        case 7:
          message.fees = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateContract {
    const message = { ...baseMsgCreateContract } as MsgCreateContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.dealId !== undefined && object.dealId !== null) {
      message.dealId = String(object.dealId);
    } else {
      message.dealId = "";
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
      message.ownereta = String(object.ownereta);
    } else {
      message.ownereta = "";
    }
    if (object.expiry !== undefined && object.expiry !== null) {
      message.expiry = String(object.expiry);
    } else {
      message.expiry = "";
    }
    if (object.fees !== undefined && object.fees !== null) {
      message.fees = Number(object.fees);
    } else {
      message.fees = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateContract): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.dealId !== undefined && (obj.dealId = message.dealId);
    message.consumer !== undefined && (obj.consumer = message.consumer);
    message.desc !== undefined && (obj.desc = message.desc);
    message.ownereta !== undefined && (obj.ownereta = message.ownereta);
    message.expiry !== undefined && (obj.expiry = message.expiry);
    message.fees !== undefined && (obj.fees = message.fees);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateContract>): MsgCreateContract {
    const message = { ...baseMsgCreateContract } as MsgCreateContract;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.dealId !== undefined && object.dealId !== null) {
      message.dealId = object.dealId;
    } else {
      message.dealId = "";
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
      message.ownereta = "";
    }
    if (object.expiry !== undefined && object.expiry !== null) {
      message.expiry = object.expiry;
    } else {
      message.expiry = "";
    }
    if (object.fees !== undefined && object.fees !== null) {
      message.fees = object.fees;
    } else {
      message.fees = 0;
    }
    return message;
  },
};

const baseMsgCreateContractResponse: object = { idValue: "" };

export const MsgCreateContractResponse = {
  encode(
    message: MsgCreateContractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.idValue !== "") {
      writer.uint32(10).string(message.idValue);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateContractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateContractResponse,
    } as MsgCreateContractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.idValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateContractResponse {
    const message = {
      ...baseMsgCreateContractResponse,
    } as MsgCreateContractResponse;
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = String(object.idValue);
    } else {
      message.idValue = "";
    }
    return message;
  },

  toJSON(message: MsgCreateContractResponse): unknown {
    const obj: any = {};
    message.idValue !== undefined && (obj.idValue = message.idValue);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateContractResponse>
  ): MsgCreateContractResponse {
    const message = {
      ...baseMsgCreateContractResponse,
    } as MsgCreateContractResponse;
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue;
    } else {
      message.idValue = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateTran(request: MsgCreateTran): Promise<MsgCreateTranResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateContract(
    request: MsgCreateContract
  ): Promise<MsgCreateContractResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateTran(request: MsgCreateTran): Promise<MsgCreateTranResponse> {
    const data = MsgCreateTran.encode(request).finish();
    const promise = this.rpc.request(
      "upalchowdhury.supplychain.supplychain.Msg",
      "CreateTran",
      data
    );
    return promise.then((data) =>
      MsgCreateTranResponse.decode(new Reader(data))
    );
  }

  CreateContract(
    request: MsgCreateContract
  ): Promise<MsgCreateContractResponse> {
    const data = MsgCreateContract.encode(request).finish();
    const promise = this.rpc.request(
      "upalchowdhury.supplychain.supplychain.Msg",
      "CreateContract",
      data
    );
    return promise.then((data) =>
      MsgCreateContractResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
