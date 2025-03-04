package zapkratos

import (
	"path/filepath"

	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
)

type Options struct {
	ModuleKeyName string
}

func NewOptions() *Options {
	return &Options{
		ModuleKeyName: "module",
	}
}

func (T *Options) WithModuleKeyName(moduleKeyName string) *Options {
	T.ModuleKeyName = moduleKeyName
	return T
}

type ZapLog struct {
	zap *zaplog.Zap
	opt *Options
}

func NewZapLog(zap *zaplog.Zap, opt *Options) *ZapLog {
	return &ZapLog{
		zap: zap,
		opt: opt,
	}
}

func (A *ZapLog) GetZap() *zaplog.Zap {
	return A.zap
}

func (A *ZapLog) SubZap() *zaplog.Zap {
	return A.GetZap().SubZap2(A.opt.ModuleKeyName, filepath.Base(runpath.Skip(1)))
}
