package alipay

import "net/url"

// TradePagePayRequest 统一收单下单并支付页面接口
func (a *AliClient) TradePagePayRequest(requestParam TradePagePayRequestParams) (result string, urlResult *url.URL, err error) {
	requestDataMap := make(map[string]interface{})
	requestDataMap["biz_content"] = a.SetDataToBizContent(requestParam,false)
	requestDataMap["notify_url"] = requestParam.NotifyUrl
	requestDataMap["app_auth_token"] = requestParam.AppAuthToken
	return a.HandlerPageRequest("POST", "alipay.trade.page.pay", requestDataMap)
}

// TradePagePayRequestParams 统一收单下单并支付页面接口请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay
type TradePagePayRequestParams struct {
	OtherRequestParams

	OutTradeNo          string               `json:"out_trade_no"`                    // 商户订单号。由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	TotalAmount         string               `json:"total_amount"`                    // 订单总金额。单位为元，精确到小数点后两位，取值范围：[0.01,100000000] 。
	Subject             string               `json:"subject"`                         // 订单标题。 注意：不可使用特殊字符，如 /，=，& 等。
	ProductCode         string               `json:"product_code,omitempty"`          // 产品码。 商家和支付宝签约的产品码。 枚举值（点击查看签约情况）：QUICK_MSECURITY_PAY：无线快捷支付产品；CYCLE_PAY_AUTH：周期扣款产品。默认值为QUICK_MSECURITY_PAY。
	Body                string               `json:"body,omitempty"`                  // 订单附加信息。如果请求时传递了该参数，将在异步通知、对账单中原样返回，同时会在商户和用户的pc账单详情中作为交易描述展示
	QrPayMode           string               `json:"qr_pay_mode,omitempty"`           // PC扫码支付的方式。 支持前置模式和跳转模式。前置模式是将二维码前置到商户的订单确认页的模式。需要商户在自己的页面中以 iframe 方式请求支付宝页面。具体支持的枚举值请查看文档
	QrcodeWidth         string               `json:"qrcode_width,omitempty"`          // 商户自定义二维码宽度。 注：qr_pay_mode=4时该参数有效
	GoodsDetail         []*GoodsDetailParams `json:"goods_detail,omitempty"`          // 订单包含的商品列表信息，json格式，其它说明详见商品明细说明
	TimeExpire          string               `json:"time_expire,omitempty"`           // 订单绝对超时时间。格式为yyyy-MM-dd HH:mm:ss。注：time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。
	TimeExpress         string               `json:"time_express,omitempty"`          // 建议使用time_expire字段。 订单相对超时时间。从买家确认支付时间开始计算。该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：5m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。默认值为15d。注：1. 无线支付场景最小值为5m，低于5m支付超时时间按5m计算。2. time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。
	RoyaltyInfo         *RoyaltyInfo         `json:"royalty_info,omitempty"`          // 描述分账信息，json格式。
	SubMerchant         *SubMerchant         `json:"sub_merchant,omitempty"`          // 二级商户信息。 直付通模式和机构间连模式下必传，其它场景下不需要传入。
	SettleInfo          *SettleInfoParams    `json:"settle_info,omitempty"`           // 描述结算信息，json格式。
	ExtendParams        *ExtendParamsParams  `json:"extend_params,omitempty"`         // 业务扩展参数
	BusinessParams      string               `json:"business_params,omitempty"`       // 商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，格式为json格式
	PromoParams         string               `json:"promo_params,omitempty"`          // 优惠参数 注：仅与支付宝协商后可用
	IntegrationType     string               `json:"integration_type,omitempty"`      // 请求后页面的集成方式。 枚举值：ALIAPP：支付宝钱包内PCWEB：PC端访问默认值为PCWEB。
	RequestFromUrl      string               `json:"request_from_url,omitempty"`      // 请求来源地址。如果使用ALIAPP的集成方式，用户中途取消支付会返回该地址。
	AgreementSignParams *AgreementSignParams `json:"agreement_sign_params,omitempty"` // 签约参数，支付后签约场景使用
	StoreId             string               `json:"store_id,omitempty"`              // 商户门店编号。 指商户创建门店时输入的门店编号。
	EnablePayChannels   string               `json:"enable_pay_channels,omitempty"`   // 指定支付渠道，多个渠道以逗号分割。用户只能使用此处指定渠道进行支付。 与disable_pay_channels互斥，支持传入的值：渠道列表。注意：如果传入了指定支付渠道，则用户只能用指定内的渠道支付，包括营销渠道也要指定才能使用。若所有指定渠道用户都不可使用，将导致用户无法支付，慎用。
	DisablePayChannels  string               `json:"disable_pay_channels,omitempty"`  // 禁用渠道,用户不可用指定渠道支付，多个渠道以逗号分割 注，与enable_pay_channels互斥
	MerchantOrderNo     string               `json:"merchant_order_no,omitempty"`     // 商户的原始订单号
	ExtUserInfo         *ExtUserInfo         `json:"ext_user_info,omitempty"`         // 外部指定买家
	InvoiceInfo         *InvoiceInfo         `json:"invoice_info,omitempty"`          // 开票信息
}

