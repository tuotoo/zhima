package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaCreditScoreIndustryGetRequest struct {
	OpenId        string
	ProductCode   string
	TransactionId string

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditScoreIndustryGetRequest) InitBizParams(openId, productCode, transactionId string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["open_id"] = openId
	m.OpenId = openId

	m.BizParams["product_code"] = productCode
	m.ProductCode = productCode

	m.BizParams["transaction_id"] = transactionId
	m.TransactionId = transactionId
}

func (m *ZhimaCreditScoreIndustryGetRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaCreditScoreIndustryGetRequest) GetApiMethodName() string {
	return "zhima.credit.score.industry.get"
}

func (m *ZhimaCreditScoreIndustryGetRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaCreditScoreIndustryGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditScoreIndustryGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditScoreIndustryGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditScoreIndustryGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditScoreIndustryGetRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaCreditScoreIndustryGetRequest) GetFileParams() map[string]string {
	return m.FileParams
}
