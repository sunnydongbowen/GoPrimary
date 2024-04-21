package bookgoexpertpro

import (
	"fmt"
	"testing"
)

func TestByteToString(t *testing.T) {
	 b:= []byte{101,102,103}
	 s:=string(b)
	 fmt.Println(s)
}

func TestStringtoByte(t *testing.T) {
	s:="abc"
	b:=[]byte(s)
	fmt.Println(b)

}