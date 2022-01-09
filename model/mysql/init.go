package mysql

import (
	"fmt"
	"smart-server/model/conf"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
	DB, err = gorm.Open(mysql.Open(conf.MysqlConf.Dsn()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
	})

	if err != nil {
		return fmt.Errorf("models.Setup err: %v", err)
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)  //用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxOpenConns(100) //设置打开数据库连接的最大数量
	return nil
}
