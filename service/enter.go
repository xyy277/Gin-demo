package service

import (
	"gsafety/service/test"
)

type ServiceGroup struct {
	TestServiceGroup test.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
