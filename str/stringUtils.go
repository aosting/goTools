package str

/**********************************************
** @Des: stringUtils
** @Author: zhangxueyuan 
** @Date:   2018-09-11 10:44:52
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-09-11 10:44:52
***********************************************/

import (
	"bytes"
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

//字符串高效拼接
func Append(substring ... string) string {
	var buffer bytes.Buffer

	for i := 0; i < len(substring); i++ {
		buffer.WriteString(substring[i])
	}
	return buffer.String()
}

//切片转换成字符串, 并指定分隔符号
func ParseSlice(slices []string, separate string) string {
	buffer := &bytes.Buffer{}
	for index, s := range slices {
		buffer.WriteString(s)
		if index != len(slices)-1 {
			buffer.WriteString(separate)
		}
	}
	return buffer.String()
}

func Str2Float64(arg string) (num float64) {
	num, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		num = 0
	}
	return num
}

//字符转换为十进制数字,空串转换为0
func Str2Int(arg string) (num int) {
	num, err := strconv.Atoi(arg)
	if err != nil {
		num = 0
	}
	return num
}

func Str2Int32(arg string) (result int32) {
	num, _ := strconv.ParseInt(arg, 10, 32)
	return int32(num)
}

//字符转换为十进制数字
func Str2Integer(arg string) (num int, err error) {
	if "" != arg {
		num, err = strconv.Atoi(arg)
	} else {
		ifo := "Str2Integer() string is null ERROR!"
		err = errors.New(ifo)
	}
	return num, err
}

//随机数
func RandomInt(num int) int {
	return rand.Intn(65536) % num
}

//以tab分割参数. 组合成字符串返回
func SeparateParamsByTab(args ...string) string {
	return SeparateParams("\t", args...)
}

//根据指定字符,分割参数,组合成串  并返回
func SeparateParams(sep string, args ...string) string {
	buffer := &bytes.Buffer{}
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

//点分十进制ip转化为数字
func Ip2Long(ip string) (uint64, error) {
	ipSeg := strings.Split(ip, ".")
	if len(ipSeg) != 4 {
		return 0, errors.New("ip format error!")
	}
	b0, err := strconv.ParseUint(ipSeg[0], 10, 64)
	b1, err := strconv.ParseUint(ipSeg[1], 10, 64)
	b2, err := strconv.ParseUint(ipSeg[2], 10, 64)
	b3, err := strconv.ParseUint(ipSeg[3], 10, 64)
	if err != nil {
		return 0, err
	}
	return b0<<24 | b1<<16 | b2<<8 | b3, nil
}

//ip的数字形式转化为点分十进制
func Long2Ip(ipLong uint64) string {
	temp := &bytes.Buffer{}
	temp.WriteString(strconv.FormatUint(ipLong>>24, 10))
	temp.WriteString(".")
	temp.WriteString(strconv.FormatUint((ipLong&0x00FFFFFF)>>16, 10))
	temp.WriteString(".")
	temp.WriteString(strconv.FormatUint((ipLong&0x0000FFFF)>>8, 10))
	temp.WriteString(".")
	temp.WriteString(strconv.FormatUint(ipLong&0x000000FF, 10))
	return temp.String()
}

func Int2String(in int) string {
	return strconv.Itoa(in)
}
func Int642String(in int64) string {
	return strconv.FormatInt(in, 10)
}

func Int322String(in int32) string {
	return strconv.Itoa(int(in))
}

//切片转成字符串，用来生成sql的in的参数
func Slice2Strig(slices []string) string {
	plantids := ""

	for index, _ := range slices {
		if index == 0 {
			plantids = slices[index]
		} else {
			if slices[index] != "" {
				plantids = plantids + "," + slices[index]
			}
		}
	}
	return plantids
}
