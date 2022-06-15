package test

import "gsafety/global"

type Test struct {
	global.GS_BASE_MODEL
	TestName string `json:"testName" form:"testName"`
}

// 可指定具体表名 解决表名复数问题
func (Test) TableName() string {
	return "test"
}
