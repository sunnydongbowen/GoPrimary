/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-20 19:50:17
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-20 21:27:54
 * @FilePath: /GoPrimary/qimi/day01/swtich1_test.go
 * @Description:
 *
 */
package qimi

import (
	"fmt"
	"testing"
)

func TestSwitch1(t *testing.T) {
	finger := 4
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("中指")
	case 3:
		fmt.Println("食指")
	case 4:
		fmt.Println("无名指")
	default:
		fmt.Println("无效输入")
	}
}



func TestSwitch2(t *testing.T) {
	num := 2 
	switch num {
	case 2,4,6,8:
		fmt.Println("偶数")
	case 3,5,7,9:
		fmt.Println("奇数")
	default:
		fmt.Println("未知")
	}
}

// 注意执行顺序
func TestSwtich(t *testing.T) {
	i := 1
	switch i {
	case 2:
		fmt.Println("2")
	case 1:
		fmt.Println("1")
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("sorry")
	}
}


