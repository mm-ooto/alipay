package alipay

// FundTransUniTransfer 单笔转账接口
func (a *Client) FundTransUniTransfer(requestParam FundTransUniTransferRequestParams) (
	responseParam FundTransUniTransferResponseParams, err error) {
	if err = a.HandlerRequest("POST", &requestParam, &responseParam); err != nil {
		return
	}
	return
}
