/*
 * @Author: bowen && bowen815@126.com
 * @Date: 2024-10-17 11:34:58
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-21 16:05:04
 * @FilePath: /GoPrimary/qimi/day01/str_test.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package qimi

import (
	"fmt"
	"strings"
	"testing"
)

func TestStr(t *testing.T) {
	filename := "c:\\go\\primary"
	fmt.Println(filename)

	// 原样输出
	s1 := "\"加油\""
	fmt.Println(s1, len(s1))

	s2 := "加油"
	fmt.Println(s2, len(s2))
	fmt.Println("jade " + "new year")
	res := fmt.Sprintf("%s%s", "jade1", "newyear")
	fmt.Print(res)
}


func TestStr2(t *testing.T) {
	slice1:=[]string{"a","b","c"}
	fmt.Println(strings.Join(slice1,"-"))
	}

//字符和字符串

func TestStr3(t *testing.T) {
	s1:="中"
	s2:='中'
	fmt.Println(s1)
	fmt.Println(s2)
}


func TestStr4(t *testing.T) {
	s1:="你好"
	fmt.Println([]rune(s1))
	fmt.Println([]byte(s1))

	// byte
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v(%c)",s1[i],s1[i])
	}
	fmt.Println()

	// rune
	for _,r := range s1 {
		fmt.Printf("%v(%c)",r,r)
	}
}

// 修改字符串
func TestStr5(t *testing.T) {
	s1:="hello"
	//s1[0]='a'  这行代码会报错，字符串是不可变的
	//fmt.Println(s1)

	//byte切片转字符串
	byteSlice := []byte(s1)
	byteSlice[0] = 'a'
	fmt.Println(string(byteSlice))

	//rune切片转字符串
	runeSlice := []rune(s1)
	runeSlice[0] = '你'
	fmt.Println(string(runeSlice))
    
}


