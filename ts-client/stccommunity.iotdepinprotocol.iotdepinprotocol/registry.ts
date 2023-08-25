import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeleteEventPb } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgDeleteKv } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgDeleteDevice } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgUpdateEventPb } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgUpdateDevice } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgCreateDevice } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgCreateEventPb } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgCreateKv } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgUpdateKv } from "./types/iotdepinprotocol/iotdepinprotocol/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgDeleteEventPb", MsgDeleteEventPb],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgDeleteKv", MsgDeleteKv],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgDeleteDevice", MsgDeleteDevice],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgUpdateEventPb", MsgUpdateEventPb],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgUpdateDevice", MsgUpdateDevice],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgCreateDevice", MsgCreateDevice],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgCreateEventPb", MsgCreateEventPb],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgCreateKv", MsgCreateKv],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgUpdateKv", MsgUpdateKv],
    
];

export { msgTypes }