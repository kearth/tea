package tea

import (
	"github.com/kearth/tea/frame"
	"github.com/kearth/tea/frame/tctx"
)

const (
	// Version 版本号
	Version = "0.0.1"
)

// 喝一杯茶
func WorkWillBeDone() {
	frame.GetSomeTea(Version).Drink(tctx.New())
}
