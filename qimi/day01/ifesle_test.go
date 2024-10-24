package qimi

import (
	"fmt"
	"testing"
)

func TestIF(t *testing.T) {
	if score := 89; score >= 90 {
		fmt.Println("优秀")
	} else if score < 90 && score >= 80 {
		fmt.Println("良好")
		fmt.Print(score)
	}
	//报错
	//fmt.Print(score)
}

func TestIF2(t *testing.T) {
	score := 11
	if score < 60 {
		fmt.Println("不及格")
	}
	fmt.Println(score)
}
