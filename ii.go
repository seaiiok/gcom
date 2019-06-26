package ii

import (
	"ii/iicmd"
	"ii/iios"
	"ii/iitime"
)

type ii struct {
	CmdV1  iicmd.V1
	OsV1   iios.V1
	TimeV1 iitime.V1
}

//New ...
//
//控制台执行类 IIcmd
//
//系统操作类 IIos
//
//时间操作类 IItime
//
func New() *ii {
	return &ii{}
}
