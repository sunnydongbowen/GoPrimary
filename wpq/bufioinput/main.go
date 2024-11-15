/*
 * @Author: sunnydongbowen
 * @Date: 2024-11-15 16:27:06
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-11-15 16:32:09
 * @FilePath: /GoPrimary/wpq/bufioinput/main.go
 * @Description:
 *
 */
package main

import "bufio"
import "os"
import "fmt"

func main() {
	reader:=bufio.NewReader(os.Stdin)

	line,_,_:=reader.ReadLine()
	data:=string(line)
	fmt.Println(data)
}