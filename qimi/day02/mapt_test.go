/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-24 11:23:50
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-24 11:50:55
 * @FilePath: /GoPrimary/qimi/day02/mapt_test.go
 * @Description:
 *
 */
package  day02

import "testing"
import "fmt"

func TestMap(t *testing.T) {
    var m map[string]int

	// must init map or will panic
	m = make(map[string]int)
    m["a"] = 1
    m["b"] = 2
    m["c"] = 3
	fmt.Println(m)

    if v, ok := m["a"]; ok {
        
    }
}