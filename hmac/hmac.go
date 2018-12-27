package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"github.com/aosting/goTools/str"
	"github.com/aosting/goTools/aes2"
	"errors"
)

/**********************************************
** @Des: hmac
** @Author: zhangxueyuan 
** @Date:   2018-09-14 15:30:50
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-09-14 15:30:50
***********************************************/

//测试的
var (
	encSecret  = "5396b9e2ca6b400c854611f98b63ac01"
	initSecret = "6e8a0638c5f04612aed43e4c0fdc2467"
)

/**
https://developers.google.com/authorized-buyers/rtb/response-guide/decrypt-price?hl=zh-CN#decryption_scheme

// Add any required base64 padding (= or ==).
final_message_valid_base64 = AddBase64Padding(final_message)

// Web-safe decode, then base64 decode.
enc_price = WebSafeBase64Decode(final_message_valid_base64)

// Message is decoded but remains encrypted.
(iv, p, sig) = enc_price // Split up according to fixed lengths.
price_pad = hmac(e_key, iv)
price = p <xor> price_pad

conf_sig = hmac(i_key, price || iv)
success = (conf_sig == sig)
*/
func Set(enc, init string) {
	encSecret = enc
	initSecret = init
}

func Decoding(price string) (float64, error) {
	if str.IsBlank(price) {
		return 0, errors.New("decodestring is null")
	}
	ivs, err := aes2.Base64URLDecode(price)
	if err != nil {
		return 0, err
	}
	if len(ivs) != 28 {
		return 0, errors.New("Decoding price error size!=28")
	}

	iv := ivs[0:16]
	enc_price := ivs[16:24]
	signature := ivs[24:28]

	pad := hmac.New(sha1.New, []byte(encSecret))
	_, err = pad.Write(iv)
	if err != nil {
		return 0, err
	}
	price_pad := pad.Sum(nil)
	tmpprice, _ := safeXORBytes(price_pad, enc_price)

	conf_tmp := hmac.New(sha1.New, []byte(initSecret))
	_, err = conf_tmp.Write(tmpprice)
	_, err1 := conf_tmp.Write(iv)
	if err != nil || err1 != nil {
		return 0, err
	}
	conf_sig := conf_tmp.Sum(nil)
	if campareSign(signature, conf_sig) {
		return float64(bytesToInt64(tmpprice)) / float64(1000000), nil
	} else {
		return 0, errors.New("sign error")
	}

}

func safeXORBytes(a, b []byte) ([]byte, int) {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	dst := make([]byte, n)
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return dst, n
}

func campareSign(decodesige []byte, confsign []byte) bool {

	if len(decodesige) > len(confsign) {
		return false
	} else {
		for i, _ := range decodesige {
			if decodesige[i] != confsign[i] {
				return false
			}
		}
		return true
	}
}
func bytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
