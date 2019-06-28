package gcom

import (
	"gcom/gcmd"
	"gcom/gos"
	"gcom/gtime"
)

//go 通用帮助类
type goCommon struct {
	GCmd  gcmd.I
	GOs   gos.I
	GTime gtime.I
}

//New ...
//
//控制台执行类 gCmd
//
//系统操作类 gOs
//
//时间操作类 gTime
//
func New() *goCommon {
	return &goCommon{}
}
