package zapkratos

import "github.com/go-kratos/kratos/v2/log"

func (A *ZapOpt) GetLogger(msgTop string) log.Logger {
	return NewLogImp(A.GetZap().LOG, msgTop)
}

func (A *ZapOpt) NewLogger(msgTop string) log.Logger {
	return NewLogImp(A.GetZap().LOG, msgTop)
}

func (A *ZapOpt) GetHelper(msgTop string) *log.Helper {
	return log.NewHelper(A.GetLogger(msgTop))
}

func (A *ZapOpt) NewHelper(msgTop string) *log.Helper {
	return log.NewHelper(A.GetLogger(msgTop))
}
