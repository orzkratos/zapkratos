package zapkratos

import (
	"testing"

	"github.com/yyle88/zaplog"
)

func TestZapOpt_GetHelper(t *testing.T) {
	zapOpt := NewZapOpt(zaplog.LOGGER, NewOptions())

	helper := zapOpt.GetHelper("test-get-helper")

	helper.Info("woca", "[a]", "[b]", "[c]")
	helper.Infow("k", "v", "k1", "v2")
}

func TestZapOpt_NewHelper(t *testing.T) {
	zapOpt := NewZapOpt(zaplog.LOGGER, NewOptions())

	helper := zapOpt.NewHelper("test-new-helper")

	helper.Info("test-message")
	helper.Infow("test", "message")
}
