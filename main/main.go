package main

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"ipmsg"
	"ipmsg/IpMsgCore"
	"ipmsg/logger"
)

//var host = "192.168.182.1"
var host = "192.168.182.1"

func OnMsgIn(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
	im.SendMessage(pkg.SenderAddr, pkg.AdditionalSection)
}

// runIpMsg 运行IPMSG
func runIpMsg() *ipmsg.IpMsg {
	im, err := ipmsg.NewIpMsg("Test", host, ipmsg.IPMSG_DEFAT_PORT)
	if err != nil {
		panic(err)
	}
	im.BindDecoder(simplifiedchinese.GBK.NewDecoder().String)
	im.BindUserManager(new(IpMsgCore.UserManager))
	im.SetEncoding("utf-8")

	im.BindEvent(ipmsg.IPMSG_SENDMSG, OnMsgIn)

	im.EntryBroadCast()
	defer im.Run()
	return im
}

func main() {
	_ = runIpMsg()
	logger.Warning("no packageHandler for cmd")
}
