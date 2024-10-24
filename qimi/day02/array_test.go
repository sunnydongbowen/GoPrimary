/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-22 16:03:39
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-23 14:56:39
 * @FilePath: /GoPrimary/qimi/day02/array_test.go
 * @Description:
 *
 */
package day02
import (
	 "fmt" 
"testing")

func TestArray(t *testing.T) {
	// arr2 := [3]int{1, 3, 5}
	// arr3 := [...]int{2, 4, 6, 8, 10}
	// var grid [4][5]int

	// fmt.Println(arr1, arr2, arr3, grid)

	// int type array
	var a [3]int
	fmt.Println(a)
	a= [3]int{1, 2, 3}
	fmt.Println(&a[0])
	a[0]=4
	
	fmt.Println(&a[0])

	// print the type of array
	fmt.Printf("%T\n", a)

    //bool type array
	var c [2]bool
	fmt.Println(c)
	c[0] = true

	// wrong c[2]=true,because out of range
	//c[2]=false

	// use len() to get the last index of array
	fmt.Println(c[len(c)-1])

	// string type array
	var d [2]string
	fmt.Println(d)
	fmt.Printf("%#v\n", d)
	d[0] = "hello"
	fmt.Printf("%#v\n", d)
}


func TestArray1(t *testing.T) {
	// use index to init array
	var X99=[100]int{98:-2,99:-1}
	fmt.Println(X99)

	// use short hand to init array
	var yy=[]string{"hello","world"}
	// print the type of array
	fmt.Printf("%T\n", yy)
	fmt.Println(yy)

	// don't know how many elements in array
	var zz=[...]string{"hello","world"}
	fmt.Printf("%T\n", zz)

	// use index to init array
	var z=[...]int{1:7,7:9}
	fmt.Println(z)

	//support use const to init array
	const size=10
	var arr=[size]int{1,2,3}
	fmt.Println(arr)

	//use for to iterate array
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	//use for range to iterate array
	for index,value := range arr{
		fmt.Println(index,value)
	}
}


//multi deimensional array
func TestArray2(t *testing.T) {
	var arr=[3][4]int{
		{0,1,2,3},
		{4,5,6,7},
		{8,9,10,11},
	}
	fmt.Println(arr)
	fmt.Println(arr[2][3])
	fmt.Println(arr[2][2])
	fmt.Println(arr[1][2])
}


func TestArray3(t *testing.T) {
	// three is the length of array. 
	// two is the length of sub array
	a:=[3][2]string{
		{"hello1","world"},
		{"hello2","world"},
		{"hello3","world"},
	}
	fmt.Println(a)
}

func TestArray4(t *testing.T) {
	// three is the length of array. 
	// two is the length of sub array
	a:=[...][2]string{
		{"hello1","world"},
		{"hello2","world"},
		{"hello3","world"},
	}

	fmt.Println(a)
	// the second is the length of sub array is not allowd to be omitted,for example
	// b:=[...][...]string{
	    
	// }
}

func TestArray5(t *testing.T) {
	// copy array 
	a:=[...]int{1,2,3,4,5}
	b:=a
	a[0]=100
	fmt.Println(a)
	// b is a copy of a,so the change of a will not affect b
	fmt.Println(b)
	
}

func TestArray6(t *testing.T) {
    a:=[3]int{1,2,3}
	//f3(a)
	// array is just passed by value.so the change of a in f3 will not affect a
	fmt.Println(a)
	// the change of a in f3 will affect b because b is a copy of a
	a=f3(a)
	fmt.Println(a)
}

func f3(v [3]int) ([3]int){
	v[0] = 200
	fmt.Println("f3",v)
	return v
}


