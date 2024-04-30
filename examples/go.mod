module example/local

go 1.22

//toolchain go1.22.0

replace github.com/kearth/tea => ../../tea

require github.com/kearth/tea v0.0.0-00010101000000-000000000000

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/spf13/cast v1.6.0 // indirect
)
