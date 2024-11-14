/*
 * @Author: sunnydongbowen
 * @Date: 2024-11-02 16:37:17
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-11-03 20:27:54
 * @FilePath: /GoPrimary/qimi/day05/print_test.go
 * @Description:,
 *
 */
package main

// 测试文件必须以_test.go结尾

import (
	"encoding/json";
	"fmt";
	"testing";
	//"qimi/day04/student"
)

func TestPrint1(t *testing.T) {
	fmt.Println("hello world")     // 打印并换行
	fmt.Print("hello qimi")        // 打印不换行
	fmt.Printf("hello %s", "qimi") // 格式化输出

}

type student struct {
	Name string `json:"name"` // JSON 标签
	Age  int    `json:"age"`
}

// func (s student) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(map[string]interface{}{
// 		"name": s.name,
// 		"age":  s.age,
// 	})
// }

func TestPrint2(t *testing.T) {
	s1 := student{Name: "qimi", Age: 18}
	//fmt.Println(s1)
	//fmt.Printf("%v\n", s1)
	//fmt.Printf("%+v\n", s1) // when print struct, %+v will print field name
	//fmt.Printf("%#v\n", s1) // when print struct, %#v will print field name and type
	
	//the field of s1 is not exported, so we can't use json.Marshal if Name and Age are not exported like "name" and "age"
	bs, err := json.Marshal(s1) // 将结构体转换为字节数组
	fmt.Printf("%b\n", bs) // 打印二进制
	fmt.Println(bs)        // 打印字节数组
	fmt.Println(string(bs)) // 打印字符串
	//fmt.Printf("%s\n", string(bs)) // when print byte array, %s will print string format
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("bs:%s,err:%v\n", bs,err) // when print struct, %s will print json format
}

func TestPrint3(t *testing.T) {
	var b = []byte{'h', 'e', 'l', 'l', 'o', 'q', 'i', 'm', 'i'}
	fmt.Println(string(b))
	fmt.Printf("%s\n", b) // %s is the same as string(b)
	fmt.Println(b) // %v is the same as fmt.Println(b).byte code will be printed
}


func TestPrint4(t *testing.T) {
    n:=12.345
    fmt.Printf("%f\n",n)
    fmt.Printf("%.2f\n",n)//保留两位小数
    fmt.Printf("%10.2f\n",n)//保留两位小数，总宽度为10，不足用空格填充
    fmt.Printf("%010.2f\n",n)//保留两位小数，总宽度为10，不足用0填充
}


func TestPrint5(t *testing.T) {
	s:="hello"
	fmt.Printf("%s\n",s)
	fmt.Printf("%10s\n",s)  //总宽度为10，不足用空格填充
	fmt.Printf("%-10s\n",s) //总宽度为10，不足用空格填充，左对齐
	fmt.Printf("%10.2s\n",s) //总宽度为10，不足用空格填充，保留两位小数
	fmt.Printf("%.2s\n",s) //保留两位小数
    
	
}