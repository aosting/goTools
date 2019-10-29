package crypto2

import (
	"encoding/hex"
	"crypto/sha256"
)

/**********************************************
** @Des: sha2
** @Author: zhangxueyuan 
** @Date:   2019-10-29 16:19:28
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2019-10-29 16:19:28
***********************************************/

func Sha2(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
