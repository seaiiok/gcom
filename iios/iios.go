package iios

import (
	"runtime"
)

type V1 struct {
}

//判断目标是否为微软Windows系统
func (o *V1) IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

//查看目标操作系统
func (o *V1) OsName() string {
	return runtime.GOOS
}
