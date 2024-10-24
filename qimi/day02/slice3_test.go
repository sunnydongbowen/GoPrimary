package day02

import (
	"fmt"
	"testing"
)


func TestSlice8(t *testing.T) {
	var s=[]string{"北京",}
	s= append(s,"上海")
	fmt.Println(s)

	s1:=[]string{"广州","深圳"}
	s=append(s,s1...)
	fmt.Println(s)

	var s3 []string
	// not need initialize s3,but if you use copy s3 to s, you need initialize s3
	s3=append(s3,s1...)
	fmt.Println(s3)
}

func TestSlice9(t *testing.T) {
	s1:=[]string{"北京","上海","广州","深圳"}
	// go slice delete element, use append,not delete, just append
	// is different from python
	s1 = append(s1[:2],s1[3:]...)
	fmt.Println(s1)
}

func TestSlice10(t *testing.T) {
	// 默认情况下，append函数会在切片的尾部追加元素
    var a=make([]int,5,10)
	fmt.Println(a,len(a),cap(a))
	for i:=0;i<10;i++{
		a=append(a,i)
	}
	fmt.Println(a,len(a),cap(a))
}

func TestSlice11(t *testing.T) {
    name:="bob"
	s:=fmt.Sprintf("%s",name)
	fmt.Println(s)
}

