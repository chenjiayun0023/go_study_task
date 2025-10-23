package controller

import (
	"go_study/task4/service"

	"github.com/gin-gonic/gin"
)

var PostCtl *PostController

type PostController struct {
	postService *service.PostService
}

func NewPostController() *PostController {
	PostCtl = &PostController{
		postService: service.PostSve,
	}
	return PostCtl
}

func (ctl *PostController) CreatePost(c *gin.Context) {
	ctl.postService.CreatePost(c)
}

func (ctl *PostController) PostPage(c *gin.Context) {
	ctl.postService.PostPage(c)
}

func (ctl *PostController) UpdatePost(c *gin.Context) {
	ctl.postService.UpdatePost(c)
}

func (ctl *PostController) DeletePost(c *gin.Context) {
	ctl.postService.DeletePost(c)
}
