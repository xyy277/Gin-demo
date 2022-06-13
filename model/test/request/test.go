package request

import (
	"gsafety/model/test"

	"gsafety/model/common/request"
)

type TestSearch struct {
	test.Test
	request.PageInfo
}
