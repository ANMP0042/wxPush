/**
 * @Author: YMBoom
 * @Description:
 * @File:  push_test
 * @Version: 1.0.0
 * @Date: 2022/06/14 14:45
 */
package src

import (
	"fmt"
	"testing"
)

func TestPush_Send(t *testing.T) {
	p, err := NewPush("e54b4aa1c331cc4c5092f4e8ddaf111212176be99")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	p.Send(Args{Title: "这是title", Content: "111"})
}
