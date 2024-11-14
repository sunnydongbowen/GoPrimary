/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-28 20:13:15
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-29 11:13:33
 * @FilePath: /GoPrimary/qimi/day04/struct1_test.go
 * @Description: 
 * 
 */
package day04

import (	"fmt"

	"testing"	)


//define struct student
type Student struct {
	Name string
	Age int
	Score float32
	mapScore map[string]float32
}

func TestStruct1(t *testing.T) {
    var stu1 Student
    stu1.Name = "Tom"
    stu1.Age = 18
    stu1.Score = 98.5
    fmt.Println(stu1)

	stu2 := Student{
		Name: "Jerry",
		Age: 17,
		Score: 99.0,
		mapScore: map[string]float32{
		    "语文":23.5,
		    "数学":99.0,
		    "英语":88.0,
		},
	}
	fmt.Println(stu2)
}

func TestStruct2(t *testing.T) {
	p1:=new(int)
    s1:=new(Student)
    fmt.Printf("p1的类型是：%T\n",p1)
    fmt.Printf("s1的类型是：%T\n",s1)
}


func TestStruct3(t *testing.T) {
    stu2:=Student{}
	fmt.Printf("stu2的类型是：%T\n",stu2)

	// like java, we can use stu2.Name to set the value of Name
	// this is similar to stu3:=new(Student)
	stu3:=&Student{} 
	fmt.Printf("stu3的类型是：%T\n",stu3)
	(*stu3).Name="Tom"
	// go help us to simplify the code, we can use stu3.Name instead of (*stu3).Name
	stu3.Age=18
	fmt.Println(stu3.Name,stu3.Age)

	//
	stu4:=*new(Student)
	fmt.Printf("stu4的类型是：%T\n",stu4)
	stu4.Name="Jerry"
	fmt.Println(stu4.Name)

	// nil pointer dereference because we did not initialize the pointer
	// var stu5 *Student  //nil pointer
	// stu5.Name="Tom"    //as according to the nil pointer dereference, this will cause a runtime error
	// fmt.Println(stu5.Name)
}


func TestStruct4(t *testing.T) {
	stu2:=Student{
		"Tom",
		18,
		22.0,
		map[string]float32{
		    "语文":23.5,
		    "数学":99.0,
		    "英语":88.0,
		},
}
	fmt.Println(stu2)
}