package ipmsg

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

type Package struct {
	//"1:100:shirouzu:jupiter:32:Hello"
	Buf               []byte
	Ver               []byte //版本
	PacketNo          uint32 //包编号
	SenderName        string //发送者名字
	SenderHost        string //发送主机
	CommandNo         uint32 //命令编号
	AdditionalSection []byte //附加信息区域
}

func (pkg Package) Marshal() []byte {
	data := fmt.Sprintf("%d:%d:%s:%s:%d:%s", IPMSG_VERSION, time.Now().Unix(), pkg.SenderName, pkg.SenderHost, pkg.CommandNo, pkg.AdditionalSection)
	return []byte(data)
}

func (pkg Package) UnMarshal(data []byte) (*Package, error) {
	pkg.Buf = data
	s := bytes.SplitN(data, []byte{':'}, 6)
	pkg.Ver = s[0]
	pkgNo, err := strconv.Atoi(string(s[1]))
	if err != nil {
		return nil, err
	}
	pkg.PacketNo = uint32(pkgNo)
	pkg.SenderName = string(s[2])
	pkg.SenderHost = string(s[3])
	cmdNo, err := strconv.Atoi(string(s[4]))
	if err != nil {
		return nil, err
	}
	pkg.CommandNo = uint32(cmdNo)
	pkg.AdditionalSection = s[5]
	return &pkg, nil
}

//UnMarshal
func UnMarshal(data []byte) (*Package, error) {
	return new(Package).UnMarshal(data)
}
