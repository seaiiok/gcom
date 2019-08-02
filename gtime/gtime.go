package gtime

import (
	"time"
)

//时间格式
const (
	TimeLayout = "2006-01-02 15:04:05"
)

// TimeString2Unix 格式时间<2006-01-02 15:04:05>转化为时间戳
func TimeString2Unix(timeString string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeLayout, timeString, loc)

	return theTime.Unix()
}

// Unix2TimeString 时间戳转化为格式时间<2006-01-02 15:04:05>
func Unix2TimeString(unix int64) string {
	return time.Unix(unix, 0).Format(TimeLayout)
}
