package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SQL语句练习

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）
、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

type student struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Age   int
	Grade string
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := InitDB()
	db.AutoMigrate(&student{})

	//插入数据
	studentI := student{Name: "张三", Age: 20, Grade: "三年级"}
	result := db.Create(&studentI)
	fmt.Println(result.RowsAffected)

	//查询
	var studentS student
	db.Debug().Where("age > ?", 18).Find(&studentS)
	fmt.Println(studentS)

	//更新
	db.Debug().Model(&student{}).Where("name = ?", "张三").Update("grade", "四年级")

	//删除
	db.Debug().Delete(&student{}, "age < ?", 15)
}
