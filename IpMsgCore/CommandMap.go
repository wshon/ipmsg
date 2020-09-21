package IpMsgCore

import (
	"ipmsg"
)

func NewCmdMap() (cmdMap map[ipmsg.CommandType]func(*ipmsg.IpMsg, *ipmsg.Package)) {
	cmdMap = make(map[ipmsg.CommandType]func(*ipmsg.IpMsg, *ipmsg.Package))
	cmdMap[ipmsg.IPMSG_BR_ENTRY] = onIpMsgBrEntry
	cmdMap[ipmsg.IPMSG_BR_EXIT] = onIpMsgBrExit
	cmdMap[ipmsg.IPMSG_ANSENTRY] = onIpMsgAnsEntry
	cmdMap[ipmsg.IPMSG_SENDMSG] = onIpMsgSendMsg
	cmdMap[ipmsg.IPMSG_RECVMSG] = onIpMsgRecvMsg
	cmdMap[ipmsg.IPMSG_NOOPERATION] = onIpMsgNoOperation
	return cmdMap
}
