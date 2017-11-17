# go-zmxy
zmxy(芝麻信用) SDK for golang

---
参照官方SDK进行设计，通过创建对应的RequestObject访问对应接口。

### 安装
```
go get http://github.com/delostik/go-zmxy
```

### 使用
```
    // 初始化
    gateWay := "https://zmopenapi.zmxy.com.cn/sandbox.do"
    client, err := zmxy.NewZhimaClient(gateWay, appId, charset, privKeyPath, pubKeyPath)
    // do something
    ...
    ...
    // 以查询芝麻分为例
    req := new(requests.ZhimaCreditScoreGetRequest)
    req.InitBizParams(openId, productCode, transactionId)
    data, err := client.Execute(req)
    // do something
```


### 目前支持请求列表：
芝麻分  
ZhimaCreditScoreGetRequest  
ZhimaCreditScoreBriefGetRequest  
ZhimaCreditScoreCobuildGetRequest  
ZhimaCreditScoreIndustryGetRequest  
  
认证和授权  
ZhimaAuthEngineMultiauthRequest  
ZhimaAuthEngineOrganizationauthRequest  
ZhimaAuthEngineProtocolRequest  
ZhimaAuthEngineSmsauthRequest  
ZhimaAuthFaceInitRequest  
ZhimaAuthFaceVerifyRequest  
ZhimaAuthInfoAuthorizeRequest  
ZhimaAuthInfoAuthqueryRequest  
ZhimaAuthInfoFreezeRequest  
ZhimaAuthZhimaCustCertifyInitialRequest  
  
SsdataFindataCommonJumpurlQueryRequest  
