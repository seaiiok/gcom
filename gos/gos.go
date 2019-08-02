package gos

import (
	"runtime"
)

// IsWindows 判断目标是否为微软Windows系统
func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

// OsName 查看目标操作系统
func OsName() string {
	return runtime.GOOS
}
