package IpMsgCore

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"ipmsg"
	"ipmsg/logger"
)

func onIpMsgBrEntry(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
	logger.Debug("pkg IPMSG_BR_ENTRY from [%s]", pkg.SenderAddr)
	addUser(pkg)
	im.SendEntryAnswer(pkg.SenderAddr)
}

func onIpMsgBrExit(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
	logger.Debug("pkg IPMSG_BR_EXIT from [%s]", pkg.SenderAddr)
	delUser(pkg)
}

func onIpMsgAnsEntry(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
	logger.Debug("pkg IPMSG_ANSENTRY from [%s]", pkg.SenderAddr)
	addUser(pkg)
}

func onIpMsgSendMsg(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
	message, _ := simplifiedchinese.GBK.NewDecoder().Bytes(pkg.AdditionalSection)
	logger.Info("pkg IPMSG_SENDMSG from [%s]# %s\n", pkg.SenderName, message)
	if pkg.CommandNo.CheckOpt(ipmsg.IPMSG_SENDCHECKOPT) {
		im.SendMessageReceived(pkg.SenderAddr, pkg.PacketNo)
	}
	if pkg.CommandNo.CheckOpt(ipmsg.IPMSG_SECRETEXOPT) {
		im.SendMessageRead(pkg.SenderAddr, pkg.PacketNo)
	}
	if pkg.CommandNo.CheckOpt(ipmsg.IPMSG_FILEATTACHOPT) {
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
	logger.Info("%s have received your ipMsg!\n", pkg.SenderName)
}

func onIpMsgNoOperation(im *ipmsg.IpMsg, pkg *ipmsg.Package) {
}
