package http4

import (
	"net/http"
	"github.com/aosting/goTools/plog"
	"io/ioutil"
	"time"
	"net/url"
	"strings"
)

/**********************************************
** @Des: client  方便服务状态统计，有超时，4xx,5xx 等
** @Author: zhangxueyuan 
** @Date:   2019-08-07 18:52:13
** @Last Modified by:   zhangxueyuan 
** @Last Modified time: 2019-08-07 18:52:13
***********************************************/

type HTTP_STATUS int64

const (
	SERVER_OK                  HTTP_STATUS = 0
	SERVER_TIME_OUT            HTTP_STATUS = 1
	SERVER_OTHER_ERROR         HTTP_STATUS = 2
	SERVER_4XX                 HTTP_STATUS = 3
	SERVER_5XX                 HTTP_STATUS = 4
	SERVER_RESPONSE_READ_ERROR HTTP_STATUS = 5
	REQUEST_ERROR              HTTP_STATUS = 10
)
const (
	CONTENT_JSON     = "application/json"
	CONTENT_PROTOBUF = "application/x-protobuf"
)
const _LABEL_ = "[_netServer_]"

func IsTimeout(err error) bool {
	if err != nil {
		if f, ok := err.(*url.Error); ok {
			if f.Timeout() {
				return true
			}
		}
	}
	return false
}

func handleResponse(resp *http.Response, err error) ([]byte, HTTP_STATUS, error) {
	if err != nil {
		plog.INFO(_LABEL_, "handleResponse: ", err)
		if IsTimeout(err) {
			return nil, SERVER_TIME_OUT, err
		}
		return nil, SERVER_OTHER_ERROR, err
	}
	if resp.StatusCode/100 == 4 {
		return nil, SERVER_4XX, err
	}
	if resp.StatusCode/100 == 5 {
		return nil, SERVER_5XX, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		plog.INFO(_LABEL_, "func VisitURL(url string) ioutil.ReadAll:", err)
		return nil, SERVER_RESPONSE_READ_ERROR, err
	}
	defer resp.Body.Close()

	return body, SERVER_OK, nil
}

func VisitUrl(url string, duration time.Duration) ([]byte, HTTP_STATUS, error) {
	client := http.Client{Timeout: duration}
	return handleResponse(client.Get(url))
}

//contentType: application/json  application/x-protobuf
func VisitUrlPost(url string, duration time.Duration, contentType string, body string) ([]byte, HTTP_STATUS, error) {
	client := http.Client{Timeout: duration}
	return handleResponse(client.Post(url, contentType, strings.NewReader(body)))
}

func VisitUrlWithHeaders(url string, duration time.Duration, headers map[string]string) ([]byte, HTTP_STATUS, error) {
	client := http.Client{Timeout: duration}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		plog.INFO(_LABEL_, "func VisitUrlWithHeaders(url string) NewRequest:", err)
		return nil, REQUEST_ERROR, err
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	return handleResponse(client.Do(request))
}

func VisitUrlPostWithHeaders(url string, duration time.Duration, body string, headers map[string]string) ([]byte, HTTP_STATUS, error) {
	client := http.Client{Timeout: duration}
	request, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		plog.INFO(_LABEL_, "func VisitUrlPostWithHeaders(url string) NewRequest:", err)
		return nil, REQUEST_ERROR, err
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	return handleResponse(client.Do(request))
}
