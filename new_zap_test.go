package zapkratos

import (
	"testing"

	"github.com/yyle88/zaplog"
)

func TestNewZapKratos(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions())

	subLog := zapKratos.SubZap()

	subLog.LOG.Info("abc")
	subLog.SUG.Info("xyz")
}

func TestOptions_WithModuleKeyName(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions().WithModuleKeyName("module"))

	subLog := zapKratos.SubZap()

	subLog.LOG.Info("abc")
	subLog.SUG.Info("xyz")
}
