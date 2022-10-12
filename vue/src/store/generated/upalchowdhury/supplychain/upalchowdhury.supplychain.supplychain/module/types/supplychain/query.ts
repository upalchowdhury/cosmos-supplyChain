/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../supplychain/params";
import { ConCounter } from "../supplychain/con_counter";
import { Actorscontract } from "../supplychain/actorscontract";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { ContractCounter } from "../supplychain/contract_counter";
import { Newcontract } from "../supplychain/newcontract";

export const protobufPackage = "upalchowdhury.supplychain.supplychain";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetConCounterRequest {}

export interface QueryGetConCounterResponse {
  ConCounter: ConCounter | undefined;
}

export interface QueryGetActorscontractRequest {
  index: string;
}

export interface QueryGetActorscontractResponse {
  actorscontract: Actorscontract | undefined;
}

export interface QueryAllActorscontractRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllActorscontractResponse {
  actorscontract: Actorscontract[];
  pagination: PageResponse | undefined;
}

export interface QueryGetContractCounterRequest {}

export interface QueryGetContractCounterResponse {
  ContractCounter: ContractCounter | undefined;
}

export interface QueryGetNewcontractRequest {
  index: string;
}

export interface QueryGetNewcontractResponse {
  newcontract: Newcontract | undefined;
}

