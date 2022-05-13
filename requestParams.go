package alipay

import "net/url"

type RequestParams interface {
	// GetOtherParams 获取非公共参数,具体如下：returnUrl,notifyUrl,appAuthToken,apiMethodName,bizContent
	GetOtherParams() url.Values
	// GetNeedEncrypt 是否需要对biz_content内容加密，加密算法为AES
	GetNeedEncrypt() bool
}

// OtherRequestParams 其它特殊的请求参数
type OtherRequestParams struct {
	NeedEncrypt  bool   `json:"-"` // 是否需要对内容biz_content进行加密
	ReturnUrl    string `json:"-"`
	NotifyUrl    string `json:"-"` // 支付宝服务器主动通知商户服务器里指定的页面http/https路径，例如：http://api.test.alipay.net/atinterface/receive_notify.htm
	AppAuthToken string `json:"-"` // 详见应用授权概述：https://opendocs.alipay.com/isv/10467/xldcyq
}
