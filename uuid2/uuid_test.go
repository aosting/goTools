package uuid2

import (
	"testing"

)

/**********************************************
** @Des: uuid_test
** @Author: zhangxueyuan 
** @Date:   2019-04-10 18:07:43
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2019-04-10 18:07:43
***********************************************/

func TestGenerateUUID(t *testing.T) {
	t.Log(GenerateUUID())
}
