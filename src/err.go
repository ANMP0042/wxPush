/**
 * @Author: YMBoom
 * @Description:
 * @File:  msg
 * @Version: 1.0.0
 * @Date: 2022/06/14 14:19
 */
package src

import "errors"

const (
	errMsgNotFound = "未找到提示信息"
	tokenIsNull    = "请先前往https://pushplus.hxtrip.com/获取token"
	contentIsNull  = "请填写内容"
)

func newErr(msg string) error {
	if msg == "" {
		msg = errMsgNotFound
	}
	return errors.New(msg)
}
