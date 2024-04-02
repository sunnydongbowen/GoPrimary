package structt

import (
	"fmt"
	"testing"
)

type Animal struct{
	Name string
}

func (s *Animal) SetName(name string){
	s.Name=name
}

type Cat struct{
	Animal;
}

type Dog struct{
	 a Animal
}

func TestEmbedddedFoo(t *testing.T){
	c:=Cat{}

	c.SetName("Cat")
	fmt.Printf("Name : %s\n",c.Name)  

	c.Animal.SetName("cat")
	fmt.Printf("Name : %s\n",c.Name)

	c.Animal.Name="Cat1"
	fmt.Printf("Name : %s\n",c.Name) 

	c.Name="cat1"
	fmt.Printf("Name : %s\n",c.Name) 

}


