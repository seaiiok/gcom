package ii

import (
	"ii/iicmd"
	"ii/iios"
	"ii/iitime"
)

type ii struct {
	IIcmd  iicmd.IIcmd
	IIos   iios.IIos
	IItime iitime.IItime
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
