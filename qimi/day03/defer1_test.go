package day03

import (
	"fmt"
	"testing"
)

// like stack,first in last out
func TestDefer1(t *testing.T) {
	defer fmt.Println("end1")
	defer fmt.Println("end2")
	defer fmt.Println("end3")
	fmt.Println("start")
}

func TestDefer2(t *testing.T) {
	fmt.Println("start")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	// notice that defer is executed after the function return so end will be printed first
	fmt.Println("end")
}

func TestDefer3(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// return value is named
func deferDemo2() (x int) {
	defer func() {
		x++
	}()
	// return x,x=1,not 1,x is named and changed
	return 1
}

func TestDeferdemo2(t *testing.T) {
	fmt.Println(deferDemo2())
	fmt.Println(deferDemo3())
}

func deferDemo3() int {
	x := 1
	defer func() {
		x++
	}()
	// return x is x, not x++
	return x
}
