/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-25 15:54:11
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-25 15:54:14
 * @FilePath: /GoPrimary/qimi/day03/map1_test.go
 * @Description:
 *
 */
package day03

import (
	"fmt"
	"testing"
)
func TestMap1(t *testing.T) {
	m := make(map[string] []int)

	s:=[]int{1,2}
	m["one"]=s
	fmt.Println(m)

	s=append(s,3)
	fmt.Println(m["one"])
	fmt.Println	(s)
}
