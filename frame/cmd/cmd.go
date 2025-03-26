package cmd

import "context"

type Cmd struct {
	// 命令名称
}

func (c *Cmd) Init(ctx context.Context) error {
	// 启动参数
	// debug := flag.Bool("debug", false, "调试模式")
	// rootDir := flag.String("root_dir", "", "根目录")
	// flag.Parse()

	// // 开启调试模式
	// if *debug {
	// 	env.DebugMode()
	// }
	// if *rootDir != "" {
	// 	env.SetRootDir(*rootDir)
	// }
	return nil
}
