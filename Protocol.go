package ipmsg

/*  IP Messenger Communication Protocol version 3.0 define  */
/*  macro  */
//go:generate stringer -type=CmdType
type CmdType uint32

//go:generate stringer -type=FileType
type FileType uint32

//go:generate stringer -type=AttrType
type AttrType uint32

//go:generate stringer -type=CmdFlag
type CmdFlag uint32

//go:generate stringer -type=SvrFlag
type SvrFlag CmdFlag

//go:generate stringer -type=EncFlag
type EncFlag CmdFlag

//go:generate stringer -type=FileFlag
type FileFlag CmdFlag

// GetCmd 获取command
//goland:noinspection GoReceiverNames
func (cmd CmdType) GetCmd() CmdType {
	return cmd & 0x000000ff
}

//goland:noinspection GoReceiverNames
func (opt CmdFlag) CheckOpt(cmd CmdType) bool {
	return uint32(cmd)&uint32(opt) != 0
}

//goland:noinspection GoSnakeCaseUsage,GoNameStartsWithPackageName,GoUnusedConst,SpellCheckingInspection
const (
	/*  header  */

	IPMSG_VERSION    = 0x0001
	IPMSG_DEFAT_PORT = 0x0979

	/*  command  基本命令字(32位命令字的低8位) */

	IPMSG_NOOPERATION     CmdType = 0x00000000 //No Operation 没有任何操作
	IPMSG_BR_ENTRY        CmdType = 0x00000001 //Entry to service (Start-up with a Broadcast command) 上线（开始于广播此命令）
	IPMSG_BR_EXIT         CmdType = 0x00000002 //Exit from service (End with a Broadcast command) 下线（结束于广播此命令）
	IPMSG_ANSENTRY        CmdType = 0x00000003 //Notify a new entry 通报新上线
	IPMSG_BR_ABSENCE      CmdType = 0x00000004 //Change absence mode 更改为离开状态
	IPMSG_BR_ISGETLIST    CmdType = 0x00000010 //Search valid sending SenderHost members 搜寻有效的主机用户
	IPMSG_OKGETLIST       CmdType = 0x00000011 //Host list sending notice 主机列表发送通知
	IPMSG_GETLIST         CmdType = 0x00000012 //Host list sending request 主机列表发送请求
	IPMSG_ANSLIST         CmdType = 0x00000013 //Host list sending 主机列表发送
	IPMSG_BR_ISGETLIST2   CmdType = 0x00000018
	IPMSG_SENDMSG         CmdType = 0x00000020 //Message transmission 消息传送
	IPMSG_RECVMSG         CmdType = 0x00000021 //Message receiving check 接收消息确认
	IPMSG_READMSG         CmdType = 0x00000030 //Message open notice 消息打开通知
	IPMSG_DELMSG          CmdType = 0x00000031 //Message discarded notice 消息丢弃通知
	IPMSG_ANSREADMSG      CmdType = 0x00000032 //Message open confirmation notice(added from version-8) 消息打开确认通知（版本8中加入）
	IPMSG_GETINFO         CmdType = 0x00000040 //Get IPMSG version info 获取IPMSG版本信息
	IPMSG_SENDINFO        CmdType = 0x00000041 //Send IPMSG version info 发送IPMSG版本信息
	IPMSG_GETABSENCEINFO  CmdType = 0x00000050 //Get absence sentence 获取离线判定
	IPMSG_SENDABSENCEINFO CmdType = 0x00000051 //Send absence sentence 发送离线判定
	IPMSG_GETFILEDATA     CmdType = 0x00000060 //File Transfer request by TCP 基于TCP的文件传送请求
	IPMSG_RELEASEFILES    CmdType = 0x00000061 //Discard attachment file 丢弃（取消）附件文件的接收
	IPMSG_GETDIRFILES     CmdType = 0x00000062 //Attachment hierarchical file request 文件夹传送请求
	IPMSG_GETPUBKEY       CmdType = 0x00000072 //RSA Public Key Acquisition 公钥获取
	IPMSG_ANSPUBKEY       CmdType = 0x00000073 //RSA Public Key Response 公钥响应

	/*  option for all command  */

	IPMSG_ABSENCEOPT SvrFlag = 0x00000100 //Absence mode(Member recognition command) 离开状态（用户识别命令）
	IPMSG_SERVEROPT  SvrFlag = 0x00000200 //Server(Reserved) 服务器（保留）
	IPMSG_DIALUPOPT  SvrFlag = 0x00010000 //Send individual member recognition command 发送个人用户识别命令

	/*  option for send command  */

	IPMSG_SENDCHECKOPT  CmdFlag = 0x00000100 //Transmission check 传送检查(需要对方返回确认信息)
	IPMSG_SECRETOPT     CmdFlag = 0x00000200 //Sealed message 封闭信息
	IPMSG_BROADCASTOPT  CmdFlag = 0x00000400 //Broadcast message 广播信息
	IPMSG_MTICASTOPT    CmdFlag = 0x00000800 //Multi-cast(Multiple casts selection) 多播
	IPMSG_NOPOPUPOPT    CmdFlag = 0x00001000 //obsolete (No longer valid) （不可用）
	IPMSG_AUTORETOPT    CmdFlag = 0x00002000 //Automatic response(Ping-pong protection) 自动回复
	IPMSG_RETRYOPT      CmdFlag = 0x00004000 //Re-send flag(Use when acquiring HOSTLIST) 重发位（在获取HOSTLIST时使用）
	IPMSG_PASSWORDOPT   CmdFlag = 0x00008000 //Lock 锁
	IPMSG_NOLOGOPT      CmdFlag = 0x00020000 //No log files 无日志文件
	IPMSG_NEWMUTIOPT    CmdFlag = 0x00040000 //obsolete New version multi-cast(reserved) 新版本多播
	IPMSG_NOADDLISTOPT  CmdFlag = 0x00080000 //Notice to the members outside of BR_ENTRY 不在线用户通知
	IPMSG_READCHECKOPT  CmdFlag = 0x00100000 //Sealed message check(added from ver8 ) 封闭信息检查（版本8中加入）
	IPMSG_FILEATTACHOPT CmdFlag = 0x00200000 //传送文件选项
	IPMSG_ENCRYPTOPT    CmdFlag = 0x00400000 //密码 （第10版）
	IPMSG_UTF8OPT       CmdFlag = 0x00800000 //使用UTF-8（第10版）
	IPMSG_CAPUTF8OPT    CmdFlag = 0x01000000 //消息使用 UTF-8（第10版）
	IPMSG_ENCEXTMSGOPT  CmdFlag = 0x04000000 //附件加密请求
	IPMSG_CLIPBOARDOPT  CmdFlag = 0x08000000 //全部使用UTF-8 （第10版）
	IPMSG_SECRETEXOPT           = IPMSG_READCHECKOPT | IPMSG_SECRETOPT

	/* encryption/capability flags for encrypt command */

	IPMSG_RSA_512       EncFlag = 0x00000001
	IPMSG_RSA_1024      EncFlag = 0x00000002
	IPMSG_RSA_2048      EncFlag = 0x00000004
	IPMSG_RC2_40        EncFlag = 0x00001000
	IPMSG_RC2_128       EncFlag = 0x00004000
	IPMSG_RC2_256       EncFlag = 0x00008000
	IPMSG_BLOWFISH_128  EncFlag = 0x00020000
	IPMSG_BLOWFISH_256  EncFlag = 0x00040000
	IPMSG_AES_256       EncFlag = 0x00100000
	IPMSG_PACKETNO_IV   EncFlag = 0x00800000
	IPMSG_ENCODE_BASE64 EncFlag = 0x01000000
	IPMSG_SIGN_MD5      EncFlag = 0x10000000
	IPMSG_SIGN_SHA1     EncFlag = 0x20000000

	/* compatibilty for Win beta version */

	IPMSG_RC2_40OLD         EncFlag = 0x00000010 // for beta1-4 only
	IPMSG_RC2_128OLD        EncFlag = 0x00000040 // for beta1-4 only
	IPMSG_BLOWFISH_128OLD   EncFlag = 0x00000400 // for beta1-4 only
	IPMSG_RC2_128OBSOLETE   EncFlag = 0x00004000
	IPMSG_RC2_256OBSOLETE   EncFlag = 0x00008000
	IPMSG_BLOWFISH_256OBSOL EncFlag = 0x00040000
	IPMSG_AES_128OBSOLETE   EncFlag = 0x00080000
	IPMSG_UNAMEEXTOPTOBSOLT EncFlag = 0x02000000
	IPMSG_SIGN_MD5OBSOLETE  EncFlag = 0x10000000
	IPMSG_RC2_40ALL                 = IPMSG_RC2_40 | IPMSG_RC2_40OLD
	IPMSG_RC2_128ALL                = IPMSG_RC2_128 | IPMSG_RC2_128OLD
	IPMSG_BLOWFISH_128ALL           = IPMSG_BLOWFISH_128 | IPMSG_BLOWFISH_128OLD

	/* file types for fileattach command */

	IPMSG_FILE_REGULAR   FileType = 0x00000001
	IPMSG_FILE_DIR       FileType = 0x00000002
	IPMSG_FILE_RETPARENT FileType = 0x00000003 // return parent directory
	IPMSG_FILE_SYMLINK   FileType = 0x00000004
	IPMSG_FILE_CDEV      FileType = 0x00000005 // for UNIX
	IPMSG_FILE_BDEV      FileType = 0x00000006 // for UNIX
	IPMSG_FILE_FIFO      FileType = 0x00000007 // for UNIX
	IPMSG_FILE_RESFORK   FileType = 0x00000010 // for Mac
	IPMSG_FILE_CLIPBOARD FileType = 0x00000020 // for Windows Clipboard

	/* file attribute options for fileattach command */

	IPMSG_FILE_RONLYOPT    FileFlag = 0x00000100
	IPMSG_FILE_HIDDENOPT   FileFlag = 0x00001000
	IPMSG_FILE_EXHIDDENOPT FileFlag = 0x00002000 // for MacOS X
	IPMSG_FILE_ARCHIVEOPT  FileFlag = 0x00004000
	IPMSG_FILE_SYSTEMOPT   FileFlag = 0x00008000

	/* extend attribute types for fileattach command */

	IPMSG_FILE_UID          AttrType = 0x00000001
	IPMSG_FILE_USERNAME     AttrType = 0x00000002 // uid by string
	IPMSG_FILE_GID          AttrType = 0x00000003
	IPMSG_FILE_GROUPNAME    AttrType = 0x00000004 // gid by string
	IPMSG_FILE_CLIPBOARDPOS AttrType = 0x00000008 //
	IPMSG_FILE_PERM         AttrType = 0x00000010 // for UNIX
	IPMSG_FILE_MAJORNO      AttrType = 0x00000011 // for UNIX devfile
	IPMSG_FILE_MINORNO      AttrType = 0x00000012 // for UNIX devfile
	IPMSG_FILE_CTIME        AttrType = 0x00000013 // for UNIX
	IPMSG_FILE_MTIME        AttrType = 0x00000014
	IPMSG_FILE_ATIME        AttrType = 0x00000015
	IPMSG_FILE_CREATETIME   AttrType = 0x00000016
	IPMSG_FILE_CREATOR      AttrType = 0x00000020 // for Mac
	IPMSG_FILE_FILETYPE     AttrType = 0x00000021 // for Mac
	IPMSG_FILE_FINDERINFO   AttrType = 0x00000022 // for Mac
	IPMSG_FILE_ACL          AttrType = 0x00000030
	IPMSG_FILE_ALIASFNAME   AttrType = 0x00000040 // alias fname
	IPMSG_FILE_UNICODEFNAME AttrType = 0x00000041 // UNICODE fname

	FILELIST_SEPARATOR = '\a'
	HOSTLIST_SEPARATOR = '\a'
	HOSTLIST_DUMMY     = "\b"

	/*  end of IP Messenger Communication Protocol version 1.2 define  */

	/*  IP Messenger for Windows  internal define  */

	IPMSG_REVERSEICON    = 0x0100
	IPMSG_TIMERINTERVAL  = 500
	IPMSG_ENTRYMINSEC    = 5
	IPMSG_GETLIST_FINISH = 0

	IPMSG_BROADCAST_TIMER    = 0x0101
	IPMSG_SEND_TIMER         = 0x0102
	IPMSG_LISTGET_TIMER      = 0x0104
	IPMSG_LISTGETRETRY_TIMER = 0x0105
	IPMSG_ENTRY_TIMER        = 0x0106
	IPMSG_DUMMY_TIMER        = 0x0107
	IPMSG_RECV_TIMER         = 0x0108
	IPMSG_ANS_TIMER          = 0x0109

	IPMSG_NICKNAME = 1
	IPMSG_FLNAME   = 2

	IPMSG_NAMESORT        = uint32(0x00000000)
	IPMSG_IPADDRSORT      = uint32(0x00000001)
	IPMSG_HOSTSORT        = uint32(0x00000002)
	IPMSG_NOGROUPSORTOPT  = uint32(0x00000100)
	IPMSG_ICMPSORTOPT     = uint32(0x00000200)
	IPMSG_NOKANJISORTOPT  = uint32(0x00000400)
	IPMSG_ALLREVSORTOPT   = uint32(0x00000800)
	IPMSG_GROUPREVSORTOPT = uint32(0x00001000)
	IPMSG_SUBREVSORTOPT   = uint32(0x00002000)
)
