package app

import "context"

type App interface {
	LoadConfig(configPath string) *AppConfig
	Init(ctx context.Context, appConfig *AppConfig)
	Start() error
}
