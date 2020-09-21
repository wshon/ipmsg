package IpMsgCore

import (
	"ipmsg"
	"ipmsg/logger"
	"net"
)

//根据用户名返回IP地址
func getAddrByName(name string) *net.UDPAddr {
	return nil
}

func addUser(pkg *ipmsg.Package) (user *ipmsg.UserInfo) {
	user = &ipmsg.UserInfo{
		Name: pkg.SenderName,
		Host: pkg.SenderHost,
		Addr: pkg.SenderAddr,
	}
	logger.Debug("add user [%s]", user)
	return user
}

func delUser(pkg *ipmsg.Package) {
	user := &ipmsg.UserInfo{
		Name: pkg.SenderName,
		Host: pkg.SenderHost,
		Addr: pkg.SenderAddr,
	}
	logger.Debug("del user [%s]", user)
}
