package core

import (
	"fmt"
	"gsafety/global"
	"gsafety/initialize"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.GS_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	// 初始化数据库
	initialize.Gorm()
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GS_CONFIG.System.Addr)
	// init
	s := func(address string, router *gin.Engine) server {
		return &http.Server{
			Addr:           address,
			Handler:        router,
			ReadTimeout:    20 * time.Second,
			WriteTimeout:   20 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
	}(address, Router)

	global.GS_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
                                 ^__^       /                
                         _______/(oo)  ___ /                    
                     /\\/(       /(__)                        
                         | w----|\\                          
                        /\\     |/    
	
	welcome to gin-gsafety
	version:v0.1
	email:zhoujiajun@gsafety.com
	default docs:http://127.0.0.1%s/swagger/index.html
	
`, address)
	global.GS_LOG.Error(s.ListenAndServe().Error())
}
