package ipmsg

import "net"

type IUserManager interface {
	Init()
	AddUser(pkg *Package) interface{ IUserInfo }
	DelUser(pkg *Package)
	//根据用户名返回IP地址
	GetAddrByName(name string) *net.UDPAddr
}

type IUserInfo interface {
	GetName()
	GetAddr() *net.UDPAddr
}

type IFileManager interface {
	GetAddrByName(name string) *IFileManager
	//向链表(接收或发送链表)中添加一个文件
	//flag : SENDFILE(发送链表) RECVFILE(接收链表)
	addFile(temp *IFileManager, flag int)

	//从链表(接收或发送链表)中删除一个文件
	//flag : SENDFILE(发送链表) RECVFILE(接收链表)
	delFile(temp *IFileManager, flag int)

	//在接收链表中按照序号(id)查找文件
	findFile(id int) *IFileManager

	//根据包标号和文件序号从发送链表中获取文件
	getFileInfo(pkgnum int, filenum int) *IFileManager

	//获取文件属性存放在fileopt中,并将文件信息存入发送链表中
	getFileOpt(fileopt string, filename string, pkgnum int, num int)

	//打印接收文件链表
	fileList()
}
