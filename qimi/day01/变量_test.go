/*
 * @Author: bowen && bowen815@126.com
 * @Date: 2024-10-17 19:14:21
 * @LastEditors: sunnydongbowen bowen815@126.com
 * @LastEditTime: 2024-10-20 16:32:24
 * @FilePath: /GoPrimary/qimi/变量_test.go
 * @Description:
 *
 * Copyright (c) 2024 by ${git_name_email}, All Rights Reserved.
 */
package qimi

import (
	"fmt"
	"testing"
)


var (
	age int   //0
	isOk bool   //false
	sex  string
)

func TestVar1(t *testing.T){
	var name string
	fmt.Println(name)

	fmt.Println(age)
	fmt.Println(isOk)

	fmt.Println("========赋值后========")
	age =100
	isOk=true

	fmt.Println(age)
	fmt.Println(isOk)

}

func foo() (int,string){
	return 6,"qimi"
}

func TestVar2(t *testing.T){
	x,_:=foo()
	fmt.Println(x)

}

  