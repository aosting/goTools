package md52

import (
	"crypto/sha1"
	"encoding/hex"
)

/**********************************************
** @Des: sha1
** @Author: zhangxueyuan 
** @Date:   2019-10-18 15:38:59
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2019-10-18 15:38:59
***********************************************/


func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}