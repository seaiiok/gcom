package gcmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"syscall"
)

var OFF bool = false

const (
	Info = 8
	Ok   = 9
	Warn = 14
	Err  = 12
)

//常用颜色枚举
const (
	FontGrey = iota + 8
	FontBlue1
	FontGreen
	FontBlue2
	FontRed
	FontPurple
	FontYellow
	FontWhite
)

//控制台执行命令
func ExecCommand(name string, arg ...string) (output string) {

	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err.Error()
	}

	defer stdout.Close()

	err = cmd.Start()
	if err != nil {
		return err.Error()
	}

	defer cmd.Wait()

	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return err.Error()
	} else {
		return string(opBytes)
	}
}

//Println 终端输出,增加颜色参数，类似fmt.Println
func Println(color int, a ...interface{}) (n int, err error) {
	if OFF {
		return
	}

	if isWindows() {
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		proc := kernel32.NewProc("SetConsoleTextAttribute")
		handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(color))
		n, err = fmt.Println(a...)
		handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
		CloseHandle := kernel32.NewProc("CloseHandle")
		CloseHandle.Call(handle)
		return
	} else {
		n, err = fmt.Println(a...)
		err = errors.New("is not windows os system")
		return
	}

}

//Printf 终端格式输出,增加颜色参数，类似fmt.Printf
func Printf(color int, format string, a ...interface{}) (n int, err error) {
	if OFF {
		return
	}

	if isWindows() {
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		proc := kernel32.NewProc("SetConsoleTextAttribute")
		handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(color))
		n, err = fmt.Printf(format, a...)
		handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
		CloseHandle := kernel32.NewProc("CloseHandle")
		CloseHandle.Call(handle)
		return
	} else {
		n, err = fmt.Printf(format, a...)
		err = errors.New("is not windows os system")
		return
	}

}

func isWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}
