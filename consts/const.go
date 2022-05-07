package consts

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

	ContentType = "application/x-www-form-urlencoded;charset=utf-8"

	// AlipayCertSn 支付宝公钥证书序列号
	AlipayCertSn = "alipay_cert_sn"
	Sign         = "sign"

	// SuccessCode 接口调用成功时的返回码
	SuccessCode = "10000"

	// EncryptTypeAes 加密类型
	EncryptTypeAes = "AES"
)
