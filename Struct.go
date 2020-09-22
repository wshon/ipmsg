package ipmsg

import (
	"net"
	"os"
)

type UserInfo struct {
	Addr *net.UDPAddr
	Name string
	Host string
	Info string
}

type FileInfo struct {
	os.FileInfo
	UserName string
	Pkgnum   interface{}
	Num      interface{}
	Name     string
}

const (
	SENDFILE = 0
	RECVFILE = 1
)
