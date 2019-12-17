package myutil

import (
	"fmt"
	"time"
)

func Util_time(tm string) string {

	length := len(tm)
	if length == 16 {
		return tm
	} else if length < 16 {
		for i := length; i < 16; i++ {
			tm = tm + "0"
		}
		return tm
	} else if length > 16 {
		return tm[:15]
	}
	return ""
}

func Util_strtotime(tm string) string {
	formatTimeStr := tm
	formatTime, err := time.Parse("2006-01-02 15:04:05", formatTimeStr)
	if err == nil {
		Time := fmt.Sprintf("%v", formatTime.Unix())
		//fmt.Println(Util_time(Time)) //打印结果：2017-04-11 13:33:37 +0000 UTC
		return Util_time(Time)
	}
	return ""
}
