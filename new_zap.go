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
	zap     *zaplog.Zap
	options *Options
}

func NewZapKratos(zap *zaplog.Zap, options *Options) *ZapKratos {
	return &ZapKratos{
		zap:     zap,
		options: options,
	}
}

func (A *ZapKratos) GetZap() *zaplog.Zap {
	return A.zap
}

func (A *ZapKratos) SubZap() *zaplog.Zap {
	return A.GetZap().SubZap2(A.options.ModuleKeyName, filepath.Base(runpath.Skip(1)))
}
