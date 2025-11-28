

# Tea Framework Examples

## 项目简介

本目录包含 Tea Framework 的完整示例项目，展示了如何使用 Tea Framework 构建 Web 应用。Tea Framework 是一个基于 Go 语言的轻量级 Web 框架，提供了简洁的 API 和丰富的功能。

## 目录结构

```
examples/
├── app/                    # 应用程序代码
│   ├── api/                # API 接口实现
│   │   └── demo/           # 示例 API
│   ├── infra/              # 基础设施层
│   │   ├── ecode/          # 错误码定义
│   │   ├── out/            # 输出工具
│   │   └── structs/        # 数据结构
│   ├── load/               # 加载器
│   └── page/               # 页面处理
│       └── demo/           # 示例页面
├── config/                 # 配置文件
│   ├── server.toml         # 服务器配置
│   ├── log.toml            # 日志配置
│   └── var.toml            # 变量配置
├── manifest/               # 应用清单
│   ├── cmd/                # 命令入口
│   │   └── server/         # 服务器入口
│   │       ├── main.go     # 主函数
│   │       ├── server.toml # 命令特定配置
│   │       └── var.toml    # 命令特定变量
│   └── bin/                # 编译输出目录
├── resource/               # 资源文件
│   ├── css/                # CSS 样式
│   ├── public/             # 公共静态资源
│   │   ├── html/           # HTML 页面
│   │   └── file/           # 示例文件
│   └── template/           # 模板文件
├── go.mod                  # Go 模块文件
├── go.sum                  # 依赖校验文件
└── README.md               # 项目说明文件
```

## 功能说明

### 1. API 接口示例

框架提供了简洁的 API 接口定义和处理方式，支持 RESTful API 开发。示例包含：

- **Hello API**: 简单的问候接口，展示基本的请求响应处理
- **文件下载 API**: 展示如何处理文件下载
- **流式响应 API**: 展示如何实现 SSE (Server-Sent Events) 流式响应

### 2. 页面处理示例

框架支持 HTML 页面渲染，示例包含：

- **Welcome 页面**: 展示基本的 HTML 页面渲染
- **Stream 页面**: 展示与流式 API 交互的前端实现

### 3. 中间件支持

框架提供了灵活的中间件机制，示例中使用了响应处理中间件。

### 4. 静态文件服务

支持静态文件路由配置，可轻松提供 CSS、JavaScript 等静态资源。

### 5. 配置管理

使用 TOML 格式的配置文件，支持环境变量和配置覆盖。

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 运行示例

```bash
cd examples
cd manifest/cmd/server
go run main.go
```

或直接运行：

```bash
cd examples
tea run
```

### 3. 访问示例

启动服务后，可以访问以下地址：

- 欢迎页面: http://localhost:9106/page/demo/v2/welcome
- 流式示例: http://localhost:9106/page/demo/v1/stream
- API 接口: http://localhost:9106/api/demo/v1/say (POST)

## 如何扩展

### 添加新的 API

1. 在 `app/api/` 目录下创建新的 API 处理函数
2. 在 `app/load/router.go` 中注册新的路由

### 添加新的页面

1. 在 `app/page/` 目录下创建新的页面处理函数
2. 在 `resource/public/html/` 目录下创建对应的 HTML 文件
3. 在 `app/load/router.go` 中注册新的路由

### 自定义中间件

在 `app/load/router.go` 中使用 `AddMiddleware` 方法添加自定义中间件。

## 配置说明

主要配置文件位于 `config/server.toml`，包含以下关键配置：

```toml
# 服务配置
[server]
name = "Tea Demo"          # 服务器名称
version = "1.0.0"          # 服务器版本
mode = "normal"            # 服务器运行模式 - debug, normal
ip = "127.0.0.1"           # 服务器IP地址
port = 9106                # 服务器端口号
root_dir = "."             # 服务器根目录
resources_dir = "resource" # 服务器资源目录
server_type = "http"       # 服务器类型 - 默认 http
```

## 框架版本

当前示例使用的 Tea Framework 版本为 0.0.9。

## 许可证

请参考项目根目录的许可证文件。