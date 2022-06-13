package v1

import (
	test "gsafety/api/v1/test"
)

type ApiGroup struct {
	TestApiGroup test.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
