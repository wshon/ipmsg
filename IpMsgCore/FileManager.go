package IpMsgCore

import (
	"ipmsg"
)

//向链表(接收或发送链表)中添加一个文件
//flag : SENDFILE(发送链表) RECVFILE(接收链表)
func add_file(temp *ipmsg.FileInfo, flag int) {}

//从链表(接收或发送链表)中删除一个文件
//flag : SENDFILE(发送链表) RECVFILE(接收链表)
func del_file(temp *ipmsg.FileInfo, flag int) {}

//在接收链表中按照序号(id)查找文件
func find_file(id int) *ipmsg.FileInfo { return nil }

//根据包标号和文件序号从发送链表中获取文件
func getfileinfo(pkgnum int, filenum int) *ipmsg.FileInfo { return nil }

//获取文件属性存放在fileopt中,并将文件信息存入发送链表中
func getfileopt(fileopt string, filename string, pkgnum int, num int) {}

//打印接收文件链表
func file_list() {}
