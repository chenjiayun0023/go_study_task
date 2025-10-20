package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL 驱动
	"github.com/jmoiron/sqlx"
)

//Sqlx入门

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Book struct {
	ID     int `gorm:"primaryKey"`
	Title  string
	Author string
	Price  float64
}

type Result struct {
	Title  string
	Author string
}

func initDB() (*sqlx.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %w", err)
	}
	fmt.Println("数据库连接成功!")
	return db, nil
}

// 创建books表并插入测试数据
func setupBookData(db *sqlx.DB) error {
	// 创建 books 表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(200) NOT NULL,
		author VARCHAR(100) NOT NULL,
		price DECIMAL(8,2) NOT NULL,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("创建books表失败: %w", err)
	}
	fmt.Println("books 表创建成功!")

	// 清空现有数据（可选，用于测试）
	_, err = db.Exec("DELETE FROM books")
	if err != nil {
		return fmt.Errorf("清空数据失败: %w", err)
	}

	// 插入测试数据
	insertSQL := `
	INSERT INTO books (title, author, price, created_at) VALUES
	('Go语言编程', '张三', 68.50, '2024-01-10'),
	('深入理解MySQL', '李四', 89.00, '2024-01-11'),
	('Python数据分析', '王五', 75.00, '2024-01-12'),
	('Java核心技术', '赵六', 45.00, '2024-01-13'),
	('JavaScript权威指南', '钱七', 120.00, '2024-01-14'),
	('算法导论', '孙八', 99.80, '2024-01-15'),
	('设计模式', '周九', 35.00, '2024-01-16'),
	('系统架构设计', '吴十', 150.00, '2024-01-17'),
	('Go并发编程实战', '张三', 78.00, '2024-01-18'),
	('数据库系统概念', '李四', 110.00, '2024-01-19');`

	_, err = db.Exec(insertSQL)
	if err != nil {
		return fmt.Errorf("插入测试数据失败: %w", err)
	}

	fmt.Println("书籍测试数据插入成功!")
	return nil
}

func main() {
	// 初始化数据库连接
	db, err := initDB()
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()

	// 创建表并插入测试数据
	err = setupBookData(db)
	if err != nil {
		log.Fatal("初始化书籍数据失败:", err)
	}

	var result []Result
	err = db.Select(&result, "SELECT title, author FROM books WHERE price > ?", 50)
	if err != nil {
		log.Fatal("查询价格大于50元的书籍失败:", err)
	} else {
		fmt.Println(result)
	}
}
