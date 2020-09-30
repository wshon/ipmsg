package main

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"ipmsg"
	"ipmsg/IpMsgCore"
	"ipmsg/logger"
)

func PackageHandler(im *ipmsg.Base) {
	for {
		pkg, _ := im.ReadPackage()
		switch pkg.CommandNo.GetCmd() {
		case ipmsg.IPMSG_BR_ENTRY:
		case ipmsg.IPMSG_BR_EXIT:
		case ipmsg.IPMSG_ANSENTRY:
		case ipmsg.IPMSG_SENDMSG:
		case ipmsg.IPMSG_RECVMSG:
		case ipmsg.IPMSG_NOOPERATION:
			// 无操作忽略
		default:
			logger.Info("recv unknown from [%s]#\n%s\n%X", pkg.SenderName, pkg, pkg)
		}
	}
}
func runBase() {
	im, err := ipmsg.NewIpMsgBase("Test", "172.18.60.209", ipmsg.IPMSG_DEFAT_PORT)
	if err != nil {
		panic(err)
	}
	pkg := im.NewPackage(ipmsg.IPMSG_BR_ENTRY, im.SenderName)
	_ = im.SendPackage(im.BroadCastAddr, pkg)
	im.BindHandler(PackageHandler)
	im.Run()
}

func runIpMsg() {
	im, err := ipmsg.NewIpMsg("Test", "0.0.0.0", ipmsg.IPMSG_DEFAT_PORT)
	if err != nil {
		panic(err)
	}
	im.BindDecoder(simplifiedchinese.GBK.NewDecoder().String)
	im.BindUserManager(new(IpMsgCore.UserManager))
	im.EntryBroadCast()
	im.Run()
}

func main() {
	runIpMsg()
}
