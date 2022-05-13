package alipay

import (
	"encoding/json"
	"github.com/mm-ooto/alipay/consts"
	"net/url"
)

// OpenAppAlipaycertDownloadRequest 应用支付宝公钥证书下载
func (a *AliClient) OpenAppAlipaycertDownloadRequest(requestParam OpenAppAlipaycertDownloadRequestParams) (
	responseParam OpenAppAlipaycertDownloadResponseParams, err error) {
	//requestDataMap := make(map[string]interface{})
	//requestDataMap["biz_content"] = a.SetDataToBizContent(requestParam, false)
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// OpenAppAlipaycertDownloadRequestParams 应用支付宝公钥证书下载请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_9/alipay.open.app.alipaycert.download
type OpenAppAlipaycertDownloadRequestParams struct {
	OtherRequestParams
	AlipayCertSn string `json:"alipay_cert_sn"` // 支付宝公钥证书序列号
}

func (o *OpenAppAlipaycertDownloadRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(consts.NotifyUrlFiled, o.NotifyUrl)
	urlValue.Add(consts.AppAuthTokenFiled, o.AppAuthToken)
	urlValue.Add(consts.ApiMethodNameFiled, "alipay.open.app.alipaycert.download")
	bytes, _ := json.Marshal(o)
	urlValue.Add(consts.BizContentFiled, string(bytes))
	return urlValue
}

func (o *OpenAppAlipaycertDownloadRequestParams) GetNeedEncrypt() bool {
	return o.NeedEncrypt == true
}

// OpenAppAlipaycertDownloadResponseParams 应用支付宝公钥证书下载响应参数
type OpenAppAlipaycertDownloadResponseParams struct {
	Data struct {
		CommonResParams
		AlipayCertContent string `json:"alipay_cert_content"` // 公钥证书Base64后的字符串
	} `json:"alipay_open_app_alipaycert_download_response"`
	Sign string `json:"sign"` // 签名
}
