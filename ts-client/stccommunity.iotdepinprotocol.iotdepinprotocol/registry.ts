import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateKv } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgCreateDevice } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgUpdateKv } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgDeleteDevice } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgUpdateDevice } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgDeleteKv } from "./types/iotdepinprotocol/iotdepinprotocol/tx";
import { MsgOracleOperator } from "./types/iotdepinprotocol/iotdepinprotocol/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgCreateKv", MsgCreateKv],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgCreateDevice", MsgCreateDevice],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgUpdateKv", MsgUpdateKv],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgDeleteDevice", MsgDeleteDevice],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgUpdateDevice", MsgUpdateDevice],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgDeleteKv", MsgDeleteKv],
    ["/stccommunity.iotdepinprotocol.iotdepinprotocol.MsgOracleOperator", MsgOracleOperator],
    
];

export { msgTypes }