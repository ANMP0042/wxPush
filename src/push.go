/**
 * @Author: YMBoom
 * @Description:
 * @File:  push
 * @Version: 1.0.0
 * @Date: 2022/06/14 14:15
 */
package src

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/url"
)

const (
	gateway         = "http://www.pushplus.plus/send"
	defaultTemplate = "html"
	sucCode         = 200
)

// 推送
type push struct {
	token string
	Args
}

type Args struct {
	Title    string
	Content  string
	Template string
	Topic    string
}

type Result struct {
	Code int    `xml:"code"`
	Msg  string `xml:"msg"`
}

func NewPush(token string) (*push, error) {
	if token == "" {
		return nil, newErr(tokenIsNull)
	}

	return &push{
		token: token,
	}, nil
}

func (p *push) Send(a Args) error {
	if a.Content == "" {
		return newErr(contentIsNull)
	}

	if a.Template == "" {
		a.Template = defaultTemplate
	}

	p.Args = a

	req := p.req()

	reqUrl := fmt.Sprintf("%s?%s", gateway, a.encode(p.token))
	resp, err := req.SetUrl(reqUrl).Do()
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	fmt.Println(string(b))
	if err != nil {
		return err
	}
	var result Result
	if err = xml.Unmarshal(b, &result); err != nil {
		return err
	}

	if result.Code != sucCode {
		return newErr(result.Msg)
	}
	return nil
}

func (p *push) req() *IRequest {
	request := NewIRequest(nil)
	request.AddHeader("content-type", "application/xhtml+xml")
	return request
}
func (a Args) encode(token string) string {
	val := url.Values{}
	val.Add("token", token)
	val.Add("content", a.Content)
	val.Add("template", a.Template)

	if a.Title != "" {
		val.Add("title", a.Title)
	}

	if a.Topic != "" {
		val.Add("topic", a.Topic)
	}
	return val.Encode()
}
