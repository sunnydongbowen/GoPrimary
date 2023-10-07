package time

import (
	"fmt"
	"testing"
	"time"
)
func TestXxx(t *testing.T) {
	cht:=time.After(30*time.Second)  //这个是一个chan
	fmt.Printf("value = %v,type=%T\n",cht,cht)

	time.Sleep(20*time.Millisecond)
	fmt.Println("主goroutine阻塞20ms")
	time.Sleep(100 * time.Millisecond)
	
	//fmt.Println("cht",<-cht) //会阻塞，会paic
	fmt.Print(cht)
}