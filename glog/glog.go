// user "\r\n"
// +build only windows

package glog

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	logname  = "log "
	enter    = "\r\n\r\n"
	LOGALL   = 0
	LOGINFO  = 1
	LOGWARN  = 2
	LOGERROR = 3
)

var (
	lastmouth = "200601"
	lastday   = "2006-01-02"
	once      = sync.Once{}
	file      = &os.File{}
	sw        = &sync.RWMutex{}
	path      = ""
	level     = 0
)

func Init(logpath string, loglevel int) {
	sw.Lock()
	defer sw.Unlock()

	path = logpath
	level = loglevel

	pmouth := time.Now().Format("200601")

	if !strings.EqualFold(pmouth, lastmouth) {
		file.Close()

		extime := time.Now().Format("200601")
		logfile := path + "/" + logname + extime + ".log"
		if _, err := os.Stat(path); !os.IsExist(err) {
			os.MkdirAll(path, 0644)
		}

		file, _ = os.OpenFile(logfile, os.O_CREATE|os.O_APPEND, 0644)
		lastmouth = pmouth
	}

}

func Info(v ...interface{}) {
	if level >= LOGALL && level <= LOGINFO {
		Init(path, level)
		logTextFormat("INFO", fmt.Sprintf("%s", v...))
	}
}

func Warn(v ...interface{}) {
	if level >= LOGALL && level <= LOGWARN {
		Init(path, level)
		logTextFormat("WARN", fmt.Sprintf("%s", v...))
	}
}

func Error(v ...interface{}) {
	if level >= LOGALL && level <= LOGERROR {
		Init(path, level)
		logTextFormat("ERRS", fmt.Sprintf("%s", v...))
	}
}

func logTextFormat(levelinfo string, v string) {
	sw.Lock()
	defer sw.Unlock()

	pday := time.Now().Format("2006-01-02")
	if !strings.EqualFold(pday, lastday) {
		file.WriteString("=========================== [" + pday + "] ===========================" + enter)
		lastday = pday
	}
	time := time.Now().Format("15:04:05.000000")
	file.WriteString("[" + levelinfo + "] " + time + " " + v + enter)
}

func Close() {
	if file != nil {
		file.Close()
	}
}
