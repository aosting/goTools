package str2

import (
	"strings"
	"strconv"
	"errors"
	"bytes"
	"net/http"
)

/**********************************************
** @Des: ip
** @Author: zhangxueyuan 
** @Date:   2018-12-27 15:41:24
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-12-27 15:41:24
***********************************************/

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
	//var temp strings.Builder
	var temp bytes.Buffer
	temp.WriteString(strconv.FormatUint(ipLong>>24, 10))
	temp.WriteString(".")
	temp.WriteString(strconv.FormatUint((ipLong&0x00FFFFFF)>>16, 10))
	temp.WriteString(".")
	temp.WriteString(strconv.FormatUint((ipLong&0x0000FFFF)>>8, 10))
	temp.WriteString(".")
	temp.WriteString(strconv.FormatUint(ipLong&0x000000FF, 10))
	return temp.String()
}

//获取客户端ip地址
func GetClientIp(r *http.Request) (ip string) {
	ip = r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("x-forwarded-for")
	}
	if ip == "" {
		ip = r.Header.Get("Remote_addr")
	}
	if ip == "" {
		ip = r.Header.Get("Proxy-Client-IP")
	}
	if ip == "" {
		ip = r.Header.Get("WL-Proxy-Client-IP")
	}
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}