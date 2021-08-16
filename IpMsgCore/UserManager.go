package IpMsgCore

import (
	"ipmsg"
	"ipmsg/logger"
	"net"
	"strings"
)

type UserManager struct {
	userList map[string]*ipmsg.UserInfo
}

func (u *UserManager) Init() {
	u.userList = make(map[string]*ipmsg.UserInfo)
}

func (u *UserManager) AddUser(pkg *ipmsg.Package) interface{ ipmsg.IUserInfo } {
	user := &ipmsg.UserInfo{
		Addr:     pkg.SenderAddr,
		HostName: pkg.SenderHost,
	}
	senderName := strings.Split(pkg.SenderName, "<")
	if len(senderName) > 1 {
		user.IdCode = senderName[1]
	}
	user.UserName = senderName[0]
	extraInfo := strings.Split(pkg.AdditionalSection, "\x00")
	if len(extraInfo) > 2 {
		user.IdCode = extraInfo[2]
	}
	if len(extraInfo) > 1 {
		user.GroupName = extraInfo[1]
	}
	user.NickName = extraInfo[0]
	if user.IdCode == "" {
		user.IdCode = user.UserName
	}
	u.userList[user.Addr.String()] = user
	logger.Debug("add user [%+v]", user)
	return user
}

func (u *UserManager) DelUser(pkg *ipmsg.Package) {
	user := &ipmsg.UserInfo{
		NickName: pkg.SenderName,
		HostName: pkg.SenderHost,
		Addr:     pkg.SenderAddr,
	}
	delete(u.userList, pkg.SenderAddr.String())
	logger.Debug("del user [%s]", user)
}

func (u *UserManager) GetAddrByName(name string) *net.UDPAddr {
	return nil
}
