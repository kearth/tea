package cmd

import (
	"flag"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/internal/bootstrap"
)

var (
	// 实例
	instance = &Cmd{}
)

// Cmd 命令行组件
type Cmd struct {
	container.Unit
}

// Instance 获取实例
func Instance() *Cmd {
	return instance
}

// Init 初始化
func (c *Cmd) Init(ctx kctx.Context) error {
	c.Unit = container.NewUnit("Cmd")
	debug := flag.Bool("debug", false, "调试模式")
	rootDir := flag.String("root_dir", "", "根目录")
	flag.Parse()

	// 开启调试模式
	if *debug {
		bootstrap.LoadEnvInfo().SetDebug()
	}

	// 设置根目录
	if *rootDir != "" {
		bootstrap.LoadEnvInfo().SetRootDir(*rootDir)
	}
	return nil
}
