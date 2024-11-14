/*
 * @Author: sunnydongbowen
 * @Date: 2024-11-07 14:13:23
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-11-07 14:50:38
 * @FilePath: /GoPrimary/qimi/day05/interface5_test.go
 * @Description:
 *
 */
package main

import (
	"fmt"
	"testing"
)

type Any interface {
}

func TestAny(t *testing.T) {
	// var x interface{}
	var x Any
	fmt.Println(x) // <nil>

	x=100
	x=new(int)
	fmt.Println(x) // 0xc0000a8000

	m:=map[string]interface{}{
		"name":"sunny",
		"age":18,
		"married":false,
	}
	fmt.Println(m) // map[age:18 married:false name:sunny]
}
