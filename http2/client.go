package http2

import (
	"net/http"
	"time"
	"io/ioutil"
	"github.com/aosting/goTools/plog"
	"strings"
)

/**********************************************
** @Des: client 生产环境必须设置超时时间
** @Author: zhangxueyuan 
** @Date:   2018-12-27 15:50:51
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2018-12-27 15:50:51
						2019-04-10 16:58:51 增加时间统计
***********************************************/

const _LABEL_ = "[_netServer_]"


func VisitUrl(url string, duration time.Duration) ([]byte, int, error) {
	client := http.Client{Timeout: duration}
	response, error := client.Get(url)
	if error != nil {
		plog.INFO(_LABEL_, "func VisitURL(url string) http.Get:", error)
		return nil, -1, error
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		plog.INFO(_LABEL_, "func VisitURL(url string) ioutil.ReadAll:", err)
		return nil, -1, error
	}
	defer response.Body.Close()
	return body, response.StatusCode, nil
}

//contentType: application/json  application/x-protobuf
func VisitUrlPost(url string, duration time.Duration, contentType string, body string) ([]byte, int, error) {
	client := http.Client{Timeout: duration}
	response, err := client.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		return nil, -1, err

	}
	defer response.Body.Close()
	var data []byte
	data = make([]byte, 0, 1024*2)
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, -1, err
	}
	return data, response.StatusCode, nil
}

func VisitUrlWithHeaders(url string, duration time.Duration, headers map[string]string) ([]byte, int, error) {
	client := http.Client{Timeout: duration}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		plog.INFO(_LABEL_, "func VisitUrlWithHeaders(url string) NewRequest:", err)
		return nil, -1, err
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	response, err := client.Do(request)
	if err != nil {
		plog.INFO(_LABEL_, "func VisitUrlWithHeaders(url string) Do:", err)
		return nil, -1, err
	}
	if response == nil {
		plog.INFO(_LABEL_, "func VisitUrlWithHeaders(url string) Do response is null")
		return nil, -1, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		plog.INFO(_LABEL_, "func VisitUrlWithHeaders(url string) ioutil.ReadAll:", err)
		return nil, -1, err
	}
	defer response.Body.Close()
	return body, response.StatusCode, nil
}

//headers   "Content-Type": "application/json"
//headers   "Content-Type": "application/x-protobuf"

func VisitUrlPostWithHeaders(url string, duration time.Duration, body string, headers map[string]string) ([]byte, int, error) {
	client := http.Client{Timeout: duration}

	request, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		plog.INFO(_LABEL_, "func VisitUrlPostWithHeaders(url string) NewRequest:", err)
		return nil, -1, err
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	response, err := client.Do(request)
	if err != nil {
		plog.INFO(_LABEL_, "func VisitUrlPostWithHeaders(url string) Do:", err)
		return nil, -1, err
	}
	if response == nil {
		plog.INFO(_LABEL_, "func VisitUrlPostWithHeaders(url string) Do response is null")
		return nil, -1, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		plog.INFO(_LABEL_, "func VisitUrlPostWithHeaders(url string) ioutil.ReadAll:", err)
		return nil, -1, err
	}
	defer response.Body.Close()
	return data, response.StatusCode, nil
}
