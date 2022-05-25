package alipay

import (
	"encoding/json"
	"github.com/mm-ooto/alipay/consts"
	"net/url"
)

// TradePrecreateRequest 统一收单线下交易预创建
func (a *AliClient) TradePrecreateRequest(requestParam TradePrecreateRequestParams) (responseParam TradePrecreateResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

// TradePrecreateRequestParams 统一收单线下交易预创建请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.precreate?scene=common
type TradePrecreateRequestParams struct {
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

func (t *TradePrecreateRequestParams) GetOtherParams() url.Values {
	urlValue := url.Values{}
	urlValue.Add(consts.NotifyUrlFiled, t.NotifyUrl)
	urlValue.Add(consts.AppAuthTokenFiled, t.AppAuthToken)
	urlValue.Add(consts.ApiMethodNameFiled, "alipay.trade.precreate")
	bytes, _ := json.Marshal(t)
	urlValue.Add(consts.BizContentFiled, string(bytes))
	return urlValue
}

func (t *TradePrecreateRequestParams) GetNeedEncrypt() bool {
	return t.NeedEncrypt == true
}

// TradePrecreateResponseParams 统一收单线下交易预创建响应参数
type TradePrecreateResponseParams struct {
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
