package day04

import (	"fmt";"testing")
type Aniamal struct {
    name string
}

type Dog struct {
	leg int
    Aniamal
}

type cat struct {
	leg int
	Aniamal
}

func (a Aniamal) move() {
	fmt.Printf("%s move\n",a.name)    
}

func (d Dog) wang() {
	fmt.Printf("%s wang\n",d.name)
    
}

func (c cat) miao() {
	fmt.Printf("%s miao\n",c.name)
}

func TestStruct5(t *testing.T) {
    d := Dog{leg:4,Aniamal:Aniamal{name:"dog"}}
    d.move()
    d.wang()
    c := cat{leg:4,Aniamal:Aniamal{name:"cat"}}
    c.move()
    c.miao()
}