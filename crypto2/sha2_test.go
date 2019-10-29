package crypto2

import "testing"

/**********************************************
** @Des: sha2_test
** @Author: zhangxueyuan 
** @Date:   2019-10-29 16:20:46
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2019-10-29 16:20:46
***********************************************/

func TestSha2(t *testing.T) {
	t.Log(Sha2("data"))
}