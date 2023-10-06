package slice

import (
	"fmt"
	"testing"
	"unsafe"
	"reflect"
) 


func TestSlice(t *testing.T) {
	a:=make([]int,0,2)
	//sh1:=unsafe.Pointer(&a)
	for i:=0;i<50;i++{
		a=append(a, i)
		if i%5==0{
			fmt.Println(unsafe.SliceData(a))
		}
	}
}

// 与上面代码类似的
func TestAppend(t *testing.T) {
	a := make([]int, 0, 2)
	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	for i := 0; i < 50; i++ {
		a = append(a, i)
		if i%5 == 0 {
			//fmt.Printf("第%d次扩容后的地址%p\n", i, &a[0]) // 这里不能写&a，&a是不变的。&a打印的控制结构，肯定无变化的
			// 控制结构有一个指向buffer的指针，有个len，还有个cap,变化的是buffer,用unsafe拿
			//fmt.Println(unsafe.Pointer(&a))
			fmt.Println(sh1.Data)
		}
	}
}