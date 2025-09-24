package routes

import (
	"gin-template/internal/app/config"

	democtr "gin-template/internal/demo/controller"
	demorepo "gin-template/internal/demo/repository"
	demosvc "gin-template/internal/demo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes 初始化依赖，注册路由
func SetupRoutes(cfg *config.Config, router *gin.Engine, db *gorm.DB) {
	// 初始化依赖
	// 初始化仓库层
	demoRepo := demorepo.NewDemoRepository(db)
	// 初始化服务层
	demoSvc := demosvc.NewDemoService(demoRepo)
	// 初始化控制器层
	demoController := democtr.NewDemoController(demoSvc)

	// 初始化路由
	api := router.Group("/api")
	{
		// 测试路由
		api.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "测试",
			})
		})
		// demo 模块路由
		demo := api.Group("/demo")
		{
			demo.GET("", demoController.ListDemo)
			demo.GET("/page", demoController.ListDemoPage)
			demo.GET("/:id", demoController.GetDemoByID)
			demo.POST("", demoController.CreateDemo)
			demo.POST("/batch", demoController.BatchCreateDemo)
			demo.PUT("/:id", demoController.UpdateDemo)
			demo.DELETE("/soft/:id", demoController.SoftDeleteDemo)
			demo.DELETE("/hard/:id", demoController.DeleteDemo)
		}
	}
}
