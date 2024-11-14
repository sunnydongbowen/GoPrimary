/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-29 19:47:10
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-30 11:47:34
 * @FilePath: /GoPrimary/qimi/day04/struct4_test.go
 * @Description:	
 *
 */
package day04

import "testing"
import "fmt"	

type Address struct {
	Province string
	City     string
}

type student struct {
	Name    string
	Age     int
	Address Address  // 匿名字段，那么默认Student就包含了address的所有字段
}

func TestStruct10(t *testing.T) {
    var s student
    s.Name = "pprof.cn"
    s.Age = 18
    s.Address.Province = "北京"
    s.Address.City = "北京"
	fmt.Println(s)
}

func TestStruct11(t *testing.T) {
    var s student
    s = student{
        Name:    "pprof.cn",
        Age:     18,
        Address: Address{
            Province: "北京",
            City:     "北京",
        },
    }
    fmt.Println(s)
}

