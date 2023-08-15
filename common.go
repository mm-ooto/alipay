package alipay

import (
	"encoding/json"
	"net/url"
)

// SystemOauthToken 换取授权访问令牌
func (a *Client) SystemOauthToken(requestParam SystemOauthTokenRequestParams) (
	responseParam SystemOauthTokenResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// AuthTokenApp 换取应用授权令牌
func (a *Client) AuthTokenApp(requestParam AuthTokenAppRequestParams) (
	responseParam AuthTokenAppResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// AsyncNotify 处理异步通知回调
// isLifeNotify 是否是生活号通知
func (a *Client) AsyncNotify(rawBody string, isLifeNotify ...struct{}) (notifyResult TradeNotificationParams, err error) {
	var urlValues url.Values
	// 调用url.ParseQuery来获取到参数列表，url.ParseQuery会自动完成url decode
	urlValues, err = url.ParseQuery(rawBody)
	if err != nil {
		return
	}

	isLifeIsNo := false
	if len(isLifeNotify) > 0 {
		isLifeIsNo = true
	}

	// 异步验签
	_, err = a.AsyncNotifyVerifySign(urlValues, isLifeIsNo)
	if err != nil {
		return
	}
	if err = json.Unmarshal([]byte(rawBody), &notifyResult); err != nil {
		return
	}

	return
}

// AppAliPayCertDownload 应用支付宝公钥证书下载
func (a *Client) AppAliPayCertDownload(requestParam AppAliPayCertDownloadRequestParams) (
	responseParam AppAliPayCertDownloadResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}
