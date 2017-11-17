package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaAuthEngineMultiauthRequest struct {
	AuthAppId      string // 外部商户应用ID
	AuthMerchantId string // 外部商户ID
	UserId         string // 支付宝用户ID

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthEngineMultiauthRequest) InitBizParams(authAppId, authMerchantId, userId string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["auth_app_id"] = authAppId
	m.AuthAppId = authAppId

	m.BizParams["auth_merchant_id"] = authMerchantId
	m.AuthMerchantId = authMerchantId

	m.BizParams["user_id"] = userId
	m.UserId = userId
}

func (m *ZhimaAuthEngineMultiauthRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaAuthEngineMultiauthRequest) GetApiMethodName() string {
	return "zhima.auth.engine.multiauth"
}

func (m *ZhimaAuthEngineMultiauthRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaAuthEngineMultiauthRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthEngineMultiauthRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthEngineMultiauthRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthEngineMultiauthRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthEngineMultiauthRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaAuthEngineMultiauthRequest) GetFileParams() map[string]string {
	return m.FileParams
}
