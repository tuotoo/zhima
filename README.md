# go-zmxy
zmxy(芝麻信用) SDK for golang

---
参照官方SDK进行设计，通过创建对应的RequestObject访问对应接口。

### 安装
```
go get http://github.com/tuotoo/zhima
```

### 使用
```
    // 初始化
    gateWay := "https://zmopenapi.zmxy.com.cn/openapi.do"
    client, err := zmxy.NewZhimaClient(gateWay, appId, charset, privKeyPath, pubKeyPath)
    // do something
    ...
    ...
    // 以查询芝麻分为例
    req := new(Requests.ZhimaCreditScoreGetRequest)
    req.InitBizParams(openId, productCode, transactionId)
    data, err := client.Execute(req)
    // do something
```


### 目前支持请求列表：
ZhimaCreditXuejiVerifyRequest  
ZhimaCustomerCertificationResultSyncRequest  
ZhimaCreditEpAssetSealGetRequest  
ZhimaCreditEpCloudatlasApplyRequest  
ZhimaCreditKkcreditAbcscoreQueryRequest  
ZhimaCreditKkcreditAbscoreQueryRequest  
ZhimaCreditEpAssetFreezeGetRequest  
ZhimaCreditDasGetRequest  
ZhimaCreditShareriskGetRequest  
ZhimaMerchantSceneCreditpayRiskevalQueryRequest  
ZhimaCreditKkcreditAcscoreQueryRequest  
ZhimaMerchantCloseloopDataUploadRequest  
ZhimaCreditEpOperationalreportGetRequest  
ZhimaCreditEpTaxOweGetRequest  
ZhimaMerchantSceneCreditpayBizriskevalQueryRequest  
ZhimaCreditEpTaxInvoiceQueryRequest  
ZhimaMerchantExpandApplyRequest  
ZhimaCustomerCertifyCheckRequest  
ZhimaMerchantSingleDataUploadRequest  
ZhimaCreditEpNegativeQueryRequest  
ZhimaMerchantCreditlifeWithholdCancelRequest  
ZhimaCreditEpContactNameGetRequest  
ZhimaMerchantEnterpriseQueryRequest  
ZhimaCreditCollectionSupportRequest  
ZhimaProfileStatusFeedbackRequest  
ZhimaCustomerCertificationInitializeRequest  
ZhimaCreditMobileRainGetRequest  
ZhimaCreditEpWatchlistGetRequest  
ZhimaCreditMsxfOnlinejdscoreQueryRequest  
ZhimaMerchantCreditlifePreauthCancelRequest  
ZhimaCreditWatchlistBriefGetRequest  
ZhimaCustomerBorrowRestoreRequest  
ZhimaCreditVehicleVerifyRequest  
ZhimaCreditAntifraudVerifyRequest  
ZhimaCreditContactGetRequest  
SsdataFindataCommonJumpurlQueryRequest  
ZhimaCreditIvsDetailGetRequest  
ZhimaOpenLogFeedbackRequest  
ZhimaMerchantExpandQueryRequest  
ZhimaCustomerCertificationQueryRequest  
ZhimaCreditEpLawsuitRecordGetRequest  
ZhimaCreditDriverVerifyRequest  
ZhimaMerchantEnterpriseApplyRequest  
ZhimaDataBatchFeedbackRequest  
ZhimaAuthFaceVerifyRequest  
ZhimaCreditPeLawsuitRecordGetRequest  
ZhimaCreditCardLimitGetRequest  
ZhimaCreditCardPermitRequest  
ZhimaCustomerContractInitializeRequest  
ZhimaAuthEngineProtocolRequest  
ZhimaCreditEpCloudatlasGetRequest  
ZhimaMerchantSceneCreditpaySignriskevalQueryRequest  
ZhimaMerchantCreditlifeWithholdQueryRequest  
ZhimaAuthFaceInitRequest  
ZhimaCreditEpCommercialPunishmentQueryRequest  
ZhimaCreditScoreGetRequest  
ZhimaCreditEpQualityPunishmentQueryRequest  
ZhimaCustomerCertifyInitialRequest  
ZhimaCreditPassinfoGetRequest  
ZhimaCreditTrueidentityGetRequest  
ZhimaCreditAntifraudRiskListRequest  
ZhimaCreditXueliGetRequest  
ZhimaCreditModulescoreQueryRequest  
ZhimaAuthInfoAuthorizeRequest  
ZhimaCustomerCertificationCertifyRequest  
ZhimaCreditCardVerifyRequest  
ZhimaCreditScoreIndustryGetRequest  
ZhimaCreditPeLawsuitDetailGetRequest  
ZhimaMerchantCreditlifeFundRefundRequest  
ZhimaCreditScoreBriefGetRequest  
ZhimaCreditEpLawsuitDetailGetRequest  
ZhimaMerchantImageUploadRequest  
ZhimaCreditPeLawsuitDetailQueryRequest  
ZhimaCreditEpRoleGetRequest  
ZhimaCustomerPerformanceFeedbackRequest  
ZhimaDataSingleFeedbackRequest  
ZhimaCreditIvsGetRequest  
ZhimaCreditXueliVerifyRequest  
ZhimaMerchantDataUploadInitializeRequest  
ZhimaAuthInfoFreezeRequest  
ZhimaCreditSkynetriskGetRequest  
ZhimaCreditContactAnalyzeQueryRequest  
ZhimaCustomerCreditrecordSummaryQueryRequest  
ZhimaCreditReportFinanceGetRequest  
ZhimaCreditEpIndividualScoreGetRequest  
ZhimaAuthInfoAuthqueryRequest  
ZhimaAuthZhimaCustCertifyInitialRequest  
ZhimaCreditEpRiskIndexGetRequest  
ZhimaAuthEngineOrganizationauthRequest  
ZhimaCustomerCertifyApplyRequest  
ZhimaCreditEpEnvprotectionPunishmentQueryRequest  
ZhimaCreditEpContactPhoneGetRequest  
ZhimaCustomerEpCertificationInitializeRequest  
ZhimaMerchantCreditlifeFundPayRequest  
ZhimaDataFeedbackurlQueryRequest  
ZhimaCustomerCertificationMaterialCertifyRequest  
ZhimaCustomerCertificationInfoQueryRequest  
ZhimaMerchantCreditlifeRiskApplyRequest  
ZhimaCreditEpCaseSeriesQueryRequest  
ZhimaCreditEpScoreGetRequest  
ZhimaAuthEngineSmsauthRequest  
ZhimaCreditEpInfoGetRequest  
ZhimaCreditBqsDefaultscoreQueryRequest  
ZhimaCreditEpBriefScoreQueryRequest  
ZhimaCreditReportVisaGetRequest  
ZhimaCustomerEpCertificationCertifyRequest  
ZhimaMerchantCreditlifePreauthUnfreezeRequest  
ZhimaCreditAntifraudIntegrationQueryRequest  
ZhimaCreditRiskEvaluateRuleQueryRequest  
ZhimaCreditWatchlistiiGetRequest  
ZhimaAuthEngineMultiauthRequest  
ZhimaCreditAntifraudScoreGetRequest  
ZhimaCreditHetroneDasscoreQueryRequest  
ZhimaCustomerBorrowScanRequest  
ZhimaCreditScoreCobuildGetRequest  
ZhimaMerchantOrderConfirmRequest  
ZhimaCreditRiskEvaluateQueryRequest  
ZhimaCustomerEpCertificationQueryRequest  
