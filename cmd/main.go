package main

import (
	"gin-template/internal/app/config"
	"gin-template/internal/app/database"
	"gin-template/internal/app/middleware"
	"gin-template/internal/app/routes"
	"gin-template/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 初始化日志记录器
	utils.InitLogger()

	// 加载配置
	cfg, err := config.LoadConfig("../config.yaml")
	if err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}

	// 初始化数据库
	db, err := database.NewDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	sqlDB, _ := db.DB() // 获取底层的 SQL 数据库连接
	defer sqlDB.Close() // 确保程序退出时关闭数据库连接，释放资源

	// 设置Gin模式
	// 生成环境设置为发布模式，发布模式的主要特性：
	// 关闭调试日志，仅保留关键错误信息
	// 启用响应数据压缩，提高传输效率
	// 禁用堆栈跟踪输出，增强安全性
	// 性能优化，减少不必要的运行时检查
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建一个最基础的路由引擎实例，不包含任何默认中间件
	router := gin.New()
	PORT := cfg.App.Port

	// 注册中间件
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.RequestIdInject())

	// 初始化依赖及注册路由
	routes.SetupRoutes(cfg, router, db)

	// 启动服务器
	logrus.Infof("服务器运行在端口 %d", PORT)
	if err := http.ListenAndServe(":"+strconv.Itoa(PORT), router); err != nil {
		logrus.Fatalf("服务器启动失败: %v", err)
	}
}
