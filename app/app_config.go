package app

import "github.com/kearth/tea/std/conf"

var (
	appConfigPath = "conf/app.toml"
)

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
