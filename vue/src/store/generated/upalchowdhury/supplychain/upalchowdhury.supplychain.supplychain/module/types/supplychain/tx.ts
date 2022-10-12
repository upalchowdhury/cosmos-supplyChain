/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "upalchowdhury.supplychain.supplychain";

export interface MsgCreateTran {
  creator: string;
  vendor: string;
  commission: string;
}

export interface MsgCreateTranResponse {
  idValue: string;
}

const baseMsgCreateTran: object = { creator: "", vendor: "", commission: "" };

export const MsgCreateTran = {
  encode(message: MsgCreateTran, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.vendor !== "") {
      writer.uint32(18).string(message.vendor);
    }
    if (message.commission !== "") {
      writer.uint32(26).string(message.commission);
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
          message.commission = reader.string();
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
      message.commission = String(object.commission);
    } else {
      message.commission = "";
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
      message.commission = "";
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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateTran(request: MsgCreateTran): Promise<MsgCreateTranResponse>;
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
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
