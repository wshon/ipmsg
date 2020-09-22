package ipmsg

import (
	"ipmsg/logger"
	"net"
	"strconv"
)

type IpMsg struct {
	*Base
	handler func(*IpMsg)
	cmdMap  map[CmdType]func(*IpMsg, *Package)
}

func NewIpMsg(user string, host string, port int) (im *IpMsg, err error) {
	ipMsgBase, err := NewIpMsgBase(user, host, port)
	if err != nil {
		return nil, err
	}
	im = &IpMsg{
		Base:   ipMsgBase,
		cmdMap: make(map[CmdType]func(*IpMsg, *Package)),
	}
	return im, nil
}

func (im *IpMsg) BindHandler(handler func(*IpMsg)) {
	im.handler = handler
}

func (im *IpMsg) BindCommandMap(cmdMap map[CmdType]func(*IpMsg, *Package)) {
	im.cmdMap = cmdMap
}

func (im *IpMsg) BindCommand(cmdNo CmdType, cmd func(*IpMsg, *Package)) {
	im.cmdMap[cmdNo] = cmd
}

func (im *IpMsg) Run() {
	if im.handler != nil {
		im.handler(im)
	} else {
		im.defaultHandler()
	}
}

func (im *IpMsg) defaultHandler() {
	for {
		pkg, _ := im.ReadPackage()
		if cmd, ok := im.cmdMap[pkg.CommandNo.GetCmd()]; ok {
			cmd(im, pkg)
		} else {
			logger.Warning("no handler for cmd [%s]", pkg.CommandNo.GetCmd())
		}
	}
}

//上线广播
func (im *IpMsg) EntryBroadCast() {
	pkg := im.newPackage(IPMSG_BR_ENTRY, im.SenderName)
	_ = im.sendPackage(im.broadCastAddr, pkg)
}

//下线广播
func (im *IpMsg) ExitBroadCast() {
	pkg := im.newPackage(IPMSG_BR_EXIT, im.SenderName)
	_ = im.sendPackage(im.broadCastAddr, pkg)
}

func (im *IpMsg) SendEntryAnswer(addr *net.UDPAddr) {
	pkg := im.newPackage(IPMSG_ANSENTRY, im.SenderName)
	_ = im.sendPackage(addr, pkg)
}

func (im *IpMsg) SendMessageReceived(addr *net.UDPAddr, packetNo uint32) {
	pkg := im.newPackage(IPMSG_RECVMSG, strconv.Itoa(int(packetNo)))
	_ = im.sendPackage(addr, pkg)
}

func (im *IpMsg) SendMessageRead(addr *net.UDPAddr, packetNo uint32) {
	pkg := im.newPackage(IPMSG_READMSG.WithFlag(IPMSG_READCHECKOPT), strconv.Itoa(int(packetNo)))
	_ = im.sendPackage(addr, pkg)
}
