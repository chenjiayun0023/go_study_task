package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL 驱动
	"github.com/jmoiron/sqlx"
)

//Sqlx入门

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	return db
}

// 创建表并插入测试数据
func setupTestData(db *sqlx.DB) error {
	// 创建 employees 表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS employees (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100) NOT NULL,
		department VARCHAR(50) NOT NULL,
		salary DECIMAL(10,2) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("创建表失败: %w", err)
	}
	fmt.Println("employees 表创建成功!")

	// 清空现有数据（可选，用于测试）
	_, err = db.Exec("DELETE FROM employees")
	if err != nil {
		return fmt.Errorf("清空数据失败: %w", err)
	}

	// 插入测试数据
	insertSQL := `
	INSERT INTO employees (name, department, salary) VALUES
	('张三', '技术部', 15000.00),
	('李四', '技术部', 18000.00),
	('王五', '技术部', 22000.00),
	('赵六', '销售部', 12000.00),
	('钱七', '销售部', 13000.00),
	('孙八', '人事部', 10000.00),
	('周九', '技术部', 25000.00),
	('吴十', '财务部', 11000.00);`

	_, err = db.Exec(insertSQL)
	if err != nil {
		return fmt.Errorf("插入测试数据失败: %w", err)
	}

	fmt.Println("测试数据插入成功!")
	return nil
}

func main() {
	db := InitDB()
	defer db.Close()

	// 创建表并插入测试数据
	err := setupTestData(db)
	if err != nil {
		log.Fatal("初始化测试数据失败:", err)
	}

	var employees []Employee
	err = db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatal("查询所有部门为\"技术部\"的员工信息失败:", err)
	} else {
		fmt.Println(employees)
	}

	var employee Employee
	err = db.Get(&employee, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatal("查询工资最高的员工信息失败:", err)
	} else {
		fmt.Println(employee)
	}
}
