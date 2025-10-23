package router

import (
	"go_study/task4/config"
	"go_study/task4/controller"
	"go_study/task4/middleware"
	"go_study/task4/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.EnvConfig) *gin.Engine {
	router := gin.Default()

	// 全局中间件
	router.Use(
		middleware.CustomRecovery(), // 恢复中间件放在最外层
		middleware.CorsMiddleware(), // 跨域中间件（紧接恢复之后）
		middleware.CustomLogger(),   // 基础日志
		middleware.DetailedLogger(), // 详细接口日志
	)

	// 健康检查路由（不记录详细日志）
	//router.GET("/health", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"status": "ok"})
	//})

	// 初始化实例
	initBean(db, cfg)

	// 路由组
	api := router.Group("/api/v1")
	{
		// 白名单
		white := api.Group("/w")
		{
			white.POST("/register", controller.UserCtl.Register)
			white.POST("/login", controller.LoginCtl.Login)
		}

		auth := api.Group("/a")
		auth.Use(middleware.JWTAuth())
		{
			// 文章路由
			post := auth.Group("/post")
			{
				post.POST("/createPost", controller.PostCtl.CreatePost)
				post.POST("/postPage", controller.PostCtl.PostPage)
				post.POST("/updatePost", controller.PostCtl.UpdatePost)
				post.POST("/deletePost", controller.PostCtl.DeletePost)
			}
			// 评论路由
			comment := auth.Group("/comment")
			{
				comment.POST("/createComment", controller.CommentCtl.CreateComment)
				comment.POST("/commentPage", controller.CommentCtl.CommentPage)
			}
		}
	}

	// 404 处理
	router.NoRoute(func(c *gin.Context) {
		log.Printf("[404] %s %s from %s", c.Request.Method, c.Request.URL.Path, c.ClientIP())
		c.JSON(404, gin.H{
			"error": "接口不存在",
			"code":  "NOT_FOUND",
		})
	})

	return router
}

func initBean(db *gorm.DB, cfg *config.EnvConfig) {
	// 统一创建Service
	service.NewUserService(db, cfg)
	service.NewPostService(db, cfg)
	service.NewCommentService(db, cfg)

	// 统一创建Controller
	controller.NewLoginController()
	controller.NewUserController()
	controller.NewPostController()
	controller.NewCommentController()
}
