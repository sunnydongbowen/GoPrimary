/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-25 19:47:12
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-28 21:18:06
 * @FilePath: /GoPrimary/qimi/day03/funcc_test.go
 * @Description:
 *
 */
package day03

import (
	"fmt"
	"testing"

)

func sum(a, b int) int {
	return a + b
}

func f2(a, b int) (int,int) {
	return a + b, a - b 
}

func f1(a, b int) (sum int, sub int) {
    sum = a + b
	sub = a - b
	return sum,sub
}

func TestF1(t *testing.T) {
	sum, sub := f1(1, 2)
	fmt.Println(sum, sub)
}

func f3(string) {
	fmt.Println("hello")
}


func TestF3(t *testing.T) {
	f3("hello") 
}

func TestSum(t *testing.T) {
	result := sum(1, 2)
	fmt.Println(result)	
	// if result != 3 {
	// 	t.Errorf("sum(1, 2) = %d; want 3", result)
	// }
	sum,sub:= f2(1, 2)
	fmt.Println(sum,sub)
}
    

// variable arguments must be at the last
func sum3(y int,x ...int) int {		
	sum:=0
	for _, v := range x {
		sum += v
	}
	return sum 
}

func TestSum3(t *testing.T) {
	fmt.Println(sum3(0,1, 2, 3, 4, 5))
}

func someFunc(x string) []int {
	if x==""{
		// return nil slice = return empty slice
		return nil
		// equal to return []int{}
		//return []int{}
	}
	return []int{1,2,3}   
}

func TestSomeFunc(t *testing.T) {
	fmt.Println(someFunc("hello"))
	fmt.Println(someFunc(""))
}

func f4() (m map[string]int) {
	m = make(map[string]int)
	m["hello"] = 100
	return
}

func TestF4(t *testing.T) {
	// return value directly
	fmt.Println(f4())
}
