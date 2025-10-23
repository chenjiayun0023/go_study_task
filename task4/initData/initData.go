package initData

import (
	"fmt"
	"go_study/task4/model"

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
