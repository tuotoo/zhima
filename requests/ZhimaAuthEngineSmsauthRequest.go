package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaAuthEngineSmsauthRequest struct {
	BizExtParams  string // 业务扩展字段,页面授权接口需要传入auth_code,channelType,stateauth_code授权码,短信唤起支付宝客户端接口的值为M_SMSchannelType渠道类型,每个授权码支持不同的渠道类型appsdk:sdk接入apppc:商户pc页面接入api:后台api接入windows:支付宝服务窗接入app:商户app接入state是商户自定义的数据,页面授权接口会原样把这个数据返回个商户
	IdentityParam string // 身份参数,短信唤起支付宝客户端授权需要传入姓名+证件类型+证件号码+手机号
	IdentityType  string // 身份类型,短信唤起支付宝客户端接口的身份识别类型为2:按照姓名+证件类型+证件号码+手机号进行授权

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthEngineSmsauthRequest) InitBizParams(bizExtParams, identityParam, identityType string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["biz_params"] = bizExtParams
	m.BizExtParams = bizExtParams

	m.BizParams["identity_param"] = identityParam
	m.IdentityParam = identityParam

	m.BizParams["identity_type"] = identityType
	m.IdentityType = identityType
}

func (m *ZhimaAuthEngineSmsauthRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaAuthEngineSmsauthRequest) GetApiMethodName() string {
	return "zhima.auth.engine.smsauth"
}

func (m *ZhimaAuthEngineSmsauthRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaAuthEngineSmsauthRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthEngineSmsauthRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthEngineSmsauthRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthEngineSmsauthRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthEngineSmsauthRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaAuthEngineSmsauthRequest) GetFileParams() map[string]string {
	return m.FileParams
}
