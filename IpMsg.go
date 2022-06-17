package ipmsg

import (
	"ipmsg/logger"
	"net"
	"strconv"
	"strings"
)

type IpMsg struct {
	*Base

	encoding       string
	packageHandler func(*IpMsg)
	cmdMap         map[CmdType]func(*IpMsg, *Package)
	eventMap       map[CmdType]func(*IpMsg, *Package)
	userManager    IUserManager
}

func NewIpMsg(user string, host string, port int) (im *IpMsg, err error) {
	ipMsgBase, err := NewIpMsgBase(user, host, port)
	if err != nil {
		return nil, err
	}
	im = &IpMsg{
		Base:        ipMsgBase,
		cmdMap:      NewCmdMap(),
		eventMap:    make(map[CmdType]func(*IpMsg, *Package)),
		userManager: &UserManager{},
	}
	return im, nil
}

// 绑定报文处理器
func (im *IpMsg) BindHandler(handler func(*IpMsg)) {
	im.packageHandler = handler
	im.cmdMap = nil
}

// 启动服务
func (im *IpMsg) Run() {
	if im.packageHandler == nil {
		im.packageHandler = defaultHandler
	}
	im.userManager.Init()
	im.packageHandler(im)
}

// 内置报文处理器
func defaultHandler(im *IpMsg) {
	for {
		pkg, _ := im.ReadPackage()
		if cmd, ok := im.cmdMap[pkg.CommandNo.GetCmd()]; ok {
			cmd(im, pkg)
		} else {
			logger.Warning("no packageHandler for cmd [%s]", pkg.CommandNo.GetCmd())
		}
		if event, ok := im.eventMap[pkg.CommandNo.GetCmd()]; ok {
			event(im, pkg)
		}
	}
}

// 绑定报文处理函数
func (im *IpMsg) BindCommand(cmdNo CmdType, cmd func(*IpMsg, *Package)) {
	if &im.packageHandler != nil {
		panic("packageHandler has been modified")
	}
	im.cmdMap[cmdNo] = cmd
}

// 绑定报文处理函数组
func (im *IpMsg) BindCommandMap(cmdMap map[CmdType]func(*IpMsg, *Package)) {
	if &im.packageHandler != nil {
		panic("packageHandler has been modified")
	}
	im.cmdMap = cmdMap
}

// 绑定用户管理器
func (im *IpMsg) BindUserManager(userManager IUserManager) {
	im.userManager = userManager
}

// 绑定事件处理器
func (im *IpMsg) BindEvent(cmdNo CmdType, cmd func(*IpMsg, *Package)) {
	im.eventMap[cmdNo] = cmd
}

// 设置报文编码
func (im *IpMsg) SetEncoding(encoding string) {
	im.encoding = strings.ToLower(encoding)
}

// 发送上线广播
func (im *IpMsg) EntryBroadCast() {
	pkg := im.NewPackage(IPMSG_BR_ENTRY, im.SenderName)
	_ = im.SendPackage(im.BroadCastAddr, pkg)
	_ = im.SendPackage(im.BroadCastV6Addr, pkg)
}

// 发送下线广播
func (im *IpMsg) ExitBroadCast() {
	pkg := im.NewPackage(IPMSG_BR_EXIT, im.SenderName)
	_ = im.SendPackage(im.BroadCastAddr, pkg)
	_ = im.SendPackage(im.BroadCastV6Addr, pkg)
}

func (im *IpMsg) SendEntryAnswer(addr *net.UDPAddr) {
	pkg := im.NewPackage(IPMSG_ANSENTRY, im.SenderName)
	_ = im.SendPackage(addr, pkg)
}

func (im *IpMsg) SendMessage(addr *net.UDPAddr, text string) {
	pkg := im.NewPackage(IPMSG_SENDMSG, text)
	pkg.SetFlag(IPMSG_SENDCHECKOPT)
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
