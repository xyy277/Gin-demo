package test

import (
	v1 "gsafety/api/v1"

	"github.com/gin-gonic/gin"
)

type TestKRouter struct{}

func (e *TestKRouter) InitTestKRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test")
	exaTestApi := v1.ApiGroupApp.TestApiGroup.TestApi
	{
		testRouter.PUT("testPut", exaTestApi.TestPut)          // test put
		testRouter.DELETE("testDelete", exaTestApi.TestDelete) // test delete
	}
}
