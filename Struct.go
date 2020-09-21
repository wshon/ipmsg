package ipmsg

import (
	"net"
	"os"
)

type UserInfo struct {
	Name string
	Host string
	Addr *net.UDPAddr
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
