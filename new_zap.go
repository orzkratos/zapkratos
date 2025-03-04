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

type ZapKratos struct {
	zap *zaplog.Zap
	opt *Options
}

func NewZapKratos(zap *zaplog.Zap, opt *Options) *ZapKratos {
	return &ZapKratos{
		zap: zap,
		opt: opt,
	}
}

func (A *ZapKratos) GetZap() *zaplog.Zap {
	return A.zap
}

func (A *ZapKratos) SubZap() *zaplog.Zap {
	return A.GetZap().SubZap2(A.opt.ModuleKeyName, filepath.Base(runpath.Skip(1)))
}
