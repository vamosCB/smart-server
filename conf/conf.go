package conf

import (
	"log"

	"gopkg.in/ini.v1"
)

// Server 服务配置
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
}

// Mysql mysql数据库配置结构
type Mysql struct {
	User     string
	Password string
	Host     string
	Name     string
}

// Redis redis配置结构
type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

var RedisConf = &Redis{}
var ServerConf = &Server{}
var MysqlConf = &Mysql{}

// ConfigRegister 解析文件，注册配置
func ConfigRegister(file string) {
	cfg, err := ini.Load(file)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse app.ini: %v", err)
	}

	if err := cfg.Section("server").MapTo(ServerConf); err != nil {
		log.Fatalf("setting.Setup, fail to parse app.ini: %v", err)
	}

	if err := cfg.Section("mysql").MapTo(MysqlConf); err != nil {
		log.Fatalf("setting.Setup, fail to parse app.ini: %v", err)
	}

	if err := cfg.Section("redis").MapTo(RedisConf); err != nil {
		log.Fatalf("setting.Setup, fail to parse app.ini: %v", err)
	}
}
