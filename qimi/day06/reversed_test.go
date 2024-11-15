package main

import (
	"fmt"
	"sort"
	"testing"
)


func reverseDemo() {
     var x = []int{3,1 , 2, 5, 7, 8, 9, 4, 6}
	 sort.Ints(x) // 排序

	 fmt.Println(x)
     // only get interface through reverse method , so we need to convert it to the type we want
	 x2:=sort.Reverse(sort.IntSlice(x))
	 sort.Reverse(x2)
	 fmt.Println(x2)

    //  for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
    //      x[i], x[j] = x[j], x[i]
    //  }
}
func TestReverse(t *testing.T) {
	reverseDemo()
}