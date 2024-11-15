/*
 * @Author: sunnydongbowen
 * @Date: 2024-11-15 16:13:30
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-11-15 16:39:59
 * @FilePath: /GoPrimary/wpq/fmtinpt/main.go
 * @Description:
 *
 */
package main

import (
	"fmt"
)

func main() {
	// var name string
	// var age int

	// fmt.Println("请输入姓名：")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)  // _,_= 表示忽略返回值
	// fmt.Printf("姓名：%s，年龄：%d\n", name, age)


	// fmt.Println("请输入姓名和年龄：")
	// fmt.Scanln(&name, &age)// wait for input  over when enter 
	// fmt.Printf("姓名：%s，年龄：%d\n", name, age)

	// 很少用，很麻烦
	// fmt.Scanf("我叫%s，今年%d岁\n")
	// fmt.Printf("姓名：%s，年龄：%d\n", name, age)
    
    var message string
	fmt.Scanln(&message)
	fmt.Println(message)
}
