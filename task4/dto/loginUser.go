package dto

import (
	"go_study/task4/model"
	"time"
)

type LoginUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserRsp struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

func UserToLoginUserRsp(user model.User, token string) LoginUserRsp {
	return LoginUserRsp{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
	}
}
