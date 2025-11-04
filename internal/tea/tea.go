package tea

import (
	"fmt"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/klog"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/frame/server"
	"github.com/kearth/tea/internal/bootstrap"
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
	klog.Print(ctx, "*  TTTTT  EEEEE   AAA   FFFFF  RRRR    AAA   M   M  EEEEE  W   W   OOO    RRRR   K   K    *")
	klog.Print(ctx, "*    T    E      A   A  F      R   R  A   A  M M M  E      W   W  O   O   R   R  K  K     *")
	klog.Print(ctx, "*    T    EEE    AAAAA  FFFFF  RRR    AAAAA  M M M  EEE    W W W  O   O   RRR    KKK      *")
	klog.Print(ctx, "*    T    E      A   A  F      R  R   A   A  M   M  E      W W W  O   O   R  R   K  K     *")
	klog.Print(ctx, "*    T    EEEEE  A   A  F      R   R  A   A  M   M  EEEEE  W   W   OOO    R   R  K   K    *")
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
		bootstrap.Env().SetVersion(t.Version), // 环境
		server.NewHTTPServer(),                // HTTP 服务器
		bootstrap.LoadInstance(),              // 加载器
	}

	for _, u := range units {
		if err := u.Setup(ctx); err != nil {
			klog.Print(ctx, fmt.Sprintf("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< [%s][%s] error:%e <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<", u.Role(), u.Name(), err))
			return err // 发生错误时停止
		}
		klog.Print(ctx, fmt.Sprintf("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< [%s][%s] success <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<", u.Role(), u.Name()))
	}
	return nil
}
