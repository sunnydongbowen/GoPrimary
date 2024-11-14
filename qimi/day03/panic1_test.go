/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-27 11:24:20
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-27 19:42:09
 * @FilePath: /GoPrimary/qimi/day03/panic1_test.go
 * @Description:
 *
 */
package day03

import (
	"fmt"
	"testing"
)

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	panic("panic in C")
}

func TestPanic(t *testing.T) {
	funcA()
	funcB()
	funcC()
}

func f11() {
	
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic！")
		}
	}()
	var m map[string]int
	m["one"] = 1
	// if panic here , recover will not catch it
	fmt.Println(m["one"])
	// panic("panic in f11")
}

func TestRecover(t *testing.T) {
	f11()
}
