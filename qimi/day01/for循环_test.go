/*
 * @Author: bowen && bowen815@126.com
 * @Date: 2024-10-20 18:46:59
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-21 11:19:19
 * @FilePath: /GoPrimary/qimi/day01/for循环_test.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package qimi

import (
	"fmt"
	"testing"
)

func TestForr(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func TestFor1(t *testing.T) {
	i := 1
	for {
		if i < 10 {
			fmt.Println(i)
			i++
		} else {
			break
		}
	}
}

func TestFor2(t *testing.T) {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func TestForrang(t *testing.T) {
	s := "Golang"
	for i, v := range s {
		//fmt.Printf("%v:%v \n",i,v)
		//fmt.Printf("%d %s\n",i,string(v))
		// 这里无需准换
		fmt.Printf("%d %c\n", i, v)
	}
}

func TestGoto1(t *testing.T) {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breaklabel
			}
			fmt.Println(i, "-", j)
		}
	}
breaklabel:
	fmt.Println("结束for循环")
}
