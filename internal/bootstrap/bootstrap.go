package bootstrap

import (
	"fmt"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kerr"
	"github.com/kearth/tea/frame/container"
)

// Install 安装组件
func Install(ctx kctx.Context, c []container.Component, sf func(string)) error {
	for _, v := range c {
		if err := v.Init(ctx); err != nil {
			return kerr.New(100001, fmt.Sprintf("[%s] init error", v.Name())).WithDisplay(err.Error())
		}
		sf(fmt.Sprintf("[Component][%s] init success", v.Name()))
	}
	return nil
}
