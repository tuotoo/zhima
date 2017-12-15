package zhima

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tuotoo/zhima/interfaces"
	"github.com/tuotoo/zhima/utils"
	"net/url"
)

type ZhimaClient struct {
	AppId string

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

func (m *ZhimaClient) GeneratePageRedirectInvokeUrl(request interfaces.ZhimaRequest) (string, error) {
	bizParamStr := m.GetBizParamStr(request)
	httpParams := m.GetSystemParams(request)

	enBizParam, err := utils.EncryptRSA([]byte(bizParamStr), m.zhimaPublicKey)
	if err != nil {
		return "", err
	}
	httpParams["params"] = base64.StdEncoding.EncodeToString(enBizParam)

	sign, err := utils.Sign(bizParamStr, m.bizPrivateKey)
	if err != nil {
		return "", err

	}
	httpParams["sign"] = sign
	urlvalue := url.Values{}
	for key, value := range httpParams {
		urlvalue.Add(key, value)
	}
	return fmt.Sprintf("%s?%s", m.GatewayUrl, urlvalue.Encode()), nil
}

func (m *ZhimaClient) Execute(request interfaces.ZhimaRequest) (string, error) {
	bizParamStr := m.GetBizParamStr(request)
	httpParams := m.GetSystemParams(request)

	enBizParam, err := utils.EncryptRSA([]byte(bizParamStr), m.zhimaPublicKey)
	if err != nil {
		return "", err
	}
	httpParams["params"] = base64.StdEncoding.EncodeToString(enBizParam)

	sign, err := utils.Sign(bizParamStr, m.bizPrivateKey)
	if err != nil {
		return "", err

	}
	httpParams["sign"] = sign

	respData, err := utils.PostRequest(m.GatewayUrl, &request, &httpParams)
	if err != nil {
		return "", err
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

func (m *ZhimaClient) GetBizParamStr(request interfaces.ZhimaRequest) string {
	return utils.BuildQueryString(request.GetApiParams(), true)
}

func (m *ZhimaClient) GetSystemParams(request interfaces.ZhimaRequest) map[string]string {
	if m.Charset == "" {
		m.Charset = "UTF-8"
	}
	version := request.GetApiVersion()
	if version == "" {
		version = m.ApiVersion
	}
	sysParams := make(map[string]string)

	//TODO: make it struct

	sysParams["app_id"] = m.AppId
	sysParams["version"] = version
	sysParams["method"] = request.GetApiMethodName()
	sysParams["charset"] = m.Charset
	sysParams["scene"] = request.GetScene()
	sysParams["channel"] = request.GetChannel()
	sysParams["platform"] = request.GetPlatform()
	sysParams["ext_params"] = request.GetExtParams()
	return sysParams
}


//解密验签
//1.先解密，再验签
//2.如果验签成功，则返回解密后的值
//3.如果验签失败，则抛出异常
//:param encryptedResponse: 返回值中的加密串
//:param sign: 返回值中的签名串
//:return str: 解密后的明文
func (m *ZhimaClient) DecryptAndVerifySign(encryptedResponse, sign string) (string,error) {
	pass := utils.VerifySign(encryptedResponse,sign,m.bizPrivateKey)
	if !pass {
		return "",errors.New("签名验证失败")
	}
	resp,err := utils.DecryptRSA(encryptedResponse,m.bizPrivateKey)
	return resp,err
}

func NewZhimaClient(gatewayUrl, appId, charset, bizPrivateKeyPath, zhimaPublicKeyPath string) (*ZhimaClient, error) {
	client := &ZhimaClient{
		AppId:      appId,
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
