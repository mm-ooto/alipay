package alipay

// TradeDataserviceBillDownloadurlQuery 查询对账单下载地址
func (a *AliClient) TradeDataserviceBillDownloadurlQuery(requestParam TradeDataserviceBillDownloadURLQueryRequestParams) (
	responseParam TradeDataserviceBillDownloadURLQueryResponseParams, err error) {
	requestDataMap := make(map[string]interface{})
	requestDataMap["biz_content"] = a.SetDataToBizContent(requestParam, requestParam.NeedEncrypt)
	requestDataMap["app_auth_token"] = requestParam.AppAuthToken
	if err = a.HandlerRequest("POST", "alipay.data.dataservice.bill.downloadurl.query", requestParam.NeedEncrypt, requestDataMap, &responseParam); err != nil {
		return
	}
	return
}

// TradeDataserviceBillDownloadURLQueryRequestParams 查询对账单下载地址接口请求参数
// 文档地址：https://docs.open.alipay.com/api_15/alipay.data.dataservice.bill.downloadurl.query
type TradeDataserviceBillDownloadURLQueryRequestParams struct {
	OtherRequestParams

	BillType string `json:"bill_type"` // 必选 账单类型，商户通过接口或商户经开放平台授权后其所属服务商通过接口可以获取以下账单类型：trade、signcustomer；trade指商户基于支付宝交易收单的业务账单；signcustomer是指基于商户支付宝余额收入及支出等资金变动的帐务账单。
	BillDate string `json:"bill_date"` // 必选 账单时间：日账单格式为yyyy-MM-dd，最早可下载2016年1月1日开始的日账单；月账单格式为yyyy-MM，最早可下载2016年1月开始的月账单。
}

// TradeDataserviceBillDownloadURLQueryResponseParams 查询对账单下载地址接口响应参数
type TradeDataserviceBillDownloadURLQueryResponseParams struct {
}
