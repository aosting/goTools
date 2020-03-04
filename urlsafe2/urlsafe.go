package urlsafe2

import (
	"encoding/base64"
	"strings"
)

/**********************************************
** @Des: urlsafe
** @Author: zhangxueyuan 
** @Date:   2020-03-04 19:37:54
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2020-03-04 19:37:54
***********************************************/

func Base64URLDecode(data string) ([]byte, error) {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	return base64.URLEncoding.DecodeString(data)
}

func Base64UrlSafeEncode(source []byte) string {
	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	bytearr := base64.StdEncoding.EncodeToString(source)
	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
	safeurl = strings.Replace(safeurl, "+", "-", -1)
	safeurl = strings.Replace(safeurl, "=", "", -1)
	return safeurl
}
