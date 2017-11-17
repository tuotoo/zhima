package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaAuthFaceVerifyRequest struct {
	BizType string // 活体认证类型，目前有认证和授权两种类型；默认为授权类型
	Images  string
	Token   string // 标识一次请求流水

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthFaceVerifyRequest) InitBizParams(bizType, images, token string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["biz_type"] = bizType
	m.BizType = bizType

	m.BizParams["images"] = images
	m.Images = images

	m.BizParams["token"] = token
	m.Token = token
}

func (m *ZhimaAuthFaceVerifyRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaAuthFaceVerifyRequest) GetApiMethodName() string {
	return "zhima.auth.face.verify"
}

func (m *ZhimaAuthFaceVerifyRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaAuthFaceVerifyRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthFaceVerifyRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthFaceVerifyRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthFaceVerifyRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthFaceVerifyRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaAuthFaceVerifyRequest) GetFileParams() map[string]string {
	return m.FileParams
}
