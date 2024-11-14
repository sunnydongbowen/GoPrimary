/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-25 15:41:46
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-25 15:41:51
 * @FilePath: /GoPrimary/algorithm/sort_test.go
 * @Description:
 *
 */
package algorithm

import (	"testing")

// sort a int slice
func Sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func TestSort(t *testing.T) {
	var arr = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	Sort(arr)
	t.Log(arr)
}