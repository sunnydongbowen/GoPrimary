package mapt

import (
	"fmt"
	"testing"
)

func TestInitByLiteral(t *testing.T){
	m:=map[string]int{
		"apple":2,
		"banana":3,
	}
	for k,v :=range m{
		fmt.Printf("%s-%d\n",k,v)
	}
}

func TestInitByMake(t *testing.T) {
	m:=make(map[string]int,10)
	m["apple"] = 2
	m["banana"]=3

	for k, v := range m {
		fmt.Printf("%s-%d\n",k,v)
	}
}

func TestMapCrud(t *testing.T){
	m:=make(map[string]string,10)
	m["apple"]="red"
	m["apple"]="green"
	delete(m,"apple")
	v,exist:=m["apple"]
	if exist{
		fmt.Printf("apple-%s\n",v)
	}
}

func TestEmptyMap(t *testing.T) {
	var m1 map[string]int
	m2:=make(map[string]int)
	fmt.Printf("len(m1)=%d\n",len(m1))
	fmt.Printf("len(m2)=%d\n",len(m2))
}