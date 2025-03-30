package cmd

import (
	"flag"

	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/tctx"
	"github.com/kearth/tea/internal/envinfo"
)

var (
	// 接口检查
	_ container.Component = (*Cmd)(nil)

	// 实例
	instance = &Cmd{}
)

// Cmd 命令行组件
type Cmd struct {
	container.BaseObject
}

// Instance 获取实例
func Instance() *Cmd {
	return instance
}

// Init 初始化
func (c *Cmd) Init(ctx tctx.Context) error {
	c.SetName("Cmd")
	debug := flag.Bool("debug", false, "调试模式")
	rootDir := flag.String("root_dir", "", "根目录")
	flag.Parse()

	// 开启调试模式
	if *debug {
		envinfo.Instance().SetDebug()
	}

	// 设置根目录
	if *rootDir != "" {
		envinfo.Instance().SetRootDir(*rootDir)
	}
	return nil
}
