package zapkratos

import (
	"testing"

	"github.com/yyle88/zaplog"
)

func TestNewZapLog(t *testing.T) {
	zapLog := NewZapLog(zaplog.LOGGER, NewOptions())

	subLog := zapLog.SubZap()

	subLog.LOG.Info("abc")
	subLog.SUG.Info("xyz")
}

func TestNewZapLog_WithModuleKeyName(t *testing.T) {
	zapLog := NewZapLog(zaplog.LOGGER, NewOptions().WithModuleKeyName("module"))

	subLog := zapLog.SubZap()

	subLog.LOG.Info("abc")
	subLog.SUG.Info("xyz")
}
