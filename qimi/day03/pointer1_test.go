/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-27 16:15:16
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-28 10:31:20
 * @FilePath: /GoPrimary/qimi/day03/pointer1_test.go
 * @Description:
 *
 */
package day03

import "testing"
import "fmt"

func TestP1(t *testing.T)  {
	s:="hello"
	prts1:=&s
//prts1是一个指针变量，指向s的内存地址
    fmt.Println(prts1)
	// get value from pointer
	fmt.Println(*prts1)
	// change value from pointer
	*prts1="world"
	fmt.Println(s)

	fmt.Printf("prts1的类型是%T\n", prts1)
}