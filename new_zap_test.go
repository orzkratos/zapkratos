package zapkratos

import (
	"testing"

	"github.com/yyle88/zaplog"
)

func TestNewZapLog(t *testing.T) {
	zapOpt := NewZapOpt(zaplog.LOGGER, NewOptions())

	subLog := zapOpt.SubZap()

	subLog.LOG.Info("abc")
	subLog.SUG.Info("xyz")
}

func TestNewZapLog_WithModuleKeyName(t *testing.T) {
	zapOpt := NewZapOpt(zaplog.LOGGER, NewOptions().WithModuleKeyName("module"))

	subLog := zapOpt.SubZap()

	subLog.LOG.Info("abc")
	subLog.SUG.Info("xyz")
}