// RoyaltyInfo 分账信息
type RoyaltyInfo struct {
	RoyaltyType        string              `json:"royalty_type,omitempty"` // 分账类型 卖家的分账类型，目前只支持传入ROYALTY（普通分账类型）
	RoyaltyDetailInfos *RoyaltyDetailInfos `json:"royalty_detail_infos"`   // 分账明细的信息，可以描述多条分账指令，json数组。
}

// RoyaltyDetailInfos 分账明细信息
type RoyaltyDetailInfos struct {
	SerialNo         int     `json:"serial_no,omitempty"`         // 分账序列号，表示分账执行的顺序，必须为正整数
	TransInType      string  `json:"trans_in_type,omitempty"`     // 接受分账金额的账户类型
	BatchNo          string  `json:"batch_no"`                    // 分账批次号 分账批次号。目前需要和转入账号类型为bankIndex配合使用。
	OutRelationId    string  `json:"out_relation_id,omitempty"`   // 商户分账的外部关联号，用于关联到每一笔分账信息，商户需保证其唯一性。 如果为空，该值则默认为“商户网站唯一订单号+分账序列号”
	TransOutType     string  `json:"trans_out_type"`              // 要分账的账户类型。 目前只支持userId：支付宝账号对应的支付宝唯一用户号。默认值为userId。
	TransOut         string  `json:"trans_out"`                   // 如果转出账号类型为userId，本参数为要分账的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
	TransIn          string  `json:"trans_in"`                    // 如果转入账号类型为userId，本参数为接受分账金额的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。  如果转入账号类型为bankIndex，本参数为28位的银行编号（商户和支付宝签约时确定）。如果转入账号类型为storeId，本参数为商户的门店ID。
	Amount           float64 `json:"amount"`                      // 分账的金额，单位为元
	Desc             float64 `json:"desc,omitempty"`              // 分账描述信息
	AmountPercentage string  `json:"amount_percentage,omitempty"` // 分账的比例，值为20代表按20%的比例分账

}

// SubMerchant 二级商户信息
type SubMerchant struct {
	MerchantId   string `json:"merchant_id"`             // 间连受理商户的支付宝商户编号，通过间连商户入驻后得到。间连业务下必传，并且需要按规范传递受理商户编号。
	MerchantType string `json:"merchant_type,omitempty"` // 商户id类型，
}

// AgreementSignParams 签约参数，支付后签约场景使用
type AgreementSignParams struct {
	PersonalProductCode string              `json:"personal_product_code"`           // 个人签约产品码，商户和支付宝签约时确定。
	SignScene           string              `json:"sign_scene,omitempty"`            // 协议签约场景，商户和支付宝签约时确定。 当传入商户签约号external_agreement_no时，场景不能为默认值DEFAULT|DEFAULT。
	ExternalAgreementNo string              `json:"external_agreement_no,omitempty"` // 商户签约号，代扣协议中标示用户的唯一签约号（确保在商户系统中唯一）。 格式规则：支持大写小写字母和数字，最长32位。商户系统按需传入，如果同一用户在同一产品码、同一签约场景下，签订了多份代扣协议，那么需要指定并传入该值。
	ExternalLogonId     string              `json:"external_logon_id,omitempty"`     // 用户在商户网站的登录账号，用于在签约页面展示，如果为空，则不展示
	SignValidityPeriod  string              `json:"sign_validity_period,omitempty"`  // 当前用户签约请求的协议有效周期。 整形数字加上时间单位的协议有效期，从发起签约请求的时间开始算起。目前支持的时间单位：1. d：天2. m：月如果未传入，默认为长期有效。
	ThirdPartyType      string              `json:"third_party_type,omitempty"`      // 签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。 取值范围：1. PARTNER（平台商户）;2. MERCHANT（集团商户），集团下子商户可共享用户签约内容;默认为PARTNER。
	SignMerchant        *SignMerchantParams `json:"sign_merchant,omitempty"`         // 此参数用于传递子商户信息，无特殊需求时不用关注。目前商户代扣、海外代扣、淘旅行信用住产品支持传入该参数（在销售方案中“是否允许自定义子商户信息”需要选是）。
	BuckleAppId         string              `json:"buckle_app_id,omitempty"`         // 商户在芝麻端申请的appId
	BuckleMerchantId    string              `json:"buckle_merchant_id,omitempty"`    // 商户在芝麻端申请的merchantId
	PromoParams         string              `json:"promo_params,omitempty"`          // 签约营销参数，此值为json格式；具体的key需与营销约定
}

// InvoiceInfo 开票信息
type InvoiceInfo struct {
	KeyInfo *InvoiceKeyInfo `json:"key_info"` // 开票关键信息
	Details string          `json:"details"`  // 开票内容 注：json数组格式
}

// InvoiceKeyInfo 开票关键信息
type InvoiceKeyInfo struct {
	IsSupportInvoice    bool   `json:"is_support_invoice"`    // 该交易是否支持开票
	InvoiceMerchantName string `json:"invoice_merchant_name"` // 开票商户名称：商户品牌简称|商户门店简称
	TaxNum              string `json:"tax_num"`               // 税号
}
