package main

import (
	"fmt"
	"time"
)

func addNumerToChan(chanName chan int){
	for   {
		chanName<-1
		time.Sleep(1*time.Second)
	}
}

func readFromChan(){
	var chan1=make(chan int,10)
	var chan2=make(chan int,10)
	go addNumerToChan(chan1)
	go addNumerToChan(chan2)
	for{
		select{
		case e:=<-chan1:
			//readChan(chan1)
			fmt.Printf("get element from chan1 %d\n",e)
		case e:=<-chan2:
			fmt.Printf("get element from chan2 %d\n",e)
		default:
			fmt.Printf("No element in chan1 and chan2\n")
			time.Sleep(1*time.Second)
							}
	}
}