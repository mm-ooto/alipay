# golang alipay

## 说明：支持普通公钥模式和公钥证书模式

## 初始化
```go
func TestName(t *testing.T) {
    aliClient, err := alipay.NewAliClient(appId, aliPublicKey, appPrivateKey, "RSA2", false)
        if err != nil {
        fmt.Printf("NewAliClient error:%s\n", err.Error())
        return
    }
    return
}
```

## 从证书/证书内容中加载相关的证书序列号
```go
    aliClient.LoadAppCertSN("certPath","certContent")// 加载应用公钥证书序列号SN
    aliClient.LoadAliCertSN("certPath","certContent")// 加载支付宝公钥证书序列号SN
    aliClient.LoadAlipayRootCertSn("certRootPath","certRootContent")// 加载支付宝根证书序列号SN
```

## 用法
```go
func TestName(t *testing.T) {
    aliClient, err := alipay.NewAliClient(appId, aliPublicKey, appPrivateKey, "RSA2", false)
        if err != nil {
        fmt.Printf("NewAliClient error:%s\n", err.Error())
        return
    }
    req := alipay.TradePagePayRequestParams{
            OtherRequestParams:alipay.OtherRequestParams{
                NeedEncrypt:  false,
                ReturnUrl:    "",
                NotifyUrl:    "",
                AppAuthToken: "",
            },
            NotifyUrl:   "",
            OutTradeNo:  "20220817010101004",
            TotalAmount: "0.01",
            Subject:     "统一收单下单并支付页面接口",
    }
    _, urlRe, err := aliClient.TradePagePayRequest(req)
    if err != nil {
        t.Log(err)
        return
    }
    t.Log(urlRe)
    bytes, _ := json.Marshal(urlRe)
    t.Log("返回值：", string(bytes))
}
```

## 异步通知
### 异步通知方法
```go
    aliClient.HandlerAsyncNotify(rawBody, isLifeIsNo) // 具体参数含义查看方法说明
```
### 异步通知示例
```go
    router := gin.Default()
	router.POST("/notify", func(c *gin.Context) {
		// 从http body中读取参数字符串
		body, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusInternalServerError, "fail")
			fmt.Printf("GetRawData err:%s\n", err.Error())
			return
		}
		fmt.Printf("rawBody：%s", string(body))
		if _, err = aliClient2.HandlerAsyncNotify(string(body), false); err != nil {
			c.String(http.StatusInternalServerError, "fail") // 输出fail，表示消息获取失败，支付宝会重新发送消息到异步地址
			return
		}
		c.String(200, "success") // 输出success是表示消息获取成功，支付宝就会停止发送异步
	})
	router.Run(":8080")
```


### 参考文档：
* 沙箱账号：https://open.alipay.com/develop/sandbox/account
* 数据签名和验签：https://opendocs.alipay.com/common/02kf5q ；  https://opendocs.alipay.com/common/02mse2
* 支付API文档：https://opendocs.alipay.com/apis
* 异步通知说明：https://opensupport.alipay.com/support/helpcenter/193/201602472200
* 异步通知参数说明：https://opendocs.alipay.com/open/203/105286
* 授权回调地址：https://opendocs.alipay.com/common/02qjlq
