package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type student struct {
	Name string  `json:"name"`
 	age  int     `json:"age`
}

func TestJsonDemo(t *testing.T){
	var s=student{
		Name:"tom",
		age:18,
	}
	//json序列化
	data,err:=json.Marshal(s)
	if err!=nil{
		fmt.Println("json marshal failed,err:",err)
		return
	}
	fmt.Printf("json str:%s\n",data)
}