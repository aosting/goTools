package time2

/**********************************************
** @Des: time
** @Author: zhangxueyuan 
** @Date:   2018-12-27 16:44:50
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-12-27 16:44:50
***********************************************/

import (
	"time"
)

//当前时间的基础上增加指定小时
func AddHour(h int64) time.Time {
	t := time.Duration(h) * time.Hour
	return time.Now().Add(t)
}

func Millisecond(t time.Time) int64 {
	return (t.UnixNano() / 1e6)
}

func Unix2Str(sr int64) (string) {
	local, _ := time.LoadLocation("Asia/Chongqing")
	return time.Unix(sr, 0).In(local).Format("1504")
}

func Str2Unix(ts string) (int64, error) {
	t, error := time.Parse("2006-01-02 15:04:05", ts)
	if error != nil {
		return 0, error
	}
	return t.Unix(), error
}

func Str2UnixCN(ts string) (int64, error) {
	local, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return 0, err
	}
	t, error := time.ParseInLocation("2006-01-02 15:04:05", ts, local)
	if error != nil {
		return 0, error
	}
	return t.Unix(), nil
}

//china location
func LocationTimeCN() (time.Time) {
	local, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return time.Now()
	}
	return time.Now().In(local)
}
