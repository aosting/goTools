package crypto2

import (
	"crypto/md5"
	"encoding/hex"
)

/**********************************************
** @Des: md5
** @Author: zhangxueyuan 
** @Date:   2018-12-27 14:43:53
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-12-27 14:43:53
***********************************************/


func Md5(b string) (tp string) {
	h := md5.New()
	h.Write([]byte(b))
	x := h.Sum(nil)
	y := make([]byte, 32)
	hex.Encode(y, x)
	return string(y)
}