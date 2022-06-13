package test

import (
	v1 "gsafety/api/v1"

	"github.com/gin-gonic/gin"
)

type TestGRouter struct{}

func (e *TestGRouter) InitTestGRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test")
	exaTestApi := v1.ApiGroupApp.TestApiGroup.TestApi
	{
		testRouter.POST("testPost", exaTestApi.TestPost)      // test post
		testRouter.GET("testGet", exaTestApi.TestGet)         // test get
		testRouter.GET("getTestList", exaTestApi.GetTestList) // test get list
	}
}
