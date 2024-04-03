package bookhizhan

import (
	"fmt"
	"testing"
	"unicode/utf8"
)


func TestChangeString(t *testing.T) {
	
	angle:="I Love you"
	angleByte:=[]byte(angle)

	for i := 3; i < 8; i++ {
		angleByte[i]=' '
	}
	fmt.Println(string(angleByte))
}

func TestStringLen(t *testing.T) {
	fmt.Println(utf8.RuneCountInString("反骨"))
	fmt.Println(utf8.RuneCountInString("反骨，hhh"))
}