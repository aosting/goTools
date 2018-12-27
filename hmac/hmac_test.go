package hmac

import (
	"testing"
)

/**********************************************
** @Des: 测试代码用例
** @Author: zhangxueyuan 
** @Date:   2018-09-14 17:15:26
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-09-14 17:15:26
***********************************************/

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDecoding(t *testing.T) {

	var Prices = [] struct {
		pirce  string
		result float64
	}{
		{
			"AACnZjHF6_AAAAFd0CcGgXlvO9SiDZdE4WFjxQ",
			1000.555,
		}, {
			"AACngidzJvAAAAFd0CjbmjEZG8p-cQX3Bivejg",
			800.1,
		}, {
			"AACnlNbLZTAAAAFd0CoVF18Sq4_W8Vc6YEAnWw",
			10,
		}, {
			"AACnojEyBvsAAAFd0Cr1H0h5--oQ1vFKUpVerg",
			1,
		},
	}

	t.Log("Use demo price to testDecoding.")
	{

		encSecret = "c546946f87ee40578d32155d86bc3f21"
		initSecret = "780f1b12d777424bb4818d43060606b1"
		for _, u := range Prices {
			tmp := Decoding(u.pirce)
			if u.result == tmp {
				t.Logf(" should hava a %f with input %s. %v", u.result, u.pirce, checkMark)
			} else {
				t.Errorf("should have a %f, but is %f  with input %s. %v", u.result, tmp, u.pirce, ballotX)
			}
		}
	}
}
