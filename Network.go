package ipmsg

import (
	"ipmsg/logger"
	"net"
	"strconv"
)

//创建 TCP Server 套接字
func createTcpServer(host string, port int) (*net.TCPListener, error) {
	logger.Debug("create tcp listener")
	udpAddr, err := net.ResolveTCPAddr("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		logger.Fatal("%s", err)
		return nil, err
	}
	//监听端口
	udpConn, err := net.ListenTCP("tcp", udpAddr)
	if err != nil {
		logger.Fatal("%s", err)
		return nil, err
	}
	return udpConn, nil
}

//创建 UDP Server 套接字
func createUdpServer(host string, port int) (*net.UDPConn, error) {
	logger.Debug("create udp listener")
	udpAddr, err := net.ResolveUDPAddr("udp", host+":"+strconv.Itoa(port))
	if err != nil {
		logger.Fatal("%s", err)
		return nil, err
	}
	//监听端口
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		logger.Fatal("%s", err)
		return nil, err
	}
	return udpConn, nil
}
