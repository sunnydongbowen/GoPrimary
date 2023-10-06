package algorithm

import ("testing"
"fmt"
)

func Sum(a []int)  int{
	sum:=0
	for _, v := range a {
		sum +=v
	}
	return sum
}

func Sum2(a []int) int{
	sum:=0
	for i := 0; i < len(a); i++ {
		sum +=a[i]
	}

	return sum
}
func TestSum(t *testing.T) {
	//4个7
	a:=[]int{1,2,3,4,5,6,7}
	fmt.Println(Sum2(a))
}

