[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/zapkratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/zapkratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/zapkratos)](https://pkg.go.dev/github.com/orzkratos/zapkratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/zapkratos/main.svg)](https://coveralls.io/github/orzkratos/zapkratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/zapkratos.svg)](https://github.com/orzkratos/zapkratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/zapkratos)](https://goreportcard.com/report/github.com/orzkratos/zapkratos)

# zapkratos

Zap integration that connects Uber Zap with Kratos microservice framework, providing structured logging with high performance.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

- 🚀 Simple Integration - Replace Kratos default logging with just a few lines
- 📊 Structured Logging - Use Uber Zap's fast structured logging
- ⚡ High Performance - Benefit from Zap's zero-allocation design
- 🔧 Flexible Config - Build config with custom options
- 🎯 Module Tracking - Auto add module info to logs
- ⚙️ Adaptation Pattern - Bridge Zap and Kratos with ease

## Installation

```bash
go get github.com/orzkratos/zapkratos
```

## Quick Start

```go
package main

import (
    "github.com/go-kratos/kratos/v2"
    "github.com/orzkratos/zapkratos"
    "github.com/yyle88/zaplog"
)

func main() {
    // Create ZapKratos instance
    zapKratos := zapkratos.NewZapKratos(
        zaplog.LOGGER,
        zapkratos.NewOptions(),
    )

    // Get logging with module context
    zapLog := zapKratos.SubZap()
    zapLog.LOG.Info("application starting...")

    // Use in Kratos app
    app := kratos.New(
        kratos.Name("my-service"),
        kratos.Logger(zapKratos.NewLogger("app")),
    )

    if err := app.Run(); err != nil {
        zapLog.LOG.Fatal("app run failed", zap.Error(err))
    }
}
```

## Complete Examples

See [zapkratos-demos](https://github.com/orzkratos/zapkratos-demos) to view complete integration in actual Kratos projects:

- **[demo1kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo1kratos)** - Basic integration with HTTP and gRPC
- **[demo2kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo2kratos)** - Advanced usage with Wire DI

The demos show:
- Integration in main.go and Wire setup
- Usage across biz/service/data tiers
- HTTP and gRPC setup with zapkratos
- Log Helper usage in business logic
## API Reference

### ZapKratos

Main struct that wraps Zap logging and provides Kratos-compatible interfaces.

```go
type ZapKratos struct {
    // Contains filtered or unexported fields
}
```

#### Constructor

```go
func NewZapKratos(zap *zaplog.Zap, options *Options) *ZapKratos
```

Creates new ZapKratos instance with the given Zap and options.

#### Methods

```go
func (A *ZapKratos) GetZap() *zaplog.Zap
```

Returns the underlying Zap instance.

```go
func (A *ZapKratos) SubZap() *zaplog.Zap
```

Creates child Zap with calling module info.

```go
func (A *ZapKratos) GetLogger(msgCaption string) log.Logger
```

Creates Kratos log.Logger with the given caption.

```go
func (A *ZapKratos) NewLogger(msgCaption string) log.Logger
```

Same as GetLogger, creates the Kratos log.Logger.

```go
func (A *ZapKratos) GetHelper(msgCaption string) *log.Helper
```

Creates Kratos log.Helper with the given caption.

```go
func (A *ZapKratos) NewHelper(msgCaption string) *log.Helper
```

Same as GetHelper, creates the Kratos log.Helper.

### Options

Config options to ZapKratos.

```go
type Options struct {
    ModuleKeyName string // Module field key name in log output
}
```

#### Constructor

```go
func NewOptions() *Options
```

Creates Options with default settings (module field = "module").

#### Methods

```go
func (T *Options) WithModuleKeyName(moduleKeyName string) *Options
```

Sets custom module field name in a chainable fashion.

### LogImp

Implementation that implements the Kratos log.Logger using Zap.

```go
type LogImp struct {
    // Contains filtered or unexported fields
}
```

#### Constructor

```go
func NewLogImp(zapLog *zap.Logger, msgCaption string) log.Logger
```

Creates LogImp wrapping the given Zap instance.
## Dependencies

- `github.com/go-kratos/kratos/v2` - Kratos microservice framework
- `go.uber.org/zap` - Uber Zap structured logging
- `github.com/yyle88/zaplog` - Zap management package
- `github.com/yyle88/runpath` - Runtime path utilities
- `github.com/yyle88/erero` - Error handling utilities

## Related Projects

**Frameworks:**
- [Kratos](https://github.com/go-kratos/kratos) - Go microservice framework
- [Zap](https://github.com/uber-go/zap) - Uber's structured logging

**zapkratos Ecosystem:**
- [zapkratos](https://github.com/orzkratos/zapkratos) - This project
- [zapkratos-demos](https://github.com/orzkratos/zapkratos-demos) - Demo projects
  - [demo1kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo1kratos) - Basic integration
  - [demo2kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo2kratos) - Advanced usage

**zapzhkratos Ecosystem (Chinese):**
- [zapzhkratos](https://github.com/orzkratos/zapzhkratos) - Chinese version with Chinese function names
- [zapzhkratos-demos](https://github.com/orzkratos/zapzhkratos-demos) - Chinese version demos
  - [demo1kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo1kratos) - 基础集成
  - [demo2kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo2kratos) - 高级用法

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE) file

---

## 💬 Contact & Feedback

**Issues & Feedback:**

- 🐛 **Bug reports?** Open an issue and describe the problem with reproduction steps
- ✨ **Feature ideas?** Open an issue to discuss the implementation approach
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/zapkratos.git`).
3. **Navigate**: Navigate to the cloned project (`cd zapkratos`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/zapkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/zapkratos)
