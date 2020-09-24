package ipmsg

import (
	"bytes"
	"fmt"
	"net"
	"reflect"
	"strconv"
)

type Package struct {
	//"1:100:shirouzu:jupiter:32:Hello"
	Buf               []byte
	Ver               string  //版本
	PacketNo          uint32  //包编号
	SenderName        string  //发送者名字
	SenderHost        string  //发送者主机
	CommandNo         CmdType //命令编号
	AdditionalSection string  //附加信息区域

	SenderAddr *net.UDPAddr //发送者地址
}

// String implements fmt.Stringer interface
func (pkg *Package) String() string {
	return fmt.Sprintf("Ver:%s, PacketNo: %d, SenderName: %s, SenderHost: %s, CommandNo: %d, AdditionalSection: %s", pkg.Ver, pkg.PacketNo, pkg.SenderName, pkg.SenderHost, pkg.CommandNo, pkg.AdditionalSection)
}

// Format implements fmt.Formatter interface
func (pkg *Package) Format(state fmt.State, verb rune) {
	switch verb {
	case 's', 'q':
		val := pkg.String()
		if verb == 'q' {
			val = fmt.Sprintf("%q", val)
		}
		_, _ = fmt.Fprint(state, val)
	case 'v':
		if state.Flag('#') {
			// Emit type before
			_, _ = fmt.Fprintf(state, "%T", pkg)
		}
		_, _ = fmt.Fprint(state, "{")
		typ := reflect.TypeOf(Package{})
		val := reflect.ValueOf(*pkg)
		for i := 0; i < typ.NumField(); i++ {
			name := typ.Field(i).Name
			if name == "Buf" {
				continue
			}
			if state.Flag('#') || state.Flag('+') {
				_, _ = fmt.Fprintf(state, "%s:", name)
			}
			fld := val.FieldByName(name)
			if name == "AdditionalSection" && fld.Len() > 0 {
				_, _ = fmt.Fprintf(state, "%X", []byte(fld.String()))
			} else {
				_, _ = fmt.Fprint(state, fld)
			}
			if i < typ.NumField()-1 {
				_, _ = fmt.Fprint(state, " ")
			}
		}
		_, _ = fmt.Fprint(state, "}")
	}
}

// Marshal
func (pkg *Package) Marshal() []byte {
	data := fmt.Sprintf("%s:%d:%s:%s:%d:%s", pkg.Ver, pkg.PacketNo, pkg.SenderName, pkg.SenderHost, pkg.CommandNo, pkg.AdditionalSection)
	return []byte(data)
}

// UnMarshal
func (pkg *Package) UnMarshal(data []byte) (*Package, error) {
	pkg.Buf = bytes.Trim(data, "\x00")
	s := bytes.SplitN(pkg.Buf, []byte{':'}, 6)
	pkg.Ver = string(s[0])
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
	pkg.CommandNo = CmdType(cmdNo)
	pkg.AdditionalSection = string(s[5])
	return pkg, nil
}

func (pkg *Package) SetFlag(flag CmdFlag) {
	pkg.CommandNo = pkg.CommandNo | CmdType(flag)
}

func (pkg *Package) ClearFlag(flag CmdFlag) {
	pkg.CommandNo = pkg.CommandNo & CmdType(^flag)
}

func (pkg *Package) CheckFlag(flag interface{}) bool {
	switch flag.(type) {
	case CmdFlag:
		return uint32(pkg.CommandNo)&uint32(flag.(CmdFlag)) != 0
	}
	return false
}

// UnMarshal
func UnMarshal(data []byte) (*Package, error) {
	return new(Package).UnMarshal(data)
}
