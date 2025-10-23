# 个人博客系统后端

# 运行环境
- Go版本：1.25.1
- 数据库: MySQL 5.7.17-log
- 操作系统: Windows 11

# 安装步骤
## 1、克隆项目
- git clone <项目仓库地址>
- cd your-project-name
## 2、安装 Go 依赖
- go mod tidy
- 或者指定下载
- go get -u gorm.io/gorm
- go get -u gorm.io/driver/mysql
- go get -u github.com/gin-gonic/gin
- go get -u github.com/golang-jwt/jwt/v5
- go get -u github.com/joho/godotenv

# 环境配置
- envConfig.go 配置文件
- .env 配置文件

# 启动方式
- 方式一：直接运行 go run main.go （开发环境）

# 中间件执行顺序流程图
请求进入  
↓  
CustomRecovery() [defer注册]  
↓  
CorsMiddleware() [处理跨域]  
↓  
CustomLogger() [记录开始时间、生成RequestID]  
↓  
DetailedLogger() [记录请求体、包装ResponseWriter]  
↓  
JWTAuth() [仅对/auth路由] [Token验证]  
↓  
接口业务逻辑 [你的Controller方法]  
↓  
DetailedLogger() [记录响应体、耗时]  
↓  
CustomLogger() [记录响应状态码、总耗时]  
↓  
CorsMiddleware() [添加CORS响应头]  
↓  
CustomRecovery() [如果发生panic，在这里捕获]  
↓  
响应返回客户端  

# 在线API接口调试
https://s.apifox.cn/9aea8886-a229-48c6-a407-0b0f61361758/365733359e0