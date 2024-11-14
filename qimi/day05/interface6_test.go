/*
 * @Author: sunnydongbowen
 * @Date: 2024-11-08 14:54:17
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-11-08 14:57:02
 * @FilePath: /GoPrimary/qimi/day05/interface6_test.go
 * @Description:
 *
 */
package main

import (	"testing";
	"fmt";
)

func TestInterface11(t *testing.T) {
	var x interface{}
	x=100
	fmt.Printf("type:%T\n", x)

	x="hello"
	fmt.Printf("type:%T\n", x)

	x=true
	fmt.Printf("type:%T\n", x)

	x=3.14
	fmt.Printf("type:%T\n", x)	

	x=struct{}{}
	fmt.Printf("type:%T\n", x)

	x=new(int)
	fmt.Printf("type:%T\n", x)

	

}

func TestInterface12(t *testing.T) {
	var x interface{}
	x=100
	// if ok -> true,说明x是bool类型
	// if true ,说明x是bool类型.change value to bool
	b,ok:=x.( bool)
	fmt.Println(b,ok)
    
	x="hello"
	s,ok:=x.(string)
	fmt.Println(s)
    
}

func justifyType(x interface{}) {
	switch v:=x.(type) {
	case int:
		fmt.Printf("x is int, value is %d\n", v)
	case string:
		fmt.Printf("x is string, value is %s\n", v)
	case bool:
		fmt.Printf("x is bool, value is %t\n", v)
	default:
		fmt.Println("unsupport type!")
}
}
func TestInterface13(t *testing.T) {
	justifyType(10)
	justifyType("hello")
	justifyType(false)
	justifyType(3.14)
	foo(Dog{Name:"旺财", Age:3})
}

type Dog struct {
	Name string
	Age int
}

func foo(x interface{}) {
	switch v:=x.(type) {
	case Dog:
		fmt.Printf("v.Name=%s\n", v.Name)
	default:
		fmt.Println("unsupport type!")
	}
}