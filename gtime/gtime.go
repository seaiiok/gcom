package gtime

import (
	"gcom/gconst"
	"time"
)

//时间操作类 版本1.0
type I struct{}

// 格式时间<2006-01-02 15:04:05>转化为时间戳
func (g *I) TimeString2Unix(timeString string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(gconst.TimeLayout, timeString, loc)

	return theTime.Unix()
}

// 时间戳转化为格式时间<2006-01-02 15:04:05>
func (g *I) Unix2TimeString(unix int64) string {
	return time.Unix(unix, 0).Format(gconst.TimeLayout)
}
