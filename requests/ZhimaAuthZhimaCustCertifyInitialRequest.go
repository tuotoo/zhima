package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaAuthZhimaCustCertifyInitialRequest struct {
	BizExtParams  string // 业务扩展属性：来源类型source_type必传，比如：1.web场景中，传入{“source_type”:"pc"}2.移动端场景中，传入{“source_type”:"h5"}
	ContractFlag  string // 与芝麻信用签订的合约外标，即使合约改签或续签该值不会发生变化
	IdentityParam string // 不同身份类型的参数列表，json字符串的key-value格式：如：identity_type= "CERT_AND_MOBILE";identity_param={"certNo":"xxx", "name":"张三", "certType":"IDENTITY_CARD"};
	IdentityType  string // 身份标识类型（后续可以扩展）：BY_CERTNO_AND_NAME:按照身份证+姓名（+手机号）进行授权
	ProductCode   string // 当前使用的产品码
	State         string // 芝麻认证过程中的冗余字段，在认证申请时传入，在结果页面回调中原样透传给商户端。【建议使用方式】用于商户端唯一标记发起认证的用户信息，在接收到芝麻信用认证结果回调后确定用户
	TransactionId string // 商户传入的业务流水号。此字段由商户生成，需确保唯一性，用于定位每一次请求，后续按此流水进行对帐。生成规则: 固定30位数字串，前17位为精确到毫秒的时间

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthZhimaCustCertifyInitialRequest) InitBizParams(bizExtParams, contractFlag, identityParam, identityType, productCode, state, transactionId string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["biz_params"] = bizExtParams
	m.BizExtParams = bizExtParams

	m.BizParams["contract_flag"] = contractFlag
	m.ContractFlag = contractFlag

	m.BizParams["identity_param"] = identityParam
	m.IdentityParam = identityParam

	m.BizParams["identity_type"] = identityType
	m.IdentityType = identityType

	m.BizParams["product_code"] = productCode
	m.ProductCode = productCode

	m.BizParams["state"] = state
	m.State = state

	m.BizParams["transaction_id"] = transactionId
	m.TransactionId = transactionId
}

func (m *ZhimaAuthZhimaCustCertifyInitialRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaAuthZhimaCustCertifyInitialRequest) GetApiMethodName() string {
	return "zhima.auth.zhima.cust.certify.initial"
}

func (m *ZhimaAuthZhimaCustCertifyInitialRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaAuthZhimaCustCertifyInitialRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthZhimaCustCertifyInitialRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthZhimaCustCertifyInitialRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthZhimaCustCertifyInitialRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthZhimaCustCertifyInitialRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaAuthZhimaCustCertifyInitialRequest) GetFileParams() map[string]string {
	return m.FileParams
}
