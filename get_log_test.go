package zapkratos

import (
	"testing"

	"github.com/yyle88/zaplog"
)

func TestZapOpt_GetHelper(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions())

	helper := zapKratos.GetHelper("test-get-helper")

	helper.Info("woca", "[a]", "[b]", "[c]")
	helper.Infow("k", "v", "k1", "v2")
}

func TestZapOpt_NewHelper(t *testing.T) {
	zapKratos := NewZapKratos(zaplog.LOGGER, NewOptions())

	helper := zapKratos.NewHelper("test-new-helper")

	helper.Info("test-message")
	helper.Infow("test", "message")
}
