package tea

import (
	"fmt"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/klog"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/internal/bootstrap"
	"github.com/kearth/tea/internal/server"
)

// Tea 茶
type Tea struct {
	container.Unit
	Version string // 版本号
}

// New
func New(version string) *Tea {
	return &Tea{
		Version: version,
		Unit:    container.NewUnit("Tea"),
	}
}

// Run
func (t *Tea) Run(ctx kctx.Context) {
	// 初始化日志
	klog.Init()
	// 开始
	klog.Print(ctx, "********************************* Tea Framework Begin *************************************")
	klog.Print(ctx, "*  TTTTT  EEEEE  AAAAA                                                                    *")
	klog.Print(ctx, "*    T    E      A   A                                                                    *")
	klog.Print(ctx, "*    T    EEE    AAAAA                                                                    *")
	klog.Print(ctx, "*    T    E      A   A                                                                    *")
	klog.Print(ctx, "*    T    EEEEE  A   A                                                                    *")
	klog.Print(ctx, "*******************************************************************************************")
	// 安装服务单元
	if err := t.SetupUnit(ctx); err != nil {
		klog.Error(ctx, err.Error())
		return
	}
	klog.Print(ctx, "********************************** Tea Framework End **************************************")
}

// SetupUnit 安装服务单元
func (t *Tea) SetupUnit(ctx kctx.Context) error {
	units := []container.Unit{
		bootstrap.EnvInfoInstance(), // 环境
		// cmd.Instance(),              // 命令行
		bootstrap.Loads(),       // 加载器
		&server.ServerPackage{}, // HTTP 服务器
	}

	for _, u := range units {
		if err := u.Setup(ctx); err != nil {
			klog.Print(ctx, fmt.Sprintf("[%s][%s] setup failed, err:%e", u.Role(), u.Name(), err))
			return err // 发生错误时停止
		}
		klog.Print(ctx, fmt.Sprintf("[%s][%s] setup success", u.Role(), u.Name()))
	}
	return nil
}
