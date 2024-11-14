package main

import (	"fmt";
"os"
	"testing"
	"bufio"

)

func TestFileWrite(t *testing.T) {
    // 1. 打开文件
    file, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("open file err=", err)
        return
	}
    defer file.Close()

    // 2. 写入文件
    str := "hello,bowen\n"
	// n is the number of bytes written
    n, err := file.WriteString(str)
    if err != nil {		
        fmt.Println("write string err=", err)
        return
}
    fmt.Printf("写入 %d 个字节\n", n)
	str2 := "hello,bowen2\n"
    n, err = file.Write([]byte(str2))
    if err != nil {
        fmt.Println("write string err=", err)
        return
    }
    fmt.Printf("写入 %d 个字节\n", n)
}	


func TestFileWrite2(t *testing.T) {
   f, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        fmt.Println("open file err=", err)
        return
    }	
	defer f.Close()

	writer := bufio.NewWriter(f)
    n, err := writer.WriteString("hello,buffio\n")
    if err != nil {
        fmt.Println("write string err=", err)
        return
    }
	// if we don't call Flush(), the data will not be written to the file immediately
	//writer.Flush()
    fmt.Printf("写入 %d 个字节\n", n)
	
	for i := 0; i < 10; i++ {
		writer.WriteString("hello,world\n") // 写入缓存
	}
	writer.Flush() // 将缓存数据写入文件
}

