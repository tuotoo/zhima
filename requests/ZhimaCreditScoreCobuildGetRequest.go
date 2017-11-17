package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaCreditScoreCobuildGetRequest struct {
	ApplyDate     string // 申请日期，可空
	BizExtParams  string // 业务拓展参数
	Mobile        string // 手机号，可空
	OpenId        string
	ProductCode   string
	TransactionId string

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditScoreCobuildGetRequest) InitBizParams(applyDate, bizExtParams, mobile, openId, productCode, transactionId string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["apply_date"] = applyDate
	m.ApplyDate = applyDate

	m.BizParams["biz_params"] = bizExtParams
	m.BizExtParams = bizExtParams

	m.BizParams["mobile"] = mobile
	m.Mobile = mobile

	m.BizParams["open_id"] = openId
	m.OpenId = openId

	m.BizParams["product_code"] = productCode
	m.ProductCode = productCode

	m.BizParams["transaction_id"] = transactionId
	m.TransactionId = transactionId
}

func (m *ZhimaCreditScoreCobuildGetRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaCreditScoreCobuildGetRequest) GetApiMethodName() string {
	return "zhima.credit.score.cobuild.get"
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaCreditScoreCobuildGetRequest) GetFileParams() map[string]string {
	return m.FileParams
}
