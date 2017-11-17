package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaAuthFaceInitRequest struct {
	ApiKey   string // 应用的标识
	AuthMsg  string // 参数的加密串
	BizType  string // 用于标识使用人离岸业务的类型是授权或者认证，默认为授权类型
	BundleId string // 不同商户的bundle_id,用来做缓存
	Token    string // 请求的sessionId

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthFaceInitRequest) InitBizParams(apiKey, authMsg, bizType, bundleId, token string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["api_key"] = apiKey
	m.ApiKey = apiKey

	m.BizParams["auth_msg"] = authMsg
	m.AuthMsg = authMsg

	m.BizParams["biz_type"] = bizType
	m.BizType = bizType

	m.BizParams["bundle_id"] = bundleId
	m.BundleId = bundleId

	m.BizParams["token"] = token
	m.Token = token
}

func (m *ZhimaAuthFaceInitRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaAuthFaceInitRequest) GetApiMethodName() string {
	return "zhima.auth.face.init"
}

func (m *ZhimaAuthFaceInitRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaAuthFaceInitRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthFaceInitRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthFaceInitRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthFaceInitRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthFaceInitRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaAuthFaceInitRequest) GetFileParams() map[string]string {
	return m.FileParams
}
