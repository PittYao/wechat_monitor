package main

import (
	"github.com/PittYao/wechat_monitor/etc"
	"github.com/PittYao/wechat_monitor/model"
	"github.com/PittYao/wechat_monitor/xmysql"
	"testing"
)

func Test_ReadJson(t *testing.T) {
	model.UnmarshalJson()
}

func Test_ReadJson2Db(t *testing.T) {
	// 读取yaml配置
	etc.C = etc.NewConfig()

	// 启动gorm
	xmysql.G = xmysql.NewXGorm()

	// 读取数据
	data := model.UnmarshalJson()

	// 处理数据
	model.DataHandler(data)
}
