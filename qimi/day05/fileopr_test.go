/*
 * @Author: sunnydongbowen
 * @Date: 2024-11-03 21:04:05
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-11-05 08:55:42
 * @FilePath: /GoPrimary/qimi/day05/fileopr_test.go
 * @Description:
 *
 */
package main

import (
	"testing";
	"fmt";
	"os";
	"io"
	"bufio"
	"io/ioutil"
)			

func TestFile1(t *testing.T) {
    f,err:=os.Open("test.txt")
    if err!=nil{
        fmt.Println("open file err",err)
        return
    }
    defer f.Close()
	// read file
	var data [1024]byte
	n,err:=f.Read(data[:])  // 读取文件内容到data中
	if err!=nil{
		fmt.Println("read file err",err)
		return
	}
	fmt.Println("read file success",string(data[:n]))
}

func TestFile2(t *testing.T) {
	file,err:=os.Open("test.txt")
	if err!=nil{
		fmt.Println("open file err",err)
		return
	}
    defer file.Close()
    // for 循环读取文件内容
	var content  []byte
	var tmp = make([]byte,128)
	for{
		n,err:=file.Read(tmp)
		if err==io.EOF{
			fmt.Println("read file end")
			break
		}
		if err!=nil{
			fmt.Println("read file err",err)
			return
		}
		content=append(content,tmp[:n]...)
		//fmt.Println(string(content[:n]))
	}
	fmt.Println(string(content))
}


//bufio 包,提供缓冲区,提高读取效率
func TestFile3(t *testing.T) {
	file,err:=os.Open("test.txt")
	if err!=nil{
		fmt.Println("open file err",err)
		return
	}
	defer file.Close()
	// for 循环读取文件内容,no need to use define tmp slice
	reader:=bufio.NewReader(file)
	for{
		line,err:=reader.ReadString('\n')
		if err==io.EOF{
			fmt.Println("read file end")
			break
		}
		if err!=nil{
			fmt.Println("read file err",err)
			return
		}
		fmt.Print(line)
	}  
}

// ioutil 包,提供读取文件内容到字节数组中
func TestFile4(t *testing.T) {
	data,err:=ioutil.ReadFile("test.txt")
	if err!=nil{
		fmt.Println("read file err",err)
		return
	}
	fmt.Println(string(data))
}