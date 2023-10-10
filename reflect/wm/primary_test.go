package wm

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func ReflectPrint(val any){
	fmt.Println("type: ",reflect.TypeOf(val), " value: ",reflect.ValueOf(val) , " kind",reflect.ValueOf(val).Kind() )
}


func TestPrimary(t *testing.T){
	num:=3.14
	ReflectPrint(num)

	//
	ti:=time.Now()
	ReflectPrint(ti)

	var v1 interface{}
	v1="hello world"
	ReflectPrint(v1)

	v1=true
	ReflectPrint(v1)

	v1=2
	ReflectPrint(v1)

	v1=make([]string,0)
	




}