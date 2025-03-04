package zapkratos

import "github.com/go-kratos/kratos/v2/log"

func (A *ZapLog) GetLogger(msgTop string) log.Logger {
	return NewLogImp(A.GetZap().LOG, msgTop)
}

func (A *ZapLog) NewLogger(msgTop string) log.Logger {
	return NewLogImp(A.GetZap().LOG, msgTop)
}

func (A *ZapLog) GetHelper(msgTop string) *log.Helper {
	return log.NewHelper(A.GetLogger(msgTop))
}

func (A *ZapLog) NewHelper(msgTop string) *log.Helper {
	return log.NewHelper(A.GetLogger(msgTop))
}
