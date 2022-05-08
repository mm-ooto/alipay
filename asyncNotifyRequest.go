package alipay

import (
	"encoding/json"
	"net/url"
)

// HandlerAsyncNotify 处理异步通知回调
// isLifeIsNo 是否是生活号通知
func (a *AliClient) HandlerAsyncNotify(rawBody string, isLifeIsNo bool) (notifyResult TradeNotificationParams, err error) {
	var urlValues url.Values
	// 调用url.ParseQuery来获取到参数列表，url.ParseQuery会自动完成url decode
	urlValues, err = url.ParseQuery(rawBody)
	if err != nil {
		return
	}

	// 异步验签
	_, err = a.AsyncNotifyVerifySign(urlValues, isLifeIsNo)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(rawBody), &notifyResult)
	return
}
