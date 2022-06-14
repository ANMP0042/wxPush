/**
 * @Author: YMBoom
 * @Description:
 * @File:  http
 * @Version: 1.0.0
 * @Date: 2022/06/14 15:44
 */
package src

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type IRequest struct {
	Req    *http.Request
	Client *http.Client
}

// NewIRequest 实例化
func NewIRequest(body io.Reader) *IRequest {
	req, _ := http.NewRequest("GET", "", body)
	iReq := &IRequest{
		Req: req,
		Client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	return iReq
}

func (i *IRequest) AddHeader(name, value string) *IRequest {
	i.Req.Header.Add(name, value)
	return i
}

func (i *IRequest) SetMethod(method string) *IRequest {
	i.Req.Method = method
	return i
}

func (i *IRequest) SetUrl(reqUrl string) *IRequest {
	u, _ := url.Parse(reqUrl)
	i.Req.URL = u
	return i
}

func (i *IRequest) Do() (*http.Response, error) {
	return i.Client.Do(i.Req)
}
