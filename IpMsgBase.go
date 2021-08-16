package ipmsg

import (
	"ipmsg/logger"
	"net"
	"strconv"
	"strings"
	"time"
)

type Base struct {
	SenderName    string
	SenderHost    string
	SenderPort    int
	BroadCastAddr *net.UDPAddr

	packageHandler func(*Base)
	messageDecoder func(s string) (string, error)
	tcpNet         *net.TCPListener
	udpNet         *net.UDPConn
}

func NewIpMsgBase(user string, host string, port int) (im *Base, err error) {
	im = &Base{
		SenderName: user,
		SenderHost: host,
		SenderPort: port,
		BroadCastAddr: &net.UDPAddr{
			IP:   net.IPv4bcast,
			Port: port,
		},
	}
	logger.Debug("start init network")
	im.tcpNet, err = createTcpServer(im.SenderHost, im.SenderPort)
	if err != nil {
		return nil, err
	}
	im.udpNet, err = createUdpServer(im.SenderHost, im.SenderPort)
	if err != nil {
		return nil, err
	}
	logger.Debug("init network success")
	return im, nil
}

func (im *Base) BindDecoder(decoder func(s string) (string, error)) {
	im.messageDecoder = decoder
}

func (im *Base) BindHandler(handler func(*Base)) {
	im.packageHandler = handler
}

func (im *Base) Run() {
	im.packageHandler(im)
}

func (im *Base) NewPackage(CommandNo CmdType, AdditionalSection string) *Package {
	return &Package{
		Ver:               strconv.Itoa(IPMSG_VERSION),
		PacketNo:          uint32(time.Now().Unix()),
		SenderName:        im.SenderName,
		SenderHost:        im.SenderHost,
		CommandNo:         CommandNo,
		AdditionalSection: AdditionalSection,
	}
}

func (im *Base) SendPackage(addr *net.UDPAddr, pkg *Package) error {
	logger.Debug("send pkg [%s] to [%s]", pkg.CommandNo.GetCmd(), addr)
	pkg.Buf = pkg.Marshal()
	logger.Trace("send new pkg %+v", pkg)
	logger.Trace("send pkg ext_data {%s}", strings.Replace(pkg.AdditionalSection, "\n", "\\n", -1))
	_, err := im.udpNet.WriteToUDP(pkg.Buf, addr)
	return err
}

func (im *Base) ReadPackage() (*Package, error) {
	buf := make([]byte, 512)
	n, addr, _ := im.udpNet.ReadFromUDP(buf)
	pkg, _ := UnMarshal(buf[:n])
	logger.Debug("recv pkg [%s] from [%s]", pkg.CommandNo.GetCmd(), addr)
	pkg.SenderAddr = addr
	logger.Trace("recv new pkg %+v", pkg)
	if im.messageDecoder != nil {
		pkg.AdditionalSection, _ = im.messageDecoder(pkg.AdditionalSection)
	}
	logger.Trace("recv pkg ext_data {%s}", strings.Replace(pkg.AdditionalSection, "\n", "\\n", -1))
	return pkg, nil
}
