package initData

import (
	"fmt"
	"go_study/task3/gorm/model"
	"gorm.io/gorm"
)

// 创建数据库表
func CreateTables(db *gorm.DB) error {
	// 自动迁移创建表
	err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		return fmt.Errorf("自动迁移失败: %w", err)
	}

	fmt.Println("数据库表创建成功!")
	return nil
}

// 插入测试数据
func InsertTestData(db *gorm.DB) error {
	// 开始事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 创建用户
	users := []model.User{
		{ID: 1, Username: "alice"},
		{ID: 2, Username: "bob"},
		{ID: 3, Username: "charlie"},
	}
	if err := tx.Create(&users).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("创建用户失败: %w", err)
	}
	fmt.Println("用户数据插入成功!")

	// 2. 创建文章
	posts := []model.Post{
		// alice 的文章
		{ID: 1, Title: "Go语言入门教程", UserID: 1},
		{ID: 2, Title: "GORM使用指南", UserID: 1},
		{ID: 3, Title: "Web开发最佳实践", UserID: 1},

		// bob 的文章
		{ID: 4, Title: "MySQL性能优化", UserID: 2},
		{ID: 5, Title: "Redis缓存实战", UserID: 2},

		// charlie 的文章
		{ID: 6, Title: "Docker容器化部署", UserID: 3},
		{ID: 7, Title: "微服务架构设计", UserID: 3},
	}
	if err := tx.Create(&posts).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("创建文章失败: %w", err)
	}
	fmt.Println("文章数据插入成功!")

	// 3. 创建评论
	comments := []model.Comment{
		// 对文章1的评论
		{ID: 1, Content: "这篇教程写得很详细，对我帮助很大！", PostID: 1},
		{ID: 2, Content: "示例代码很实用，可以直接用在项目中", PostID: 1},
		{ID: 3, Content: "期待作者更新更多Go相关的内容", PostID: 1},

		// 对文章2的评论
		{ID: 4, Content: "GORM确实很方便，大大提高了开发效率", PostID: 2},
		{ID: 5, Content: "关联查询的部分讲得很清楚", PostID: 2},

		// 对文章3的评论
		{ID: 6, Content: "实战经验很宝贵，谢谢分享", PostID: 3},

		// 对文章4的评论
		{ID: 7, Content: "数据库优化技巧很实用", PostID: 4},
		{ID: 8, Content: "索引优化的部分对我启发很大", PostID: 4},
		{ID: 9, Content: "希望能看到更多实际案例", PostID: 4},

		// 对文章5的评论
		{ID: 10, Content: "Redis的使用场景总结得很好", PostID: 5},

		// 对文章6的评论
		{ID: 11, Content: "Docker部署确实很方便", PostID: 6},
		{ID: 12, Content: "容器化的优势很明显", PostID: 6},

		// 对文章7的评论 - 这个文章暂时没有评论
	}
	if err := tx.Create(&comments).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("创建评论失败: %w", err)
	}
	fmt.Println("评论数据插入成功!")

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	fmt.Println("所有测试数据插入成功!")
	return nil
}
