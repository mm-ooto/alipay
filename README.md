# Alipay SDK for Golang

## 说明：支持普通公钥模式和公钥证书模式

## 加签的两种模式（所需要的信息）
### 公钥证书模式
`AppID`、`应用的私钥`、`应用的公钥证书文件`、`支付宝公钥证书文件`、`支付宝根证书文件`

### 公钥模式
`AppId`、`应用的私钥`、`应用的公钥`、`支付宝公钥`

## 使用步骤：
1. 创建一个Client实例
2. 设置request参数并发起API请求，方法（Client.HandlerRequest）

## 初始化
```go
func TestName(t *testing.T) {
    aliClient, err := alipay.NewClient(appId, aliPublicKey, appPrivateKey, "RSA2", false)
        if err != nil {
        fmt.Printf("NewClient error:%s\n", err.Error())
        return
    }
    return
}
```

## 从证书/证书内容中加载相关的证书序列号（certPath，certContent 二选一）
```go
    aliClient.LoadAppCertSN("certPath","certContent")// 加载应用公钥证书序列号SN
    aliClient.LoadAliCertSN("certPath","certContent")// 加载支付宝公钥证书序列号SN
    aliClient.LoadAlipayRootCertSn("certRootPath","certRootContent")// 加载支付宝根证书序列号SN
```

## 参考示例
```go
func TestTradePagePay(t *testing.T) {
    aliClient, err := alipay.NewClient(appId, aliPublicKey, appPrivateKey, "RSA2", false)
        if err != nil {
        fmt.Printf("NewClient error:%s\n", err.Error())
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
    _, urlRe, err := aliClient.TradePagePay(req)
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
    aliClient.AsyncNotify(rawBody, isLifeNotify) // 具体参数含义查看方法说明
```
### 异步通知示例
```go
func main() {
    router := gin.Default()
    // 异步通知回调示例：
	router.POST("/notify", func(c *gin.Context) {
		// 从http body中读取参数字符串
		body, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusInternalServerError, "fail")
			fmt.Printf("GetRawData err:%s\n", err.Error())
			return
		}
		fmt.Printf("rawBody：%s", string(body))
		if _, err = aliClient2.AsyncNotify(string(body)); err != nil {
			c.String(http.StatusInternalServerError, "fail") // 输出fail，表示消息获取失败，支付宝会重新发送消息到异步地址
			return
		}
		c.String(200, "success") // 输出success是表示消息获取成功，支付宝就会停止发送异步
	})
	router.Run(":8080")
}
```



## 目前已实现的接口
* 换取应用授权令牌：alipay.AuthTokenApp()
* 换取授权访问令牌：alipay.SystemOauthToken()
* app支付接口2.0：
* 统一收单下单并支付页面：alipya.TradePagePay()
* 统一收单交易创建：alipay.TradeCreate()
* 统一收单交易撤销：alipay.TradeCancel()
* 统一收单交易关闭：alipay.TradeClose()
* 统一收单交易退款查询：alipay.TradeFastPayRefundQuery()
* 统一收单线下交易预创建：alipay.TradePreCreate()
* 统一收单线下交易查询：alipay.TradeQuery()
* 统一收单交易退款：alipay.TradeRefund()
* 应用支付宝公钥证书下载：alipay.AppAliPayCertDownload()
* 单笔转账：alipay.FundTransUniTransfer()
* 查询对账单下载地址：alipay.TradeBillDownloadUrlQuery()
* 处理异步通知回调：alipay.AsyncNotify()


## 参考文档：
* 沙箱账号：https://open.alipay.com/develop/sandbox/account
* 数据签名和验签：https://opendocs.alipay.com/common/02kf5q ；  https://opendocs.alipay.com/common/02mse2
* 支付API文档：https://opendocs.alipay.com/apis
* 异步通知说明：https://opensupport.alipay.com/support/helpcenter/193/201602472200
* 异步通知参数说明：https://opendocs.alipay.com/open/203/105286
* 授权回调地址：https://opendocs.alipay.com/common/02qjlq
