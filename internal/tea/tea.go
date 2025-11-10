package tea

import (
	"fmt"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/klog"
	"github.com/kearth/tea/frame/container"
	"github.com/kearth/tea/internal/bootstrap"
)

// Tea 茶
type Tea struct {
	container.Unit
	Version string // 版本号
	Load    func()
}

// New
func New(version string, load func()) *Tea {
	return &Tea{
		Version: version,
		Unit:    container.NewUnit("Tea").SetRole(container.RoleFramework),
		Load:    load,
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

func (t *Tea) PrintSucc(ctx kctx.Context, u container.Unit) {
	klog.Print(ctx, fmt.Sprintf("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< [%s][%s] success <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<", u.Role(), u.Name()))
}

func (t *Tea) PrintError(ctx kctx.Context, u container.Unit, err error) {
	klog.Print(ctx, fmt.Sprintf("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< [%s][%s] error:%e <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<", u.Role(), u.Name(), err))
}

func (t *Tea) PrintPanic(ctx kctx.Context, u container.Unit, err interface{}) {
	klog.Error(ctx, fmt.Sprintf("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< [%s][%s] panic:%v <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<", u.Role(), u.Name(), err))
}

// SetupUnit 安装服务单元
func (t *Tea) SetupUnit(ctx kctx.Context) error {
	defer func() {
		// 捕获异常
		if err := recover(); err != nil {
			t.PrintPanic(ctx, t, err)
		}
	}()

	//  初始化环境
	if err := bootstrap.Env().SetVersion(t.Version).Setup(ctx); err != nil {
		t.PrintError(ctx, bootstrap.Env(), err)
		return err // 发生错误时停止
	}
	t.PrintSucc(ctx, bootstrap.Env())

	// 加载自定义初始化
	if t.Load != nil {
		t.Load()
	}

	// 初始化加载器
	units := t.NeedLoadUnits()
	for _, u := range units {
		if err := u.Setup(ctx); err != nil {
			t.PrintError(ctx, u, err)
			return err // 发生错误时停止
		}
		t.PrintSucc(ctx, u)
	}

	bootstrap.LoadInstance().Run(ctx)
	return nil
}

// NeedLoadUnits 需要加载的服务单元
func (t *Tea) NeedLoadUnits() []container.Unit {
	units := []container.Unit{}
	units = append(units, bootstrap.LoadInstance()) // 加载器
	return units
}
