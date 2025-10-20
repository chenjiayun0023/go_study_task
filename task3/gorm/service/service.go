package service

import (
	"errors"
	"fmt"
	"go_study/task3/gorm/model"

	"gorm.io/gorm"
)

// 查询某个用户发布的所有文章及其对应的评论信息
func QueryUserPostsWithComments(db *gorm.DB, username string) error {
	var user model.User

	// 使用预加载一次性获取用户、文章和评论
	err := db.Preload("Posts.Comments").
		Where("username = ?", username).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("用户 %s 不存在\n", username)
			return nil
		}
		return fmt.Errorf("查询失败: %w", err)
	}

	if len(user.Posts) == 0 {
		fmt.Printf("用户 %s 没有发布任何文章\n", username)
		return nil
	}

	fmt.Printf("用户 %s 发布的所有文章及评论:\n", username)
	for i, post := range user.Posts {
		fmt.Printf("%d. 文章: 《%s》\n", i+1, post.Title)
		fmt.Printf("   评论数量: %d\n", len(post.Comments))
		for j, comment := range post.Comments {
			fmt.Printf("   %d.%d %s\n", i+1, j+1, comment.Content)
		}
		if len(post.Comments) == 0 {
			fmt.Printf("   暂无评论\n")
		}
		fmt.Println()
	}

	return nil
}

// 查询评论数量最多的文章信息
func QueryMostCommentedPost(db *gorm.DB) error {
	var result struct {
		PostID       int    `gorm:"column:post_id"`
		Title        string `gorm:"column:title"`
		UserID       int    `gorm:"column:user_id"`
		Username     string `gorm:"column:username"`
		CommentCount int    `gorm:"column:comment_count"`
	}

	// 使用子查询找到评论数量最多的文章
	err := db.Table("posts").
		Select(`
            posts.id as post_id,
            posts.title as title,
            posts.user_id as user_id,
            users.username as username,
            COUNT(comments.id) as comment_count
        `).
		Joins("JOIN users ON users.id = posts.user_id").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id, posts.title, posts.user_id, users.username").
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error
	if err != nil {
		return fmt.Errorf("查询失败: %w", err)
	}

	fmt.Printf("评论数量最多的文章:\n")
	fmt.Printf("  文章ID: %d\n", result.PostID)
	fmt.Printf("  标题: 《%s》\n", result.Title)
	fmt.Printf("  作者: %s (ID: %d)\n", result.Username, result.UserID)
	fmt.Printf("  评论数量: %d\n", result.CommentCount)

	return nil
}

// 根据评论ID删除
func DeleteComment(db *gorm.DB, commentID int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 先查询评论对象，确保钩子函数能获取到完整数据
		var comment model.Comment
		if err := tx.First(&comment, commentID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("评论不存在，ID: %d", commentID)
			}
			return fmt.Errorf("查询评论失败: %w", err)
		}

		fmt.Printf("准备删除评论: ID=%d, 文章ID=%d\n", comment.ID, comment.PostID)

		// 删除具体的评论对象（不是通过ID）
		result := tx.Delete(&comment)
		if result.Error != nil {
			return fmt.Errorf("删除评论失败: %w", result.Error)
		}

		fmt.Printf("评论删除成功，ID: %d\n", commentID)
		return nil
	})
}
