package initialize

import (
	_ "gsafety/docs"
	"gsafety/global"
	"gsafety/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)

	Router := gin.Default()
	testRouter := router.RouterGroupApp.Test

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GS_LOG.Info("register swagger handler")

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		testRouter.InitTestGRouter(PublicGroup) // testG
		testRouter.InitTestKRouter(PublicGroup) // testK

	}

	global.GS_LOG.Info("router register success")
	return Router
}
