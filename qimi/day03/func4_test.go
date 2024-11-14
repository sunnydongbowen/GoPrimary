package day03

import ("testing"
	"fmt"
	)


// func add(x,y int) int {
//     return x+y
// }

func f17(a,b int,add func(int,int) int) int{
    return add(a,b)
}

func TestFunc4(t *testing.T) {
    res:=f17(1,2,add)
	fmt.Println(res)
}
