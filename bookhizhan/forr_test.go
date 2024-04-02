package shizhan

import (
	"fmt"
	"testing"
)

func TestForSLice(t *testing.T){
	for i, v := range []int{1,2,3,4,5} {
		fmt.Println(i,v)	
	}
}

func TestForMap(t *testing.T){
	var m map[int]string
	m = map[int]string{
		1:"Go",
		2:"Python",
		3:"Shell",
	}
	for k, v := range m{
		fmt.Println(k,v)	
	}
}

func TestString(t *testing.T) {
	s:="你好,GO"
	for i, v := range s {
		//fmt.Println(i,v)
		fmt.Println(i,string(v))
	}
}