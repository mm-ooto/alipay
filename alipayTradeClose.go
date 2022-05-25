package alipay

import (
	"encoding/json"
	"github.com/mm-ooto/alipay/consts"
	"net/url"
)

// TradeClose 统一收单交易关闭接口
func (a *AliClient) TradeClose(requestParam TradeCloseRequestParams) (responseParam TradeCloseResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeCloseRequestParams 统一收单交易关闭接口请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.close
type TradeCloseRequestParams struct {
	OtherRequestParams

	TradeNo    string `json:"trade_no,omitempty"`     // 该交易在支付宝系统中的交易流水号。最短 16 位，最长 64 位。和out_trade_no不能同时为空，如果同时传了 out_trade_no和 trade_no，则以 trade_no为准。
	OutTradeNo string `json:"out_trade_no,omitempty"` // 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_n
	OperatorId string `json:"operator_id,omitempty"`  // 商家操作员编号 id，由商家自定义。
}

func (t *TradeCloseRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(consts.NotifyUrlFiled, t.NotifyUrl)
	urlValue.Add(consts.AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(consts.ApiMethodNameFiled, "alipay.trade.close")
	bytes, _ := json.Marshal(t)
	urlValue.Add(consts.BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradeCloseRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
}

// TradeCloseResponseParams 统一收单交易关闭接口响应参数
type TradeCloseResponseParams struct {
	Data struct {
		CommonResParams
		TradeNo    string `json:"trade_no,omitempty"`     // 支付宝交易号
		OutTradeNo string `json:"out_trade_no,omitempty"` // 创建交易传入的商户订单号
	} `json:"alipay_trade_close_response"`
	Sign string `json:"sign"` // 签名
}
