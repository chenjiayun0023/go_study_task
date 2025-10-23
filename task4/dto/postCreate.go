package dto

import "go_study/task4/model"

type PostCreateReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint
}

func PostCreateReqToPost(req PostCreateReq) model.Post {
	return model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}
}
