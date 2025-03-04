package zapkratos

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yyle88/zaplog"
)

func TestNewLogImp(t *testing.T) {
	logImp := NewLogImp(zaplog.LOG, "test-new-log-imp")

	helper := log.NewHelper(logImp)

	helper.Infow("k", "v", "k1", "v2")
	helper.Infow("k", "v", "v2")
}
