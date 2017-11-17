package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaCreditScoreBriefGetRequest struct {
	AdmittanceScore string
	CertNo          string
	CertType        string // IDENTITY_CARD, PASSPORT, ALIPAY_USER_ID
	Name            string
	ProductCode     string
	TransactionId   string

	interfaces.ZhimaRequestParams
}

func (m *ZhimaCreditScoreBriefGetRequest) InitBizParams(admittanceScore, certNo, certType, name, productCode, transactionId string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["admittance_score"] = admittanceScore
	m.AdmittanceScore = admittanceScore

	m.BizParams["cert_no"] = certNo
	m.CertNo = certNo

	m.BizParams["cert_type"] = certType
	m.CertNo = certType

	m.BizParams["name"] = name
	m.Name = name

	m.BizParams["product_code"] = productCode
	m.ProductCode = productCode

	m.BizParams["transaction_id"] = transactionId
	m.TransactionId = transactionId
}

func (m *ZhimaCreditScoreBriefGetRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaCreditScoreBriefGetRequest) GetApiMethodName() string {
	return "zhima.credit.score.brief.get"
}

func (m *ZhimaCreditScoreBriefGetRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaCreditScoreBriefGetRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaCreditScoreBriefGetRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaCreditScoreBriefGetRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaCreditScoreBriefGetRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaCreditScoreBriefGetRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaCreditScoreBriefGetRequest) GetFileParams() map[string]string {
	return m.FileParams
}
