package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/kearth/tea/frame/errs"
)

// ParseTOML 解析配置文件
func ParseTOML(path string) (*gcfg.Config, error) {
	if !gfile.Exists(path) {
		return nil, errs.ConfigFileNotExists
	}
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(path)
	return g.Cfg(), nil
}