export interface QueryAllNewcontractRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNewcontractResponse {
  newcontract: Newcontract[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetConCounterRequest: object = {};

export const QueryGetConCounterRequest = {
  encode(
    _: QueryGetConCounterRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetConCounterRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetConCounterRequest,
    } as QueryGetConCounterRequest;
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

  fromJSON(_: any): QueryGetConCounterRequest {
    const message = {
      ...baseQueryGetConCounterRequest,
    } as QueryGetConCounterRequest;
    return message;
  },

  toJSON(_: QueryGetConCounterRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetConCounterRequest>
  ): QueryGetConCounterRequest {
    const message = {
      ...baseQueryGetConCounterRequest,
    } as QueryGetConCounterRequest;
    return message;
  },
};

const baseQueryGetConCounterResponse: object = {};

export const QueryGetConCounterResponse = {
  encode(
    message: QueryGetConCounterResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ConCounter !== undefined) {
      ConCounter.encode(message.ConCounter, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetConCounterResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetConCounterResponse,
    } as QueryGetConCounterResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ConCounter = ConCounter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetConCounterResponse {
    const message = {
      ...baseQueryGetConCounterResponse,
    } as QueryGetConCounterResponse;
    if (object.ConCounter !== undefined && object.ConCounter !== null) {
      message.ConCounter = ConCounter.fromJSON(object.ConCounter);
    } else {
      message.ConCounter = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetConCounterResponse): unknown {
    const obj: any = {};
    message.ConCounter !== undefined &&
      (obj.ConCounter = message.ConCounter
        ? ConCounter.toJSON(message.ConCounter)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetConCounterResponse>
  ): QueryGetConCounterResponse {
    const message = {
      ...baseQueryGetConCounterResponse,
    } as QueryGetConCounterResponse;
    if (object.ConCounter !== undefined && object.ConCounter !== null) {
      message.ConCounter = ConCounter.fromPartial(object.ConCounter);
    } else {
      message.ConCounter = undefined;
    }
    return message;
  },
};

const baseQueryGetActorscontractRequest: object = { index: "" };

export const QueryGetActorscontractRequest = {
  encode(
    message: QueryGetActorscontractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetActorscontractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetActorscontractRequest,
    } as QueryGetActorscontractRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetActorscontractRequest {
    const message = {
      ...baseQueryGetActorscontractRequest,
    } as QueryGetActorscontractRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetActorscontractRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetActorscontractRequest>
  ): QueryGetActorscontractRequest {
    const message = {
      ...baseQueryGetActorscontractRequest,
    } as QueryGetActorscontractRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetActorscontractResponse: object = {};

export const QueryGetActorscontractResponse = {
  encode(
    message: QueryGetActorscontractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.actorscontract !== undefined) {
      Actorscontract.encode(
        message.actorscontract,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetActorscontractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetActorscontractResponse,
    } as QueryGetActorscontractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.actorscontract = Actorscontract.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetActorscontractResponse {
    const message = {
      ...baseQueryGetActorscontractResponse,
    } as QueryGetActorscontractResponse;
    if (object.actorscontract !== undefined && object.actorscontract !== null) {
      message.actorscontract = Actorscontract.fromJSON(object.actorscontract);
    } else {
      message.actorscontract = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetActorscontractResponse): unknown {
    const obj: any = {};
    message.actorscontract !== undefined &&
      (obj.actorscontract = message.actorscontract
        ? Actorscontract.toJSON(message.actorscontract)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetActorscontractResponse>
  ): QueryGetActorscontractResponse {
    const message = {
      ...baseQueryGetActorscontractResponse,
    } as QueryGetActorscontractResponse;
    if (object.actorscontract !== undefined && object.actorscontract !== null) {
      message.actorscontract = Actorscontract.fromPartial(
        object.actorscontract
      );
    } else {
      message.actorscontract = undefined;
    }
    return message;
  },
};

const baseQueryAllActorscontractRequest: object = {};

export const QueryAllActorscontractRequest = {
  encode(
    message: QueryAllActorscontractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllActorscontractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllActorscontractRequest,
    } as QueryAllActorscontractRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllActorscontractRequest {
    const message = {
      ...baseQueryAllActorscontractRequest,
    } as QueryAllActorscontractRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllActorscontractRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllActorscontractRequest>
  ): QueryAllActorscontractRequest {
    const message = {
      ...baseQueryAllActorscontractRequest,
    } as QueryAllActorscontractRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllActorscontractResponse: object = {};

export const QueryAllActorscontractResponse = {
  encode(
    message: QueryAllActorscontractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.actorscontract) {
      Actorscontract.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllActorscontractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllActorscontractResponse,
    } as QueryAllActorscontractResponse;
    message.actorscontract = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.actorscontract.push(
            Actorscontract.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllActorscontractResponse {
    const message = {
      ...baseQueryAllActorscontractResponse,
    } as QueryAllActorscontractResponse;
    message.actorscontract = [];
    if (object.actorscontract !== undefined && object.actorscontract !== null) {
      for (const e of object.actorscontract) {
        message.actorscontract.push(Actorscontract.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllActorscontractResponse): unknown {
    const obj: any = {};
    if (message.actorscontract) {
      obj.actorscontract = message.actorscontract.map((e) =>
        e ? Actorscontract.toJSON(e) : undefined
      );
    } else {
      obj.actorscontract = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllActorscontractResponse>
  ): QueryAllActorscontractResponse {
    const message = {
      ...baseQueryAllActorscontractResponse,
    } as QueryAllActorscontractResponse;
    message.actorscontract = [];
    if (object.actorscontract !== undefined && object.actorscontract !== null) {
      for (const e of object.actorscontract) {
        message.actorscontract.push(Actorscontract.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetContractCounterRequest: object = {};

export const QueryGetContractCounterRequest = {
  encode(
    _: QueryGetContractCounterRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetContractCounterRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetContractCounterRequest,
    } as QueryGetContractCounterRequest;
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

  fromJSON(_: any): QueryGetContractCounterRequest {
    const message = {
      ...baseQueryGetContractCounterRequest,
    } as QueryGetContractCounterRequest;
    return message;
  },

  toJSON(_: QueryGetContractCounterRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetContractCounterRequest>
  ): QueryGetContractCounterRequest {
    const message = {
      ...baseQueryGetContractCounterRequest,
    } as QueryGetContractCounterRequest;
    return message;
  },
};

const baseQueryGetContractCounterResponse: object = {};

export const QueryGetContractCounterResponse = {
  encode(
    message: QueryGetContractCounterResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ContractCounter !== undefined) {
      ContractCounter.encode(
        message.ContractCounter,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetContractCounterResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetContractCounterResponse,
    } as QueryGetContractCounterResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ContractCounter = ContractCounter.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetContractCounterResponse {
    const message = {
      ...baseQueryGetContractCounterResponse,
    } as QueryGetContractCounterResponse;
    if (
      object.ContractCounter !== undefined &&
      object.ContractCounter !== null
    ) {
      message.ContractCounter = ContractCounter.fromJSON(
        object.ContractCounter
      );
    } else {
      message.ContractCounter = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetContractCounterResponse): unknown {
    const obj: any = {};
    message.ContractCounter !== undefined &&
      (obj.ContractCounter = message.ContractCounter
        ? ContractCounter.toJSON(message.ContractCounter)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetContractCounterResponse>
  ): QueryGetContractCounterResponse {
    const message = {
      ...baseQueryGetContractCounterResponse,
    } as QueryGetContractCounterResponse;
    if (
      object.ContractCounter !== undefined &&
      object.ContractCounter !== null
    ) {
      message.ContractCounter = ContractCounter.fromPartial(
        object.ContractCounter
      );
    } else {
      message.ContractCounter = undefined;
    }
    return message;
  },
};

const baseQueryGetNewcontractRequest: object = { index: "" };

export const QueryGetNewcontractRequest = {
  encode(
    message: QueryGetNewcontractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNewcontractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNewcontractRequest,
    } as QueryGetNewcontractRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNewcontractRequest {
    const message = {
      ...baseQueryGetNewcontractRequest,
    } as QueryGetNewcontractRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetNewcontractRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNewcontractRequest>
  ): QueryGetNewcontractRequest {
    const message = {
      ...baseQueryGetNewcontractRequest,
    } as QueryGetNewcontractRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetNewcontractResponse: object = {};

export const QueryGetNewcontractResponse = {
  encode(
    message: QueryGetNewcontractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.newcontract !== undefined) {
      Newcontract.encode(
        message.newcontract,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetNewcontractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetNewcontractResponse,
    } as QueryGetNewcontractResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.newcontract = Newcontract.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNewcontractResponse {
    const message = {
      ...baseQueryGetNewcontractResponse,
    } as QueryGetNewcontractResponse;
    if (object.newcontract !== undefined && object.newcontract !== null) {
      message.newcontract = Newcontract.fromJSON(object.newcontract);
    } else {
      message.newcontract = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNewcontractResponse): unknown {
    const obj: any = {};
    message.newcontract !== undefined &&
      (obj.newcontract = message.newcontract
        ? Newcontract.toJSON(message.newcontract)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNewcontractResponse>
  ): QueryGetNewcontractResponse {
    const message = {
      ...baseQueryGetNewcontractResponse,
    } as QueryGetNewcontractResponse;
    if (object.newcontract !== undefined && object.newcontract !== null) {
      message.newcontract = Newcontract.fromPartial(object.newcontract);
    } else {
      message.newcontract = undefined;
    }
    return message;
  },
};

const baseQueryAllNewcontractRequest: object = {};

export const QueryAllNewcontractRequest = {
  encode(
    message: QueryAllNewcontractRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllNewcontractRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNewcontractRequest,
    } as QueryAllNewcontractRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNewcontractRequest {
    const message = {
      ...baseQueryAllNewcontractRequest,
    } as QueryAllNewcontractRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNewcontractRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNewcontractRequest>
  ): QueryAllNewcontractRequest {
    const message = {
      ...baseQueryAllNewcontractRequest,
    } as QueryAllNewcontractRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNewcontractResponse: object = {};

export const QueryAllNewcontractResponse = {
  encode(
    message: QueryAllNewcontractResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.newcontract) {
      Newcontract.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllNewcontractResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllNewcontractResponse,
    } as QueryAllNewcontractResponse;
    message.newcontract = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.newcontract.push(Newcontract.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNewcontractResponse {
    const message = {
      ...baseQueryAllNewcontractResponse,
    } as QueryAllNewcontractResponse;
    message.newcontract = [];
    if (object.newcontract !== undefined && object.newcontract !== null) {
      for (const e of object.newcontract) {
        message.newcontract.push(Newcontract.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNewcontractResponse): unknown {
    const obj: any = {};
    if (message.newcontract) {
      obj.newcontract = message.newcontract.map((e) =>
        e ? Newcontract.toJSON(e) : undefined
      );
    } else {
      obj.newcontract = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNewcontractResponse>
  ): QueryAllNewcontractResponse {
    const message = {
      ...baseQueryAllNewcontractResponse,
    } as QueryAllNewcontractResponse;
    message.newcontract = [];
    if (object.newcontract !== undefined && object.newcontract !== null) {
      for (const e of object.newcontract) {
        message.newcontract.push(Newcontract.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a ConCounter by index. */
  ConCounter(
    request: QueryGetConCounterRequest
  ): Promise<QueryGetConCounterResponse>;
  /** Queries a Actorscontract by index. */
  Actorscontract(
    request: QueryGetActorscontractRequest
  ): Promise<QueryGetActorscontractResponse>;
  /** Queries a list of Actorscontract items. */
  ActorscontractAll(
    request: QueryAllActorscontractRequest
  ): Promise<QueryAllActorscontractResponse>;
  /** Queries a ContractCounter by index. */
  ContractCounter(
    request: QueryGetContractCounterRequest
  ): Promise<QueryGetContractCounterResponse>;
  /** Queries a Newcontract by index. */
  Newcontract(
    request: QueryGetNewcontractRequest
  ): Promise<QueryGetNewcontractResponse>;
  /** Queries a list of Newcontract items. */
  NewcontractAll(
    request: QueryAllNewcontractRequest
  ): Promise<QueryAllNewcontractResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "upalchowdhury.supplychain.supplychain.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  ConCounter(
    request: QueryGetConCounterRequest
  ): Promise<QueryGetConCounterResponse> {
    const data = QueryGetConCounterRequest.encode(request).finish();
    const promise = this.rpc.request(
      "upalchowdhury.supplychain.supplychain.Query",
      "ConCounter",
      data
    );
    return promise.then((data) =>
      QueryGetConCounterResponse.decode(new Reader(data))
    );
  }

  Actorscontract(
    request: QueryGetActorscontractRequest
  ): Promise<QueryGetActorscontractResponse> {
    const data = QueryGetActorscontractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "upalchowdhury.supplychain.supplychain.Query",
      "Actorscontract",
      data
    );
    return promise.then((data) =>
      QueryGetActorscontractResponse.decode(new Reader(data))
    );
  }

  ActorscontractAll(
    request: QueryAllActorscontractRequest
  ): Promise<QueryAllActorscontractResponse> {
    const data = QueryAllActorscontractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "upalchowdhury.supplychain.supplychain.Query",
      "ActorscontractAll",
      data
    );
    return promise.then((data) =>
      QueryAllActorscontractResponse.decode(new Reader(data))
    );
  }

  ContractCounter(
    request: QueryGetContractCounterRequest
  ): Promise<QueryGetContractCounterResponse> {
    const data = QueryGetContractCounterRequest.encode(request).finish();
    const promise = this.rpc.request(
      "upalchowdhury.supplychain.supplychain.Query",
      "ContractCounter",
      data
    );
    return promise.then((data) =>
      QueryGetContractCounterResponse.decode(new Reader(data))
    );
  }

  Newcontract(
    request: QueryGetNewcontractRequest
  ): Promise<QueryGetNewcontractResponse> {
    const data = QueryGetNewcontractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "upalchowdhury.supplychain.supplychain.Query",
      "Newcontract",
      data
    );
    return promise.then((data) =>
      QueryGetNewcontractResponse.decode(new Reader(data))
    );
  }

  NewcontractAll(
    request: QueryAllNewcontractRequest
  ): Promise<QueryAllNewcontractResponse> {
    const data = QueryAllNewcontractRequest.encode(request).finish();
    const promise = this.rpc.request(
      "upalchowdhury.supplychain.supplychain.Query",
      "NewcontractAll",
      data
    );
    return promise.then((data) =>
      QueryAllNewcontractResponse.decode(new Reader(data))
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
