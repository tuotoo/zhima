package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaCreditScoreGetRequest struct {
	OpenId        string
	ProductCode   string
	TransactionId string

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditScoreGetRequest) InitBizParams(openId, productCode, transactionId string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)
	m.BizParams["open_id"] = openId
	m.OpenId = openId
	m.BizParams["product_code"] = productCode
	m.ProductCode = productCode
	m.BizParams["transaction_id"] = transactionId
	m.TransactionId = transactionId
}

func (m *ZhimaCreditScoreGetRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaCreditScoreGetRequest) GetApiMethodName() string {
	return "zhima.credit.score.get"
}

func (m *ZhimaCreditScoreGetRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaCreditScoreGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditScoreGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditScoreGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditScoreGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditScoreGetRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaCreditScoreGetRequest) GetFileParams() map[string]string {
	return m.FileParams
}
