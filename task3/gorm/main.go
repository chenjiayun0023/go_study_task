package main

import (
	"go_study/task3/gorm/config"
	"go_study/task3/gorm/initData"
	"go_study/task3/gorm/service"
	"log"
)

//进阶gorm

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。

题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。

题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

func main() {
	db := config.InitDB()
	// 创建表
	err := initData.CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}

	// 插入测试数据
	err = initData.InsertTestData(db)
	if err != nil {
		log.Fatal("插入测试数据失败:", err)
	}

	//查询
	service.QueryUserPostsWithComments(db, "charlie")
	service.QueryMostCommentedPost(db)

	//删除评论
	service.DeleteComment(db, 11)

}
