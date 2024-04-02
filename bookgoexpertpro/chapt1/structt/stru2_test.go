package structt

import (
	"fmt"
	"testing"
)

type Student struct{
	Name string
}

func (s Student) SetName(name string){
	s.Name=name
}

func (s *Student) UpdateName(name string){
	s.Name=name
}
func TestReceiver(t *testing.T){
	s:=Student{}
	s.SetName("aaa")
	fmt.Printf("Name:%s\n",s.Name)  //打的空

	s.UpdateName("bbb")
	fmt.Printf("Name:%s\n",s.Name)

}