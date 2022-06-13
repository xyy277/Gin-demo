package test

import (
	"context"
	"encoding/json"
	"gsafety/model/test"
	"gsafety/model/test/request"
	"strconv"

	"gsafety/global"
)

type TestService struct{}

func (testService *TestService) TestPost(id uint, name string) (data string, err error) {
	// uint => string
	err = global.GS_REDIS.Set(context.Background(), strconv.Itoa(int(id)), name, 0).Err()
	const get_back = "test post is OKay"
	return get_back, err
}

func (testService *TestService) TestGet(id string) (data string, err error) {
	get_back, err := global.GS_REDIS.Get(context.Background(), id).Result()
	return get_back, err
}

func (TestService *TestService) GetTestInfoList(info request.TestSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GS_DB.Model(&test.Test{})
	json, err := json.MarshalIndent(info, "", " ")
	global.GS_LOG.Info(string(json))
	var testList []test.Test
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TestName != "" {
		db = db.Where("test_name LIKE ?", "%"+info.TestName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&testList).Error
	return testList, total, err
}
