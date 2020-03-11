package uuid2

import (
	"github.com/twinj/uuid"
	"github.com/aosting/goTools/md52"
)

/**********************************************
** @Des: uuid2
** @Author: zhangxueyuan 
** @Date:   2018-12-27 14:43:19
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-12-27 14:43:19
***********************************************/

func GenerateUUID() string {
	u4 := uuid.NewV4()
	if u4 != nil {
		return md52.Md52(u4.String())
	} else {
		u1 := uuid.NewV1()
		u2 := uuid.NewV2(uuid.DomainGroup)
		u3 := uuid.NewV3(u1, u2)
		return md52.Md5(u3.String())
	}
}
