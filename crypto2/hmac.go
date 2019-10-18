package md52

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

/**********************************************
** @Des: hmac
** @Author: zhangxueyuan 
** @Date:   2019-10-18 15:38:13
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2019-10-18 15:38:13
***********************************************/

func Hmac(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}
