package IpMsgCore

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"ipmsg"
	"ipmsg/logger"
)

func onIpMsgBrEntry(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
	addUser(pkg)
	im.SendEntryAnswer(pkg.SenderAddr)
}

func onIpMsgBrExit(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
	delUser(pkg)
}

func onIpMsgAnsEntry(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
	addUser(pkg)
}

func onIpMsgSendMsg(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
	message, _ := simplifiedchinese.GBK.NewDecoder().String(pkg.AdditionalSection)
	logger.Info("new msg from [%s]# %s\n", pkg.SenderName, message)
	if ipmsg.IPMSG_SENDCHECKOPT.CheckOpt(pkg.CommandNo) {
		im.SendMessageReceived(pkg.SenderAddr, pkg.PacketNo)
	}
	if ipmsg.IPMSG_SECRETEXOPT.CheckOpt(pkg.CommandNo) {
		im.SendMessageRead(pkg.SenderAddr, pkg.PacketNo)
	}
	if ipmsg.IPMSG_FILEATTACHOPT.CheckOpt(pkg.CommandNo) {
		//char * p = ipMsg + strlen(ipMsg) + 1
		////printf("filemsg=%s\n",p);
		//char * fileopt = strtok(p, "\a") //fileopt指向第一个文件属性
		//do{ //循环提取文件信息
		//	IPMSG_FILE, ftemp
		//	fmt.Sscanf(fileopt, "%d:%[^:]:%lx:%lx", &ftemp.PacketNo, ftemp.selfName, &ftemp.size, &ftemp.ltime)
		//	strcpy(ftemp.user, user.selfName)
		//	ftemp.pkgnum, = pkgnum
		//	add_file(ftemp, RECVFILE)
		//	fileopt = strtok(NULL, "\a") //fileopt指向下一个文件属性
		//}
		//while(fileopt != NULL)
		//IPMSG_OUT_MSG_COLOR(
		//	printf("<<<Recv file from %s!>>>\n", user.selfName),
		//)
	}
}

func onIpMsgRecvMsg(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
	logger.Info("[%s] have received your ipMsg!\n", pkg.SenderName)
}

func onIpMsgNoOperation(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
}
