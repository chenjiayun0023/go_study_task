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

var CommentSve *CommentService

type CommentService struct {
	Db  *gorm.DB
	cfg *config.EnvConfig
}

func NewCommentService(db *gorm.DB, cfg *config.EnvConfig) *CommentService {
	CommentSve = &CommentService{Db: db, cfg: cfg}
	return CommentSve
}

func (s *CommentService) CreateComment(c *gin.Context) {
	var commentCreateReq dto.CommentCreateReq
	if err := c.ShouldBindJSON(&commentCreateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post model.Post
	if err := s.Db.First(&post, commentCreateReq.PostID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query post, " + err.Error()})
		return
	}

	// 从上下文中获取用户信息
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}

	commentCreateReq.UserID = userID.(uint)

	comment := dto.CommentCreateReqToComment(commentCreateReq)

	if err := s.Db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment, " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully"})
}

func (s *CommentService) CommentPage(c *gin.Context) {
	var commentPageReq dto.CommentPageReq
	if err := c.ShouldBindJSON(&commentPageReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 计算偏移量
	offset := (commentPageReq.Page - 1) * commentPageReq.PageSize

	// 查询数据
	var commentPageRsps = make([]dto.CommentPageRsp, 0)
	if err := s.Db.Raw("SELECT c.id, c.content, c.user_id, u.username, c.created_at, c.updated_at FROM comments c left join users u on c.user_id = u.id WHERE c.post_id = ? and c.deleted_at is null group by c.id ORDER BY c.created_at DESC LIMIT ?, ?",
		commentPageReq.PostID, offset, commentPageReq.PageSize).Scan(&commentPageRsps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query comment, " + err.Error()})
		return
	}

	// 获取总数
	var total int64
	if err := s.Db.Raw("SELECT count(1) FROM comments WHERE post_id = ? and deleted_at is null", commentPageReq.PostID).Scan(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query comment, " + err.Error()})
		return
	}

	// 计算总页数
	totalPage := int(math.Ceil(float64(total) / float64(commentPageReq.PageSize)))

	c.JSON(http.StatusOK, dto.PageResult[dto.CommentPageRsp]{
		Data:      commentPageRsps,
		Total:     total,
		Page:      commentPageReq.Page,
		PageSize:  commentPageReq.PageSize,
		TotalPage: totalPage,
	})
}
