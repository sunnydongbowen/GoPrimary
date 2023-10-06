package algorithm

import (
	"fmt"
	"testing"
)

// 用计算机程序解决生活中的问题
func TestAge(t *testing.T){
	var fa=45
	var chi=15

	for i:=0;i<15;i++{
		if fa/chi==11 && fa%chi==0{
			fmt.Printf("%d年前父亲的年龄是儿子的11倍，此时父亲%d岁，儿子%d岁",i,fa,chi)
		}
		fa --
		chi--
	}
}