package service

import (
	"go_study/task4/config"
	"go_study/task4/dto"
	"go_study/task4/model"
	"go_study/task4/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var UserSve *UserService

type UserService struct {
	Db  *gorm.DB
	cfg *config.EnvConfig
}

func NewUserService(db *gorm.DB, cfg *config.EnvConfig) *UserService {
	UserSve = &UserService{Db: db, cfg: cfg}
	return UserSve
}

func (s *UserService) Register(c *gin.Context) {
	var registerUserReq dto.RegisterUserReq
	if err := c.ShouldBindJSON(&registerUserReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUserReq.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	registerUserReq.Password = string(hashedPassword)

	user := dto.RegisterUserReqToUser(registerUserReq)

	if err := s.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user, " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (s *UserService) Login(c *gin.Context) {
	var loginUserReq dto.LoginUserReq
	if err := c.ShouldBind(&loginUserReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := s.Db.Where("username = ?", loginUserReq.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exist, " + err.Error()})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUserReq.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password, " + err.Error()})
		return
	}

	// 生成 JWT
	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token, " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.UserToLoginUserRsp(user, token))
}
