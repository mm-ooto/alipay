package alipay

// TradeCancel 统一收单交易撤销接口
func (a *AliClient) TradeCancel(requestParam TradeCancelRequestParams) (responseParam TradeCancelResponseParams, err error) {
	requestDataMap := make(map[string]interface{})
	requestDataMap["biz_content"] = a.SetDataToBizContent(requestParam, requestParam.NeedEncrypt)
	requestDataMap["notify_url"] = requestParam.NotifyUrl
	requestDataMap["app_auth_token"] = requestParam.AppAuthToken
	if err = a.HandlerRequest("POST", "alipay.trade.cancel", requestParam.NeedEncrypt, requestDataMap, &responseParam); err != nil {
		return
	}
	return
}

// TradeCancelRequestParams 统一收单交易撤销接口请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.cancel
type TradeCancelRequestParams struct {
	OtherRequestParams

	OutTradeNo string `json:"out_trade_no,omitempty"` // 原支付请求的商户订单号,和支付宝交易号不能同时为空
	TradeNo    string `json:"trade_no,omitempty"`     // 支付宝交易号，和商户订单号不能同时为空
}

// TradeCancelResponseParams 统一收单交易撤销接口响应参数
type TradeCancelResponseParams struct {
	Data struct {
		CommonResParams
		TradeNo    string `json:"trade_no,omitempty"`     // 支付宝交易号; 当发生交易关闭或交易退款时返回；
		OutTradeNo string `json:"out_trade_no,omitempty"` // 商户订单号
		RetryFlag  string `json:"retry_flag"`             // 是否需要重试
		Action     string `json:"action"`                 // 本次撤销触发的交易动作,接口调用成功且交易存在时返回。可能的返回值：close：交易未支付，触发关闭交易动作，无退款；refund：交易已支付，触发交易退款动作； 未返回：未查询到交易，或接口调用失败；
	} `json:"alipay_trade_cancel_response"`
	Sign string `json:"sign"` // 签名
}
