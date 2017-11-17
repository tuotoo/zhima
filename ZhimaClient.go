package go_zmxy

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/delostik/go-zmxy/interfaces"
	"github.com/delostik/go-zmxy/utils"
)

type ZhimaClient struct {
	Appid string

	bizPrivateKey  []byte
	zhimaPublicKey []byte

	GatewayUrl string
	Format     string
	ApiVersion string
	Charset    string
	signType   string
	interfaces.ZhimaRequestParams
}

func (m *ZhimaClient) InitPemKey(bizPrivateKeyPath, zhimaPublicKeyPath string) error {
	var err error
	m.bizPrivateKey, err = utils.GetPemKeyFile(bizPrivateKeyPath)
	if err != nil {
		return err
	}
	m.zhimaPublicKey, err = utils.GetPemKeyFile(zhimaPublicKeyPath)
	if err != nil {
		return err
	}
	return nil
}

func (m *ZhimaClient) Execute(request interfaces.ZhimaRequest) (string, error) {
	bizParamStr := m.GetBizParamStr(&request)
	httpParams := m.GetSystemParams(&request)

	enBizParam, err := utils.EncryptRSA([]byte(bizParamStr), m.zhimaPublicKey)
	if err != nil {
		return "", err
	}
	(*httpParams)["params"] = base64.StdEncoding.EncodeToString(enBizParam)

	sign, err := utils.Sign(bizParamStr, m.bizPrivateKey)
	if err != nil {
		return "", err

	}
	(*httpParams)["sign"] = sign

	respData, err := utils.PostRequest(m.GatewayUrl, &request, httpParams)
	if err != nil {
		return "", nil
	}

	// Decrypt if needed
	if respData.Encrypted {
		respData.Data, err = utils.DecryptRSA(respData.Data, m.bizPrivateKey)
		if err != nil {
			return "", err
		}
	}

	// Check format (must be json)
	var temp map[string]interface{}
	err = json.Unmarshal([]byte(respData.Data), &temp)
	if err != nil {
		return "", err
	}

	// Verify signature
	if !utils.VerifySign(respData.Data, respData.Sign, m.zhimaPublicKey) {
		return "", errors.New("verify signature failed")
	}

	return respData.Data, nil
}

func (m *ZhimaClient) GetBizParamStr(request *interfaces.ZhimaRequest) string {
	return utils.BuildQueryString((*request).GetBizParams(), true)
}

func (m *ZhimaClient) GetSystemParams(request *interfaces.ZhimaRequest) *map[string]string {
	if m.Charset == "" {
		m.Charset = "UTF-8"
	}
	version := (*request).GetApiVersion()
	if version == "" {
		version = m.ApiVersion
	}
	sysParams := make(map[string]string)

	//TODO: make it struct

	sysParams["app_id"] = m.Appid
	sysParams["version"] = version
	sysParams["method"] = (*request).GetApiMethodName()
	sysParams["charset"] = m.Charset
	sysParams["scene"] = (*request).GetScene()
	sysParams["channel"] = (*request).GetChannel()
	sysParams["platform"] = (*request).GetPlatform()
	sysParams["ext_params"] = (*request).GetExtParams()
	return &sysParams
}

func NewZhimaClient(gatewayUrl, appId, charset, bizPrivateKeyPath, zhimaPublicKeyPath string) (*ZhimaClient, error) {
	client := &ZhimaClient{
		Appid:      appId,
		GatewayUrl: gatewayUrl,
		Format:     "json",
		ApiVersion: "1.0",
		Charset:    charset,
		signType:   "RSA",
	}
	err := client.InitPemKey(bizPrivateKeyPath, zhimaPublicKeyPath)
	if err != nil {
		return nil, err
	}
	return client, nil
}
