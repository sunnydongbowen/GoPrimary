/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-21 21:03:38
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-21 21:17:11
 * @FilePath: /GoPrimary/qimi/day01/运算_test.go
 * @Description:
 *
 */
package qimi

import (
	"fmt"
	"testing"
)

// 运算符 与或非
func TestWei(t *testing.T) {
	bb1:=0b1001
	bb2:=0b1100
	fmt.Printf("%b\n",bb1&bb2)
	fmt.Printf("%b\n",bb1|bb2)
	fmt.Printf("%b\n",bb1^bb2)

	fmt.Printf("%b\n",1<<3)
	fmt.Printf("%b\n",0b1000>>2)
}

func TestYi(t *testing.T) {
    
}