package http3

import "time"

/**********************************************
** @Des: client.go
** @Author: zhangxueyuan 
** @Date:   2019-04-10 17:09:41
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2019-04-10 17:09:41
***********************************************/

type VisitGet func(url string, duration time.Duration) ([]byte, int, error)
type VisitPost func(url string, duration time.Duration, contentType string, body string) ([]byte, int, error)

func TimeSpentGet(inner VisitGet) (VisitGet, int64) {
	start := time.Now()
	return func(url string, duration time.Duration) ([]byte, int, error) {
		return inner(url, duration)
	}, time.Since(start).Nanoseconds() / 1000000
}

func TimeSpentPost(inner VisitPost) (VisitPost, int64) {
	start := time.Now()
	return func(url string, duration time.Duration, contentType string, body string) ([]byte, int, error) {
		return inner(url, duration, contentType, body)
	}, time.Since(start).Nanoseconds() / 1000000
}
