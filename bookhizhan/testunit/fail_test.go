package testunit

import (
	"fmt"
	"testing"
)

func TestFail(t *testing.T) {
	fmt.Println("before fail")
	//t.FailNow()
	t.Fail()
	fmt.Println("After fail")
}
