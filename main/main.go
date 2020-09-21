package main

import (
	"ipmsg"
	"ipmsg/IpMsgCore"
)

func main() {
	im := ipmsg.NewIpMsg("Test", "0.0.0.0", ipmsg.IPMSG_DEFAT_PORT)
	im.EntryBroadCast()
	im.BindHandler(IpMsgCore.PackageHandler)
	im.Run()
}
