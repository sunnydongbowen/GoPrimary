/*
 * @Author: sunnydongbowen
 * @Date: 2024-10-30 15:05:50
 * @LastEditors: sunnydongbowen
 * @LastEditTime: 2024-10-31 19:32:16
 * @FilePath: /GoPrimary/qimi/day04/struct6_test.go
 * @Description:
 *
 */
package day04

import "testing"
import "fmt"
import "encoding/json"

type Class struct {
	Title    string
	Student []Student    `json:"student_list" xml:"student_list"`   //tag
}


func TestStruct12(t *testing.T) {
     c:=Class{
		Title: "101",
		Student: make([]Student, 0, 200),
     }
	 for i := 0; i < 10; i++ {
		student := Student{
			Name: fmt.Sprintf("stu%02d", i),
			Age:  i,
		}
		c.Student = append(c.Student, student)
	 }
	 //fmt.Printf("%#v\n", c)
	 // json marshal
	 data, err := json.Marshal(c)
	 if err != nil {
		fmt.Printf("json marshal failed, err:%v\n", err)
		return
	 }
	 fmt.Printf("%s\n", data)

	 // json unmarshal
	 str:=`{"Title":"101","Student":[{"Name":"stu00","Age":0},{"Name":"stu02","Age":1}]}`
	 c1:=&Class{}
	 err = json.Unmarshal([]byte(str), c1)
	 if err != nil {
		fmt.Printf("json unmarshal failed, err:%v\n", err)
		return
	 }
	 fmt.Printf("%#v\n", c1)
}

