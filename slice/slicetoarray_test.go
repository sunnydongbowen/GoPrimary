package slice

import (
	"fmt"
	"testing"
	"unsafe"
) 
//切片转数组，很少用。知道就好
func TestSTA(t *testing.T){
	s:=[]int{1,2,3}
	var p=(*[3]int)(unsafe.Pointer(&s[0]))
	p[1] +=10
	fmt.Printf("%v\n",s)	
}

func TestSTA2(t *testing.T){
	s:=[]int{1,2,3}
	p:=(*[3]int)(s)
	p[1] +=10
	fmt.Printf("%v\n",s)
	// *[3]int,数组指针
	fmt.Printf("%T\n",p)



}

