package alipay

import (
	"encoding/json"
	"github.com/mm-ooto/alipay/consts"
	"net/url"
)

// TradeRefundRequest 统一收单交易退款接口
func (a *AliClient) TradeRefundRequest(requestParam TradeRefundRequestParams) (responseParam TradeRefundResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeRefundRequestParams 统一收单交易退款接口请求参数
// 文档地址：https://opendocs.alipay.com/apis/0287wa
type TradeRefundRequestParams struct {
	OtherRequestParams

	OutTradeNo              string                                `json:"out_trade_no,omitempty"`              //  商户订单号。订单支付时传入的商户订单号，商家自定义且保证商家系统中唯一。与支付宝交易号 trade_no 不能同时为空。
	TradeNo                 string                                `json:"trade_no,omitempty"`                  // 支付宝交易号。和商户订单号 out_trade_no 不能同时为空。
	RefundAmount            float64                               `json:"refund_amount"`                       // 退款金额。 需要退款的金额，该金额不能大于订单金额，单位为元，支持两位小数。
	RefundReason            string                                `json:"refund_reason,omitempty"`             // 退款原因说明。商家自定义，将在对账单的退款明细中作为备注返回，同时会在商户和用户的pc退款账单详情中展示
	OutRequestNo            string                                `json:"out_request_no,omitempty"`            // 退款请求号。标识一次退款请求，需要保证在交易号下唯一，如需部分退款，则此参数必传。
	RefundRoyaltyParameters []*OpenApiRoyaltyDetailInfoPojoParams `json:"refund_royalty_parameters,omitempty"` // 退分账明细信息。
	QueryOptions            []*string                             `json:"query_options,omitempty"`             // 查询选项
}

func (t *TradeRefundRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(consts.AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(consts.ApiMethodNameFiled, "alipay.trade.refund")
	bytes, _ := json.Marshal(t)
	urlValue.Add(consts.BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradeRefundRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
}

// TradeRefundResponseParams 统一收单交易退款接口响应参数
type TradeRefundResponseParams struct {
	Data struct {
		CommonResParams
		TradeNo              string                `json:"trade_no"`                // 支付宝交易号
		OutTradeNo           string                `json:"out_trade_no"`            // 商家订单号
		BuyerLogonId         string                `json:"buyer_logon_id"`          // 用户的登录id
		FundChange           string                `json:"fund_change"`             // 本次退款是否发生了资金变化
		RefundFee            float64               `json:"refund_fee"`              // 退款总金额。指该笔交易累计已经退款成功的金额。
		RefundDetailItemList []TradeFundBillParams `json:"refund_detail_item_list"` // 退款使用的资金渠道。只有在签约中指定需要返回资金明细，或者入参的query_options中指定时才返回该字段信息。
		StoreName            string                `json:"store_name"`              // 交易在支付时候的门店名称
		BuyerUserId          string                `json:"buyer_user_id"`           // 买家在支付宝的用户id
		SendBackFee          string                `json:"send_back_fee"`           // 本次商户实际退回金额。说明：如需获取该值，需在入参query_options中传入 refund_detail_item_list。
	} `json:"alipay_trade_refund_response"`
	Sign string `json:"sign"` // 签名
}

// OpenApiRoyaltyDetailInfoPojoParams 退分账明细信息。
type OpenApiRoyaltyDetailInfoPojoParams struct {
	RoyaltyType  string  `json:"royalty_type,omitempty"`   // 分账类型. 普通分账为：transfer; 补差为：replenish; 为空默认为分账transfer;
	TransOut     string  `json:"trans_out,omitempty"`      // 支出方账户。如果支出方账户类型为userId，本参数为支出方的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；如果支出方类型为loginName，本参数为支出方的支付宝登录号。 泛金融类商户分账时，该字段不要上送。
	TransOutType string  `json:"trans_out_type,omitempty"` // 支出方账户类型。userId表示是支付宝账号对应的支付宝唯一用户号;loginName表示是支付宝登录号； 泛金融类商户分账时，该字段不要上送。
	TransInType  string  `json:"trans_in_type,omitempty"`  // 收入方账户类型。userId表示是支付宝账号对应的支付宝唯一用户号;cardAliasNo表示是卡编号;loginName表示是支付宝登录号；
	TransIn      string  `json:"trans_in"`                 // 收入方账户。如果收入方账户类型为userId，本参数为收入方的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；如果收入方类型为cardAliasNo，本参数为收入方在支付宝绑定的卡编号；如果收入方类型为loginName，本参数为收入方的支付宝登录号；
	Amount       float64 `json:"amount,omitempty"`         // 分账的金额，单位为元
	Desc         string  `json:"desc,omitempty"`           // 分账描述
	RoyaltyScene string  `json:"royalty_scene,omitempty"`  // 可选值：达人佣金、平台服务费、技术服务费、其他
	TransInName  string  `json:"trans_in_name,omitempty"`  // 分账收款方姓名，上送则进行姓名与支付宝账号的一致性校验，校验不一致则分账失败。不上送则不进行姓名校验
}

// TradeFundBillParams 退款使用的资金渠道。
type TradeFundBillParams struct {
	FundChannel string  `json:"fund_channel"` // 交易使用的资金渠道
	Amount      float64 `json:"amount"`       // 该支付工具类型所使用的金额
	RealAmount  float64 `json:"real_amount"`  // 渠道实际付款金额
	FundType    string  `json:"fund_type"`    // 渠道所使用的资金类型,目前只在资金渠道(fund_channel)是银行卡渠道(BANKCARD)的情况下才返回该信息(DEBIT_CARD:借记卡,CREDIT_CARD:信用卡,MIXED_CARD:借贷合一卡)
}
