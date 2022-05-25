package alipay

import (
	"encoding/json"
	"github.com/mm-ooto/alipay/consts"
	"net/url"
)

// TradeFastpayRefundQueryRequest 统一收单交易退款查询
func (a *AliClient) TradeFastpayRefundQueryRequest(requestParam TradeFastpayRefundQueryRequestParams) (
	responseParam TradeFastpayRefundQueryResponseParams, err error) {
	if err = a.HandlerRequest("POST",&requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeFastpayRefundQueryRequestParams 统一收单交易退款查询请求参数
// 文档地址：https://opendocs.alipay.com/apis/0287wc
type TradeFastpayRefundQueryRequestParams struct {
	OtherRequestParams

	TradeNo      string    `json:"trade_no,omitempty"`      // 支付宝交易号。 和商户订单号不能同时为空
	OutTradeNo   string    `json:"out_trade_no,omitempty"`  // 商户订单号。 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
	OutRequestNo string    `json:"out_request_no"`          // 退款请求号。 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的商户订单号。
	QueryOptions []*string `json:"query_options,omitempty"` // 查询选项
}

func (t *TradeFastpayRefundQueryRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(consts.AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(consts.ApiMethodNameFiled, "alipay.trade.fastpay.refund.query")
	bytes, _ := json.Marshal(t)
	urlValue.Add(consts.BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradeFastpayRefundQueryRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
}

// TradeFastpayRefundQueryResponseParams 统一收单交易退款查询响应参数
type TradeFastpayRefundQueryResponseParams struct {
	Data struct {
		CommonResParams
		TradeNo              string                      `json:"trade_no"`                // 支付宝交易号
		OutTradeNo           string                      `json:"out_trade_no"`            // 创建交易传入的商户订单号
		OutRequestNo         string                      `json:"out_request_no"`          // 本笔退款对应的退款请求号。
		TotalAmount          float64                     `json:"total_amount"`            // 该笔退款所对应的交易的订单金额
		RefundAmount         float64                     `json:"refund_amount"`           // 本次退款请求，对应的退款金额
		RefundStatus         string                      `json:"refund_status"`           // 退款状态。枚举值： REFUND_SUCCESS 退款处理成功；未返回该字段表示退款请求未收到或者退款失败；注：如果退款查询发起时间早于退款时间，或者间隔退款发起时间太短，可能出现退款查询时还没处理成功，后面又处理成功的情况，建议商户在退款发起后间隔10秒以上再发起退款查询请求。
		RefundRoyaltys       []RefundRoyaltyResultParams `json:"refund_royaltys"`         // 退分账明细信息
		GmtRefundWay         string                      `json:"gmt_refund_way"`          // 退款时间。默认不返回该信息，需要在入参的query_options中指定"gmt_refund_pay"值时才返回该字段信息。格式为yyyy-MM-dd HH:mm:ss
		RefundDetailItemList []TradeFundBillParams       `json:"refund_detail_item_list"` // 本次退款使用的资金渠道；默认不返回该信息，需要在入参的query_options中指定"refund_detail_item_list"值时才返回该字段信息。
		SendBackFee          string                      `json:"send_back_fee"`           // 本次商户实际退回金额；默认不返回该信息，需要在入参的query_options中指定"refund_detail_item_list"值时才返回该字段信息。
		DepositBackInfo      DepositBackInfoParams       `json:"deposit_back_info"`       // 银行卡冲退信息。 该字段默认不返回；
		EnterprisePayInfo    EnterprisePayInfoParams     `json:"enterprise_pay_info"`     // 因公付退款信息，只有入参的query_options中指定enterprise_pay_info时才返回该字段信息
	} `json:"alipay_trade_fastpay_refund_query_response"`
	Sign string `json:"sign"` // 签名
}

// RefundRoyaltyResultParams 退分账明细信息
type RefundRoyaltyResultParams struct {
	RefundAmount  float64 `json:"refund_amount"`   // 退分账金额
	RoyaltyType   string  `json:"royalty_type"`    // 分账类型. 普通分账为：transfer;补差为：replenish;为空默认为分账transfer;
	ResultCode    string  `json:"result_code"`     // 退分账结果码
	TransOut      string  `json:"trans_out"`       // 转出人支付宝账号对应用户ID
	TransOutEmail string  `json:"trans_out_email"` // 转出人支付宝账号
	TransIn       string  `json:"trans_in"`        // 转入人支付宝账号对应用户ID
	TransInEmail  string  `json:"trans_in_email"`  // 转入人支付宝账号
}

// DepositBackInfoParams 银行卡冲退信息。
type DepositBackInfoParams struct {
	HasDepositBack     string  `json:"has_deposit_back"`      // 是否存在银行卡冲退信息。
	DbackStatus        string  `json:"dback_status"`          // 银行卡冲退状态。S-成功，F-失败，P-处理中。银行卡冲退失败，资金自动转入用户支付宝余额。
	DbackAmount        float64 `json:"dback_amount"`          // 银行卡冲退金额
	BankAckTime        string  `json:"bank_ack_time"`         // 银行响应时间，格式为yyyy-MM-dd HH:mm:ss
	EstBankReceiptTime string  `json:"est_bank_receipt_time"` // 预估银行到账时间，格式为yyyy-MM-dd HH:mm:ss
}
