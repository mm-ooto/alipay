package alipay

import "net/url"

// TradeCreate 统一收单交易创建接口
func (a *Client) TradeCreate(requestParam TradeCreateRequestParams) (responseParam TradeCreateResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeCancel 统一收单交易撤销接口
func (a *Client) TradeCancel(requestParam TradeCancelRequestParams) (responseParam TradeCancelResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeClose 统一收单交易关闭接口
func (a *Client) TradeClose(requestParam TradeCloseRequestParams) (responseParam TradeCloseResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeQuery 统一收单线下交易查询
func (a *Client) TradeQuery(requestParam TradeQueryRequestParams) (responseParam TradeQueryResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradePreCreate 统一收单线下交易预创建
func (a *Client) TradePreCreate(requestParam TradePreCreateRequestParams) (responseParam TradePreCreateResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeRefund 统一收单交易退款接口
func (a *Client) TradeRefund(requestParam TradeRefundRequestParams) (responseParam TradeRefundResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeAppPay app支付接口2.0
func (a *Client) TradeAppPay(requestParam TradeAppPayRequestParams) (result string, err error) {
	return a.HandlerSDKRequest(&requestParam)
}

// TradePagePay 统一收单下单并支付页面接口
func (a *Client) TradePagePay(requestParam TradePagePayRequestParams) (result string, urlResult *url.URL, err error) {
	return a.HandlerPageRequest("GET", &requestParam)
}

// TradeFastPayRefundQuery 统一收单交易退款查询
func (a *Client) TradeFastPayRefundQuery(requestParam TradeFastPayRefundQueryRequestParams) (
	responseParam TradeFastPayRefundQueryResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeBillDownloadUrlQuery 查询对账单下载地址
func (a *Client) TradeBillDownloadUrlQuery(requestParam TradeBillDownloadUrlQueryRequestParams) (
	responseParam TradeBillDownloadUrlQueryResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}
