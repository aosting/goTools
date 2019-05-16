package str2

/**********************************************
** @Des: stringUtils
** @Author: zhangxueyuan 
** @Date:   2018-09-11 10:44:52
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-09-11 10:44:52
***********************************************/

import (
	"strings"
	"bytes"
)

//字符串高效拼接
func Append(substring ... string) string {
	//var buffer strings.Builder
	var buffer bytes.Buffer
	l := len(substring)
	for i := 0; i < l; i++ {
		buffer.WriteString(substring[i])
	}
	return buffer.String()
}

func AppendCap(cap int, substring ... string) string {
	//var buffer strings.Builder
	var buffer bytes.Buffer
	l := len(substring)
	buffer.Grow(cap)
	for i := 0; i < l; i++ {
		buffer.WriteString(substring[i])
	}
	return buffer.String()
}


//dir use strings.Join
func ParseSlice(slices []string, separate string) string {
	return strings.Join(slices, separate)
}

//根据指定字符,分割参数,组合成串  并返回
func SeparateParams(sep string, args ...string) string {
	//var buffer strings.Builder
	var buffer bytes.Buffer

	for index, arg := range args {
		buffer.WriteString(arg)
		if index != len(args)-1 {
			buffer.WriteString(sep)
		}
	}
	return buffer.String()
}

////判断字符串是否为空串:非空返回ture, 否则返回false
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

//判断字符串是否为空串
func IsBlank(str string) bool {
	str = strings.TrimSpace(str)
	if len(str) < 1 {
		return true
	}
	return false
}



