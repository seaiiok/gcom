package gcom

import (
	"gcom/gcmd"
	"gcom/gconfig"
	"gcom/gconvert"
	"gcom/gos"
	"gcom/gtime"
)

//go 通用帮助类
type goCommon struct {
	GCmd     gcmd.I
	GOs      gos.I
	GTime    gtime.I
	GConvert gconvert.I
	GConfig  gconfig.I
}

// New ...
//
// GCmd 控制台执行类
//
// GOs 系统操作类
//
// GConvert 转换操作类
//
// GConfig 配置文件操作类
//
func New() *goCommon {
	return &goCommon{}
}
