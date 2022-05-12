package alipay

// TradeAppPayRequest app支付接口2.0
func (a *AliClient) TradeAppPayRequest(requestParam TradeAppPayRequestParams) (result string, err error) {
	requestDataMap := make(map[string]interface{})
	requestDataMap["biz_content"] = a.SetDataToBizContent(requestParam,false)
	requestDataMap["app_auth_token"] = requestParam.AppAuthToken
	return a.HandlerSDKRequest("alipay.trade.app.pay", requestDataMap)
}

// TradeAppPayRequestParams app支付接口2.0请求参数 ,omitempty
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.app.pay
type TradeAppPayRequestParams struct {
	OtherRequestParams

	OutTradeNo          string                `json:"out_trade_no"`                    // 商户订单号。由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	TotalAmount         string                `json:"total_amount"`                    // 订单总金额。单位为元，精确到小数点后两位，取值范围：[0.01,100000000] 。
	Subject             string                `json:"subject"`                         // 订单标题。 注意：不可使用特殊字符，如 /，=，& 等。
	ProductCode         string                `json:"product_code"`          // 产品码。 商家和支付宝签约的产品码。 枚举值（点击查看签约情况）：QUICK_MSECURITY_PAY：无线快捷支付产品；CYCLE_PAY_AUTH：周期扣款产品。默认值为QUICK_MSECURITY_PAY。
	Body                string                `json:"body,omitempty"`                  // 订单附加信息。如果请求时传递了该参数，将在异步通知、对账单中原样返回，同时会在商户和用户的pc账单详情中作为交易描述展示
	GoodsDetail         []*GoodsDetailParams  `json:"goods_detail,omitempty"`          // 订单包含的商品列表信息，json格式，其它说明详见商品明细说明
	TimeExpire          string                `json:"time_expire,omitempty"`           // 订单绝对超时时间。格式为yyyy-MM-dd HH:mm:ss。注：time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。
	TimeExpress         string                `json:"time_express,omitempty"`          // 建议使用time_expire字段。 订单相对超时时间。从买家确认支付时间开始计算。该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：5m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。默认值为15d。注：1. 无线支付场景最小值为5m，低于5m支付超时时间按5m计算。2. time_expire和timeout_express两者只需传入一个或者都不传，如果两者都传，优先使用time_expire。
	ExtendParams        []*ExtendParamsParams `json:"extend_params,omitempty"`         // 业务扩展参数
	PromoParams         string                `json:"promo_params,omitempty"`          // 优惠参数 注：仅与支付宝协商后可用
	PassbackParams      string                `json:"passback_params,omitempty"`       // 公用回传参数。 如果请求时传递了该参数，支付宝会在异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。
	AgreementSignParams *SignParamsParams     `json:"agreement_sign_params,omitempty"` // 签约参数。如果希望在sdk中支付并签约，需要在这里传入签约信息。 周期扣款场景 product_code 为 CYCLE_PAY_AUTH 时必填。
	StoreId             string                `json:"store_id,omitempty"`              // 商户门店编号。 指商户创建门店时输入的门店编号。
	EnablePayChannels   string                `json:"enable_pay_channels,omitempty"`   // 指定支付渠道，多个渠道以逗号分割。用户只能使用此处指定渠道进行支付。 与disable_pay_channels互斥，支持传入的值：渠道列表。注意：如果传入了指定支付渠道，则用户只能用指定内的渠道支付，包括营销渠道也要指定才能使用。若所有指定渠道用户都不可使用，将导致用户无法支付，慎用。
	SpecifiedChannel    string                `json:"specified_channel,omitempty"`     // 指定单通道，仅支持传入一个渠道。 注意：目前仅支持传入 pcredit，若由于用户原因指定渠道不可用（不能支付），允许用户选择其他渠道支付。该参数不可与花呗分期参数同时传入。
	DisablePayChannels  string                `json:"disable_pay_channels,omitempty"`  // 禁用渠道,用户不可用指定渠道支付，多个渠道以逗号分割 注，与enable_pay_channels互斥
	MerchantOrderNo     string                `json:"merchant_order_no,omitempty"`     // 商户的原始订单号
	ExtUserInfo         *ExtUserInfo          `json:"ext_user_info,omitempty"`         // 外部指定买家
}

// SignParamsParams 签约参数。
type SignParamsParams struct {
	PersonalProductCode string              `json:"personal_product_code"`           // 个人签约产品码，商户和支付宝签约时确定。
	SignScene           string              `json:"sign_scene"`                      // 协议签约场景，商户和支付宝签约时确定，商户可咨询技术支持。
	ExternalAgreementNo string              `json:"external_agreement_no,omitempty"` // 商户签约号，代扣协议中标示用户的唯一签约号（确保在商户系统中唯一）。 格式规则：支持大写小写字母和数字，最长32位。 商户系统按需传入，如果同一用户在同一产品码、同一签约场景下，签订了多份代扣协议，那么需要指定并传入该值。
	ExternalLogonId     string              `json:"external_logon_id,omitempty"`     // 用户在商户网站的登录账号，用于在签约页面展示，如果为空，则不展示
	AccessParams        *AccessParams       `json:"access_params"`                   // 请按当前接入的方式进行填充，且输入值必须为文档中的参数取值范围。
	SubMerchant         *SignMerchantParams `json:"sub_merchant,omitempty"`          // 此参数用于传递子商户信息，无特殊需求时不用关注。目前商户代扣、海外代扣、淘旅行信用住产品支持传入该参数（在销售方案中“是否允许自定义子商户信息”需要选是）。
	PeriodRuleParams    *PeriodRuleParams   `json:"period_rule_params,omitempty"`    // 周期管控规则参数period_rule_params，在签约周期扣款产品（如CYCLE_PAY_AUTH_P）时必传，在签约其他产品时无需传入。 周期扣款产品，会按照这里传入的参数提示用户，并对发起扣款的时间、金额、次数等做相应限制。
	SignNotifyUrl       string              `json:"sign_notify_url,omitempty"`       // 签约成功后商户用于接收异步通知的地址。如果不传入，签约与支付的异步通知都会发到外层notify_url参数传入的地址；如果外层也未传入，签约与支付的异步通知都会发到商户appid配置的网关地址。
}

