package test

import (
	"context"
	"smart-server/conf"
	"smart-server/model/mysql"
	"sync"
	"testing"
)

var once sync.Once

func oneTestSoup(t *testing.T) {
	once.Do(func() {
		if err := mysql.InitDB(); err != nil {
			t.Error("get mysql instance failed")
		}
	})
}
func Test_GetStudentInfoByID(t *testing.T) {
	//oneTestSoup(t)
	conf.ConfigRegister("../../../app.ini")
	t.Logf("%+v", conf.MysqlConf)
	if err := mysql.InitDB(); err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	result, err := GetStudentInfoByID(ctx, 1)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
