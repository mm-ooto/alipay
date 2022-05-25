package alipay

import (
	"github.com/mm-ooto/alipay/consts"
	"net/url"
)

// SystemOauthTokenRequest 换取授权访问令牌
func (a *AliClient) SystemOauthTokenRequest(requestParam SystemOauthTokenRequestParams) (
	responseParam SystemOauthTokenResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}

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
	urlValue.Add(consts.AppAuthTokenFiled, s.AppAuthToken)
	urlValue.Add(consts.ApiMethodNameFiled, "alipay.system.oauth.token")
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
