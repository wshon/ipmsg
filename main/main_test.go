package main

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"ipmsg"
	"ipmsg/logger"
	"testing"
)

// PackageHandler 自定义的报文处理器
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

// runBase 运行自定义报文处理器
func TestRunBase(t *testing.T) {
	im, err := ipmsg.NewIpMsgBase("Test", host, ipmsg.IPMSG_DEFAT_PORT)
	if err != nil {
		panic(err)
	}
	pkg := im.NewPackage(ipmsg.IPMSG_BR_ENTRY, im.SenderName)
	_ = im.SendPackage(im.BroadCastAddr, pkg)
	im.BindHandler(PackageHandler)
	im.Run()
}

// listenIpMsg 报文监听处理器
func listenHandler(im *ipmsg.IpMsg) {
	for {
		pkg, _ := im.ReadPackage()
		logger.Warning("no packageHandler for cmd [%s]", pkg.CommandNo.GetCmd())
	}
}

// listenIpMsg 运行报文监听
func TestListenIpMsg(t *testing.T) {
	im, err := ipmsg.NewIpMsg("Test", host, ipmsg.IPMSG_DEFAT_PORT)
	if err != nil {
		panic(err)
	}
	im.BindDecoder(simplifiedchinese.GBK.NewDecoder().String)
	im.BindHandler(listenHandler)
	im.Run()
}
