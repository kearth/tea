package tlog

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kearth/klib/kctx"
	"github.com/kearth/klib/kutil"
	"github.com/kearth/tea/frame/base"
)

// 颜色映射
var colorMaps = map[int]int{
	glog.LEVEL_DEBU: glog.COLOR_YELLOW,
	glog.LEVEL_INFO: glog.COLOR_GREEN,
	glog.LEVEL_NOTI: glog.COLOR_CYAN,
	glog.LEVEL_WARN: glog.COLOR_MAGENTA,
	glog.LEVEL_ERRO: glog.COLOR_RED,
	glog.LEVEL_CRIT: glog.COLOR_HI_RED,
	glog.LEVEL_PANI: glog.COLOR_HI_RED,
	glog.LEVEL_FATA: glog.COLOR_HI_RED,
}

// 日志
var instance = g.Log()

// 初始化日志
func init() {
	glog.SetDefaultHandler(DefaultHandler)
}

// LogInfo 日志信息
type LogInfo struct {
	Time       string
	Level      string
	LevelInt   int
	ResponseID string
	RequestId  string
	Body       []any
	Add        map[string]string
}

// Logger 获取日志实例
func Logger() *glog.Logger {
	return instance
}

// formatId 格式化 ID
func formatId(id string) string {
	return kutil.If[string](id == "" || id == base.DefaultID, "-", id)
}

// formatBody 格式化日志主体
func formatBody(body []any) string {
	if len(body) > 1 {
		m := make(map[string]any)
		for i, v := range body {
			if i%2 == 1 {
				m[gconv.String(body[i-1])] = v
			}
		}
		encoded, err := gjson.EncodeString(m)
		return kutil.If[string](err != nil, "", encoded)
	} else if len(body) == 1 {
		return gconv.String(body[0])
	}
	return ""
}

// formatAdd 格式化附加参数
func formatAdd(add map[string]string) string {
	return kutil.If[string](len(add) > 0, func() string {
		encoded, err := gjson.EncodeString(add)
		return kutil.If[string](err != nil, "", encoded)
	}(), "")
}

// String 格式化日志
func (l *LogInfo) String() string {
	if l.Level == "" {
		return formatBody(l.Body)
	}
	add := formatAdd(l.Add)
	body := kutil.If[string](add != "", add, formatBody(l.Body))
	return fmt.Sprintf(
		"%s %s {%s} {%s} %s",
		l.Time,
		color.New(color.Attribute(colorMaps[l.LevelInt])).Sprint("["+l.Level+"]"),
		formatId(l.ResponseID),
		formatId(l.RequestId),
		body)
}

// DefaultHandler 默认日志处理
func DefaultHandler(ctx context.Context, in *glog.HandlerInput) {
	newCtx := kctx.New(ctx)
	in.Buffer.WriteString((&LogInfo{
		Time:       in.Time.Format("2006-01-02 15:04:05 Z07:00"),
		Level:      in.LevelFormat,
		LevelInt:   in.Level,
		ResponseID: in.TraceId,
		RequestId:  newCtx.TraceID(),
		Body:       in.Values,
		Add:        newCtx.Values(),
	}).String())
	in.Buffer.WriteString("\n")
	in.Next(ctx)
}

// AddToCtx 追加参数
func AddToCtx(ctx context.Context, key string, val string) context.Context {
	newCtx := kctx.New(ctx)
	newCtx.Set(key, val)
	return newCtx.Context()
}

// AddMapToCtx 追加参数 - map
func AddMapToCtx(ctx context.Context, kv map[string]string) context.Context {
	newCtx := kctx.New(ctx)
	for k, v := range kv {
		newCtx.Set(k, v)
	}
	return newCtx.Context()
}

// Info 打印日志
func Info(ctx context.Context, v ...any) {
	Logger().Info(ctx, v...)
}

// // Debug 打印日志
// func Debug(ctx context.Context, v ...any) {
// 	if env.IsDebug() {
// 		Logger().Debug(ctx, v...)
// 	}
// }

// Notice 打印日志
func Notice(ctx context.Context, v ...any) {
	Logger().Notice(ctx, v...)
}

// Warning 打印日志
func Warn(ctx context.Context, v ...any) {
	Logger().Warning(ctx, v...)
}

// Error 打印日志
func Error(ctx context.Context, v ...any) {
	Logger().Error(ctx, v...)
}

// Panic 打印日志
func Panic(ctx context.Context, v ...any) {
	Logger().Panic(ctx, v...)
}
