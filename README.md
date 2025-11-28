# Tea Framework

Tea是一个轻量级、高性能的Go语言Web框架，专为快速构建HTTP服务而设计。框架基于模块化架构，提供简洁的API接口，帮助开发者高效地创建Web应用。

## 框架特点

- **轻量级设计**：核心代码简洁，易于理解和扩展
- **高性能HTTP服务器**：基于goframe/ghttp实现，支持高并发
- **路由管理**：支持路由分组、中间件和静态文件服务
- **API文档生成**：内置Swagger/OpenAPI支持，便于API文档管理
- **优雅的错误处理**：统一的错误处理机制和日志记录
- **灵活的配置系统**：支持多环境配置管理
- **单元化架构**：基于Unit的服务管理机制，便于组件管理和生命周期控制

## 安装

```bash
go get github.com/kearth/tea
```

## 快速开始

### 创建基本应用

```go
package main

import (
    "github.com/kearth/klib/kctx"
    "github.com/kearth/tea"
)

func main() {
    // 创建一个简单的Tea应用
    tea.Drink(kctx.New(), func(ctx kctx.Context) {
        // 在这里可以添加自定义初始化逻辑
    })
}
```

### 添加HTTP路由

```go
package main

import (
    "github.com/gogf/gf/v2/net/ghttp"
    "github.com/kearth/klib/kctx"
    "github.com/kearth/tea"
    "github.com/kearth/tea/frame/server/httpserver"
)

func main() {
    tea.Drink(kctx.New(), func(ctx kctx.Context) {
        // 获取HTTP路由
        router := httpserver.NewHTTPRouter()
        
        // 添加中间件
        router.AddMiddleware(func(r *ghttp.Request) {
            r.Response.Writeln("Hello from middleware!")
            r.Middleware.Next()
        })
        
        // 添加路由处理
        router.AddBind(func(r *ghttp.Request) {
            r.Response.Writeln("Hello, Tea Framework!")
        })
        
        // 创建路由组
        router.Group("/api", func(group *httpserver.Group) {
            group.Bind("/user", func(r *ghttp.Request) {
                r.Response.Writeln("User API")
            })
        })
    })
}
```

## 目录结构

Tea框架采用清晰的目录结构组织代码：

- `frame/` - 框架核心组件
  - `base/` - 基础功能和工具
  - `env/` - 环境变量和配置管理
  - `server/` - 服务器相关实现
    - `httpserver/` - HTTP服务器实现
- `internal/` - 内部实现
  - `bootstrap/` - 框架引导程序
  - `tea/` - 核心实现
- `cli/` - 命令行工具
- `examples/` - 示例项目

## 核心功能

### HTTP服务器

Tea框架提供了强大的HTTP服务器功能，支持：

- 路由管理和分组
- 中间件机制
- 静态文件服务
- Swagger/OpenAPI文档
- 优雅启动和关闭

### 配置管理

支持多环境配置，可通过配置文件或环境变量进行配置。

### 日志系统

集成了日志记录功能，支持不同级别的日志输出。

## 示例项目

框架提供了完整的示例项目，位于`examples/`目录下，展示了如何使用Tea框架构建完整的Web应用。

## 版本信息

当前版本：v0.1.0

## 许可证

详细许可证信息请参阅 [LICENSE](LICENSE) 文件。

## 特别鸣谢

本项目依赖 [GoFrame v2](https://goframe.org/) 框架，感谢 GoFrame 团队提供的优秀开源作品。

GoFrame 作为一个全功能的 Go 语言框架，为 Tea Framework 提供了强大的底层支持，特别是在 HTTP 服务器、路由管理和配置系统等方面。Tea Framework 在 GoFrame 的基础上进行了轻量级封装和扩展，旨在提供更简洁的 API 和更灵活的应用构建体验。