package main

import (
	"fmt"
	"sort"
	"testing"
	//"fmt"
)

// interface
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// reverse struct implements Interface
type reverse struct {
	Interface
}

// 实现Interface接口. and change Less method
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

func (r reverse) Swap(i, j int) {
	r.Interface.Swap(j, i)
}

func (r reverse) Len() int {
	return r.Interface.Len()
}
// 
func Reverse(data Interface) Interface {
	return &reverse{data}
}

type IntSlice []int

func (p IntSlice) Len() int {
		return len(p)
}         

func (p IntSlice) Less(i, j int) bool {
	return p[i] <p[j] 
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func TestInterface8(t *testing.T) {
	var x Interface 
	//x =sort.Reverse([]int{1, 2, 3, 4, 5})
	x = Reverse(x)
	x.Swap(0, 1)
}

func TestInterface9(t *testing.T) {
	data:=IntSlice{1,5,3,4,5}
	reversed:=Reverse(data)
	reversed.Swap(0, 1)
	if data[0] != 2 || data[1] != 1 {
		t.Errorf("After Swap(0, 1), data = %v", data)
	}
	if  !reversed.Less(0,4){
		t.Errorf("After Swap(0, 1), reversed.Less(0, 4) = false")
	}
	
}


type Address struct{
	city string
}

type Person struct{
	name string
	Address   
	Interface
}

func (a Address) showAddr() {
    fmt.Printf(a.city)
}

func (s student) showName(){
	fmt.Printf(s.Name)
}

func TestInterface10(t *testing.T) {
	var p Person
	// p.city = "北京"
	// fmt.Println(p.city)

	// p.showAddr()

	p.Address = Address{"上海"}
	p.showAddr()

	// p.showName() // wrong 
	//p.Interface = student{"张三"} //wrong
	// open source code of interface
	p.Interface=sort.IntSlice{}
	p.Len()
}