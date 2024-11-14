/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-26 19:15:14
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-26 19:20:02
 * @FilePath: /GoPrimary/qimi/day03/func5_test.go
 * @Description: 
 * 
 */
package day03

import ("testing"
	"fmt")

func adder() func(int) int {
    var x int
	// x is used in the closure
	// return a function
    return func(y int) int {
        x += y
        return x
    }
}

// a very good example of closure
func TestAdder(t *testing.T) {
    f:=adder()
    sum := f(1)
	fmt.Println(sum) // 1
    sum = f(20)
	fmt.Println(sum) // 21
    sum = f(300)
	fmt.Println(sum) // 321
}

