package load

import (
	"fmt"

	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/klog"
)

// Loader 加载器接口
type Loader func(ctx kctx.Context) (name string, err error)

// LoadAll 加载所有
func LoadAll(ctx kctx.Context) {
	// 启动服务器
	for _, loader := range []Loader{
		LoadRouter,
	} {
		if name, err := loader(ctx); err != nil {
			klog.Panic(ctx, err)
		} else {
			klog.ColorPrint(ctx, klog.Green, fmt.Sprintf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> [Loader][%s] success >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", name))
		}
	}
}
