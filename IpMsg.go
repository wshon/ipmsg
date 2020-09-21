package ipmsg

import (
	"ipmsg/logger"
	"strconv"
)

type IpMsg struct {
	*Base
	packageHandler func(*IpMsg) error
}

func NewIpMsg(user string, host string, port int) (im *IpMsg) {
	im = &IpMsg{
		Base: NewIpMsgBase(user, host, port),
	}
	return im
}

func (im *IpMsg) BindHandler(handler func(*IpMsg) error) {
	im.packageHandler = handler
}

func (im *IpMsg) Run() {
	_ = im.packageHandler(im)
}

//上线广播
func (im *IpMsg) EntryBroadCast() {
	logger.Debug("send entry broadcast to [%s]", im.broadCastAddr)
	pkg := im.newPackage(IPMSG_BR_ENTRY, []byte(im.SenderName))
	_ = im.sendPackage(im.broadCastAddr, pkg)
}

//下线广播
func (im *IpMsg) ExitBroadCast() {
	logger.Debug("send exit broadcast to [%s]", im.broadCastAddr)
	pkg := im.newPackage(IPMSG_BR_EXIT, []byte(im.SenderName))
	_ = im.sendPackage(im.broadCastAddr, pkg)
}

func (im *IpMsg) SendEntryAnswer(receiver *UserInfo) {
	logger.Debug("send entry answer to [%s]", receiver.Addr)
	pkg := im.newPackage(IPMSG_ANSENTRY, []byte(im.SenderName))
	_ = im.sendPackage(receiver.Addr, pkg)
}

func (im *IpMsg) SendMessageReceived(receiver *UserInfo, packetNo int) {
	logger.Debug("send message received to [%s]", im.broadCastAddr)
	pkg := im.newPackage(IPMSG_RECVMSG, []byte(strconv.Itoa(packetNo)))
	_ = im.sendPackage(receiver.Addr, pkg)
}
