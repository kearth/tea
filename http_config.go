package tea

import "github.com/kearth/tea/core"

// HTTPConfig
var _ core.IContainer = &HTTPConfig{}

// init
func init() {
	// register
	core.IOC().Register(new(HTTPConfig))
}

// HTTPConfig
type HTTPConfig struct {
	Port         int `toml:"Port"`
	ReadTimeout  int `toml:"ReadTimeout"`
	WriteTimeout int `toml:"WriteTimeout"`
	IdleTimeout  int `toml:"IdleTimeout"`
}

// Name
func (h *HTTPConfig) Name() string {
	return "HTTPConfig"
}

// New
func (h *HTTPConfig) New() core.IContainer {
	return h
}
