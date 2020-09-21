package ipmsg

import (
	"ipmsg/logger"
	"net"
	"strconv"
	"strings"
	"time"
)

type Base struct {
	SenderName     string
	SenderHost     string
	SenderPort     int
	packageHandler func(*Base) error
	broadCastAddr  *net.UDPAddr
	tcp            *net.TCPListener
	udp            *net.UDPConn
}

func NewIpMsgBase(user string, host string, port int) (im *Base, err error) {
	im = &Base{
		SenderName: user,
		SenderHost: host,
		SenderPort: port,
		broadCastAddr: &net.UDPAddr{
			IP:   net.IPv4bcast,
			Port: port,
		},
	}
	logger.Debug("start init network")
	im.tcp, err = createTcpServer(im.SenderHost, im.SenderPort)
	if err != nil {
		return nil, err
	}
	im.udp, err = createUdpServer(im.SenderHost, im.SenderPort)
	if err != nil {
		return nil, err
	}
	logger.Debug("init network success")
	return im, nil
}

func (im *Base) BindHandler(handler func(*Base) error) {
	im.packageHandler = handler
}

func (im *Base) Run() {
	_ = im.packageHandler(im)
}

func (im *Base) newPackage(CommandNo CmdType, AdditionalSection []byte) *Package {
	return &Package{
		Ver:               strconv.Itoa(IPMSG_VERSION),
		PacketNo:          uint32(time.Now().Unix()),
		SenderName:        im.SenderName,
		SenderHost:        im.SenderHost,
		CommandNo:         CommandNo,
		AdditionalSection: AdditionalSection,
	}
}

func (im *Base) sendPackage(addr *net.UDPAddr, pkg *Package) error {
	pkg.Buf = pkg.Marshal()
	logger.Debug("send new pkg %+v", pkg)
	logger.Debug("send pkg ext_data {%s}", strings.Replace(string(pkg.AdditionalSection), "\n", "\\n", -1))
	_, err := im.udp.WriteToUDP(pkg.Buf, addr)
	logger.Trace("send pkg [%s] to [%s]", pkg.CommandNo.GetCmd(), addr)
	return err
}

func (im *Base) ReadPackage() (*Package, error) {
	buf := make([]byte, 512)
	n, addr, _ := im.udp.ReadFromUDP(buf)
	pkg, _ := UnMarshal(buf[:n])
	logger.Trace("recv pkg [%s] from [%s]", pkg.CommandNo.GetCmd(), addr)
	pkg.SenderAddr = addr
	logger.Debug("recv new pkg %+v", pkg)
	logger.Debug("recv pkg ext_data {%s}", strings.Replace(string(pkg.AdditionalSection), "\n", "\\n", -1))
	return pkg, nil
}
