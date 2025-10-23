package dto

import "go_study/task4/model"

type CommentCreateReq struct {
	Content string `json:"content" binding:"required"`
	UserID  uint
	PostID  uint `json:"post_id" binding:"required"`
}

func CommentCreateReqToComment(req CommentCreateReq) model.Comment {
	return model.Comment{
		Content: req.Content,
		UserID:  req.UserID,
		PostID:  req.PostID,
	}
}
