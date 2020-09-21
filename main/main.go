package main

import (
	"ipmsg"
	"ipmsg/IpMsgCore"
)

func run1() {
	im := ipmsg.NewIpMsg("Test", "172.18.60.209", ipmsg.IPMSG_DEFAT_PORT)
	im.EntryBroadCast()
	im.BindHandler(IpMsgCore.PackageHandler)
	im.Run()
}
func run2() {
	im := ipmsg.NewIpMsg("Test", "172.18.60.209", ipmsg.IPMSG_DEFAT_PORT)
	im.EntryBroadCast()
	im.BindCommandMap(IpMsgCore.NewCmdMap())
	im.Run()
}

func main() {
	run2()
}
