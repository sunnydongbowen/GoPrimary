/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-25 15:28:59
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-31 19:47:31
 * @FilePath: /GoPrimary/algorithm/countword_test.go
 * @Description:
 *
 */
package algorithm

import (
	"fmt"
	"strings"
	"testing"
)

// can you write the title by yourself?
// count the number of words in a string,this is a interesting problem
func CountWord(input string) map[string]int {
	//or you can use the slpit function to split the string
	// words := strings.Split(input, " ")
	words := strings.Fields(input)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}
func TestCountWord(t *testing.T) {
	// Test case 1
	input1 := "how do you do"
	m1ap1 := CountWord(input1)
	fmt.Println(m1ap1)
	// expectedOutput1 := map[string]int{"how": 1, "do": 2, "you": 1}
}
