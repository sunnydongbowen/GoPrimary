package time

import (
	"fmt"
	"testing"
	"time"
)

var notify bool
func TestFunc(t *testing.T) {
	timer:=time.AfterFunc(30*time.Millisecond,func ()  {
		fmt.Println("我再等30ms开始执行,你可以忙别的")
	})
	time.Sleep(20*time.Millisecond)
	notify=true
	if notify{
		timer.Stop()
		fmt.Println("收到通知，任务终止")
	}
	fmt.Println("主函数开始")
}

