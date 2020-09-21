package IpMsgCore

import (
	"ipmsg"
	"ipmsg/logger"
	"net"
)

//根据用户名返回IP地址
func get_addr_by_name(name string) *net.UDPAddr {
	return nil
}
func add_user(user *ipmsg.UserInfo) {
	logger.Debug("add user [%s]", user)
}
func del_user(user *ipmsg.UserInfo) {
	logger.Debug("del user [%s]", user)
}
