package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaAuthInfoAuthqueryRequest struct {
	AuthCategory  string // 授权类型，用于识别当前查询是否授权的分流；可传参数“B2B”或“C2B”，自助商户请填写“C2B”。
	IdentityParam string // 不同身份类型传入的参数列表,json字符串的key-value格式身份类型identityType=1:{"mobileNo":"15158657683"}身份类型identityType=2:{"certNo":"330100xxxxxxxxxxxx","name":"张三","certType":"IDENTITY_CARD"}
	IdentityType  string // 身份标识类型0:按照openId查询2:按照身份证+姓名查询

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthInfoAuthqueryRequest) InitBizParams(authCategory, identityParam, identityType string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["auth_category"] = authCategory
	m.AuthCategory = authCategory

	m.BizParams["identity_param"] = identityParam
	m.IdentityParam = identityParam

	m.BizParams["identity_type"] = identityType
	m.IdentityType = identityType
}

func (m *ZhimaAuthInfoAuthqueryRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaAuthInfoAuthqueryRequest) GetApiMethodName() string {
	return "zhima.auth.info.authquery"
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaAuthInfoAuthqueryRequest) GetFileParams() map[string]string {
	return m.FileParams
}
