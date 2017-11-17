package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaAuthEngineOrganizationauthRequest struct {
	BizExtParams  string // 授权码入参，1). 若identity_Type=2时，{"auth_code":"M_P_COMPANY_CERT"} 2). 若identity_Type=5时，{"auth_code":"M_P_COMPANY_UID"}
	IdentityParam string // 证件号目前支持2种：a. 工商注册号：NATIONAL_LEGAL  b. 社会统一信用代码 : NATIONAL_LEGAL_MERGE1). 若identity_type=2时：identity_param={\"certNo\":\"440000400004160\",\"certType\":\"NATIONAL_LEGAL\",\"name\":\"中国移动通信集团广东有限公司\"}"2). 若identity_type=5时：identity_param={\"userId\":\"2088xxxxxx\"}"
	IdentityType  string // 用户身份标识类型：5： userId入参的类型标识；2： 证件号和姓名的入参的类型标识

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthEngineOrganizationauthRequest) InitBizParams(bizExtParams, identityParam, identityType string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["biz_params"] = bizExtParams
	m.BizExtParams = bizExtParams

	m.BizParams["identity_param"] = identityParam
	m.IdentityParam = identityParam

	m.BizParams["identity_type"] = identityType
	m.IdentityType = identityType
}

func (m *ZhimaAuthEngineOrganizationauthRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaAuthEngineOrganizationauthRequest) GetApiMethodName() string {
	return "zhima.auth.engine.organizationauth"
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaAuthEngineOrganizationauthRequest) GetFileParams() map[string]string {
	return m.FileParams
}
