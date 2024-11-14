package main

import "fmt"

// student manager system
type Student struct {
	id int
	name string
	age int
}
var students []Student

func main() {
	var choice int
	for {
		fmt.Println("1. show all students")
		fmt.Println("2. add student")
		fmt.Println("3. delete student")
		fmt.Println("4. update student")
		fmt.Println("5. search student")
		fmt.Println("0. exit")
		fmt.Scanln(&choice)
		if choice == 0 {
			break
		}
		switch choice {
		case 1:
			showAll()
		case 2:
		    addStudent()
		case 3:
		    deleteStudent()
	    case 4:
		    updateStudent()
		case 5:
			searchStudent()
		default:
			fmt.Println("输入错误")
		}
	}
	
}

func showAll() {
	for _, student := range students {
		fmt.Printf("学生id：%d，学生姓名：%s，学生年龄：%d\n",student.id, student.name, student.age)
	}
}

func addStudent() {
	var student Student
	//var id int
	fmt.Print("请输入学生id:")
	//id = len(students) + 1
	fmt.Scanln(&student.id)
	if student.id <=0 || student.id > 100 || student.id == 0{ 
		fmt.Println("输入错误")
		return
	}
	// if err != nil {
	// 	fmt.Println("输入错误")
	// 	return
	// }
	//student.id = id
	//fmt.Scanln(&student.id)


	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&student.name)
	
	fmt.Print("请输入学生年龄:")
	fmt.Scanln(&student.age)
	students = append(students, student)
}

func deleteStudent() {
	var id int
	fmt.Print("请输入要删除的学生id:")
	fmt.Scanln(&id)
	for i, student := range students {
		if student.id == id {
			students = append(students[:i], students[i+1:]...)
			break
		}
	}
}

func updateStudent() {
	var id int
	fmt.Print("请输入要修改的学生id:")
	fmt.Scanln(&id)
	for i, student := range students {
		if student.id == id {
			fmt.Print("请输入学生姓名:")
			fmt.Scanln(&student.name)
			fmt.Print("请输入学生年龄:")
			fmt.Scanln(&student.age)
			students[i] = student
}
	}
}

func searchStudent() {
	var id int
	fmt.Print("请输入要查询的学生id:")
	fmt.Scanln(&id)
	for _, student := range students {
		if student.id == id {
			fmt.Printf("学生id：%d，姓名：%s，年龄：%d\n", student.id, student.name, student.age)
			return
		}
	}
	fmt.Println("未找到该学生")
}