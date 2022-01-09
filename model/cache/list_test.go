package cache

import (
	"context"
	"smart-server/model/conf"
	"sync"
	"testing"
)

var once sync.Once

func oneTestSoup(t *testing.T) {
	conf.InitConfig("../../app.ini")
	once.Do(func() {
		if err := InitCache(); err != nil {
			t.Error("get redis instance failed")
		}
	})
}
func Test_Get(t *testing.T) {
	oneTestSoup(t)
	result, err := Get(context.Background(), "chenbo")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
