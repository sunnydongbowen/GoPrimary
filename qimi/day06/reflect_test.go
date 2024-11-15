package main

import (
	"fmt"
	"testing"
	"reflect"
)

func TestReflect(t *testing.T) {
	var a int = 1
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.ValueOf(a).Type())

	var b float32 = 1.0
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b))
}

