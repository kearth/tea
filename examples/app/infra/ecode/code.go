package ecode

import "github.com/kearth/klib/kerr"

var (
	// 成功
	Succ = kerr.Succ
	// 文件不存在
	FileNotFound = kerr.New(60001, "file not found")
	// 扫描错误
	ScannerError = kerr.New(60002, "scanner error")
)
