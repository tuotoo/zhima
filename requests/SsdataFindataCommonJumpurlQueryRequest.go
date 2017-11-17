package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type SsdataFindataCommonJumpurlQueryRequest struct {
	BizExtParams string
	BizNo        string

	interfaces.ZhimaRequestParams
}

func (m *SsdataFindataCommonJumpurlQueryRequest) InitBizParams(bizExtParams, bizNo string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["biz_ext_params"] = bizExtParams
	m.BizExtParams = bizExtParams

	m.BizParams["biz_no"] = bizNo
	m.BizNo = bizNo
}

func (m *SsdataFindataCommonJumpurlQueryRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*SsdataFindataCommonJumpurlQueryRequest) GetApiMethodName() string {
	return "ssdata.findata.common.jumpurl.query request"
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetScene() string {
	return m.Scene
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetChannel() string {
	return m.Channel
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetPlatform() string {
	return m.Platform
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *SsdataFindataCommonJumpurlQueryRequest) GetFileParams() map[string]string {
	return m.FileParams
}
