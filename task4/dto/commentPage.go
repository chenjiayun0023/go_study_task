package dto

import "time"

type CommentPageReq struct {
	PageInfo
	PostID uint `json:"post_id" binding:"required"`
}

type CommentPageRsp struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
