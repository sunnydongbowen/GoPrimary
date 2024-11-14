/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-24 11:23:50
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-24 15:32:17
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

	// init map with make
    var m1 = map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
	}
	fmt.Println(m1)

    if v, ok := m["a"]; ok {
        fmt.Println(v)
    }

	for k, v := range m {
        fmt.Println(k, v)
	}
}

// slice of map
func TestMapSlice(t *testing.T) {
	// s is slice which element is map
    // this type is offen used in golang。py  [{},{},{},{},{}]
    // {[],[],[],[],[]}
    var s = make([]map[string]int, 5)
    s[0] = make(map[string]int)
    s[0]["a"] = 1
    s[0]["b"] = 2
    s[0]["c"] = 3
    fmt.Println(s)
}