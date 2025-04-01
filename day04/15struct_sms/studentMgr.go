package main

import "fmt"

// 有一个物件：
//
//	1.它保存了一些数据  --> 结构体的字段
//	2.它有四个功能     --> 结构体的方法

// 造一个学生管理者
type sms struct {
	allStudents map[int64]student
}

type student struct {
	id   int64
	name string
}

func newSms() sms {
	return sms{
		allStudents: make(map[int64]student),
	}
}

func newStudent(id int64, name string) student {
	return student{
		id:   id,
		name: name,
	}
}

func isExist(id int64, allStudents map[int64]student) bool {
	if _, ok := allStudents[id]; ok {
		return true
	}
	return false
}

// 删除学生
func (s sms) deleteStudent() {
	var id int64
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)

	if !isExist(id, s.allStudents) {
		fmt.Println("学生不存在")
		return
	}
	delete(s.allStudents, id)
	fmt.Scanln("删除成功!")

}

// 增加学生
func (s sms) addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)

	if isExist(id, s.allStudents) {
		fmt.Println("学生已存在")
		return
	}

	var newStu = newStudent(id, name)
	s.allStudents[id] = newStu

	fmt.Println("添加成功!")

}

// 查看学生
func (s sms) showAllStudents() {
	for k, v := range s.allStudents {
		fmt.Printf("id: %d, name: %s\n", k, v.name)
	}
}

// 修改学生
func (s sms) editStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入要修改的学生学号：")
	fmt.Scanln(&id)

	if !isExist(id, s.allStudents) {
		fmt.Println("学生不存在")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d, 姓名: %s\n", id, s.allStudents[id].name)

	fmt.Print("请输入修改后学生姓名：")
	fmt.Scanln(&name)

	s.allStudents[id] = newStudent(id, name)
	fmt.Println("修改成功!")

}
