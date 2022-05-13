package alipay

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
