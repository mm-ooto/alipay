package alipay

const (
	// SignTypeRSA 商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2
	SignTypeRSA  = "RSA"
	SignTypeRSA2 = "RSA2"

	// ApiVersion 版本号
	ApiVersion = "1.0"

	// FormatJson 请求格式 如utf-8,gbk,gb2312等
	FormatJson = "JSON" // Json格式

	// CharSetUTF8 编码格式
	CharSetUTF8 = "UTF-8" // UTF8

	// RequestTimestampFormat 发送请求的时间的格式
	RequestTimestampFormat = "2006-01-02 15:04:05"

	// GateWalProdUrl (生产环境) 支付宝接口地址
	GateWalProdUrl = "https://openapi.alipay.com/gateway.do"
	// GateWalSandboxUrl (沙盒环境) 支付宝接口地址
	GateWalSandboxUrl = "https://openapi.alipaydev.com/gateway.do"

	ContentTypeFromUrlEncoded = "application/x-www-form-urlencoded;charset=utf-8"

	// AlipayCertSnField 支付宝公钥证书序列号
	AlipayCertSnField = "alipay_cert_sn"
	// SignFiled 签名字段
	SignFiled = "sign"
	// SignTypeFiled 签名类型字段
	SignTypeFiled = "sign_type"
	// ResponseSuffix 响应后缀
	ResponseSuffix = "_response"
	// ErrorResponse 错误响应
	ErrorResponse = "error_response"

	// SuccessCode 接口调用成功时的返回码
	SuccessCode = "10000"

	// EncryptTypeAes 加密类型
	EncryptTypeAes = "AES"

	// CertificatePrefix 证书前后缀标识
	CertificatePrefix = "-----BEGIN CERTIFICATE-----"
	CertificateSuffix = "-----END CERTIFICATE-----"

	// RSAPrivatePrefix rsa私钥前后缀标识
	RSAPrivatePrefix = "-----BEGIN RSA PRIVATE KEY-----"
	RSAPrivateSuffix = "-----END RSA PRIVATE KEY-----"

	// PublicKeyPrefix 公钥前后缀标识
	PublicKeyPrefix = "-----BEGIN PUBLIC KEY-----"
	PublicKeySuffix = "-----END PUBLIC KEY-----"
)

const (
	// ReturnUrlFiled 同步跳转地址字段
	ReturnUrlFiled = "return_url"
	// NotifyUrlFiled 异步通知地址字段
	NotifyUrlFiled = "notify_url"
	// BizContentFiled 接口请求参数集合字段
	BizContentFiled = "biz_content"
	// AppAuthTokenFiled 获取应用授权token字段，详见应用授权概述：https://opendocs.alipay.com/isv/10467/xldcyq
	AppAuthTokenFiled = "app_auth_token"
	// ApiMethodNameFiled 接口名称字段
	ApiMethodNameFiled = "method"
	// EncryptTypeField 加密类型字段
	EncryptTypeField = "encrypt_type"
)
