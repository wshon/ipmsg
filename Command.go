package ipmsg

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"ipmsg/logger"
)

func NewCmdMap() (cmdMap map[CmdType]func(*IpMsg, *Package)) {
	cmdMap = make(map[CmdType]func(*IpMsg, *Package))
	cmdMap[IPMSG_BR_ENTRY] = OnIpMsgBrEntry
	cmdMap[IPMSG_BR_EXIT] = OnIpMsgBrExit
	cmdMap[IPMSG_ANSENTRY] = OnIpMsgAnsEntry
	cmdMap[IPMSG_SENDMSG] = OnIpMsgSendMsg
	cmdMap[IPMSG_RECVMSG] = OnIpMsgRecvMsg
	cmdMap[IPMSG_NOOPERATION] = onIpMsgNoOperation
	return cmdMap
}

func OnIpMsgBrEntry(im *IpMsg, pkg *Package) {
	im.userManager.AddUser(pkg)
	im.SendEntryAnswer(pkg.SenderAddr)
}

func OnIpMsgBrExit(im *IpMsg, pkg *Package) {
	im.userManager.DelUser(pkg)
}

func OnIpMsgAnsEntry(im *IpMsg, pkg *Package) {
	im.userManager.AddUser(pkg)
}

func OnIpMsgSendMsg(im *IpMsg, pkg *Package) {
	message, _ := simplifiedchinese.GBK.NewDecoder().String(pkg.AdditionalSection)
	logger.Info("new msg from [%s]# %s\n", pkg.SenderName, message)
	if pkg.CheckFlag(IPMSG_SENDCHECKOPT) {
		im.SendMessageReceived(pkg.SenderAddr, pkg.PacketNo)
	}
	if pkg.CheckFlag(IPMSG_SECRETEXOPT) {
		im.SendMessageRead(pkg.SenderAddr, pkg.PacketNo)
	}
	if pkg.CheckFlag(IPMSG_FILEATTACHOPT) {
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

func OnIpMsgRecvMsg(im *IpMsg, pkg *Package) {
	logger.Info("[%s] have received your ipMsg!\n", pkg.SenderName)
}

func onIpMsgNoOperation(im *IpMsg, pkg *Package) {
}
