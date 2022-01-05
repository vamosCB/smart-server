package model

import (
	"log"
	"smart-server/model/mysql"
	"smart-server/model/redis"
)

// ModelRegister 数据层注册
func ModelRegister() {
	// 注册mysql
	if err := mysql.InitDB(); err != nil {
		log.Fatalf("Init Mysql Fail! err = %+v", err)
	}

	//注册redis
	if err := redis.InitRedis(); err != nil {
		log.Fatalf("Init redis Fail! err = %+v", err)

	}
}
