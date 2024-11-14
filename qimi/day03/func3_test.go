/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-26 15:28:21
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-26 16:35:34
 * @FilePath: /GoPrimary/qimi/day03/func3_test.go
 * @Description:
 *
 */
package day03
import "fmt"
import "testing"
type Myslice []int

type MyMap map[string]string

type myFunc func()

func f8() {
	var s Myslice
	s = append(s, 1)
	fmt.Println(s) 
}

func TestF8(t *testing.T) {
		f8()
}

type Fa func (x,y int) 
type Fb func(name string,score int)

//simplify func,improve code readability
func ff(a Fa, b Fb) {
	a(1,2)
	b("sunny",100)
}

type calculation func(int,int) int

 

func TestF9(t *testing.T) {
	// add := func(a,b int) int {
	// 	return a + b
	// }
	var add calculation
	add= func(a,b int) int {
		return a + b
	}
    val := add(1,2)
	fmt.Println(val)

	sub:= func(a,b int) int {
		return a - b
	}
	val = sub(1,2)
	fmt.Println(val)

	add2:=calculation(add)
	fmt.Println(add2(1,2))

	sub2 := calculation(sub)
	fmt.Println(sub2(1,2))
}

func add(a,b int) int {
	return a + b
}

func sub(a,b int) int {
	return a - b
}

func f18(x,y int,s string) func(int,int) int{
	switch s {
	case "+":
		return add
	case "-":
		return sub
	}
    return nil
}

func TestF18(t *testing.T) {
	f := f18(1,2,"+")
	fmt.Println(f(1,2))
}