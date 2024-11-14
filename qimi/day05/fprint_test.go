package main

import (
	"fmt"
	"os"
	"testing"
)


func TestFprint(t *testing.T) {

	f,err:=os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY, 0644,)//打开文件 ，如果不存在则创建
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer f.Close()
	//fmt.Fprint(f, "hello world\n")

	//fmt.Fprint(os.Stdout, "hello world\n")
	//fmt.Fprintln(os.Stdout, "hello world")
	fmt.Fprintf(f, "hello world %s\n", "qimi")

	//fmt.Fprint(os.Stdout, "hello world\n")
	//fmt.Fprintln(os.Stdout, "hello world")
}