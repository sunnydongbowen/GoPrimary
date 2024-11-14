package main

import "testing"
import "fmt"


// 
type studentPrinter interface {
	printInfo()
}

// student relize studentPrinter interface
func (s student) printInfo() {
	fmt.Println(s.Name, s.Age)
}

func TestInterface(t *testing.T) {
	var sp studentPrinter
	sp = student{"qimi", 18}
	sp.printInfo()
}

type dreamer interface {
    makedream()
}

type dog struct {}

func (d dog) makedream() {
    fmt.Println("dog make dream")
}

func (s student) makedream() {
    fmt.Println("student make dream")
}

func TestInterface2(t *testing.T) {
    var d dreamer
    d = dog{}
    d.makedream()

    d = student{"qimi", 18}
    d.makedream()
}

func TestInterface3(t *testing.T) {
    var dog1 dog
	//dog1.move()
	dog1.makedream()

	var s student
	s.makedream()
}


type Mover interface {
	move()
}


type dog1 struct {}


// value receiver
// func (d dog1) move() {
// 	fmt.Println("dog move")
// }

func (d *dog1) move() { // pointer receiver
	fmt.Println("dog move")
}
func TestInterface4(t *testing.T) {
	var m Mover
	//m = dog1{}
	//if we use value receiver, we can't use pointer
	// if we use pointer receiver, we can use pointer or value
	m = &dog1{}
	m.move()

	m = &dog1{}
	m.move() 
}