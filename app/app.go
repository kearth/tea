package app

import "context"

type App interface {
	LoadConfig(configPath string) *AppConfig
	Init(ctx context.Context, appConfig *AppConfig)
	Start() error
}

type AppConfig interface {
	Parse() error
}

type Bootstrap interface {
	Init(ctx context.Context)
	Start() error
}
