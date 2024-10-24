package algorithm

import (
	"fmt"
	"testing"
)

// find the sum is 8 in the array, return the index
func FindSum(arr []int, sum int) []int {
    for i := 0; i < len(arr); i++ {
		// every time we find a pair, we can break the loop.we use the array length - 1 to avoid the duplicate pair like [1,7] and [7,1]
		// and the inner for loop will not be executed when i == len(arr) - 1, so we can avoid the index out of range error
        for j := i + 1; j < len(arr); j++ {
            if arr[i] + arr[j] == sum {
                return []int{i, j}
            }
        }
    }
    return nil
}

func TestFindSum(t *testing.T) {
    arr := []int{1,2,3,4,5,6,7,8,9,10}
    sum := 8
	index:=FindSum(arr, sum)
	fmt.Println(index)
}