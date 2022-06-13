package test

import "gsafety/service"

type ApiGroup struct {
	TestApi
}

var (
	testService = service.ServiceGroupApp.TestServiceGroup.TestService
)
