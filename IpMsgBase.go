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

func NewIpMsgBase(user string, host string, port int) (im *Base) {
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
	im.tcp, _ = createTcpServer(im.SenderHost, im.SenderPort)
	im.udp, _ = createUdpServer(im.SenderHost, im.SenderPort)
	logger.Debug("init network success")
	return im
}

func (im *Base) BindHandler(handler func(*Base) error) {
	im.packageHandler = handler
}

func (im *Base) Run() {
	_ = im.packageHandler(im)
}

func (im *Base) newPackage(CommandNo CommandType, AdditionalSection []byte) *Package {
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
	logger.Trace("send new package %+v", pkg)
	logger.Trace("package ext_data {%s}", strings.Replace(string(pkg.AdditionalSection), "\n", "\\n", -1))
	_, err := im.udp.WriteToUDP(pkg.Buf, addr)
	return err
}

func (im *Base) ReadPackage() (*Package, error) {
	buf := make([]byte, 512)
	n, addr, _ := im.udp.ReadFromUDP(buf)
	pkg, _ := UnMarshal(buf[:n])
	pkg.SenderAddr = addr
	logger.Trace("recv new package %+v", pkg)
	logger.Trace("package ext_data {%s}", strings.Replace(string(pkg.AdditionalSection), "\n", "\\n", -1))
	return pkg, nil
}
