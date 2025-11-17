package log

import (
	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/klog"
)

func LoadLogger(ctx kctx.Context) error {
	klog.Print(ctx, "load logger")
	return nil
}
