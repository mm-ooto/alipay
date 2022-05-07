package alipay

// OpenAuthTokenAppRequest 换取应用授权令牌
func (a *AliClient) OpenAuthTokenAppRequest(requestParam OpenAuthTokenAppRequestParams) (
	responseParam OpenAuthTokenAppResponseParams, err error) {
	requestDataMap := make(map[string]interface{})
	requestDataMap["biz_content"] = a.SetDataToBizContent(requestParam,false)
	if err = a.HandlerRequest("POST", "alipay.open.auth.token.app",false, requestDataMap, &responseParam); err != nil {
		return
	}
	return
}

// OpenAuthTokenAppRequestParams 换取应用授权令牌请求参数
// 文档地址：https://opendocs.alipay.com/apis/api_9/alipay.open.auth.token.app
type OpenAuthTokenAppRequestParams struct {
	GrantType    string `json:"grant_type"`              // 授权方式。支持：1.authorization_code，表示换取使用用户授权码code换取授权令牌access_token。 2.refresh_token，表示使用refresh_token刷新获取新授权令牌。
	Code         string `json:"code,omitempty"`          // 授权码，用户对应用授权后得到。本参数在 grant_type 为 authorization_code 时必填；为 refresh_token 时不填。
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新令牌，上次换取访问令牌时得到。本参数在 grant_type 为 authorization_code 时不填；为 refresh_token 时必填，且该值来源于此接口的返回值 app_refresh_token（即至少需要通过 grant_type=authorization_code 调用此接口一次才能获取）。
}

// OpenAuthTokenAppResponseParams 换取应用授权令牌响应参数
type OpenAuthTokenAppResponseParams struct {
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
