package interfaces

type ZhimaRequest interface {
	GetApiMethodName() string
	GetBizParams() map[string]string

	GetApiVersion() string
	GetScene() string
	GetChannel() string
	GetPlatform() string
	GetExtParams() string
	GetFileParams() map[string]string
}

type ZhimaRequestParams struct {
	BizParams  map[string]string
	FileParams map[string]string
	ApiVersion string
	Scene      string
	Channel    string
	Platform   string
	ExtParams  string
}
