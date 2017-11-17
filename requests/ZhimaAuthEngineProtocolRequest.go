package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaAuthEngineProtocolRequest struct {
	BizExtParams  string // 授权码入参
	IdentityParam string // 授权的身份标识参数
	IdentityType  string // 用户的身份标识类型\n3：身份证+姓名+手机号进行支付宝识别，核身，授权\n4：身份证+姓名+手机号（可选）进行公安网的核身，授权

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthEngineProtocolRequest) InitBizParams(bizExtParams, identityParam, identityType string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["biz_params"] = bizExtParams
	m.BizExtParams = bizExtParams

	m.BizParams["identity_param"] = identityParam
	m.IdentityParam = identityParam

	m.BizParams["identity_type"] = identityType
	m.IdentityType = identityType
}

func (m *ZhimaAuthEngineProtocolRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaAuthEngineProtocolRequest) GetApiMethodName() string {
	return "zhima.auth.engine.protocol"
}

func (m *ZhimaAuthEngineProtocolRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaAuthEngineProtocolRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthEngineProtocolRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthEngineProtocolRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthEngineProtocolRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthEngineProtocolRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaAuthEngineProtocolRequest) GetFileParams() map[string]string {
	return m.FileParams
}
