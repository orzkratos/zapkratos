package zapkratos

import "github.com/go-kratos/kratos/v2/log"

func (A *ZapKratos) GetLogger(msgTop string) log.Logger {
	return NewLogImp(A.GetZap().LOG, msgTop)
}

func (A *ZapKratos) NewLogger(msgTop string) log.Logger {
	return NewLogImp(A.GetZap().LOG, msgTop)
}

func (A *ZapKratos) GetHelper(msgTop string) *log.Helper {
	return log.NewHelper(A.GetLogger(msgTop))
}

func (A *ZapKratos) NewHelper(msgTop string) *log.Helper {
	return log.NewHelper(A.GetLogger(msgTop))
}
