/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-28 17:19:30
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-31 19:40:32
 * @FilePath: /GoPrimary/qimi/day04/type1_test.go
 * @Description:
 *
 */
package day04

import (
	"fmt"
	"testing"
)

// 自定义类型
type Myint int

// 类型型别名
type youint=int

func TestType(t *testing.T) {
    var a Myint
    var b youint
    a = 10
    b = 20
	// 打印类型
    fmt.Printf("%T %T\n",a,b)
	//convert force	强制转换. convert b to Myint and then assign to a
	a=Myint(b)
	fmt.Println(a)
	fmt.Printf("%T\n",a)
}