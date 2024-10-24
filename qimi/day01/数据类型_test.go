/*
 * @Author: bowen && bowen815@126.com
 * @Date: 2024-10-20 20:53:30
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-22 15:57:09
 * @FilePath: /GoPrimary/qimi/day01/数据类型_test.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package qimi

import (
	"fmt"
	"testing"
	"math"
)


func TestBool(t *testing.T) {
    vb1 := true
	var vb2 bool
	fmt.Printf("%T value:%v\n", vb1,vb1)
	fmt.Printf("%T value:%v\n", vb2,vb2)
}


//强制类型转换
func TestConvert(t *testing.T) {
    var i1 int8=1
	fmt.Println(int64(i1))

	f1:=12.34
	fmt.Println(int64(f1))
}


//计算直角三角形斜边长
func TestTriangle(t *testing.T) {
	var a,b int=3,4
	// 这个之前也有遇到过，go语言中，math.Sqrt()函数的参数是float64类型，所以需要将a和b转换为float64类型
	var c int= int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
	fmt.Printf("%T\n",c)
}


func TestInt(t *testing.T) {
	var a int=10
	fmt.Printf("%d\n",a)
	fmt.Printf("%b\n",a)

	var b int=077
	fmt.Printf("%o\n",b)

	var c int=0xff
	fmt.Printf("%x\n",c)
	fmt.Printf("%X\n",c)
}

//进制转换
var (
	i1 int8 = 0b1001
	i2  uint16= 0o77
	i3 = 0x123
	i4 int = 123_456
)


func TestConvert2(t *testing.T) {
	var i1 int=101
	fmt.Printf("%d\n",i1)
	fmt.Printf("%b\n",i1)
	fmt.Printf("%o\n",i1)
	fmt.Printf("%x\n",i1)
    
	i2:=int8(9)
	fmt.Printf("%T\n",i2)

	i3:=0x1234567
	fmt.Printf("%d\n",i3)

	i4:=077
	fmt.Printf("%d\n",i4)
}

