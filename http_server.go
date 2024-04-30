package tea

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cast"
)

// Check Interface
var _ IContainer = &httpServer{}

// Bootstrap
type Bootstrap func(ctx Context) IError // 启动函数
// RouterFunc
type RouterFunc func(ctx Context) *HTTPRouter // 路由函数

// init
func init() {
	// register 注册
	IOC().Register(new(httpServer))
}

// httpServer HTTP服务器
type httpServer struct {
	http.Server              // 继承
	ConfigPath    string     // 配置文件路径
	BootstrapFunc Bootstrap  // 启动函数
	RouterFunc    RouterFunc // 路由函数
}

// Name
func (h *httpServer) Name() string {
	return "httpServer"
}

const (
	DefaultConfigPath = "./conf/app.toml" // 默认配置文件路径
)

// New 创建实例
func (h *httpServer) New() IContainer {
	// Default
	h.ConfigPath = DefaultConfigPath
	h.RouterFunc = func(ctx Context) *HTTPRouter {
		defaultRouter := NewHTTPRouter()
		defaultRouter.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, welcomePage)
		})
		return defaultRouter
	}
	h.BootstrapFunc = func(ctx Context) IError {
		return nil
	}
	return h
}

// initServe
func (h *httpServer) initServe(ctx Context) error {
	var err error
	httpconfig, err := IOC().Get("HTTPConfig")
	config := httpconfig.(*HTTPConfig)
	if err = Parse(h.ConfigPath, &config); err != nil {
		panic(err)
	}
	if config.Port != 0 {
		h.Addr = fmt.Sprintf(":%s", cast.ToString(config.Port))
	}
	var to int
	if to = config.ReadTimeout; to > 0 {
		h.ReadTimeout = time.Duration(to) * time.Millisecond
	}
	if to = config.WriteTimeout; to > 0 {
		h.WriteTimeout = time.Duration(to) * time.Millisecond
	}
	if to = config.IdleTimeout; to > 0 {
		h.IdleTimeout = time.Duration(to) * time.Millisecond
	}
	return nil
}

// Start 启动
func (h *httpServer) Start() IError {
	ctx, cancel := WithCancel(Background())
	defer h.Shutdown(ctx, cancel)
	if err := h.initServe(ctx); err != nil {
		return FrameworkError.Wrap(err)
	}
	if err := h.BootstrapFunc(ctx); err != nil {
		return err
	}
	h.Handler = h.RouterFunc(ctx)
	// TODO 本机信息输出
	// TODO 配置信息输出
	fmt.Println("ListenAndServe:8080")
	if err := h.Server.ListenAndServe(); err != nil {
		return FrameworkError.Wrap(err)
	}
	return nil
}

// Shutdown
func (h *httpServer) Shutdown(ctx Context, cancel context.CancelFunc) IError {
	// TODO
	return nil
}

// Bootstrap 设置启动函数
func (h *httpServer) SetBootstrap(bs Bootstrap) {
	h.BootstrapFunc = bs
}

// SetRouter 设置路由
func (h *httpServer) SetRouter(r RouterFunc) {
	h.RouterFunc = r
}

// SetConf
func (h *httpServer) SetConf(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	h.ConfigPath = path
	return nil
}

// 欢迎页面
var welcomePage = `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<title>框架欢迎页面</title>
<style>
  body {
    font-family: 'Arial', sans-serif;
    margin: 0;
    padding: 0;
    background-color: #f4f4f4;
  }
  .container {
    width: 80%;
    margin: auto;
    overflow: hidden;
  }
  header {
    background: #50b3a2;
    color: white;
    padding-top: 30px;
    min-height: 70px;
    border-bottom: #e8491d 3px solid;
  }
  header a {
    color: #ffffff;
    text-decoration: none;
    text-transform: uppercase;
    font-size: 16px;
  }
  header ul {
    padding: 0;
    list-style: none;
  }
  header nav {
    float: right;
    margin-top: 10px;
  }
  header nav li {
    display: inline;
    margin-left: 20px;
  }
  header #branding {
    float: left;
    margin-top: 10px;
  }
  header #branding h1 {
    margin: 0;
  }
  .main-content {
    padding: 20px;
    background: white;
  }
  footer {
    padding: 20px;
    background: #50b3a2;
    color: white;
    text-align: center;
    font-size: 14px;
  }
</style>
</head>
<body>
<div class="container">
  <header>
    <div id="branding">
      <h1>Go 框架欢迎页面</h1>
    </div>
    <nav>
      <ul>
        <li><a href="/">首页</a></li>
        <li><a href="/features">功能列表</a></li>
        <li><a href="/documentation">文档</a></li>
        <li><a href="/examples">示例代码</a></li>
      </ul>
    </nav>
  </header>
  <div class="main-content">
    <h2>欢迎使用 Go 框架</h2>
    <p>我们的框架旨在简化 Go 应用程序的开发。</p>
    <p>请访问我们的 <a href="/features">功能列表</a> 以了解更多信息。</p>
  </div>
  <footer>
    <p>&copy; 2023 Go 框架。保留所有权利。</p>
  </footer>
</div>
</body>
</html>
`
