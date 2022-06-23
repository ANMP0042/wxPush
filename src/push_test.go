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
	p, err := NewPush("6ec18e8d7c934a111cc8759d000ac5e05a9")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println(p.Send(Args{Title: "这是title", Content: "111"}))
}
