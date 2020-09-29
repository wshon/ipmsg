package IpMsgCore

import (
	"ipmsg"
	"ipmsg/logger"
	"net"
	"strings"
)

type UserManager struct {
}

func (u UserManager) AddUser(pkg *ipmsg.Package) interface{ ipmsg.IUserInfo } {
	user := &ipmsg.UserInfo{
		Addr: pkg.SenderAddr,
		Host: pkg.SenderHost,
		Info: pkg.AdditionalSection,
	}
	idInfo := strings.Split(pkg.SenderName, "-")
	if len(idInfo) > 1 {
		user.Id = idInfo[1]
	}
	user.Name = idInfo[0]
	logger.Debug("add user [%+v]", user)
	return user
}

func (u UserManager) DelUser(pkg *ipmsg.Package) {
	user := &ipmsg.UserInfo{
		Name: pkg.SenderName,
		Host: pkg.SenderHost,
		Addr: pkg.SenderAddr,
	}
	logger.Debug("del user [%s]", user)
}

func (u UserManager) GetAddrByName(name string) *net.UDPAddr {
	return nil
}
