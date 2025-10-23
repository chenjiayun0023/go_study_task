package service

import (
	"go_study/task4/config"
	"go_study/task4/dto"
	"go_study/task4/model"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var PostSve *PostService

type PostService struct {
	Db  *gorm.DB
	cfg *config.EnvConfig
}

func NewPostService(db *gorm.DB, cfg *config.EnvConfig) *PostService {
	PostSve = &PostService{Db: db, cfg: cfg}
	return PostSve
}

func (s *PostService) CreatePost(c *gin.Context) {
	var postCreateReq dto.PostCreateReq
	if err := c.ShouldBindJSON(&postCreateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 从上下文中获取用户信息
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}
	postCreateReq.UserID = userID.(uint)

	post := dto.PostCreateReqToPost(postCreateReq)

	if err := s.Db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post, " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

func (s *PostService) PostPage(c *gin.Context) {
	var postPageReq dto.PostPageReq
	if err := c.ShouldBindQuery(&postPageReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 计算偏移量
	offset := (postPageReq.Page - 1) * postPageReq.PageSize

	// 查询数据
	var postPageRsps = make([]dto.PostPageRsp, 0)
	if err := s.Db.Raw("SELECT id, title, content, user_id, created_at, updated_at FROM posts where deleted_at is null  ORDER BY created_at DESC LIMIT ?, ?", offset, postPageReq.PageSize).Scan(&postPageRsps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query post, " + err.Error()})
		return
	}

	// 获取总数
	var total int64
	if err := s.Db.Raw("SELECT count(1) FROM posts where deleted_at is null").Scan(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query post, " + err.Error()})
		return
	}

	// 计算总页数
	totalPage := int(math.Ceil(float64(total) / float64(postPageReq.PageSize)))

	c.JSON(http.StatusOK, dto.PageResult[dto.PostPageRsp]{
		Data:      postPageRsps,
		Total:     total,
		Page:      postPageReq.Page,
		PageSize:  postPageReq.PageSize,
		TotalPage: totalPage,
	})
}

func (s *PostService) UpdatePost(c *gin.Context) {
	var postUpdateReq dto.PostUpdateReq
	if err := c.ShouldBindJSON(&postUpdateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从上下文中获取用户信息
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}

	var post model.Post
	if err := s.Db.First(&post, postUpdateReq.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query post, " + err.Error()})
		return
	}

	if userID.(uint) != post.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No permission to update the post"})
		return
	}

	if err := s.Db.Model(&post).Select("title", "content").Updates(postUpdateReq).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post, " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func (s *PostService) DeletePost(c *gin.Context) {
	var post model.Post
	if err := s.Db.First(&post, c.Query("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query post, " + err.Error()})
		return
	}

	// 从上下文中获取用户信息
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}

	if userID.(uint) != post.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No permission to delete the post"})
		return
	}

	if err := s.Db.Delete(&model.Post{}, post.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post, " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post delete successfully"})
}
