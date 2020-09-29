package ipmsg

import (
	"ipmsg/logger"
	"net"
	"os"
)

type UserManager struct {
}

func (u UserManager) AddUser(pkg *Package) interface{ IUserInfo } {
	user := &UserInfo{}
	logger.Error("no implement [IUserManager]")
	return user
}

func (u UserManager) DelUser(pkg *Package) {
	logger.Error("no implement [IUserManager]")
}

func (u UserManager) GetAddrByName(name string) *net.UDPAddr {
	logger.Error("no implement [IUserManager]")
	return nil
}

type UserInfo struct {
	Id   string
	Addr *net.UDPAddr
	Name string
	Host string
	Info string
}

func (u UserInfo) GetName() {
	logger.Error("no implement [IUserInfo]")
}

func (u UserInfo) GetAddr() *net.UDPAddr {
	logger.Error("no implement [IUserInfo]")
	return nil
}

type FileInfo struct {
	os.FileInfo
	UserName string
	Pkgnum   interface{}
	Num      interface{}
	Name     string
}
