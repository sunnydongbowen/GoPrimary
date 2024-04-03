package bookhizhan

import (
	"errors"
	"fmt"
	"testing"
)

var errDivisionZero =errors.New("division by zero")

func div(diviend,divisor int) (int,error){

	if divisor == 0{
		return 0,errDivisionZero
	} 
	return diviend/divisor,nil

} 
func TestError(t *testing.T){
	a,err:=	div(9,0)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(a)
}