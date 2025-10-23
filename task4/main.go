package main

import (
	"go_study/task4/config"
	"go_study/task4/initData"
	"go_study/task4/router"
	"log"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化数据库
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移（仅在开发环境）
	if cfg.Env == "development" {
		initData.CreateTables(db)
	}

	// 设置路由
	router := router.SetupRouter(db, cfg)

	// 启动服务器
	log.Printf("Server starting on %s", cfg.ServerPort)
	err = router.Run(":" + cfg.ServerPort)
	if err != nil {
		panic(err)
	}
}
