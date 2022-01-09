package model

import (
	"log"
	"smart-server/model/cache"
	"smart-server/model/conf"
	"smart-server/model/mysql"
)

// ModelRegister 数据层注册
func ModelRegister() {
	// 注册配置文件
	if err := conf.InitConfig("app.ini"); err != nil {
		log.Fatalf("Init Conf Fail! err = %+v", err)
	}

	// 注册mysql
	if err := mysql.InitDB(); err != nil {
		log.Fatalf("Init Mysql Fail! err = %+v", err)
	}

	// 注册redis缓存
	if err := cache.InitCache(); err != nil {
		log.Fatalf("Init redis Fail! err = %+v", err)

	}
}
