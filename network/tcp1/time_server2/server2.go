package main

import (
	"net"
	"time"
	"fmt"
	"os"
)



func main(){
	service:=":1200"
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",service)  //读取ip地址了
	CheckError(err)
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	CheckError(err)
	for{
		conn,err:=listener.Accept()
		if err !=nil {
				continue
			}
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn){
	defer conn.Close()
	daytime:=time.Now().Format("2006-01-02 15:04:05")
	conn.Write([]byte(daytime+"\n"))
}

func CheckError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error: %s", err.Error())
		os.Exit(1)
	}
}