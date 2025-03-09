package zapkratos

import (
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yyle88/erero"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogImp 是将 zap.Logger 转换为 kratos log.Logger 的适配器
type LogImp struct {
	zapLog *zap.Logger
	msgTop string
}

// NewLogImp 创建一个新的 LogImp 的对象实现 kratos log.Logger 的接口
func NewLogImp(zapLog *zap.Logger, msgTop string) log.Logger {
	return &LogImp{
		zapLog: zapLog,
		msgTop: msgTop,
	}
}

// Log 实现 kratos log.Logger 接口
func (a *LogImp) Log(logLevel log.Level, keyvals ...interface{}) error {
	var zapLevel zapcore.Level
	switch logLevel {
	case log.LevelDebug:
		zapLevel = zap.DebugLevel
	case log.LevelInfo:
		zapLevel = zap.InfoLevel
	case log.LevelWarn:
		zapLevel = zap.WarnLevel
	case log.LevelError:
		zapLevel = zap.ErrorLevel
	case log.LevelFatal:
		zapLevel = zap.DPanicLevel
	default:
		zapLevel = zap.DebugLevel //假如不确定级别就设置为最低级别的
	}

	// 使用 zap.Logger 记录日志
	zapInstance := a.zapLog.Check(zapLevel, a.msgTop)
	if zapInstance == nil {
		return erero.Errorf("wrong-log-level-param zap=%v arg=%v", zapLevel, logLevel)
	}

	var fields = make([]zap.Field, 0, (len(keyvals)+1)/2) //避免奇数
	for idx := 0; idx < len(keyvals); idx += 2 {
		if idx+1 < len(keyvals) {
			kes, ok := keyvals[idx].(string)
			if !ok {
				kes = newBK(idx) //因为日志是有可能要打印json的，因此这里的键还是组个正确的
			}
			fields = append(fields, zap.Any(kes, keyvals[idx+1]))
		}
	}
	if len(keyvals)%2 == 1 { //说明是奇数的，说明最后1个只有值没有键，而不是只有值没有键，因为键更难漏掉
		fields = append(fields, zap.Any(newBK(len(keyvals)-1), keyvals[(len(keyvals)-1)]))
	}

	zapInstance.Write(fields...)
	return nil
}

func newBK(idx int) string {
	return "BAD_KEY_" + strconv.Itoa(idx)
}
