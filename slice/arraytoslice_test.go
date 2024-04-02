package slice

import (
	"fmt"
	"testing"
) 


func TestATS(t *testing.T) {
	array :=[3]int{1,3,4}
	sl:=array[:]
	sl[2] +=1
	fmt.Printf("%v\n",sl)
}

