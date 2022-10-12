/* eslint-disable */
import { Params } from "../supplychain/params";
import { ConCounter } from "../supplychain/con_counter";
import { Actorscontract } from "../supplychain/actorscontract";
import { ContractCounter } from "../supplychain/contract_counter";
import { Newcontract } from "../supplychain/newcontract";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "upalchowdhury.supplychain.supplychain";

/** GenesisState defines the supplychain module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  conCounter: ConCounter | undefined;
  actorscontractList: Actorscontract[];
  contractCounter: ContractCounter | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  newcontractList: Newcontract[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.conCounter !== undefined) {
      ConCounter.encode(message.conCounter, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.actorscontractList) {
      Actorscontract.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.contractCounter !== undefined) {
      ContractCounter.encode(
        message.contractCounter,
        writer.uint32(34).fork()
      ).ldelim();
    }
    for (const v of message.newcontractList) {
      Newcontract.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.actorscontractList = [];
    message.newcontractList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.conCounter = ConCounter.decode(reader, reader.uint32());
          break;
        case 3:
          message.actorscontractList.push(
            Actorscontract.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.contractCounter = ContractCounter.decode(
            reader,
            reader.uint32()
          );
          break;
        case 5:
          message.newcontractList.push(
            Newcontract.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.actorscontractList = [];
    message.newcontractList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.conCounter !== undefined && object.conCounter !== null) {
      message.conCounter = ConCounter.fromJSON(object.conCounter);
    } else {
      message.conCounter = undefined;
    }
    if (
      object.actorscontractList !== undefined &&
      object.actorscontractList !== null
    ) {
      for (const e of object.actorscontractList) {
        message.actorscontractList.push(Actorscontract.fromJSON(e));
      }
    }
    if (
      object.contractCounter !== undefined &&
      object.contractCounter !== null
    ) {
      message.contractCounter = ContractCounter.fromJSON(
        object.contractCounter
      );
    } else {
      message.contractCounter = undefined;
    }
    if (
      object.newcontractList !== undefined &&
      object.newcontractList !== null
    ) {
      for (const e of object.newcontractList) {
        message.newcontractList.push(Newcontract.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.conCounter !== undefined &&
      (obj.conCounter = message.conCounter
        ? ConCounter.toJSON(message.conCounter)
        : undefined);
    if (message.actorscontractList) {
      obj.actorscontractList = message.actorscontractList.map((e) =>
        e ? Actorscontract.toJSON(e) : undefined
      );
    } else {
      obj.actorscontractList = [];
    }
    message.contractCounter !== undefined &&
      (obj.contractCounter = message.contractCounter
        ? ContractCounter.toJSON(message.contractCounter)
        : undefined);
    if (message.newcontractList) {
      obj.newcontractList = message.newcontractList.map((e) =>
        e ? Newcontract.toJSON(e) : undefined
      );
    } else {
      obj.newcontractList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.actorscontractList = [];
    message.newcontractList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.conCounter !== undefined && object.conCounter !== null) {
      message.conCounter = ConCounter.fromPartial(object.conCounter);
    } else {
      message.conCounter = undefined;
    }
    if (
      object.actorscontractList !== undefined &&
      object.actorscontractList !== null
    ) {
      for (const e of object.actorscontractList) {
        message.actorscontractList.push(Actorscontract.fromPartial(e));
      }
    }
    if (
      object.contractCounter !== undefined &&
      object.contractCounter !== null
    ) {
      message.contractCounter = ContractCounter.fromPartial(
        object.contractCounter
      );
    } else {
      message.contractCounter = undefined;
    }
    if (
      object.newcontractList !== undefined &&
      object.newcontractList !== null
    ) {
      for (const e of object.newcontractList) {
        message.newcontractList.push(Newcontract.fromPartial(e));
      }
    }
    return message;
  },
};

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
