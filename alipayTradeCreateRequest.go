package alipay

import (
	"encoding/json"
	"github.com/mm-ooto/alipay/consts"
	"net/url"
)

// TradeCreateRequest 统一收单交易创建接口
func (a *AliClient) TradeCreateRequest(requestParam TradeCreateRequestParams) (responseParam TradeCreateResponseParams, err error) {
	//requestDataMap := make(map[string]interface{})
	//requestDataMap["biz_content"] = a.SetDataToBizContent(requestParam, requestParam.NeedEncrypt)
	//requestDataMap["notify_url"] = requestParam.NotifyUrl
	//requestDataMap["app_auth_token"] = requestParam.AppAuthToken
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradeCreateRequestParams 统一收单交易创建接口请求参数
// 文档地址：https://opendocs.alipay.com/apis/02890n
type TradeCreateRequestParams struct {
	OtherRequestParams

	OutTradeNo           string                     `json:"out_trade_no"`                    // 商户订单号。由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	TotalAmount          float64                    `json:"total_amount"`                    // 订单总金额，单位为元，精确到小数点后两位，取值范围为 [0.01,100000000]，金额不能为 0。如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】
	Subject              string                     `json:"subject"`                         // 订单标题。 注意：不可使用特殊字符，如 /，=，& 等。
	ProductCode          string                     `json:"product_code"`                    // 销售产品码。如果签约的是当面付快捷版，则传 OFFLINE_PAYMENT；其它支付宝当面付产品传 FACE_TO_FACE_PAYMENT；不传则默认使用 FACE_TO_FACE_PAYMENT。
	SellerId             string                     `json:"seller_id,omitempty"`             // 卖家支付宝用户 ID。 当需要指定收款账号时，通过该参数传入，如果该值为空，则默认为商户签约账号对应的支付宝用户ID。 收款账号优先级规则：门店绑定的收款账户>请求传入的seller_id>商户签约账号对应的支付宝用户ID； 注：直付通和机构间联场景下seller_id无需传入或者保持跟pid一致；如果传入的seller_id与pid不一致，需要联系支付宝小二配置收款关系；
	BuyerId              string                     `json:"buyer_id,omitempty"`              // 买家支付宝用户ID。 2088开头的16位纯数字，小程序场景下获取用户ID请参考：用户授权; 其它场景下获取用户ID请参考：网页授权获取用户信息; 注：交易的买家与卖家不能相同。
	Body                 string                     `json:"body,omitempty"`                  // 订单附加信息。	如果请求时传递了该参数，将在异步通知、对账单中原样返回，同时会在商户和用户的pc账单详情中作为交易描述展示
	GoodsDetail          []*GoodsDetailParams       `json:"goods_detail,omitempty"`          // 订单包含的商品列表信息，为 JSON 格式，其它说明详见商品明细说明
	TimeExpire           string                     `json:"time_expire,omitempty"`           // 订单绝对超时时间。 格式为yyyy-MM-dd HH:mm:ss。注：time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。
	TimeoutExpress       string                     `json:"timeout_express,omitempty"`       // 订单相对超时时间。从交易创建时间开始计算。 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。 当面付场景默认值为3h。注：time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。
	SettleInfo           *SettleInfoParams          `json:"settle_info,omitempty"`           // 描述结算信息，json格式
	ExtendParams         *ExtendParamsParams        `json:"extend_params,omitempty"`         // 业务扩展参数
	BusinessParams       *BusinessParamsParams      `json:"business_params,omitempty"`       // 商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，格式为json格式
	DiscountableAmount   float64                    `json:"discountable_amount,omitempty"`   // 可打折金额。参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围为 [0.01,100000000]。如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
	UndiscountableAmount float64                    `json:"undiscountable_amount,omitempty"` // 不可打折金额。不参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果同时传入了【可打折金额】、【不可打折金额】和【订单总金额】，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】。如果订单金额全部参与优惠计算，则【可打折金额】和【不可打折金额】都无需传入。
	StoreId              string                     `json:"store_id,omitempty"`              // 商户门店编号。 指商户创建门店时输入的门店编号。
	OperatorId           string                     `json:"operator_id,omitempty"`           // 商户操作员编号。
	TerminalId           string                     `json:"terminal_id,omitempty"`           // 商户机具终端编号。
	ReceiverAddressInfo  *ReceiverAddressInfoParams `json:"receiver_address_info,omitempty"` // 收货人及地址信息
}

func (t *TradeCreateRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(consts.NotifyUrlFiled, t.NotifyUrl)
	urlValue.Add(consts.AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(consts.ApiMethodNameFiled, "alipay.trade.create")
	bytes, _ := json.Marshal(t)
	urlValue.Add(consts.BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradeCreateRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
}


// TradeCreateResponseParams 统一收单交易创建接口响应参数
type TradeCreateResponseParams struct {
	Data struct {
		CommonResParams
		OutTradeNo string `json:"out_trade_no"` // 商户订单号
		TradeNo    string `json:"trade_no"`     // 支付宝交易号
	} `json:"alipay_trade_create_response"`
	Sign string `json:"sign"` // 签名
}

// SettleInfoParams 描述结算信息
type SettleInfoParams struct {
	SettleDetailInfos []*SettleDetailInfoParams `json:"settle_detail_infos"`          // 结算详细信息，json数组，目前只支持一条。
	SettlePeriodTime  string                    `json:"settle_period_time,omitempty"` // 该笔订单的超期自动确认结算时间，到达期限后，将自动确认结算。此字段只在签约账期结算模式时有效。取值范围：1d～365d。d-天。 该参数数值不接受小数点。
}

// SettleDetailInfoParams 结算详细信息
type SettleDetailInfoParams struct {
	TransInType      string `json:"trans_in_type"`                // 结算收款方的账户类型。cardAliasNo：结算收款方的银行卡编号; userId：表示是支付宝账号对应的支付宝唯一用户号; loginName：表示是支付宝登录号；defaultSettle：表示结算到商户进件时设置的默认结算账号，结算主体为门店时不支持传defaultSettle；
	TransIn          string `json:"trans_in"`                     // 结算收款方。当结算收款方类型是cardAliasNo时，本参数为用户在支付宝绑定的卡编号；结算收款方类型是userId时，本参数为用户的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；当结算收款方类型是loginName时，本参数为用户的支付宝登录号；当结算收款方类型是defaultSettle时，本参数不能传值，保持为空。
	SummaryDimension string `json:"summary_dimension,omitempty"`  // 结算汇总维度，按照这个维度汇总成批次结算，由商户指定。 目前需要和结算收款方账户类型为cardAliasNo配合使用
	SettleEntityId   string `json:"settle_entity_id,omitempty"`   // 结算主体标识。当结算主体类型为SecondMerchant时，为二级商户的SecondMerchantID；当结算主体类型为Store时，为门店的外标。
	SettleEntityType string `json:"settle_entity_type,omitempty"` // 结算主体类型。 二级商户:SecondMerchant;商户或者直连商户门店:Store
	Amount           string `json:"amount"`                       // 结算的金额，单位为元。在创建订单和支付接口时必须和交易金额相同。在结算确认接口时必须等于交易金额减去已退款金额。
}

// BusinessParamsParams 商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景
type BusinessParamsParams struct {
	CampusCard      string `json:"campus_card,omitempty"`       // 校园卡编号
	CardType        string `json:"card_type,omitempty"`         // 虚拟卡卡类型
	ActualOrderTime string `json:"actual_order_time,omitempty"` // 实际订单时间，在乘车码场景，传入的是用户刷码乘车时间，格式：2019-05-14 09:18:55
	GoodTaxes       string `json:"good_taxes,omitempty"`        // 商户传入的交易税费。需要落地风控使用
}

// LogisticsDetailParams 物流信息
type LogisticsDetailParams struct {
	LogisticsType string `json:"logistics_type,omitempty"` // 物流类型, POST 平邮, EXPRESS 其他快递, VIRTUAL 虚拟物品, EMS EMS, DIRECT 无需物流。
}

// ReceiverAddressInfoParams 收货人及地址信息
type ReceiverAddressInfoParams struct {
	Name         string `json:"name,omitempty"`          // 收货人的姓名
	Address      string `json:"address,omitempty"`       // 收货地址
	Mobile       string `json:"mobile,omitempty"`        // 收货人手机号
	Zip          string `json:"zip,omitempty"`           // 收货地址邮编
	DivisionCode string `json:"division_code,omitempty"` // 中国标准城市区域码
}
