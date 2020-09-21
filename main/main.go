package main

import (
	"ipmsg"
	"ipmsg/IpMsgCore"
)

func run1() {
	im, err := ipmsg.NewIpMsg("Test", "172.17.0.225", ipmsg.IPMSG_DEFAT_PORT)
	if err != nil {
		panic(err)
	}
	im.EntryBroadCast()
	im.BindHandler(IpMsgCore.PackageHandler)
	im.Run()
}
func run2() {
	im, err := ipmsg.NewIpMsg("Test", "172.17.0.225", ipmsg.IPMSG_DEFAT_PORT)
	if err != nil {
		panic(err)
	}
	im.EntryBroadCast()
	im.BindCommandMap(IpMsgCore.NewCmdMap())
	im.Run()
}

func main() {
	run2()
}
