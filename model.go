package alipay

import (
	"encoding/json"
	"net/url"
)

type RequestParams interface {
	// GetOtherParams 除公共参数以外的其它请求参数,如：returnUrl,notifyUrl,appAuthToken,apiMethodName,bizContent
	GetOtherParams() url.Values
	// GetNeedEncrypt 是否需要对biz_content内容加密，加密算法为AES
	GetNeedEncrypt() bool
}

// OtherRequestParams 其它特殊的请求参数
type OtherRequestParams struct {
	NeedEncrypt  bool   `json:"-"` // 是否需要对内容biz_content进行加密
	ReturnUrl    string `json:"-"`
	NotifyUrl    string `json:"-"` // 支付宝服务器主动通知商户服务器里指定的页面http/https路径，例如：http://api.test.alipay.net/atinterface/receive_notify.htm
	AppAuthToken string `json:"-"` // 详见应用授权概述：https://opendocs.alipay.com/isv/10467/xldcyq
}

// CommonResParams 公共响应参数
type CommonResParams struct {
	Code    string `json:"code"`     // 网关返回码,详见文档：https://doc.open.alipay.com/docs/doc.htm?treeId=291&articleId=105806&docType=1
	Msg     string `json:"msg"`      // 网关返回码描述,详见文档：https://doc.open.alipay.com/docs/doc.htm?treeId=291&articleId=105806&docType=1
	SubCode string `json:"sub_code"` // 业务返回码，参见具体的API接口文档
	SubMsg  string `json:"sub_msg"`  // 业务返回码描述，参见具体的API接口文档
}

// CommonReqParams 公共请求参数
type CommonReqParams struct {
	AppId     string `json:"app_id"`           // 支付宝分配给开发者的应用ID
	Method    string `json:"method"`           // 接口名称，例如：alipay.trade.precreate
	Format    string `json:"format,omitempty"` // 格式，仅支持JSON
	Charset   string `json:"charset"`          // 请求使用的编码格式，如utf-8,gbk,gb2312等
	SignType  string `json:"sign_type"`        // 商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2
	Sign      string `json:"sign"`             // 商户请求参数的签名串，详见：https://opendocs.alipay.com/open/291/105974
	Timestamp string `json:"timestamp"`        // 发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version   string `json:"version"`          // 调用的接口版本，固定为：1.0
	OtherRequestParams
	AppCertSn        string `json:"app_cert_sn,omitempty"`         // 应用公钥证书 SN（如果使用证书签名则需要带入），注意：如果使用公钥证书签名则需要在请求参数中将"appCertSN"和"alipayRootCertSn"传入，SN 值是通过解析 X.509 证书文件中签发机构名称（name）以及内置序列号（serialNumber），将二者拼接后的字符串计算 MD5 值获取，可参考开放平台 SDK 源码
	AlipayRootCertSn string `json:"alipay_root_cert_sn,omitempty"` // 支付宝根证书 SN（如果使用证书签名则需要带入）

	BizContent string `json:"biz_content,omitempty"` // 请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
}

///////////////////////////////////////////////////////////////////////////////////////

// SystemOauthTokenRequestParams 换取授权访问令牌请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_9/alipay.system.oauth.token
type SystemOauthTokenRequestParams struct {
	OtherRequestParams
	GrantType    string `json:"grant_type"`              // 授权方式。支持：1.authorization_code，表示换取使用用户授权码code换取授权令牌access_token。 2.refresh_token，表示使用refresh_token刷新获取新授权令牌。
	Code         string `json:"code,omitempty"`          // 授权码，用户对应用授权后得到。本参数在 grant_type 为 authorization_code 时必填；为 refresh_token 时不填。
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新令牌，上次换取访问令牌时得到。本参数在 grant_type 为 authorization_code 时不填；为 refresh_token 时必填，且该值来源于此接口的返回值 app_refresh_token（即至少需要通过 grant_type=authorization_code 调用此接口一次才能获取）。
}

