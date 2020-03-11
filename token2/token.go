package token2

/**********************************************
** @Des: token
** @Author: zhangxueyuan 
** @Date:   2019-11-14 14:58:32
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2019-11-14 14:58:32
***********************************************/

type TokenPool struct {
	Limit int
	Pool  chan struct{}
}

func NewTokenPool(limit int) TokenPool {
	tokenPool := TokenPool{
		Limit: limit,
		Pool:  make(chan struct{}, limit),
	}
	for i := 0; i < limit; i++ {
		tokenPool.Pool <- struct{}{}
	}
	return tokenPool
}

func (tp *TokenPool) Cap() int {
	if tp == nil {
		return 0
	}
	return tp.Limit
}

func (tp *TokenPool) Len() int {
	if tp == nil {
		return 0
	}
	return len(tp.Pool)
}

func (tp *TokenPool) Set(s struct{}) {
	if tp == nil {
		return
	}
	tp.Pool <- s
}

//this function maybe block the goroutine
func (tp *TokenPool) GetToken() {
	if tp == nil {
		return
	}
	<-tp.Pool
}
