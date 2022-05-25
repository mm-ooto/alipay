package alipay

import (
	"encoding/json"
	"github.com/mm-ooto/alipay/consts"
	"net/url"
)

// FundTransUniTransferRequest 单笔转账接口
func (a *AliClient) FundTransUniTransferRequest(requestParam FundTransUniTransferRequestParams) (
	responseParam FundTransUniTransferResponseParams, err error) {
	if err = a.HandlerRequest("POST",&requestParam, &responseParam); err != nil {
		return
	}
	return
}

// FundTransUniTransferRequestParams 单笔转账接口请求参数
// 文档地址：https://opendocs.alipay.com/open/02byuo?scene=ca56bca529e64125a2786703c6192d41
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
	urlValue.Add(consts.ApiMethodNameFiled, "alipay.fund.trans.uni.transfer")
	bytes, _ := json.Marshal(f)
	urlValue.Add(consts.BizContentFiled, string(bytes))
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
