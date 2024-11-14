package main

import (
	"fmt"
	"testing"
)

type WashingMachine interface {
    wash()
    dry()
}

type dryer struct {}

func (d dryer) dry() {
    fmt.Println("dry")
}

type haier struct {
	dryer //embedding dryer
}

func (h haier) wash() {
    fmt.Println("wash")
}

func TestInterface5(t *testing.T) {
    var a WashingMachine
    a = haier{}
    a.wash()
    a.dry() //dryer的dry方法被嵌入到了haier中
}

