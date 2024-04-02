package sllice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceExpress(t *testing.T){
	orderLen:=5
	order:=make([]uint16,2*orderLen)
	pollorder:=order[:orderLen:orderLen]
	lockorder:=order[orderLen:][:orderLen:orderLen]
	fmt.Println("len(pollorder):" ,len(pollorder))
	fmt.Println("cap(pollorder):" ,cap(pollorder))
	fmt.Println("len(lockorder):" ,len(lockorder))
	fmt.Println("cap（lockorder):" ,cap(lockorder))
}

func TestSimpleExpress(t *testing.T) {
	a:=[5]int{1,2,3,4,5}
	b:=a[2:4]
	for k,v :=range b{
		fmt.Println(k,v)
	}
	fmt.Println(len(b))
}

func TestSliceCap(t *testing.T) {
	baseSlice:=make([]int,0,10)
	newslice:=baseSlice[2:5]
	fmt.Printf("%v",newslice)
}

func TestSliceString(t *testing.T) {
	baseStr:="Hello World"
	fmt.Printf("baseStr:%s\n",baseStr)
	fmt.Printf("baseStr Type：%s\n",reflect.TypeOf(baseStr))
	
	newStr:=baseStr[0:5]
	fmt.Printf("newStr:%s\n",newStr)
	fmt.Printf("newStr Type：%s\n",reflect.TypeOf(newStr))
}

func TestAppend(t *testing.T) {
	a:=[5]int{1,2,3,4,5}
	b:=a[1:4:4]
	fmt.Println(b)
	b=append(b, 0)
	fmt.Println(b)
	fmt.Println(a)
}

func TestSliceExtend(t *testing.T) {
	s:=make([]int,0,10)
	s1:=append(s,1,2,3)
	s2:=append(s1,4)
	fmt.Println(&s1[0]==&s2[0])
}


func TestSCap(t *testing.T) {
	var array [10]int
	var slice = array[5:6]
	fmt.Printf("len{slice} = %d\n", len(slice))
	fmt.Printf("cap{slice} = %d\n",cap(slice))
	fmt.Println(&slice[0]==&array[5])

}

func TestSlice(t *testing.T) {
	sliceA:=make([]int,5,10)
	sliceB:=sliceA[0:5]
	sliceC:=sliceA[0:5:5]
	fmt.Printf("lenSliceA = %d\n",len(sliceA))
	fmt.Printf("lenSliceB = %d\n",len(sliceB))
	fmt.Printf("capSliceB = %d\n",cap(sliceB)) //10
	fmt.Printf("capSliceC = %d\n",cap(sliceC))
}


