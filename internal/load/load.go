package load

import (
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/t"
	"github.com/kearth/tea/frame/tctx"
	"github.com/kearth/tea/frame/tlog"
	"github.com/kearth/tea/frame/utils"
)

var (
	// 接口
	_ container.Object = (*Load)(nil)
	// instance 实例
	instance = &Load{}
)

// Load 加载器
type Load struct {
	container.BaseObject
}

// Instance 获取实例
func Instance() *Load {
	return instance
}

// Init 初始化
func (l *Load) Init(ctx tctx.Context) error {
	l.SetName("Load")
	var err error
	// 解析步骤
	s := t.GetServer()
	for _, step := range []container.Step{
		container.NewStep("Set", s.Set),
		container.NewStep("Start", s.Start),
		container.NewStep("Stop", s.Stop),
	} {
		if err = step.Run(ctx); err != nil {
			tlog.Error(ctx, utils.SPF("[Step][%s] error:%e", step.Name(), err))
			return err
		}
		tlog.Notice(ctx, utils.SPF("[Step][%s] success", step.Name()))
	}
	return nil
}
