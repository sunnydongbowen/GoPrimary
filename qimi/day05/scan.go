/*
 * @Author: sunnydongbowen
 * @Date: 2024-11-03 19:36:19
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-11-03 20:26:24
 * @FilePath: /GoPrimary/qimi/day05/scan.go
 * @Description: 
 * 
 */
package main	

import ( 
	"fmt")



func Scan() {
	var (
		name   string
		age    int
		height float32
	)
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Print("请输入年龄：")
	fmt.Scan(&age)
	fmt.Print("请输入身高：")
	fmt.Scan(&height)
}
func scan2() {
	var (
		name   string
		age    int
		height float32
	)
	fmt.Print("请输入姓名、年龄、身高：")
	fmt.Scan(&name, &age, &height)
	fmt.Printf("姓名：%s，年龄：%d，身高：%f\n", name, age, height)

}

//fmt.Scanf()函数，可以按照指定的格式读取字符串
//格式要求：1:%s,2:%d,3:%f，很严格！必须严格按照格式输入
//如果输入的格式不对，就会报错。一般别这么用，会被用户骂死！
func scan3() {
	var (
		name   string
		age    int
		height float32
	)
	fmt.Scanf("1:%s,2:%d,3:%f",&name, &age, &height)
	fmt.Printf("姓名：%s，年龄：%d，身高：%f\n", name, age, height)
}	
	
func scan4() {
	var (
		name   string
		age    int
		height float32
	)
	fmt.Print("请输入姓名、年龄、身高：")
	fmt.Scanln(&name, &age, &height)
	fmt.Printf("姓名：%s，年龄：%d，身高：%.2f\n", name, age, height)
}

func main() {
	//Scan()
	//scan2()
	//scan3()
	//scan4()
	Bufio()
}