func (s *SystemOauthTokenRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(AppAuthTokenFiled, s.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.system.oauth.token")
	return urlValue
}

func (s *SystemOauthTokenRequestParams) GetNeedEncrypt() bool {
	return s.NeedEncrypt == true
}

// SystemOauthTokenResponseParams 换取授权访问令牌响应参数
type SystemOauthTokenResponseParams struct {
	Data struct {
		CommonResParams
		UserId       string `json:"user_id"`       // 支付宝用户的唯一标识。以2088开头的16位数字。
		AccessToken  string `json:"access_token"`  // 访问令牌。通过该令牌调用需要授权类接口
		ExpiresIn    string `json:"expires_in"`    // 访问令牌的有效时间，单位是秒。
		RefreshToken string `json:"refresh_token"` // 刷新令牌。通过该令牌可以刷新access_token
		ReExpiresIn  string `json:"re_expires_in"` // 刷新令牌的有效时间，单位是秒。
		AuthStart    string `json:"auth_start"`    // 授权token开始时间，作为有效期计算的起点
	} `json:"alipay_system_oauth_token_response"`
	Sign string `json:"sign"` // 签名
}

///////////////////////////////////////////////////////////////////////////////////////

// AuthTokenAppRequestParams 换取应用授权令牌请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_9/alipay.open.auth.token.app
type AuthTokenAppRequestParams struct {
	OtherRequestParams
	GrantType    string `json:"grant_type"`              // 授权方式。支持：1.authorization_code，表示换取使用用户授权码code换取授权令牌access_token。 2.refresh_token，表示使用refresh_token刷新获取新授权令牌。
	Code         string `json:"code,omitempty"`          // 授权码，用户对应用授权后得到。本参数在 grant_type 为 authorization_code 时必填；为 refresh_token 时不填。
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新令牌，上次换取访问令牌时得到。本参数在 grant_type 为 authorization_code 时不填；为 refresh_token 时必填，且该值来源于此接口的返回值 app_refresh_token（即至少需要通过 grant_type=authorization_code 调用此接口一次才能获取）。
}

func (o *AuthTokenAppRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(NotifyUrlFiled, o.NotifyUrl)
	urlValue.Add(AppAuthTokenFiled, o.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.open.auth.token.app")
	bytes, _ := json.Marshal(o)
	urlValue.Add(BizContentFiled, string(bytes))
	return urlValue
}

func (o *AuthTokenAppRequestParams) GetNeedEncrypt() bool {
	return o.NeedEncrypt == true
}

// AuthTokenAppResponseParams 换取应用授权令牌响应参数
type AuthTokenAppResponseParams struct {
	Data struct {
		CommonResParams
		UserId       string `json:"user_id"`        // 授权商户的user_id
		AppAuthId    string `json:"app_auth_id"`    // 授权商户的appid
		AppAuthToken string `json:"app_auth_token"` // 应用授权令牌
		RefreshToken string `json:"refresh_token"`  // 刷新令牌。通过该令牌可以刷新access_token
		ExpiresIn    string `json:"expires_in"`     // 访问令牌的有效时间，单位是秒。
		ReExpiresIn  string `json:"re_expires_in"`  // 刷新令牌的有效时间，单位是秒。
	} `json:"alipay.open.auth.token.app_response"`
	Sign string `json:"sign"` // 签名
}

///////////////////////////////////////////////////////////////////////////////////////

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
	urlValue.Add(NotifyUrlFiled, t.NotifyUrl)
	urlValue.Add(AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.trade.create")
	bytes, _ := json.Marshal(t)
	urlValue.Add(BizContentFiled, string(bytes))
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

///////////////////////////////////////////////////////////////////////////////////////

// TradeCancelRequestParams 统一收单交易撤销接口请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.cancel
type TradeCancelRequestParams struct {
	OtherRequestParams

	OutTradeNo string `json:"out_trade_no,omitempty"` // 原支付请求的商户订单号,和支付宝交易号不能同时为空
	TradeNo    string `json:"trade_no,omitempty"`     // 支付宝交易号，和商户订单号不能同时为空
}

func (t *TradeCancelRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.trade.cancel")
	bytes, _ := json.Marshal(t)
	urlValue.Add(BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradeCancelRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
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

///////////////////////////////////////////////////////////////////////////////////////

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
	urlValue.Add(NotifyUrlFiled, t.NotifyUrl)
	urlValue.Add(AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.trade.close")
	bytes, _ := json.Marshal(t)
	urlValue.Add(BizContentFiled, string(bytes))
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

///////////////////////////////////////////////////////////////////////////////////////

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
	urlValue.Add(AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.trade.query")
	bytes, _ := json.Marshal(t)
	urlValue.Add(BizContentFiled, string(bytes))
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

///////////////////////////////////////////////////////////////////////////////////////

// TradePreCreateRequestParams 统一收单线下交易预创建请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.precreate?scene=common
type TradePreCreateRequestParams struct {
	OtherRequestParams

	OutTradeNo         string               `json:"out_trade_no"`                  // 商户订单号。由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	TotalAmount        float64              `json:"total_amount"`                  // 订单总金额，单位为元，精确到小数点后两位，取值范围为 [0.01,100000000]，金额不能为 0。如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】
	Subject            string               `json:"subject"`                       // 订单标题。 注意：不可使用特殊字符，如 /，=，& 等。
	ProductCode        string               `json:"product_code"`                  // 销售产品码。如果签约的是当面付快捷版，则传 OFFLINE_PAYMENT；其它支付宝当面付产品传 FACE_TO_FACE_PAYMENT；不传则默认使用 FACE_TO_FACE_PAYMENT。
	SellerId           string               `json:"seller_id,omitempty"`           // 卖家支付宝用户 ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户 ID。不允许收款账号与付款方账号相同
	Body               string               `json:"body,omitempty"`                // 订单附加信息。	如果请求时传递了该参数，将在异步通知、对账单中原样返回，同时会在商户和用户的pc账单详情中作为交易描述展示
	GoodsDetail        []*GoodsDetailParams `json:"goods_detail,omitempty"`        // 订单包含的商品列表信息，为 JSON 格式，其它说明详见商品明细说明
	ExtendParams       *ExtendParamsParams  `json:"extend_params,omitempty"`       // 业务扩展参数
	DiscountableAmount float64              `json:"discountable_amount,omitempty"` // 可打折金额。参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围为 [0.01,100000000]。如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
	StoreId            string               `json:"store_id,omitempty"`            // 商户门店编号。 指商户创建门店时输入的门店编号。
	OperatorId         string               `json:"operator_id,omitempty"`         // 商户操作员编号。
	TerminalId         string               `json:"terminal_id,omitempty"`         // 商户机具终端编号。
	MerchantOrderNo    string               `json:"merchant_order_no,omitempty"`   // 商户原始订单号，最大长度限制 32 位
}

func (t *TradePreCreateRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(NotifyUrlFiled, t.NotifyUrl)
	urlValue.Add(AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.trade.precreate")
	bytes, _ := json.Marshal(t)
	urlValue.Add(BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradePreCreateRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
}

// TradePreCreateResponseParams 统一收单线下交易预创建响应参数
type TradePreCreateResponseParams struct {
	Data struct {
		CommonResParams
		OutTradeNo string `json:"out_trade_no"` // 商户的订单号
		QrCode     string `json:"qr_code"`      // 当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码，例如：https://qr.alipay.com/bavh4wjlxf12tper3a
	} `json:"alipay_trade_precreate_response"`
	Sign string `json:"sign"` // 签名
}

// GoodsDetailParams 订单包含的商品列表信息
type GoodsDetailParams struct {
	GoodsId        string  `json:"goods_id"`                  // 商品的编号
	GoodsName      string  `json:"goods_name"`                // 商品名称
	Quantity       int     `json:"quantity"`                  // 商品数量
	Price          float64 `json:"price"`                     // 商品单价，单位为元
	GoodsCategory  string  `json:"goods_category,omitempty"`  // 商品类目
	CategoriesTree string  `json:"categories_tree,omitempty"` // 商品类目树，从商品类目根节点到叶子节点的类目id组成，类目id值使用|分割
	ShowUrl        string  `json:"show_url,omitempty"`        // 商品的展示地址
}

// ExtendParamsParams 业务扩展参数
type ExtendParamsParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty"` // 系统商编号 该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID
	CardType             string `json:"card_type,omitempty"`               // 卡类型
	SpecifiedSellerName  string `json:"specified_seller_name,omitempty"`   // 特殊场景下，允许商户指定交易展示的卖家名称
}

///////////////////////////////////////////////////////////////////////////////////////

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
	urlValue.Add(AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.trade.refund")
	bytes, _ := json.Marshal(t)
	urlValue.Add(BizContentFiled, string(bytes))
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

///////////////////////////////////////////////////////////////////////////////////////

// TradeAppPayRequestParams app支付接口2.0请求参数 ,omitempty
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.app.pay
type TradeAppPayRequestParams struct {
	OtherRequestParams

	OutTradeNo          string                `json:"out_trade_no"`                    // 商户订单号。由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	TotalAmount         string                `json:"total_amount"`                    // 订单总金额。单位为元，精确到小数点后两位，取值范围：[0.01,100000000] 。
	Subject             string                `json:"subject"`                         // 订单标题。 注意：不可使用特殊字符，如 /，=，& 等。
	ProductCode         string                `json:"product_code"`                    // 产品码。 商家和支付宝签约的产品码。 枚举值（点击查看签约情况）：QUICK_MSECURITY_PAY：无线快捷支付产品；CYCLE_PAY_AUTH：周期扣款产品。默认值为QUICK_MSECURITY_PAY。
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

func (t *TradeAppPayRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(NotifyUrlFiled, t.NotifyUrl)
	urlValue.Add(AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.trade.app.pay")
	bytes, _ := json.Marshal(t)
	urlValue.Add(BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradeAppPayRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
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

///////////////////////////////////////////////////////////////////////////////////////

// TradePagePayRequestParams 统一收单下单并支付页面接口请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay
type TradePagePayRequestParams struct {
	OtherRequestParams

	OutTradeNo          string               `json:"out_trade_no"`                    // 商户订单号。由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	TotalAmount         string               `json:"total_amount"`                    // 订单总金额。单位为元，精确到小数点后两位，取值范围：[0.01,100000000] 。
	Subject             string               `json:"subject"`                         // 订单标题。 注意：不可使用特殊字符，如 /，=，& 等。
	ProductCode         string               `json:"product_code"`                    // 产品码。 商家和支付宝签约的产品码。 枚举值（点击查看签约情况）：QUICK_MSECURITY_PAY：无线快捷支付产品；CYCLE_PAY_AUTH：周期扣款产品。默认值为QUICK_MSECURITY_PAY。
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

func (t *TradePagePayRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(ReturnUrlFiled, t.ReturnUrl)
	urlValue.Add(NotifyUrlFiled, t.NotifyUrl)
	urlValue.Add(AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.trade.page.pay")
	bytes, _ := json.Marshal(t)
	urlValue.Add(BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradePagePayRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
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

///////////////////////////////////////////////////////////////////////////////////////

// TradeFastPayRefundQueryRequestParams 统一收单交易退款查询请求参数
// 文档地址：https://opendocs.alipay.com/apis/0287wc
type TradeFastPayRefundQueryRequestParams struct {
	OtherRequestParams

	TradeNo      string    `json:"trade_no,omitempty"`      // 支付宝交易号。 和商户订单号不能同时为空
	OutTradeNo   string    `json:"out_trade_no,omitempty"`  // 商户订单号。 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
	OutRequestNo string    `json:"out_request_no"`          // 退款请求号。 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的商户订单号。
	QueryOptions []*string `json:"query_options,omitempty"` // 查询选项
}

func (t *TradeFastPayRefundQueryRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.trade.fastpay.refund.query")
	bytes, _ := json.Marshal(t)
	urlValue.Add(BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradeFastPayRefundQueryRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
}

// TradeFastPayRefundQueryResponseParams 统一收单交易退款查询响应参数
type TradeFastPayRefundQueryResponseParams struct {
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

///////////////////////////////////////////////////////////////////////////////////////

// TradeBillDownloadUrlQueryRequestParams 查询对账单下载地址接口请求参数
// 文档地址：https://docs.open.alipay.com/api_15/alipay.data.dataservice.bill.downloadurl.query
type TradeBillDownloadUrlQueryRequestParams struct {
	OtherRequestParams

	BillType string `json:"bill_type"` // 必选 账单类型，商户通过接口或商户经开放平台授权后其所属服务商通过接口可以获取以下账单类型：trade、signcustomer；trade指商户基于支付宝交易收单的业务账单；signcustomer是指基于商户支付宝余额收入及支出等资金变动的帐务账单。
	BillDate string `json:"bill_date"` // 必选 账单时间：日账单格式为yyyy-MM-dd，最早可下载2016年1月1日开始的日账单；月账单格式为yyyy-MM，最早可下载2016年1月开始的月账单。
}

func (t *TradeBillDownloadUrlQueryRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.data.dataservice.bill.downloadurl.query")
	bytes, _ := json.Marshal(t)
	urlValue.Add(BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradeBillDownloadUrlQueryRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
}

// TradeBillDownloadUrlQueryResponseParams 查询对账单下载地址接口响应参数
type TradeBillDownloadUrlQueryResponseParams struct {
	Data struct {
		CommonResParams
		BillDownloadUrl string `json:"bill_download_url"` // 账单下载地址链接，获取连接后30秒后未下载，链接地址失效。
	} `json:"alipay_data_dataservice_bill_downloadurl_query_response"`
	Sign string `json:"sign"` // 签名

}

///////////////////////////////////////////////////////////////////////////////////////

// FundTransUniTransferRequestParams 单笔转账接口请求参数
// 文档地址：https://opendocs.alipay.com/open/02byuo
type FundTransUniTransferRequestParams struct {
	OutBizNo       string       `json:"out_biz_no"`       // 商家侧唯一订单号，由商家自定义。对于不同转账请求，商家需保证该订单号在自身系统唯一。
	TransAmount    float64      `json:"trans_amount"`     // 订单总金额，单位为元，不支持千位分隔符，精确到小数点后两位，取值范围[0.1,100000000]。
	ProductCode    string       `json:"product_code"`     // 销售产品码。单笔无密转账固定为 TRANS_ACCOUNT_NO_PWD。
	BizScene       string       `json:"biz_scene"`        // 业务场景。单笔无密转账固定为 DIRECT_TRANSFER。
	OrderTitle     string       `json:"order_title"`      // 转账业务的标题，用于在支付宝用户的账单里显示。
	PayeeInfo      *Participant `json:"payee_info"`       // 收款方信息
	Remark         string       `json:"remark,omitempty"` // 业务备注。
	BusinessParams string       `json:"business_params"`  // 转账业务请求的扩展参数，支持传入的扩展参数如下： payer_show_name_use_alias：是否展示付款方别名，可选，收款方在支付宝账单中可见。枚举支持：* true：展示别名，将展示商家支付宝在商家中心 商户信息 > 商户基本信息 页面配置的 商户别名。* false：不展示别名。默认为 false。
}

func (f *FundTransUniTransferRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(ApiMethodNameFiled, "alipay.fund.trans.uni.transfer")
	bytes, _ := json.Marshal(f)
	urlValue.Add(BizContentFiled, string(bytes))
	return urlValue
}

func (f *FundTransUniTransferRequestParams) GetNeedEncrypt() bool {
	return false
}

// Participant 收款方信息
type Participant struct {
	Identity     string `json:"identity"`       // 参与方的标识 ID。 当 identity_type=ALIPAY_USER_ID 时，填写支付宝用户 UID。示例值：2088123412341234。当 identity_type=ALIPAY_LOGON_ID 时，填写支付宝登录号。示例值：186xxxxxxxx。
	IdentityType string `json:"identity_type"`  // 参与方的标识类型，目前支持如下枚举： ALIPAY_USER_ID：支付宝会员的用户 ID，可通过 获取会员信息 能力获取。ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式。
	Name         string `json:"name,omitempty"` // 参与方真实姓名。如果非空，将校验收款支付宝账号姓名一致性。 当 identity_type=ALIPAY_LOGON_ID 时，本字段必填。若传入该属性，则在支付宝回单中将会显示这个属性。
}

// FundTransUniTransferResponseParams 单笔转账接口响应参数
type FundTransUniTransferResponseParams struct {
	Data struct {
		CommonResParams
		OutBizNo       string `json:"out_biz_no"`        // 商家订单号
		OrderId        string `json:"order_id"`          // 支付宝转账订单号
		PayFundOrderId string `json:"pay_fund_order_id"` // 支付宝支付资金流水号
		Status         string `json:"status"`            // 转账单据状态
		TransDate      string `json:"trans_date"`        // 订单支付时间，格式为yyyy-MM-dd HH:mm:ss
	} `json:"alipay_fund_trans_uni_transfer_response"`
	Sign string `json:"sign"` // 签名
}

///////////////////////////////////////////////////////////////////////////////////////

// AppAliPayCertDownloadRequestParams 应用支付宝公钥证书下载请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_9/alipay.open.app.alipaycert.download
type AppAliPayCertDownloadRequestParams struct {
	OtherRequestParams
	AlipayCertSn string `json:"alipay_cert_sn"` // 支付宝公钥证书序列号
}

func (o *AppAliPayCertDownloadRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(NotifyUrlFiled, o.NotifyUrl)
	urlValue.Add(AppAuthTokenFiled, o.AppAuthToken)
	urlValue.Add(ApiMethodNameFiled, "alipay.open.app.alipaycert.download")
	bytes, _ := json.Marshal(o)
	urlValue.Add(BizContentFiled, string(bytes))
	return urlValue
}

func (o *AppAliPayCertDownloadRequestParams) GetNeedEncrypt() bool {
	return o.NeedEncrypt == true
}

// AppAliPayCertDownloadResponseParams 应用支付宝公钥证书下载响应参数
type AppAliPayCertDownloadResponseParams struct {
	Data struct {
		CommonResParams
		AlipayCertContent string `json:"alipay_cert_content"` // 公钥证书Base64后的字符串
	} `json:"alipay_open_app_alipaycert_download_response"`
	Sign string `json:"sign"` // 签名
}

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////

// TradeNotificationParams 异步通知响应参数
// 文档：https://opendocs.alipay.com/open/203/105286
type TradeNotificationParams struct {
	AuthAppId           string `json:"auth_app_id"`           // App Id
	NotifyTime          string `json:"notify_time"`           // 通知时间
	NotifyType          string `json:"notify_type"`           // 通知类型
	NotifyId            string `json:"notify_id"`             // 通知校验ID
	AppId               string `json:"app_id"`                // 开发者的app_id
	Charset             string `json:"charset"`               // 编码格式
	Version             string `json:"version"`               // 接口版本
	SignType            string `json:"sign_type"`             // 签名类型
	Sign                string `json:"sign"`                  // 签名
	TradeNo             string `json:"trade_no"`              // 支付宝交易号
	OutTradeNo          string `json:"out_trade_no"`          // 商户订单号
	OutBizNo            string `json:"out_biz_no"`            // 商户业务号
	BuyerId             string `json:"buyer_id"`              // 买家支付宝用户号
	BuyerLogonId        string `json:"buyer_logon_id"`        // 买家支付宝账号
	SellerId            string `json:"seller_id"`             // 卖家支付宝用户号
	SellerEmail         string `json:"seller_email"`          // 卖家支付宝账号
	TradeStatus         string `json:"trade_status"`          // 交易状态
	TotalAmount         string `json:"total_amount"`          // 订单金额
	ReceiptAmount       string `json:"receipt_amount"`        // 实收金额
	InvoiceAmount       string `json:"invoice_amount"`        // 开票金额
	BuyerPayAmount      string `json:"buyer_pay_amount"`      // 付款金额
	PointAmount         string `json:"point_amount"`          // 集分宝金额
	RefundFee           string `json:"refund_fee"`            // 总退款金额
	Subject             string `json:"subject"`               // 商品的标题/交易标题/订单标题/订单关键字等，是请求时对应的参数，原样通知回来。
	Body                string `json:"body"`                  // 商品描述
	GmtCreate           string `json:"gmt_create"`            // 交易创建时间
	GmtPayment          string `json:"gmt_payment"`           // 交易付款时间
	GmtRefund           string `json:"gmt_refund"`            // 交易退款时间
	GmtClose            string `json:"gmt_close"`             // 交易结束时间
	FundBillList        string `json:"fund_bill_list"`        // 支付金额信息
	PassbackParams      string `json:"passback_params"`       // 回传参数
	VoucherDetailList   string `json:"voucher_detail_list"`   // 优惠券信息
	AgreementNo         string `json:"agreement_no"`          //支付宝签约号
	ExternalAgreementNo string `json:"external_agreement_no"` // 商户自定义签约号
}
