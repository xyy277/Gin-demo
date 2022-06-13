package router

import (
	test "gsafety/router/test"
)

type RouterGroup struct {
	Test test.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
