package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"type:varchar(50);uniqueIndex;not null"`
	PostCount int    `gorm:"default:0"` // 添加文章数量统计字段
	//一对多关系：一个用户可以发布多篇文章
	Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID    int    `gorm:"primaryKey"`
	Title string `gorm:"type:varchar(200);not null"`
	// 外键：关联用户
	UserID        int    `gorm:"not null;index"`
	CommentCount  int    `gorm:"default:0"`                      // 评论数量统计
	CommentStatus string `gorm:"type:varchar(20);default:'无评论'"` // 评论状态
	// 一对多关系：一篇文章可以有多个评论
	Comments []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID      int    `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`
	// 外键：关联文章
	PostID int `gorm:"not null;index"`
}

/**
此处在 创建/删除文章时，更新用户的文章数量统计字段，有并发问题先不考虑
      创建/删除评论时，更新文章的评论数量和状态字段，有并发问题先不考虑
*/

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量
	result := tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1))

	if result.Error != nil {
		return fmt.Errorf("更新用户文章数量失败: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("用户不存在，ID: %d", p.UserID)
	}

	fmt.Printf("用户 ID %d 的文章数量已更新\n", p.UserID)
	return nil
}

// Comment 的钩子函数 - 在创建后自动更新文章评论数量和状态
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	// 更新文章的评论数量
	result := tx.Model(&Post{}).Where("id = ?", c.PostID).
		Updates(map[string]interface{}{
			"comment_count":  gorm.Expr("comment_count + ?", 1),
			"comment_status": "有评论",
		})

	if result.Error != nil {
		return fmt.Errorf("更新文章评论数量失败: %w", result.Error)
	}

	fmt.Printf("文章 ID %d 的评论数量已增加，当前状态: 有评论\n", c.PostID)
	return nil
}

// Comment 的钩子函数 - 在删除后自动更新文章评论数量和状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 先获取文章当前的评论数量
	var post Post
	if err := tx.First(&post, c.PostID).Error; err != nil {
		return fmt.Errorf("文章不存在，ID: %d", c.PostID)
	}

	// 计算删除后的评论数量
	newCommentCount := post.CommentCount - 1
	if newCommentCount < 0 {
		newCommentCount = 0
	}

	// 根据评论数量决定状态
	var newStatus string
	if newCommentCount == 0 {
		newStatus = "无评论"
	} else {
		newStatus = "有评论"
	}

	// 更新文章的评论数量和状态
	result := tx.Model(&Post{}).Where("id = ?", c.PostID).
		Updates(map[string]interface{}{
			"comment_count":  newCommentCount,
			"comment_status": newStatus,
		})

	if result.Error != nil {
		return fmt.Errorf("更新文章评论状态失败: %w", result.Error)
	}

	fmt.Printf("文章 ID %d 的评论状态已更新: 数量=%d, 状态=%s\n",
		c.PostID, newCommentCount, newStatus)
	return nil
}
