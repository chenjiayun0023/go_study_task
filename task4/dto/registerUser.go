package dto

import "go_study/task4/model"

type RegisterUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func RegisterUserReqToUser(req RegisterUserReq) model.User {
	return model.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
}
