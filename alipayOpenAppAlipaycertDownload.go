package alipay

// OpenAppAlipaycertDownloadRequest 应用支付宝公钥证书下载
func (a *AliClient) OpenAppAlipaycertDownloadRequest(requestParam OpenAppAlipaycertDownloadRequestParams) (
	responseParam OpenAppAlipaycertDownloadResponseParams, err error) {
	requestDataMap := make(map[string]interface{})
	requestDataMap["biz_content"] = a.SetDataToBizContent(requestParam, false)
	if err = a.HandlerRequest("POST", "alipay.open.app.alipaycert.download", false, requestDataMap, &responseParam); err != nil {
		return
	}
	return
}

// OpenAppAlipaycertDownloadRequestParams 应用支付宝公钥证书下载请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_9/alipay.open.app.alipaycert.download
type OpenAppAlipaycertDownloadRequestParams struct {
	AlipayCertSn string `json:"alipay_cert_sn"` // 支付宝公钥证书序列号
}

// OpenAppAlipaycertDownloadResponseParams 应用支付宝公钥证书下载响应参数
type OpenAppAlipaycertDownloadResponseParams struct {
	Data struct {
		CommonResParams
		AlipayCertContent string `json:"alipay_cert_content"` // 公钥证书Base64后的字符串
	} `json:"alipay_open_app_alipaycert_download_response"`
	Sign string `json:"sign"` // 签名
}
