package main

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"ipmsg"
	"ipmsg/IpMsgCore"
)

func run1() {
	im, err := ipmsg.NewIpMsg("Test", "0.0.0.0", ipmsg.IPMSG_DEFAT_PORT)
	if err != nil {
		panic(err)
	}
	im.EntryBroadCast()
	im.BindHandler(IpMsgCore.PackageHandler)
	im.Run()
}
func run2() {
	conv := simplifiedchinese.GBK.NewDecoder().String

	im, err := ipmsg.NewIpMsg("Test", "0.0.0.0", ipmsg.IPMSG_DEFAT_PORT)
	if err != nil {
		panic(err)
	}
	im.BindDecoder(conv)
	im.EntryBroadCast()
	im.BindCommandMap(IpMsgCore.NewCmdMap())
	im.Run()
}

func main() {
	run2()
}
