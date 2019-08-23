package plog

import (
	"testing"
	"github.com/aosting/goTools/plog"
)

/**********************************************
** @Des: stdLogger_test
** @Author: zhangxueyuan 
** @Date:   2019-08-23 11:13:57
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2019-08-23 11:13:57
***********************************************/

func TestError(t *testing.T) {
	plog.ERROR("欧呦，我出错了！")
	plog.INFO("欧呦，我发出了一个信息！")
}
