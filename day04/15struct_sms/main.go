package main

import (
	"fmt"
	"os"
)

/*
	方法版学生管理系统
	写一个系统能够查看、新增、删除学生
*/

// 学生管理系统

// 菜单函数
func showMenu() {
	fmt.Println("欢迎来到学生管理系统!")
	fmt.Println(`
			1. 查看所有学生
			2. 新增学生
			3. 修改学生
			4. 删除学生
			5. 退出
	`)
}

func main() {
	s := newSms()
	for {
		showMenu()
		// 等待用户输入选项
		fmt.Print("请输入操作编号：")
		var input string
		fmt.Scanln(&input)
		fmt.Printf("你输入的是：%s\n", input)
		switch input {
		case "1":
			s.showAllStudents()
		case "2":
			s.addStudent()
		case "3":
			s.editStudent()
		case "4":
			s.deleteStudent()
		case "5":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}
