package tea

var _ IContainer = &HTTPConfig{}

// init
func init() {
	// register
	IOC().Register(new(HTTPConfig))
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
func (h *HTTPConfig) New() IContainer {
	return h
}
