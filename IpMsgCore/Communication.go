package IpMsgCore

import (
	"errors"
	"fmt"
	"ipmsg"
	"ipmsg/logger"
	"net"
	"os"
	"time"
)

var selfName string
var selfHost string

//接收消息线程，接收其他客户端发送的UDP数据
func PackageHandler(im *ipmsg.IpMsg) {
	for {
		pkg, _ := im.ReadPackage()
		switch pkg.CommandNo.GetCmd() {
		case ipmsg.IPMSG_BR_ENTRY:
			ipmsg.OnIpMsgBrEntry(im, pkg)
		case ipmsg.IPMSG_BR_EXIT:
			ipmsg.OnIpMsgBrExit(im, pkg)
		case ipmsg.IPMSG_ANSENTRY:
			ipmsg.OnIpMsgAnsEntry(im, pkg)
		case ipmsg.IPMSG_SENDMSG:
			ipmsg.OnIpMsgSendMsg(im, pkg)
		case ipmsg.IPMSG_RECVMSG:
			ipmsg.OnIpMsgRecvMsg(im, pkg)
		case ipmsg.IPMSG_NOOPERATION:
			// 无操作忽略
		default:
			logger.Info("recv unknown from [%s]#\n%s\n%X", pkg.SenderName, pkg, pkg)
		}
	}
}

//接收文件(参数为接收文件列表中的序号)
func RecvFile(id int) error {
	// 是否存在该文件
	p := findFile(id)
	if p == nil {
		return errors.New("no such file id")
	}
	s_addr := getAddrByName(p.UserName) //根据发送者姓名获取发送这地址
	if s_addr == nil {
		delFile(p, ipmsg.RECVFILE)
		return errors.New("recv file error: user is not online")
	}
	// 创建临时TCP client用来接收文件
	conn, err := net.Dial("tcp", s_addr.String())
	if err != nil {
		logger.Error("listen error: ", err)
		return err
	}
	defer conn.Close()

	//发送 IPMSG_GETFILEDATA
	t := time.Now().Unix()
	data := fmt.Sprintf("1:%d:%s:%s:%d:%x:%d:0", t, selfName, selfHost, ipmsg.IPMSG_GETFILEDATA, p.Pkgnum, p.Num)
	_, _ = conn.Write([]byte(data))

	// 创建文件并写入文件内容
	fmt.Println(p.Name)
	file, err := os.Create(p.Name)
	if err != nil {
		logger.Error("os.Create err:", err)
		return err
	}
	defer file.Close()

	//接收文件
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			logger.Debug("文件读取完毕")
			break
		}
		if err != nil {
			logger.Error("conn.Read err:", err)
			return err
		}
		_, _ = file.Write(buf[:n])
	}
	//从文件列表中删除接收过的文件
	delFile(p, ipmsg.RECVFILE)
	return nil
}

//发送文件的线程
func sendfile_thread() {
	//int
	//fd = tcp_fd() //获取TCP_Server套接口描述符
	//while(1)
	//{
	//	struct sockaddr_in addr = {
	//		AF_INET
	//	}
	//		unsigned int addrlen = sizeof(addr)
	//		int clifd = accept(fd, (struct sockaddr*)&addr, &addrlen)
	//		if ( clifd<0 )
	//	{
	//		perror("accept")
	//		exit(1)
	//	}
	//		while(1) // 发送多个文件
	//	{
	//		IPMSG_FILE *p = NULL
	//		FILE *fp = NULL
	//		IPMSG_USER user
	//		long pkgnum = 0
	//		char edition[100] = ""
	//		long oldpkgnum = 0
	//		long CommandNo = 0
	//		int filenum = 0
	//		char buf[1400] = ""
	//		int sendsize = 0
	//		//接收IPMSG_GETFILEDATA
	//		if (recv(clifd, buf, sizeof(buf), 0)==0)
	//		break
	//		sscanf(buf, "%[^:]:%ld:%[^:]:%[^:]:%ld:%lx:%x", edition, &pkgnum, user.selfName, user.SenderHost, &CommandNo, \
	//		&oldpkgnum, &filenum)
	//		//是否是IPMSG_GETFILEDATA
	//		if ((GET_MODE(CommandNo)&IPMSG_GETFILEDATA)!=IPMSG_GETFILEDATA)
	//		break
	//		//获取之前发送的文件信息
	//		if ((p = getFileInfo(oldpkgnum, filenum))==NULL)
	//	{
	//		return NULL
	//	}
	//		if ( (fp = fopen(p->selfName, "r"))==NULL )
	//	{
	//		IPMSG_OUT_MSG_COLOR(
	//		printf("senderror: no such file: %s\n", p->selfName)
	//	)
	//		return NULL
	//	}
	//		do //发送文件
	//	{
	//		int size = fread(buf, 1, sizeof(buf), fp)
	//		send(clifd, buf, size, 0)
	//		sendsize += size
	//	}while(sendsize < p->size)
	//		fclose(fp)            //关闭文件
	//		del_file(p, SENDFILE) //从发送文件链表中删除文件
	//	}                         //end wile1 // 循环发送多个文件
	//		close(clifd)          //关闭套接口等待下个用户连接
	//	} //end while
	//	return NULL
}
