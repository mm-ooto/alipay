# golang alipay

## 说明：目前只实现了普通公钥模式，公钥证书模式还未完全实现

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


### 参考文档：
* 数据签名和验签：https://opendocs.alipay.com/common/02kf5q
* 支付API文档：https://opendocs.alipay.com/apis
* 沙箱账号：https://open.alipay.com/develop/sandbox/account
