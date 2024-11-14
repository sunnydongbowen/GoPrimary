package main

import "fmt"

type Student struct {
	ID  int64
	Name string
	Age int
}


type Manager struct {
	StuInfo map[int64]*Student
}


func main() {
	// 初始化Manager
	mgr := Manager{
		StuInfo: make(map[int64]*Student, 8),
	}
	var choice int
	
	for {
		fmt.Println("1.展示所有学生信息")
		fmt.Println("2.添加学生信息")
		fmt.Println("3.删除学生信息")
		fmt.Println("4.更新学生信息")
		fmt.Println("5.查询学生信息")
		fmt.Println("0.退出程序")

		fmt.Scanln(&choice)
		switch choice {
		case 1:
			mgr.showAll()
		case 2:
			mgr.addStudent()
		case 3:
			mgr.deleteStudent() 
		case 4:
			mgr.updateStudent()
		case 5:
			mgr.searchStudent()
		case 0:	
			return
	}
}
}

func (st *Manager) showAll() {
	fmt.Println("展示所有学生信息:")
	for _, student := range st.StuInfo {
		fmt.Printf("id:%d,name:%s, age:%d,\n",student.ID, student.Name, student.Age)
	}
}

func (st *Manager) addStudent() {
	fmt.Println("添加学生:")
	var (
		id int64
		name string
		age int
	)
	fmt.Print("请输入学生id,姓名,年龄:")
	fmt.Scanln(&id,&name, &age)
	if _,ok:= st.StuInfo[id]; ok{
		fmt.Println("该学生id已存在")
		return
	}
	student := &Student{
		ID: id,
		Name: name,
		Age: age,
	}
	st.StuInfo[student.ID] = student
	fmt.Println("添加成功")
}

func (st *Manager) deleteStudent() {
	fmt.Println("删除学生:")
	var id int64
	fmt.Println("请输入学生id:")
	fmt.Scanln(&id)
	if _,ok:=st.StuInfo[id]; !ok{
		fmt.Println("该学生不存在")
		return
	}
	delete(st.StuInfo, id)
	fmt.Println("删除成功")
}


func (st *Manager) updateStudent() {
	fmt.Println("更新学生信息:")
	var (id int64
		name string
		age int)
		
	fmt.Println("请输入学生id:")
	fmt.Scanln(&id)
	if _,ok:=st.StuInfo[id]; !ok{
		fmt.Println("该学生不存在")
		return
	}
	fmt.Println("请输入学生姓名和年龄:")
	fmt.Scanln(&name,&age)
	st.StuInfo[id].Name = name
	st.StuInfo[id].Age = age
	fmt.Println("更新成功")
}

func (st *Manager) searchStudent() {
	fmt.Println("查询学生:")
	var id int64
	fmt.Println("请输入学生id:")
	fmt.Scanln(&id)
	if _,ok:=st.StuInfo[id]; !ok{
		fmt.Println("该学生不存在")
		return
	}
	fmt.Printf("学生信息:id:%d,name:%s,age:%d\n", st.StuInfo[id].ID,st.StuInfo[id].Name, st.StuInfo[id].Age)
}