/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-31 14:55:42
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-31 19:37:18
 * @FilePath: /GoPrimary/qimi/day04/struct7_test.go
 * @Description:
 *
 */
package day04

import (
	"fmt"
	"testing"
)

type Person1 struct {
	Name  string
	Age   int
	dream []string  
}

func (p *Person1) SetDream(dream []string) {
	tmp:= make([]string, len(dream))
	copy(tmp, dream)
	p.dream = tmp
	//p.dream = dream
}


func (p *Person1) SetDream2(dream []string) {
	p.dream = dream
}

func TestStruct13(t *testing.T) {
	p := Person1{
		Name: "sunny",
		Age:  18,
	}
	
	p.SetDream([]string{"dream1", "dream2"})
	fmt.Println(p.dream)

	
	p.SetDream([]string{"睡觉", "喝水"})
	fmt.Println(p.dream)
}


func TestStruct14(t *testing.T) {
	p:= Person1{
		Name: "sunny",
		Age:  18,
	}
	data:= []string{"dream1", "dream2"}
	p.SetDream2(data)
	fmt.Println(p.dream)

	// if we set the data[0] to "睡觉", it will not change the dream of p
	data[0] = "睡觉"
	//p.SetDream(data)
	fmt.Println(p.dream)
}