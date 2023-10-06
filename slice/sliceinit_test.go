package slice

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s []int

	s=[]int{1,2,3}
	fmt.Println(s)

    //注意这里的打印
	sm:=make([]int, 10,20)
	fmt.Println(sm)

	sm=[]int{1,2,3}
	fmt.Println(sm)


	sbool:=make([]bool,2,4)
	fmt.Println(sbool)

	// chans切片,注意这个的声明方式
	chans:=make([]chan int,2,4)
	fmt.Println(chans)
}

//slice 扩容
func TestStringSlice(t *testing.T){
	s:=make([]string,0)
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
    fmt.Printf("%p\n",s)

	s = append(s, "雨花台")
	s=append(s, "jiangning")
	s=append(s, "guli")
	fmt.Printf("%p\n",s)
	s=append(s, "gulou")
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
	fmt.Printf("%p\n",s)

	s=append(s, "NJ")
	s=append(s, "BJ")
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
	fmt.Printf("%p\n",s)

	//注意这个
	ss:= []string{"北京", "上海", "深圳", "杭州"}
	s = append(s, ss...)
	fmt.Println(s)
	fmt.Printf("%p\n",s)
}

