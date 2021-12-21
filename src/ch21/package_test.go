package ch21

import (
	cm "github.com/easierway/concurrent_map"
	"github.com/panguangyao/go_learning/src/ch21/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetNum(6))
}

//适应 go get 拉取包，无需提交src目录
func TestRemotePackage(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
