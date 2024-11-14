package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
);
func Bufio() {
	//bufio包实现了带缓存的IO操作
	//bufio包主要提供了3个类型，Reader、Writer、Scanner
	//Reader类型提供了带缓存区的读操作
	//Writer类型提供了带缓存区的写操作
	//Scanner类型用于扫描文本，可以指定分隔符，逐个读取内容，通常用于逐行读取文件
	//bufio包还提供了其他一些函数，如NewReader、NewWriter、NewScanner等，用于创建Reader、Writer、Scanner对象
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败：", err)
		return
	}
	str=strings.TrimSpace(str)
	//str = strings.Trim(str, "\r\n")
	fmt.Println("输入的内容是", str)
}