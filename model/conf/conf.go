package conf

import (
	"fmt"

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
	MaxRetries  int
	DialTimeout int
	IdleTimeout int
}

var RedisConf = &Redis{}
var ServerConf = &Server{}
var MysqlConf = &Mysql{}

// InitConfig 解析文件，注册配置
func InitConfig(file string) error {
	cfg, err := ini.Load(file)
	if err != nil {
		return err
	}

	if err := cfg.Section("server").MapTo(ServerConf); err != nil {
		return err
	}

	if err := cfg.Section("mysql").MapTo(MysqlConf); err != nil {
		return err
	}

	if err := cfg.Section("redis").MapTo(RedisConf); err != nil {
		return err
	}
	return nil
}

// Dsn 返回mysql连接串
func (mysql *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysql.User,
		mysql.Password,
		mysql.Host,
		mysql.Name)
}
