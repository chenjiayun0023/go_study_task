package controller

import (
	"go_study/task4/service"

	"github.com/gin-gonic/gin"
)

var CommentCtl *CommentController

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController() *CommentController {
	CommentCtl = &CommentController{
		commentService: service.CommentSve,
	}
	return CommentCtl
}

func (ctl *CommentController) CreateComment(c *gin.Context) {
	ctl.commentService.CreateComment(c)
}

func (ctl *CommentController) CommentPage(c *gin.Context) {
	ctl.commentService.CommentPage(c)
}
