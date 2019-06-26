package iitime

import (
	"ii/iiconst"
	"time"
)

type V1 struct {
}

//New ...return V1
func New() *V1 {
	return &V1{}
}

// 格式时间<2006-01-02 15:04:05>转化为时间戳
func (t *V1) TimeString2Unix(timeString string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(iiconst.TimeLayout, timeString, loc)

	return theTime.Unix()
}

// 时间戳转化为格式时间<2006-01-02 15:04:05>
func (t *V1) Unix2TimeString(unix int64) string {
	return time.Unix(unix, 0).Format(iiconst.TimeLayout)
}