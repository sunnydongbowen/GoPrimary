/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-16 11:33:49
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-22 09:46:40
 * @FilePath: /GoPrimary/algorithm/findnumber_test.go
 * @Description: 七米day1的题目
 *
 */
package algorithm

import (
	"fmt"
	"testing"
)

// 10w零1个数字，除了一个数字出现一次，其他数字都出现两次，找出这个数字
func FindNumber(nums []int) int {
	var res int

	for _, num := range nums {
		res ^= num
	}
	return res
}

func TestFindNumber(t *testing.T) {
	nums := []int{13, 13, 1, 1, 3, 5, 5, 7, 7, 9, 9, 11, 11, 2, 2, 4, 4, 6, 6, 8, 8, 10, 10, 12, 12}
	if len(nums)%2 == 0 {
		return
	}
	res := nums[0]
	for _, v := range nums[1:] {
		res ^= v
	}
	fmt.Print(res)
}
