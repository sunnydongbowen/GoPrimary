package day02

import (
	"fmt"
	"testing"
)
func TestSlice(t *testing.T) {
	//just define slice  not initialized slice
	var s1 []int
	var s2[]bool
	var s3 []string

    // is different from array. 
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3=%v len(s3)=%d cap(s3)=%d\n", s3, len(s3), cap(s3))

	 var a1 [2]int
	 fmt.Println(a1)

	 fmt.Println(s1 == nil)
	 fmt.Println(s2 == nil)
	 fmt.Println(s3 == nil)
	 //fmt.Println(a1 == nil)
}

// 
func TestSliceInit(t *testing.T) {
	a1:= [5]int{1,2,3,4,5}
	s:=a1[1:3]
	fmt.Println(s)
	fmt.Println(len(s))
	// the capacity of slice is 4 because the capacity of array[2:5] is 4。the capacity of slice is the length of the array 
	//from  where the slice is defined
	fmt.Println(cap(s))

	// array how to convert to slice
	s1:=a1[:4]
	fmt.Println	(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
    
	s2:=a1[:]
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}


func TestSliceMake(t *testing.T) {
    a:=[5]int{1,2,3,4,5}
	// why the higher bound is 5? beacuse the capacity of slice is 5 and the length is 5， so the higher bound is 5
    // actuaclly the index of slice is from 0 to 4
	s1:=a[:5]
    fmt.Println(s1)
	a[0]=100
	fmt.Println(s1)

	str:="hello"
	// contain 0 to 2,not contain 3
	s2:=str[:3]
	// s2 is string not slice
	fmt.Println(s2)
	fmt.Printf("%T\n",s2)
	//fmt.Println(s2,len(s2),cap(s2))

}

func TestSlice3(t *testing.T) {
	a:=[]int{1,2,3,4,5}
	fmt.Printf("%T\n",a)

	// the capacity of slice is max-low
	// the len of slice is high-low
	s1:=a[1:2:3]
	fmt.Println(s1,len(s1),cap(s1))
    
}
