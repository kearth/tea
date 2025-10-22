package bootstrap

import (
	"github.com/kearth/klib/kerr"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/tctx"
	"github.com/kearth/tea/frame/utils"
)

// Install 安装组件
func Install(ctx tctx.Context, c []container.Component, sf func(string)) error {
	for _, v := range c {
		if err := v.Init(ctx); err != nil {
			return kerr.New(100001, utils.SPF("[%s] init error", v.Name())).WithDisplay(err.Error())
		}
		sf(utils.SPF("[Component][%s] init success", v.Name()))
	}
	return nil
}
