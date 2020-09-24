package ipmsg

import (
	"ipmsg/logger"
	"net"
	"strconv"
)

type IpMsg struct {
	*Base
	UserManager    IUserManager
	packageHandler func(*IpMsg)
	cmdMap         map[CmdType]func(*IpMsg, *Package)
}

func NewIpMsg(user string, host string, port int) (im *IpMsg, err error) {
	ipMsgBase, err := NewIpMsgBase(user, host, port)
	if err != nil {
		return nil, err
	}
	im = &IpMsg{
		Base:   ipMsgBase,
		cmdMap: NewCmdMap(),
	}
	return im, nil
}

func (im *IpMsg) BindCommandMap(cmdMap map[CmdType]func(*IpMsg, *Package)) {
	im.cmdMap = cmdMap
}

func (im *IpMsg) BindCommand(cmdNo CmdType, cmd func(*IpMsg, *Package)) {
	im.cmdMap[cmdNo] = cmd
}

func (im *IpMsg) BindHandler(handler func(*IpMsg)) {
	im.packageHandler = handler
}

func (im *IpMsg) Run() {
	if im.packageHandler != nil {
		im.packageHandler(im)
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
			logger.Warning("no packageHandler for cmd [%s]", pkg.CommandNo.GetCmd())
		}
	}
}

//上线广播
func (im *IpMsg) EntryBroadCast() {
	pkg := im.NewPackage(IPMSG_BR_ENTRY, im.SenderName)
	_ = im.SendPackage(im.BroadCastAddr, pkg)
}

//下线广播
func (im *IpMsg) ExitBroadCast() {
	pkg := im.NewPackage(IPMSG_BR_EXIT, im.SenderName)
	_ = im.SendPackage(im.BroadCastAddr, pkg)
}

func (im *IpMsg) SendEntryAnswer(addr *net.UDPAddr) {
	pkg := im.NewPackage(IPMSG_ANSENTRY, im.SenderName)
	_ = im.SendPackage(addr, pkg)
}

func (im *IpMsg) SendMessageReceived(addr *net.UDPAddr, packetNo uint32) {
	pkg := im.NewPackage(IPMSG_RECVMSG, strconv.Itoa(int(packetNo)))
	_ = im.SendPackage(addr, pkg)
}

func (im *IpMsg) SendMessageRead(addr *net.UDPAddr, packetNo uint32) {
	pkg := im.NewPackage(IPMSG_READMSG, strconv.Itoa(int(packetNo)))
	pkg.SetFlag(IPMSG_READCHECKOPT)
	_ = im.SendPackage(addr, pkg)
}
