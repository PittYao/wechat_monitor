package main

import (
	"github.com/PittYao/wechat_monitor/etc"
	"github.com/PittYao/wechat_monitor/xgin"
	"github.com/PittYao/wechat_monitor/xmysql"
)

func main() {
	// 读取yaml配置
	etc.C = etc.NewConfig()

	// 启动gorm
	xmysql.G = xmysql.NewXGorm()

	// 启动gin
	xgin.NewGin()

}
