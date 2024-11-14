/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-26 12:15:38
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-26 12:22:10
 * @FilePath: /GoPrimary/qimi/day03/func2_test.go
 * @Description:
 *
 */
package day03

import "testing"
import "fmt"

var i = 10

func f6() {
	i = 20
}
func TestFunc2(t *testing.T) {
// i has been declared as a global variable,if you change the value of i in f6(),the value of i will be changed.
 f6()
 fmt.Println(i)
}

func TestFunc3(t *testing.T) {
	// i is declared as a local variable.
    i:=3
	fmt.Println(i)
}