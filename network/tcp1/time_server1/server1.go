package main

import (
	"fmt"
	"net"
	"os"
	"time"
)


func main(){
	service:=":7777"   //这是端口号
	tcpAddr,err:=net.ResolveTCPAddr("tcp4", service)
	CheckError(err)
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	CheckError(err)
	for{
		conn,err:=listener.Accept()
		if err != nil {
			continue
		}
		daytime:=time.Now().String()
		conn.Write([]byte(daytime))   //返回给客户端的时间
		conn.Close()
	}
}

func CheckError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error: %s", err.Error())
		os.Exit(1)
	}
}