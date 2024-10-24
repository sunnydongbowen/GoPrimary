/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-23 18:35:49
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-23 19:21:23
 * @FilePath: /GoPrimary/qimi/day02/slice2_test.go
 * @Description:
 *
 */
package day02

import (
	"fmt"
	"testing"
)

// use make create slice
func TestSlice4(t *testing.T) {
	var s []int
	fmt.Println(s==nil)
	
	// when use make create slice, it is not nil
	// and if cap=len,it can be ignored
	// s=make([]int, 0)
	s=make([]int, 2,4)
	fmt.Println(s, len(s), cap(s))

	s=append(s, 1,2,3)
	fmt.Println(s, len(s), cap(s))
}

func TestSlice5(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[1:3]
	fmt.Println(s2, len(s2), cap(s2))

	s3:=[]int{99:1}
	fmt.Println(s3, len(s3), cap(s3))
}


func TestSlice6(t *testing.T) {
	a:=[]int{1,2,3,4,5}
	b:=make([]int, 10,10)
	//b:=make([]int, len(a))
	// copy a to b,is differnt from append.must have len and >=len of a
	copy(b, a)
	// why b is not 1,2,3,4,5? because b is a slice, it is a reference 	type	
	fmt.Println(b)
}

func TestSlice7(t *testing.T) {
	a:=[]int{1,2,3,4,5}
	b:=make([]int, 0,10)
	b=append(b, a...)
	//b[0]=100
	fmt.Println(a)
    
}