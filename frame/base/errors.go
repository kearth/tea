package base

import (
	"github.com/kearth/klib/kerr"
)

// Base error definitions
var (
	ConfigFileNotExists   = kerr.New(100001, "Config file not exists").WithDisplay("配置文件不存在")
	DefaultServerNotFound = kerr.New(100101, "Default server not found").WithDisplay("默认服务器未找到")
	RouterNotFound        = kerr.New(100102, "Router not found").WithDisplay("路由器未找到")
	ListenerNotFound      = kerr.New(100103, "Listener not found").WithDisplay("监听器未找到")
	ResourceNotFound      = kerr.New(100104, "Resource not found").WithDisplay("资源未找到")
)
