/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-29 11:16:37
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-11-04 17:23:51
 * @FilePath: /GoPrimary/qimi/day04/struct2_test.go
 * @Description: 	结构体
 *
 */
package day04

import (
	"fmt"
	"testing"
	"unsafe"
)

type MyStruct struct {
	a int8 //1 byte
	b int8
	c int8
	d int8
}

type Mystruct2 struct {}

func TestMystruct(t *testing.T) {
	var v1 MyStruct
	//get size  of struct, it is 4 bytes. means 4 * 8 = 32 bits memory
	fmt.Printf("v1 size:%d\n", unsafe.Sizeof(v1))

	//get a memory address
	fmt.Printf("v1 address:%p\n", &v1)
	fmt.Printf("v1 a address:%p\n", &v1.a)
	fmt.Printf("v1 b address:%p\n", &v1.b)
	fmt.Printf("v1 c address:%p\n", &v1.c)
	fmt.Printf("v1 d address:%p\n", &v1.d)

	var v2  Mystruct2
	// empty struct size is 0.not use memory
	fmt.Printf("v2 size:%d\n",	unsafe.Sizeof(v2) )
    
	//這個後面會用到
	set1:=map[string]struct {}{}
	fmt.Println(set1)
}

func TestMystruct2(t *testing.T) {
    // empty struct 
	nameList:=[]string{"Tom", "Jerry", "Sunny","Tom"}
	// map key is string, value is empty struct
	// type empty struct{}
	var nameMap= make(map[string]struct{}, len(nameList))  //map[string]empty
	// get keys of map. not use value,so we use struct{} to save memory
	for _, name := range nameList {
		nameMap[name] = struct{}{}  //struct{}empty 
	}
	for key:= range nameMap {
		fmt.Println(key)
	}
}

// 和顺序有关 ，和类型有关
type MyStruct3 struct {
	a int8    //1byte
	c int8    //1byte
	b int32   //4byte
	d int64   //8
}

func TestMystruct3(t *testing.T) {
	var v1 MyStruct3
	fmt.Printf("%p\n",&v1) //get a memory address=the first address of struct
	fmt.Printf("%p\n",&v1.a)
	fmt.Printf("v1 size:%d\n", unsafe.Sizeof(v1))
}