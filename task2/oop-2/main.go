package main

import "fmt"

//面向对象

/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工的信息为姓名：%s ,年龄：%d、工号：%s", e.Name, e.Age, e.EmployeeID)
}

func main() {
	person := Person{Name: "张三", Age: 20}
	employee := Employee{Person: person, EmployeeID: "11"}

	employee.PrintInfo()
}
