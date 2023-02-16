package app

import "github.com/kearth/tea/std/conf"

var (
	appConfigPath = "conf/app.toml"
	appConf       = (*AppConfig)(nil)
)

// AppConfig
type AppConfig struct {
	AppName    string
	RunMode    string
	Port       int
	IDC        string
	HttpServer *struct {
		ReadTimeout  int
		WriteTimeout int
		IdleTimeout  int
	}
}

// SetAppConfigPath
func SetAppConfigPath(path string) {
	appConfigPath = path
}

// GetAppConfigPath
func GetAppConfigPath() string {
	return appConfigPath
}

// LoadAppConfig
func LoadAppConfig(path string) (*AppConfig, error) {
	if err := conf.Parse(path, &appConf); err != nil {
		return nil, err
	}
	return appConf, nil
}
