package app

type App interface {
	LoadConfig(string)
	Init()
	Start() error
}
