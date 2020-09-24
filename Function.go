package ipmsg

import (
	"ipmsg/logger"
	"net"
	"os"
)

type UserInfo struct {
	Addr *net.UDPAddr
	Name string
	Host string
	Info string
}

func (u UserInfo) AddUser(pkg *Package) *IUserManager {
	user := &UserInfo{
		Name: pkg.SenderName,
		Host: pkg.SenderHost,
		Addr: pkg.SenderAddr,
		Info: pkg.AdditionalSection,
	}
	logger.Debug("add user [%+v]", user)
	return user
}

func (u UserInfo) DelUser(pkg *Package) *IUserManager {
	panic("implement me")
}

func (u UserInfo) GetAddrByName(name string) *net.UDPAddr {
	panic("implement me")
}

func (u UserInfo) AddUser(pkg *Package) {
}

func (u UserInfo) DelUser(pkg *Package) {
	panic("implement me")
}

type FileInfo struct {
	os.FileInfo
	UserName string
	Pkgnum   interface{}
	Num      interface{}
	Name     string
}
