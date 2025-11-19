[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/zapkratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/zapkratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/zapkratos)](https://pkg.go.dev/github.com/orzkratos/zapkratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/zapkratos/main.svg)](https://coveralls.io/github/orzkratos/zapkratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/zapkratos.svg)](https://github.com/orzkratos/zapkratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/zapkratos)](https://goreportcard.com/report/github.com/orzkratos/zapkratos)

# zapkratos

将 Uber Zap 与 Kratos 微服务框架集成的日志适配器，提供高性能的结构化日志记录。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## ENGLISH README

[English](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 主要特性

- 🚀 简单集成 - 只需几行代码即可替换 Kratos 默认日志功能
- 📊 结构化日志 - 使用 Uber Zap 快速的结构化日志能力
- ⚡ 高性能 - 受益于 Zap 的零内存分配设计
- 🔧 灵活配置 - 构建配置使用自定义选项
- 🎯 模块追踪 - 自动添加模块信息到日志
- ⚙️ 适配模式 - 轻松桥接 Zap 和 Kratos

## 安装

```bash
go get github.com/orzkratos/zapkratos
```

## 快速开始

```go
package main

import (
    "github.com/go-kratos/kratos/v2"
    "github.com/orzkratos/zapkratos"
    "github.com/yyle88/zaplog"
)

func main() {
    // 创建 ZapKratos 实例
    zapKratos := zapkratos.NewZapKratos(
        zaplog.LOGGER,
        zapkratos.NewOptions(),
    )

    // 获取带模块上下文的日志
    zapLog := zapKratos.SubZap()
    zapLog.LOG.Info("application starting...")

    // 在 Kratos 应用中使用
    app := kratos.New(
        kratos.Name("my-service"),
        kratos.Logger(zapKratos.NewLogger("app")),
    )

    if err := app.Run(); err != nil {
        zapLog.LOG.Fatal("app run failed", zap.Error(err))
    }
}
```

## 完整示例

查看 [zapkratos-demos](https://github.com/orzkratos/zapkratos-demos) 了解实际 Kratos 项目中的完整集成：

- **[demo1kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo1kratos)** - HTTP 和 gRPC 基础集成
- **[demo2kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo2kratos)** - Wire 依赖注入高级用法

演示项目展示：
- main.go 和 Wire 配置中的集成方式
- biz/service/data 各层的使用方法
- HTTP 和 gRPC 服务中的 zapkratos 配置
- Log Helper 在业务逻辑中的使用
## API 参考

### ZapKratos

包装 Zap 日志并提供 Kratos 兼容接口的主要结构体。

```go
type ZapKratos struct {
    // 包含已过滤或未导出的字段
}
```

#### 构造函数

```go
func NewZapKratos(zap *zaplog.Zap, options *Options) *ZapKratos
```

使用给定的 Zap 和选项创建新的 ZapKratos 实例。

#### 方法

```go
func (A *ZapKratos) GetZap() *zaplog.Zap
```

返回底层的 Zap 实例。

```go
func (A *ZapKratos) SubZap() *zaplog.Zap
```

创建带有调用模块信息的子 Zap。

```go
func (A *ZapKratos) GetLogger(msgCaption string) log.Logger
```

创建带消息说明的 Kratos log.Logger。

```go
func (A *ZapKratos) NewLogger(msgCaption string) log.Logger
```

与 GetLogger 相同，创建 Kratos log.Logger。

```go
func (A *ZapKratos) GetHelper(msgCaption string) *log.Helper
```

创建带消息说明的 Kratos log.Helper。

```go
func (A *ZapKratos) NewHelper(msgCaption string) *log.Helper
```

与 GetHelper 相同，创建 Kratos log.Helper。

### Options

ZapKratos 的配置选项。

```go
type Options struct {
    ModuleKeyName string // 日志输出中的模块字段键名
}
```

#### 构造函数

```go
func NewOptions() *Options
```

创建具有默认设置的 Options（模块键 = "module"）。

#### 方法

```go
func (T *Options) WithModuleKeyName(moduleKeyName string) *Options
```

以可链式调用的方式设置自定义模块字段键名。

### LogImp

使用 Zap 实现 Kratos log.Logger 接口的适配器。

```go
type LogImp struct {
    // 包含已过滤或未导出的字段
}
```

#### 构造函数

```go
func NewLogImp(zapLog *zap.Logger, msgCaption string) log.Logger
```

创建包装给定 Zap 的 LogImp 适配器。
## 依赖项

- `github.com/go-kratos/kratos/v2` - Kratos 微服务框架
- `go.uber.org/zap` - Uber Zap 结构化日志
- `github.com/yyle88/zaplog` - Zap 管理包
- `github.com/yyle88/runpath` - 运行时路径工具
- `github.com/yyle88/erero` - 错误处理工具

## 相关项目

**框架：**
- [Kratos](https://github.com/go-kratos/kratos) - Go 微服务框架
- [Zap](https://github.com/uber-go/zap) - Uber 的结构化日志

**zapkratos 生态：**
- [zapkratos](https://github.com/orzkratos/zapkratos) - 本项目
- [zapkratos-demos](https://github.com/orzkratos/zapkratos-demos) - 演示项目
  - [demo1kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo1kratos) - 基础集成
  - [demo2kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo2kratos) - 高级用法

**zapzhkratos 生态（中文版）：**
- [zapzhkratos](https://github.com/orzkratos/zapzhkratos) - 使用中文函数名的中文版本
- [zapzhkratos-demos](https://github.com/orzkratos/zapzhkratos-demos) - 中文版演示项目
  - [demo1kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo1kratos) - 基础集成
  - [demo2kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo2kratos) - 高级用法

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证

MIT License - 查看 [LICENSE](LICENSE) 文件

---

## 💬 联系反馈

**问题和反馈：**

- 🐛 **Bug 报告？** 打开 issue 并描述问题和复现步骤
- ✨ **功能想法？** 打开 issue 讨论实现方案
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/zapkratos.git`）
3. **导航**：进入克隆的项目（`cd zapkratos`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细说明

请确保测试通过并包含相关的文档更新。

---

## 🌟 支持

欢迎通过提交 Merge Request 和报告 Issue 为此项目做贡献。

**项目支持：**

- ⭐ **给 GitHub 星标** 如果这个项目帮助了你
- 🤝 **分享给队友** 和（golang）编程朋友
- 📝 **写技术博客** 关于开发工具和工作流程 - 我们提供内容写作支持
- 🌟 **加入生态系统** - 致力于支持开源和（golang）开发场景

**用这个包快乐编程！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/zapkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/zapkratos)
