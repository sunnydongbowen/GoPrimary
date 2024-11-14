/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-28 15:09:45
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-28 15:22:29
 * @FilePath: /GoPrimary/algorithm/coin_test.go
 * @Description:
 *
 */
package algorithm

import (
	"fmt"
	"testing"
)

/* 50 coins distributed ,rule like this
a: name contains  per 'e' or 'E'，get 1 coin
b: name contains  per 'o' or 'O'，get 3 coin
c: name contains  per 'i' or 'I'，get 2 coin
d: name contains  per 'u' or 'U'，get 4 coin'
ask how many coins each user get. and left how many coins
*/

var (
	coins=50
	users = []string{
	    "Matthew",
	    "Sarah",
		"Augustus",
		"Heidi",
		"Emilie",
		"Peter",
		"Gunnar",
		"RACHEL",
		"Karl",
		"Veronica",
	}
	distribution = make(map[string]int,len(users))
)

func distributionCoins()  int{
	for _, user := range users {
		for _, c := range user {
			switch c {
			case 'e', 'E':
				distribution[user]++
			case 'o', 'O':
				distribution[user] += 3
			case 'i', 'I':
				distribution[user] += 2
			case 'u', 'U':
				distribution[user] += 4
			}
		}		
}
    for _, v := range distribution {
		coins -= v
	}
	fmt.Println(distribution)
	return coins
}

func TestCoin(t *testing.T) {
	left:=distributionCoins()
	fmt.Println("left:",left)	
	//t.Log(distribution)
}
