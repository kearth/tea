package tea

import (
	"fmt"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kerr"
	"github.com/kearth/klib/klog"
	"github.com/kearth/klib/kunit"
	"github.com/kearth/tea/internal/bootstrap"
)

// Tea 茶
type Tea struct {
	kunit.Unit
	Version string // 版本号
	Load    func()
}

// New
func New(version string, load func()) *Tea {
	return &Tea{
		Version: version,
		Unit:    kunit.NewUnit("Tea").SetRole(kunit.RoleFramework),
		Load:    load,
	}
}

// Run
func (t *Tea) Run(ctx kctx.Context) {
	// 初始化日志
	klog.Init()
	// 开始
	klog.ColorPrint(ctx, klog.Green, "********************************* Tea Framework Begin *************************************")
	klog.ColorPrint(ctx, klog.Green, "*  TTTTT  EEEEE   AAA   FFFFF  RRRR    AAA   M   M  EEEEE  W   W   OOO    RRRR   K   K    *")
	klog.ColorPrint(ctx, klog.Green, "*    T    E      A   A  F      R   R  A   A  M M M  E      W   W  O   O   R   R  K  K     *")
	klog.ColorPrint(ctx, klog.Green, "*    T    EEE    AAAAA  FFFFF  RRR    AAAAA  M M M  EEE    W W W  O   O   RRR    KKK      *")
	klog.ColorPrint(ctx, klog.Green, "*    T    E      A   A  F      R  R   A   A  M   M  E      W W W  O   O   R  R   K  K     *")
	klog.ColorPrint(ctx, klog.Green, "*    T    EEEEE  A   A  F      R   R  A   A  M   M  EEEEE  W   W   OOO    R   R  K   K    *")
	klog.ColorPrint(ctx, klog.Green, "*******************************************************************************************")
	// 安装服务单元
	if err := t.SetupUnit(ctx); err != nil {
		klog.Error(ctx, err.Error())
		return
	}
	klog.ColorPrint(ctx, klog.Green, "********************************** Tea Framework End **************************************")
}

func (t *Tea) PrintSucc(ctx kctx.Context, u kunit.Unit) {
	klog.ColorPrint(ctx, klog.Yellow, fmt.Sprintf("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< [%s][%s][%s] success <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<", u.Role(), u.Name(), u.Cost()))
}

func (t *Tea) PrintError(ctx kctx.Context, u kunit.Unit, err error) {
	klog.ColorPrint(ctx, klog.Red, fmt.Sprintf("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< [%s][%s][%s] error:%e <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<", u.Role(), u.Name(), u.Cost(), err))
}

func (t *Tea) PrintPanic(ctx kctx.Context, u kunit.Unit, err interface{}) {
	klog.ColorPrint(ctx, klog.HiRed, fmt.Sprintf("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< [%s][%s][%s] panic:%v <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<", u.Role(), u.Name(), u.Cost(), err))
	klog.Panic(ctx, err)
}

func (t *Tea) PrintUserDefine(ctx kctx.Context) {
	klog.ColorPrint(ctx, klog.Magenta, "<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< User Define Load <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
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
	envInstance := bootstrap.Env()
	envInstance.SetFn(func(ctx kctx.Context, input ...any) (output any, err kerr.Error) {
		return nil, envInstance.SetVersion(t.Version).Setup(ctx)
	})
	if _, err := envInstance.Call(ctx); err != nil {
		t.PrintError(ctx, envInstance, err)
		return err // 发生错误时停止
	}
	t.PrintSucc(ctx, envInstance)

	var nu kunit.Unit
	units := t.NeedLoadUnits()
	for _, u := range units {
		nu = u
		nu.SetFn(func(ctx kctx.Context, input ...any) (output any, err kerr.Error) {
			return nil, nu.Setup(ctx)
		})
		if _, err := nu.Call(ctx); err != nil {
			t.PrintError(ctx, nu, err)
			return err // 发生错误时停止
		}
		t.PrintSucc(ctx, nu)
	}

	// 初始化加载器
	loadInstance := bootstrap.LoadInstance()
	loadInstance.SetFn(func(ctx kctx.Context, input ...any) (output any, err kerr.Error) {
		defer func() {
			// 捕获异常
			if err := recover(); err != nil {
				t.PrintPanic(ctx, loadInstance, err)
			}
		}()
		// 加载自定义初始化
		if t.Load != nil {
			t.PrintUserDefine(ctx)
			t.Load()
			t.PrintUserDefine(ctx)
		}
		return nil, loadInstance.Setup(ctx)
	})
	if _, err := loadInstance.Call(ctx); err != nil {
		t.PrintError(ctx, loadInstance, err)
		return err // 发生错误时停止
	}
	t.PrintSucc(ctx, loadInstance)
	bootstrap.LoadInstance().Run(ctx)
	return nil
}

// NeedLoadUnits 需要加载的服务单元
func (t *Tea) NeedLoadUnits() []kunit.Unit {
	units := []kunit.Unit{}
	return units
}
