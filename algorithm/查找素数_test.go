package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func IsPrime(a int) bool{
	if a<=1 {
		return false
	}
	//注意这个for循环！从2开始包括2，i<=,否则输出结果有不小差距！
	for i := 2; i <=int(math.Sqrt(float64(a))); i++ {
		// 注意这里返回的是false，有就不是了。
		if a%i==0 {
			return false
		}
	}
	return true
}

func FindPrime(start,end int) {
	for i := start; i <=end; i++ {
		if IsPrime(i){
			fmt.Println(i)
		}
	}
}

func TestPrime(t *testing.T)  {
	FindPrime(1,100)
}