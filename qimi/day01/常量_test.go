/*
 * @Author: bowen && bowen815@126.com
 * @Date: 2024-10-20 11:11:29
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-21 19:46:34
 * @FilePath: /GoPrimary/qimi/day01/常量_test.go
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

const(
	n1=iota
	n2
	n3
)

const(
	_=iota
	KB=1<<(10*iota)
	MB=1<<(10*iota)
	GB=1<<(10*iota)
	TB=1<<(10*iota)
	PB=1<<(10*iota)
)

const (
	c1=iota
	c2 
	c3=100
	c4=iota
	c5
)
func TestIo(t *testing.T) {
	fmt.Println(n1,n2,n3)
	fmt.Println(KB,MB,GB)

	fmt.Println(c1,c2,c3,c4,c5)
}



func TestConst(t *testing.T) {
	// 常量
    fmt.Println((math.MaxFloat32))
	fmt.Println((math.MaxInt32))
	fmt.Printf("%f\n",math.MaxFloat32)
	fmt.Printf("%d\n",math.MaxInt32)
}
