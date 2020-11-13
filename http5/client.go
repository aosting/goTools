package http5

import (
	"github.com/aosting/goTools/http4"
	"github.com/aosting/goTools/plog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/**********************************************
** @Des: client  方便服务状态统计，有超时，4xx,5xx 等
** @Author: zhangxueyuan
** @Date:   2019-08-07 18:52:13
** @Last Modified by:   zhangxueyuan
** @Last Modified time: 2019-08-07 18:52:13
***********************************************/

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

func handleResponse(resp *http.Response, err error) ([]byte, int, http4.HTTP_STATUS, error) {

	statuscode := -2
	if resp != nil {
		statuscode = resp.StatusCode
	}
	if err != nil {
		plog.INFO(_LABEL_, "handleResponse: ", err)
		if IsTimeout(err) {
			return nil, statuscode, http4.SERVER_TIME_OUT, err
		}
		return nil, statuscode, http4.SERVER_OTHER_ERROR, err
	}
	if resp.StatusCode/100 == 4 {
		return nil, statuscode, http4.SERVER_4XX, err
	}
	if resp.StatusCode/100 == 5 {
		return nil, statuscode, http4.SERVER_5XX, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		plog.INFO(_LABEL_, "func VisitURL(url string) ioutil.ReadAll:", err)
		return nil, statuscode, http4.SERVER_RESPONSE_READ_ERROR, err
	}

	return body, statuscode, http4.SERVER_OK, nil
}

func VisitUrl(url string, c *http.Client, duration time.Duration) ([]byte, int, http4.HTTP_STATUS, error) {
	if c == nil {
		c = &http.Client{Timeout: duration}
	}
	return handleResponse(c.Get(url))
}

//contentType: application/json  application/x-protobuf
func VisitUrlPost(url string, c *http.Client, duration time.Duration, contentType string, body string) ([]byte, int, http4.HTTP_STATUS, error) {
	if c == nil {
		c = &http.Client{Timeout: duration}
	}
	return handleResponse(c.Post(url, contentType, strings.NewReader(body)))
}

func VisitUrlPostForm(url string, c *http.Client, duration time.Duration, data url.Values) ([]byte, int, http4.HTTP_STATUS, error) {
	if c == nil {
		c = &http.Client{Timeout: duration}
	}
	return handleResponse(c.PostForm(url, data))
}

func VisitUrlWithHeaders(url string, c *http.Client, duration time.Duration, headers map[string]string) ([]byte, int, http4.HTTP_STATUS, error) {
	if c == nil {
		c = &http.Client{Timeout: duration}
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		plog.INFO(_LABEL_, "func VisitUrlWithHeaders(url string) NewRequest:", err)
		return nil, -1, http4.REQUEST_ERROR, err
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	return handleResponse(c.Do(request))
}

func VisitUrlPostWithHeaders(url string, c *http.Client, duration time.Duration, body string, headers map[string]string) ([]byte, int, http4.HTTP_STATUS, error) {
	if c == nil {
		c = &http.Client{Timeout: duration}
	}
	request, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		plog.INFO(_LABEL_, "func VisitUrlPostWithHeaders(url string) NewRequest:", err)
		return nil, -1, http4.REQUEST_ERROR, err
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	return handleResponse(c.Do(request))
}
