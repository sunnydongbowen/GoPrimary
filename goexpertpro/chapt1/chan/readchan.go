package main

import (
	"fmt"
)

func readChan(name string ,ch chan float64 ){
	for e := range ch {
		fmt.Printf("get element from %s %f\n",name,e)
	}
}

//错误写法
// func readChan1(ch chan float64){
// 	for e := range ch {
// 		fmt.Printf("get element from %s %f\n",ch,e)
// 	}

// }
var chan1 = make(chan float64,10)
var chan2= make(chan float64,10)

func read(){
	chan1 <-1.1
	chan1 <-1.2
	chan1 <-1.3
    
	chan2 <-2.1
	chan2 <-2.2
	chan2 <-2.3

	close(chan1)
    close(chan2)

	readChan("chan1",chan1)
	readChan("chan2",chan2)

}
