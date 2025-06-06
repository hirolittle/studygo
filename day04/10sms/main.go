package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看、新增、删除学生
*/

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{id, name}
}

var allStudents map[int64]*student // 变量声明

func main() {
	allStudents = make(map[int64]*student) // 初始化，开辟内存空间
	for {
		// 1. 打印菜单
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
		4. 退出
	`)
		fmt.Print("请输入你要干啥: ")
		// 2. 等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了第 %d 个选项\n", choice)
		// 3. 执行对应的函数
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚～")
		}
	}
}

func deleteStudent() {
	// 1. 请用户输入要删除的学生的学号
	var deleteId int64
	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&deleteId)
	// 2. 去 allStudents 这个 map 中根据学号删除对应的键值对
	delete(allStudents, deleteId)
}

func addStudent() {
	// 向 allStudents 添加一个学生
	// 1. 创建一个学生
	var (
		id   int64
		name string
	)
	// 1.1 获取用户输入
	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanln(&name)
	// 1.2 造学生（调用 student 的构造函数）
	newStu := newStudent(id, name)

	// 2. 追加到 allStudents 这个 map 中
	allStudents[id] = newStu

}

func showAllStudents() {
	// 把所有的学生打印出来
	for k, v := range allStudents {
		fmt.Printf("学号：%d，姓名： %s\n", k, v.name)
	}
}
