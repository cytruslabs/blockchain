/* eslint-disable */
import { Params } from "../zestchain/params";
import { PromoCount } from "../zestchain/promo_count";
import { Promo } from "../zestchain/promo";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cytruslabs.zestchain.zestchain";

/** GenesisState defines the zestchain module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  promoCount: PromoCount | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  promoList: Promo[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.promoCount !== undefined) {
      PromoCount.encode(message.promoCount, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.promoList) {
      Promo.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.promoList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.promoCount = PromoCount.decode(reader, reader.uint32());
          break;
        case 3:
          message.promoList.push(Promo.decode(reader, reader.uint32()));
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
    message.promoList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.promoCount !== undefined && object.promoCount !== null) {
      message.promoCount = PromoCount.fromJSON(object.promoCount);
    } else {
      message.promoCount = undefined;
    }
    if (object.promoList !== undefined && object.promoList !== null) {
      for (const e of object.promoList) {
        message.promoList.push(Promo.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.promoCount !== undefined &&
      (obj.promoCount = message.promoCount
        ? PromoCount.toJSON(message.promoCount)
        : undefined);
    if (message.promoList) {
      obj.promoList = message.promoList.map((e) =>
        e ? Promo.toJSON(e) : undefined
      );
    } else {
      obj.promoList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.promoList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.promoCount !== undefined && object.promoCount !== null) {
      message.promoCount = PromoCount.fromPartial(object.promoCount);
    } else {
      message.promoCount = undefined;
    }
    if (object.promoList !== undefined && object.promoList !== null) {
      for (const e of object.promoList) {
        message.promoList.push(Promo.fromPartial(e));
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
