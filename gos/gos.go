package gos

import (
	"runtime"
)

//版本1.0
type I struct{}

//判断目标是否为微软Windows系统
func (g *I) IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

//查看目标操作系统
func (g *I) OsName() string {
	return runtime.GOOS
}
