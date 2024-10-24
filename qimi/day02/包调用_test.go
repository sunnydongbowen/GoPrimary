/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-22 15:43:20
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-22 15:45:11
 * @FilePath: /GoPrimary/qimi/day02/包调用_test.go
 * @Description: 
 * 
 */
package day02

import(
	"testing"
	"GoPrimary/qimi/day02/hello"
)


func TestHello(t *testing.T) {
	//调用hello包中的Hello函数
	hello.Hello()
}