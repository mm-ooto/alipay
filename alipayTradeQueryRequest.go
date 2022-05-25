package alipay

import (
	"encoding/json"
	"github.com/mm-ooto/alipay/consts"
	"net/url"
)

// TradeQueryRequest 统一收单线下交易查询
func (a *AliClient) TradeQueryRequest(requestParam TradeQueryRequestParams) (responseParam TradeQueryResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeQueryRequestParams 统一收单线下交易查询请求参数
// 文档地址：https://opendocs.alipay.com/apis/02890l?scene=common
type TradeQueryRequestParams struct {
	OtherRequestParams

	OutTradeNo   string    `json:"out_trade_no,omitempty"`   // 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
	TradeNo      string    `json:"trade_no,omitempty"`       // 支付宝交易号，和商户订单号不能同时为空
	OrgPid       string    `json:"org_pid,omitempty"`        // 银行间联模式下有用，其它场景请不要使用； 双联通过该参数指定需要查询的交易所属收单机构的pid;
	QueryOptIons []*string `json:"query_opt_ions,omitempty"` // 查询选项，商户传入该参数可定制本接口同步响应额外返回的信息字段，数组格式。支持枚举如下：trade_settle_info：返回的交易结算信息，包含分账、补差等信息。 fund_bill_list：交易支付使用的资金渠道。
}

func (t *TradeQueryRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(consts.AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(consts.ApiMethodNameFiled, "alipay.trade.query")
	bytes, _ := json.Marshal(t)
	urlValue.Add(consts.BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradeQueryRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
}

// TradeQueryResponseParams 统一收单线下交易查询响应参数
type TradeQueryResponseParams struct {
	Data struct {
		CommonResParams
		TradeNo               string                  `json:"trade_no"`                 // 支付宝交易号
		OutTradeNo            string                  `json:"out_trade_no"`             // 商家订单号
		BuyerLogonId          string                  `json:"buyer_logon_id"`           // 买家支付宝账号
		TradeStatus           string                  `json:"trade_status"`             // 交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款）
		TotalAmount           float64                 `json:"total_amount"`             // 交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount
		TransCurrency         string                  `json:"trans_currency"`           // 标价币种，该参数的值为支付时传入的
		SettleCurrency        string                  `json:"settle_currency"`          // 订单结算币种，对应支付接口传入的
		SettleAmount          float64                 `json:"settle_amount"`            // 结算币种订单金额
		PayCurrency           float64                 `json:"pay_currency"`             // 订单支付币种
		PayAmount             string                  `json:"pay_amount"`               // 订单币种订单金额
		SettleTransRate       string                  `json:"settle_trans_rate"`        // 结算币种兑换标价币种汇率
		TransPayRate          string                  `json:"trans_pay_rate"`           // 标价币种兑换支付币种汇率
		BuyerPayAmount        float64                 `json:"buyer_pay_amount"`         // 买家实付金额，单位为元，两位小数。该金额代表该笔交易买家实际支付的金额，不包含商户折扣等金额
		PointAmount           float64                 `json:"point_amount"`             // 积分支付的金额，单位为元，两位小数。该金额代表该笔交易中用户使用积分支付的金额，比如集分宝或者支付宝实时优惠等
		InvoiceAmount         float64                 `json:"invoice_amount"`           // 交易中用户支付的可开具发票的金额，单位为元，两位小数。该金额代表该笔交易中可以给用户开具发票的金额
		SendPayDate           string                  `json:"send_pay_date"`            // 本次交易打款给卖家的时间
		ReceiptAmount         string                  `json:"receipt_amount"`           // 实收金额，单位为元，两位小数。该金额为本笔交易，商户账户能够实际收到的金额
		StoreId               string                  `json:"store_id"`                 // 商户门店编号
		TerminalId            string                  `json:"terminal_id"`              // 商户机具终端编号
		FundBillList          FundBillListParams      `json:"fund_bill_list"`           // 交易支付使用的资金渠道。 只有在签约中指定需要返回资金明细，或者入参的query_options中指定时才返回该字段信息。
		StoreName             string                  `json:"store_name"`               // 请求交易支付中的商户店铺的名称
		BuyerUserId           string                  `json:"buyer_user_id"`            // 买家在支付宝的用户id
		IndustrySepcDetailGov string                  `json:"industry_sepc_detail_gov"` // 行业特殊信息-统筹相关
		IndustrySepcDetailAcc string                  `json:"industry_sepc_detail_acc"` // 行业特殊信息-个账相关
		ChargeAmount          string                  `json:"charge_amount"`            // 该笔交易针对收款方的收费金额； 只在银行间联交易场景下返回该信息；
		ChargeFlags           string                  `json:"charge_flags"`             // 费率活动标识
		SettlementId          string                  `json:"settlement_id"`            // 支付清算编号，用于清算对账使用； 只在银行间联交易场景下返回该信息；
		TradeSettleInfo       TradeSettleInfoParams   `json:"trade_settle_info"`        // 返回的交易结算信息，包含分账、补差等信息。 只有在query_options中指定时才返回该字段信息。
		AuthTradePayMode      string                  `json:"auth_trade_pay_mode"`      // 预授权支付模式，该参数仅在信用预授权支付场景下返回。信用预授权支付：CREDIT_PREAUTH_PAY
		BuyerUserType         string                  `json:"buyer_user_type"`          // 买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户。
		MdiscountAmount       string                  `json:"mdiscount_amount"`         // 商家优惠金额
		DiscountAmount        string                  `json:"discount_amount"`          // 平台优惠金额
		Subject               string                  `json:"subject"`                  // 订单标题；只在银行间联交易场景下返回该信息；
		SubMerchantId         string                  `json:"alipay_sub_merchant_id"`   // 间连商户在支付宝端的商户编号； 只在银行间联交易场景下返回该信息；
		ExtInfos              string                  `json:"ext_infos"`                // 交易额外信息，特殊场景下与支付宝约定返回。 json格式。
		PassbackParams        string                  `json:"passback_params"`          // 公用回传参数。 返回支付时传入的passback_params参数信息
		HbFqPayInfo           HbFqPayInfoParams       `json:"hb_fq_pay_info"`           // 若用户使用花呗分期支付，且商家开通返回此通知参数，则会返回花呗分期信息。json格式其它说明详见花呗分期信息说明。 注意：商家需与支付宝约定后才返回本参数。
		CreditPayMode         string                  `json:"credit_pay_mode"`          // 信用支付模式。表示订单是采用信用支付方式（支付时买家没有出资，需要后续履约）。
		CreditBizOrderId      string                  `json:"credit_biz_order_id"`      // 信用业务单号。信用支付场景才有值，先用后付产品里是芝麻订单号。
		EnterprisePayInfo     EnterprisePayInfoParams `json:"enterprise_pay_info"`      // 因公付支付信息
	} `json:"alipay_trade_query_response"`
	Sign string `json:"sign"` // 签名
}

// FundBillListParams 交易支付使用的资金渠道
type FundBillListParams struct {
	FundChannel string  `json:"fund_channel"` // 交易使用的资金渠道
	Amount      float64 `json:"amount"`       // 该支付工具类型所使用的金额
	RealAmount  float64 `json:"real_amount"`  // 渠道实际付款金额
}

// TradeSettleInfoParams 交易结算明细信息列表
type TradeSettleInfoParams struct {
	TradeSettleDetailList []TradeSettleDetailParams `json:"trade_settle_detail_list"`
}

// TradeSettleDetailParams 交易结算明细信息
type TradeSettleDetailParams struct {
	OperationType     string  `json:"operation_type"`      // 结算操作类型。有以下几种类型：replenish(补差)、replenish_refund(退补差)、transfer(分账)、transfer_refund(退分账)、settle(结算)、settle_refund(退结算)、on_settle(待结算)。
	OperationSerialNo string  `json:"operation_serial_no"` // 商户操作序列号。商户发起请求的外部请求号。
	OperationDt       string  `json:"operation_dt"`        // 操作日期
	TransOut          string  `json:"trans_out"`           // 转出账号
	TransIn           string  `json:"trans_in"`            // 转入账号
	Amount            float64 `json:"amount"`              // 实际操作金额，单位为元，两位小数。该参数的值为分账或补差或结算时传入
	OriTransOut       string  `json:"ori_trans_out"`       // 商户请求的转出账号
	OriTransIn        string  `json:"ori_trans_in"`        // 商户请求的转入账号
}

// HbFqPayInfoParams 用户使用花呗分期支付信息
type HbFqPayInfoParams struct {
	UserInstallNum string `json:"user_install_num"` // 用户使用花呗分期支付的分期数
}

// EnterprisePayInfoParams 因公付支付信息
type EnterprisePayInfoParams struct {
	InvoiceAmount float64 `json:"invoice_amount"` // 开票金额
}
