package ipmsg

/*  IP Messenger Communication Protocol version 3.0 define  */
/*  macro  */

// 获取command
func GET_MODE(command uint32) uint32 { return command & 0x000000ff }

// 获取选项option
func GET_OPT(command uint32) uint32 { return command & 0xffffff00 }

const (
	/*  header  */

	IPMSG_VERSION    = 0x0001
	IPMSG_DEFAT_PORT = 0x0979

	/*  command  基本命令字(32位命令字的低8位) */

	IPMSG_NOOPERATION = uint32(0x00000000) //No Operation 没有任何操作

	IPMSG_BR_ENTRY   = uint32(0x00000001) //Entry to service (Start-up with a Broadcast command) 上线（开始于广播此命令）
	IPMSG_BR_EXIT    = uint32(0x00000002) //Exit from service (End with a Broadcast command) 下线（结束于广播此命令）
	IPMSG_ANSENTRY   = uint32(0x00000003) //Notify a new entry 通报新上线
	IPMSG_BR_ABSENCE = uint32(0x00000004) //Change absence mode 更改为离开状态

	IPMSG_BR_ISGETLIST  = uint32(0x00000010) //Search valid sending SenderHost members 搜寻有效的主机用户
	IPMSG_OKGETLIST     = uint32(0x00000011) //Host list sending notice 主机列表发送通知
	IPMSG_GETLIST       = uint32(0x00000012) //Host list sending request 主机列表发送请求
	IPMSG_ANSLIST       = uint32(0x00000013) //Host list sending 主机列表发送
	IPMSG_BR_ISGETLIST2 = uint32(0x00000018)

	IPMSG_SENDMSG    = uint32(0x00000020) //Message transmission 消息传送
	IPMSG_RECVMSG    = uint32(0x00000021) //Message receiving check 接收消息确认
	IPMSG_READMSG    = uint32(0x00000030) //Message open notice 消息打开通知
	IPMSG_DELMSG     = uint32(0x00000031) //Message discarded notice 消息丢弃通知
	IPMSG_ANSREADMSG = uint32(0x00000032) //Message open confirmation notice(added from version-8) 消息打开确认通知（版本8中加入）

	IPMSG_GETINFO  = uint32(0x00000040) //Get IPMSG version info 获取IPMSG版本信息
	IPMSG_SENDINFO = uint32(0x00000041) //Send IPMSG version info 发送IPMSG版本信息

	IPMSG_GETABSENCEINFO  = uint32(0x00000050) //Get absence sentence 获取离线判定
	IPMSG_SENDABSENCEINFO = uint32(0x00000051) //Send absence sentence 发送离线判定

	IPMSG_GETFILEDATA  = uint32(0x00000060) //File Transfer request by TCP 基于TCP的文件传送请求
	IPMSG_RELEASEFILES = uint32(0x00000061) //Discard attachment file 丢弃（取消）附件文件的接收
	IPMSG_GETDIRFILES  = uint32(0x00000062) //Attachment hierarchical file request 文件夹传送请求

	IPMSG_GETPUBKEY = uint32(0x00000072) //RSA Public Key Acquisition 公钥获取
	IPMSG_ANSPUBKEY = uint32(0x00000073) //RSA Public Key Response 公钥响应

	/*  option for all command  */

	IPMSG_ABSENCEOPT    = uint32(0x00000100) //Absence mode(Member recognition command) 离开状态（用户识别命令）
	IPMSG_SERVEROPT     = uint32(0x00000200) //Server(Reserved) 服务器（保留）
	IPMSG_DIALUPOPT     = uint32(0x00010000) //Send individual member recognition command 发送个人用户识别命令
	IPMSG_FILEATTACHOPT = uint32(0x00200000) //传送文件选项
	IPMSG_ENCRYPTOPT    = uint32(0x00400000)
	IPMSG_UTF8OPT       = uint32(0x00800000)
	IPMSG_CAPUTF8OPT    = uint32(0x01000000)
	IPMSG_ENCEXTMSGOPT  = uint32(0x04000000)
	IPMSG_CLIPBOARDOPT  = uint32(0x08000000)

	/*  option for send command  */

	IPMSG_SENDCHECKOPT = uint32(0x00000100) //Transmission check 传送检查(需要对方返回确认信息)
	IPMSG_SECRETOPT    = uint32(0x00000200) //Sealed message 封闭信息
	IPMSG_BROADCASTOPT = uint32(0x00000400) //Broadcast message 广播信息
	IPMSG_MTICASTOPT   = uint32(0x00000800) //Multi-cast(Multiple casts selection) 多播
	IPMSG_NOPOPUPOPT   = uint32(0x00001000) //(No longer valid) （不可用）
	IPMSG_AUTORETOPT   = uint32(0x00002000) //Automatic response(Ping-pong protection) 自动回复
	IPMSG_RETRYOPT     = uint32(0x00004000) //Re-send flag(Use when acquiring HOSTLIST) 重发位（在获取HOSTLIST时使用）
	IPMSG_PASSWORDOPT  = uint32(0x00008000) //Lock 锁
	IPMSG_NOLOGOPT     = uint32(0x00020000) //No log files 无日志文件
	IPMSG_NEWMUTIOPT   = uint32(0x00040000) //New version multi-cast(reserved) 新版本多播
	IPMSG_NOADDLISTOPT = uint32(0x00080000) //Notice to the members outside of BR_ENTRY 不在线用户通知
	IPMSG_READCHECKOPT = uint32(0x00100000) //Sealed message check(added from ver8 ) 封闭信息检查（版本8中加入）
	IPMSG_SECRETEXOPT  = IPMSG_READCHECKOPT | IPMSG_SECRETOPT

	/*  obsolete option for send command  */

	IPMSG_NOPOPUPOPTOBSOLT  = uint32(0x00001000)
	IPMSG_NEWMULTIOPTOBSOLT = uint32(0x00040000)

	/* encryption/capability flags for encrypt command */

	IPMSG_RSA_512       = uint32(0x00000001)
	IPMSG_RSA_1024      = uint32(0x00000002)
	IPMSG_RSA_2048      = uint32(0x00000004)
	IPMSG_RC2_40        = uint32(0x00001000)
	IPMSG_RC2_128       = uint32(0x00004000)
	IPMSG_RC2_256       = uint32(0x00008000)
	IPMSG_BLOWFISH_128  = uint32(0x00020000)
	IPMSG_BLOWFISH_256  = uint32(0x00040000)
	IPMSG_AES_256       = uint32(0x00100000)
	IPMSG_PACKETNO_IV   = uint32(0x00800000)
	IPMSG_ENCODE_BASE64 = uint32(0x01000000)
	IPMSG_SIGN_MD5      = uint32(0x10000000)
	IPMSG_SIGN_SHA1     = uint32(0x20000000)

	/* compatibilty for Win beta version */

	IPMSG_RC2_40OLD         = uint32(0x00000010) // for beta1-4 only
	IPMSG_RC2_128OLD        = uint32(0x00000040) // for beta1-4 only
	IPMSG_BLOWFISH_128OLD   = uint32(0x00000400) // for beta1-4 only
	IPMSG_RC2_128OBSOLETE   = uint32(0x00004000)
	IPMSG_RC2_256OBSOLETE   = uint32(0x00008000)
	IPMSG_BLOWFISH_256OBSOL = uint32(0x00040000)
	IPMSG_AES_128OBSOLETE   = uint32(0x00080000)
	IPMSG_UNAMEEXTOPTOBSOLT = uint32(0x02000000)
	IPMSG_SIGN_MD5OBSOLETE  = uint32(0x10000000)
	IPMSG_RC2_40ALL         = (IPMSG_RC2_40 | IPMSG_RC2_40OLD)
	IPMSG_RC2_128ALL        = (IPMSG_RC2_128 | IPMSG_RC2_128OLD)
	IPMSG_BLOWFISH_128ALL   = (IPMSG_BLOWFISH_128 | IPMSG_BLOWFISH_128OLD)

	/* file types for fileattach command */

	IPMSG_FILE_REGULAR   = uint32(0x00000001)
	IPMSG_FILE_DIR       = uint32(0x00000002)
	IPMSG_FILE_RETPARENT = uint32(0x00000003) // return parent directory
	IPMSG_FILE_SYMLINK   = uint32(0x00000004)
	IPMSG_FILE_CDEV      = uint32(0x00000005) // for UNIX
	IPMSG_FILE_BDEV      = uint32(0x00000006) // for UNIX
	IPMSG_FILE_FIFO      = uint32(0x00000007) // for UNIX
	IPMSG_FILE_RESFORK   = uint32(0x00000010) // for Mac
	IPMSG_FILE_CLIPBOARD = uint32(0x00000020) // for Windows Clipboard

	/* file attribute options for fileattach command */

	IPMSG_FILE_RONLYOPT    = uint32(0x00000100)
	IPMSG_FILE_HIDDENOPT   = uint32(0x00001000)
	IPMSG_FILE_EXHIDDENOPT = uint32(0x00002000) // for MacOS X
	IPMSG_FILE_ARCHIVEOPT  = uint32(0x00004000)
	IPMSG_FILE_SYSTEMOPT   = uint32(0x00008000)

	/* extend attribute types for fileattach command */

	IPMSG_FILE_UID          = uint32(0x00000001)
	IPMSG_FILE_USERNAME     = uint32(0x00000002) // uid by string
	IPMSG_FILE_GID          = uint32(0x00000003)
	IPMSG_FILE_GROUPNAME    = uint32(0x00000004) // gid by string
	IPMSG_FILE_CLIPBOARDPOS = uint32(0x00000008) //
	IPMSG_FILE_PERM         = uint32(0x00000010) // for UNIX
	IPMSG_FILE_MAJORNO      = uint32(0x00000011) // for UNIX devfile
	IPMSG_FILE_MINORNO      = uint32(0x00000012) // for UNIX devfile
	IPMSG_FILE_CTIME        = uint32(0x00000013) // for UNIX
	IPMSG_FILE_MTIME        = uint32(0x00000014)
	IPMSG_FILE_ATIME        = uint32(0x00000015)
	IPMSG_FILE_CREATETIME   = uint32(0x00000016)
	IPMSG_FILE_CREATOR      = uint32(0x00000020) // for Mac
	IPMSG_FILE_FILETYPE     = uint32(0x00000021) // for Mac
	IPMSG_FILE_FINDERINFO   = uint32(0x00000022) // for Mac
	IPMSG_FILE_ACL          = uint32(0x00000030)
	IPMSG_FILE_ALIASFNAME   = uint32(0x00000040) // alias fname
	IPMSG_FILE_UNICODEFNAME = uint32(0x00000041) // UNICODE fname

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
