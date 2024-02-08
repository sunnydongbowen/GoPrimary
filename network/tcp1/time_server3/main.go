package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	service:=":1200"
	tcpAddr,err:=net.ResolveTCPAddr("tcp4", service)
	CheckError(err)
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	CheckError(err)
	for{
		conn,err:=listener.Accept()
		if err != nil {
			continue
		}
		//这个是关键
		go handleClient(conn)
	}
	
}

// 这里不一样了
func handleClient(conn net.Conn){
	conn.SetReadDeadline(time.Now().Add(2*time.Minute))
	request :=make([]byte,128)

	defer conn.Close()
	for {
		read_len,err:=conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if  read_len==0 {
			break
		}else if  strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime :=strconv.FormatInt(time.Now().Unix(),10)
			conn.Write([]byte(daytime))
		}else if  strings.TrimSpace(string(request[:read_len])) == "time" {
			daytime:=time.Now().Format("2006-01-02 15:04:05")
			conn.Write([]byte(daytime+"\n"))
		}
		request = make([]byte, 128)
	}
}

func CheckError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error: %s", err.Error())
		os.Exit(1)
	}
}