type AccessParams struct {
	Channel string `json:"channel"` // 目前支持以下值： 1. ALIPAYAPP （钱包h5页面签约）2. QRCODE(扫码签约)3. QRCODEORSMS(扫码签约或者短信签约)
}

type SignMerchantParams struct {
	SubMerchantId                 string `json:"sub_merchant_id,omitempty"`                  // 子商户的商户id
	SubMerchantName               string `json:"sub_merchant_name,omitempty"`                // 子商户的商户名称
	SubMerchantServiceName        string `json:"sub_merchant_service_name,omitempty"`        // 子商户的服务名称
	SubMerchantServiceDescription string `json:"sub_merchant_service_description,omitempty"` // 子商户的服务描述
}

type PeriodRuleParams struct {
	PeriodType    string  `json:"period_type"`              // 周期类型period_type是周期扣款产品必填，枚举值为DAY和MONTH。 DAY即扣款周期按天计，MONTH代表扣款周期按自然月。与另一参数period组合使用确定扣款周期，例如period_type为DAY，period=30，则扣款周期为30天；period_type为MONTH，period=3，则扣款周期为3个自然月。自然月是指，不论这个月有多少天，周期都计算到月份中的同一日期。例如1月3日到2月3日为一个自然月，1月3日到4月3日为三个自然月。注意周期类型使用MONTH的时候，计划扣款时间execute_time不允许传28日之后的日期（可以传28日），以此避免有些月份可能不存在对应日期的情况。
	Period        int     `json:"period"`                   // 周期数period是周期扣款产品必填。与另一参数period_type组合使用确定扣款周期，例如period_type为DAY，period=90，则扣款周期为90天。
	ExecuteTime   string  `json:"execute_time"`             // 首次执行时间execute_time是周期扣款产品必填，即商户发起首次扣款的时间。精确到日，格式为yyyy-MM-dd 结合其他必填的扣款周期参数，会确定商户以后的扣款计划。发起扣款的时间需符合这里的扣款计划。
	SingleAmount  float64 `json:"single_amount"`            // 单次扣款最大金额single_amount是周期扣款产品必填，即每次发起扣款时限制的最大金额，单位为元。商户每次发起扣款都不允许大于此金额。
	TotalAmount   float64 `json:"total_amount,omitempty"`   // 总金额限制，单位为元。如果传入此参数，商户多次扣款的累计金额不允许超过此金额。
	TotalPayments int     `json:"total_payments,omitempty"` // 总扣款次数。如果传入此参数，则商户成功扣款的次数不能超过此次数限制（扣款失败不计入）。
}

type ExtUserInfo struct {
	Name          string `json:"name"`            // 指定买家姓名。 注： need_check_info=T或fix_buyer=T时该参数才有效
	Mobile        string `json:"mobile"`          // 指定买家手机号。 注：该参数暂不校验
	CertType      string `json:"cert_type"`       // 指定买家证件类型。 枚举值：IDENTITY_CARD：身份证；PASSPORT：护照；OFFICER_CARD：军官证；SOLDIER_CARD：士兵证；HOKOU：户口本。如有其它类型需要支持，请与蚂蚁金服工作人员联系。注： need_check_info=T或fix_buyer=T时该参数才有效，支付宝会比较买家在支付宝留存的证件类型与该参数传入的值是否匹配。
	CertNo        string `json:"cert_no"`         // 买家证件号。注：need_check_info=T或fix_buyer=T时该参数才有效，支付宝会比较买家在支付宝留存的证件号码与该参数传入的值是否匹配。
	MinAge        string `json:"min_age"`         // 允许的最小买家年龄。 买家年龄必须大于等于所传数值注：1. need_check_info=T时该参数才有效2. min_age为整数，必须大于等于0
	FixBuyer      string `json:"fix_buyer"`       // 是否强制校验买家身份。 需要强制校验传：T;不需要强制校验传：F或者不传；当传T时，接口上必须指定cert_type、cert_no和name信息且支付宝会校验传入的信息跟支付买家的信息都匹配，否则报错。默认为不校验。
	NeedCheckInfo string `json:"need_check_info"` // 是否强制校验买家信息； 需要强制校验传：T;不需要强制校验传：F或者不传；当传T时，支付宝会校验支付买家的信息与接口上传递的cert_type、cert_no、name或age是否匹配，只有接口传递了信息才会进行对应项的校验；只要有任何一项信息校验不匹配交易都会失败。如果传递了need_check_info，但是没有传任何校验项，则不进行任何校验。默认为不校验。
}
