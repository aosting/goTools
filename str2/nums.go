package str2

import (
	"strconv"
)

/**********************************************
** @Des: nums
** @Author: zhangxueyuan 
** @Date:   2018-12-27 15:18:07
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-12-27 15:18:07
***********************************************/

func Str2Int(arg string) (num int) {
	num, _ = strconv.Atoi(arg)
	return num
}

func Str2Int32(arg string) (result int32) {
	num, _ := strconv.ParseInt(arg, 10, 32)
	return int32(num)
}
func Str2Int64(arg string) (num int64) {
	num, _ = strconv.ParseInt(arg, 10, 64)
	return num
}

func Str2Float64(arg string) (num float64) {
	num, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		num = 0
	}
	return num
}
func Str2Float32(arg string) (result float32) {
	num, err := strconv.ParseFloat(arg, 32)
	if err != nil {
		num = 0
	}
	return float32(num)
}

func Int2String(in int) string {
	return strconv.Itoa(in)
}

func Int322String(in int32) string {
	return strconv.Itoa(int(in))
}

func Int642String(in int64) string {
	return strconv.FormatInt(in, 10)
}

func Float322String(in float32) string {
	return strconv.FormatFloat(float64(in), 'E', -1, 32)
}

func Float642String(in float64) string {
	return strconv.FormatFloat(in, 'E', -1, 64)
}
