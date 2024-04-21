package main

import (
	"time"

	"github.com/pkg/profile"
)

func joinSlice() []string{
	var arr []string
	for i:=0;i<100000;i++{
		// 多次append，引起内存重新分配，性能较低
		arr = append(arr, "arr")
	}
	return arr
}

func joinSlice2() []string{
	const count = 100000
	var arr []string =make([]string, count)
	for i := 0; i < count; i++ {
		arr[i]="arr"
	}
	return arr
}
func main(){
	//开始性能分析，返回一个停止接口
	stopper:=profile.Start(profile.CPUProfile,profile.ProfilePath("."))

	//在main函数结束时停止性能分析
	defer stopper.Stop()

	//分析核心逻辑
	joinSlice2()

	//让程序至少运行1s
	time.Sleep(time.Second)
}