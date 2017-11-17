package requests

import (
	"github.com/delostik/go-zmxy/interfaces"
)

type ZhimaAuthInfoFreezeRequest struct {
	BizNo        string //用户在商户处申请业务的唯一识别码;\n每个申请仅对应一条冻结记录，若存在相同流水号的冻结周期将覆盖
	BizExtParams string // 扩展字段，json字符串的key-value格式
	BizType      string // 申请原因001: 贷款申请中, 002:信用卡申请中, 003:租车申请中, 004:信贷服务期内, 005:信贷逾期中
	EndDate      string //冻结结束时间，若该时间减去冻结开始时间的差值大于冻结周期，则该时间将赋值冻结开始时间+冻结周期
	OpenId       string // 用户在商端的身份标识
	StartDate    string // 冻结开始时间，若该时间早于系统当前时间，则会取当前时间作为冻结开始时间

	interfaces.ZhimaRequestParams
}

func (m *ZhimaAuthInfoFreezeRequest) InitBizParams(bizNo, bizExtParams, bizType, endDate, openId, startDate string) {
	m.BizParams = make(map[string]string)
	m.FileParams = make(map[string]string)

	m.BizParams["biz_no"] = bizNo
	m.BizNo = bizNo

	m.BizParams["biz_params"] = bizExtParams
	m.BizExtParams = bizExtParams

	m.BizParams["biz_type"] = bizType
	m.BizType = bizType

	m.BizParams["end_date"] = endDate
	m.EndDate = endDate

	m.BizParams["open_id"] = openId
	m.OpenId = openId

	m.BizParams["start_date"] = startDate
	m.StartDate = startDate
}

func (m *ZhimaAuthInfoFreezeRequest) init() {
	m.BizParams = make(map[string]string)
}

func (*ZhimaAuthInfoFreezeRequest) GetApiMethodName() string {
	return "zhima.auth.info.freeze"
}

func (m *ZhimaAuthInfoFreezeRequest) GetBizParams() map[string]string {
	return m.BizParams
}

func (m *ZhimaAuthInfoFreezeRequest) GetApiVersion() string {
	return m.ApiVersion
}

func (m *ZhimaAuthInfoFreezeRequest) GetScene() string {
	return m.Scene
}

func (m *ZhimaAuthInfoFreezeRequest) GetChannel() string {
	return m.Channel
}

func (m *ZhimaAuthInfoFreezeRequest) GetPlatform() string {
	return m.Platform
}

func (m *ZhimaAuthInfoFreezeRequest) GetExtParams() string {
	return m.ExtParams
}

func (m *ZhimaAuthInfoFreezeRequest) GetFileParams() map[string]string {
	return m.FileParams
}
