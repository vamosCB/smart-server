package mysql

import (
	"fmt"
	"smart-server/conf"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// BaseColumn 基础字段
type BaseColumn struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" comment:"更新时间"`
}

// InitDB 初始化mysql数据库
func InitDB() error {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.MysqlConf.User,
		conf.MysqlConf.Password,
		conf.MysqlConf.Host,
		conf.MysqlConf.Name))

	if err != nil {
		return fmt.Errorf("models.Setup err: %v", err)
	}
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	return nil
}
