package str

import "math/rand"

/**********************************************
** @Des: random
** @Author: zhangxueyuan 
** @Date:   2018-12-27 15:22:54
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-12-27 15:22:54
***********************************************/


//随机数
func RandomInt(num int) int {
	return rand.Intn(65536) % num
}
