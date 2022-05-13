package main

import (
	"encoding/json"
	"fmt"
	"github.com/mm-ooto/alipay"
	"testing"
	"time"
)

var aliClient *alipay.AliClient
var encryptKey string

func TestMain(m *testing.M) {
	var err error
	aliClient, err = NewAliClient()
	if err != nil {
		fmt.Printf("TestMain NewClient err:%s", err.Error())
		return
	}
	m.Run()
}

func NewAliClient() (aliClient *alipay.AliClient, err error) {
	appId := "2016091200490539"
	appPrivateKey := "MIIEpAIBAAKCAQEAvCqpPIlU/VM78FXnVWoO/iPbiIg/TSdZO2UmKdd+sLtVRqhy+doS08pzw/Qq5XUSL4QwVhoRShoQmGKtAAxI2YshCnk61WdHZogOlmniQlundZmflRLQgNrEnGftcuNO9xMAgwXtke1akgXDHjpkyGhl8M6yJ2cwou9wH1x2VDsW0cFVdgOT7uQD2W81CCGkPeMu5XfxKbBoKRE8OPNMsKAE2Bd6dPStmg8w3E2ZCGo8xVLH+vdSyGX30iYZHFP65WFoQt4ANJU/OgaWrIOcGdskFGTEtGeKEDoPw27LY/VD4I+2KolGalv8qcSZME+3FrmlHL+7I2ZgVKsVUHqT0QIDAQABAoIBAHmuqOSR9tkfU1qXYtMUk/97FsPTQARX1teXELfsOGx3qKzZ0AiNIrG9cWGd64OZUppRxKRZlSazdlnlLfUi/JVZ6JMKVKaedEj04WIZtQyukrt1DgLsONOrJYvzlVU/c9hJfII+eiRtNq3JdiV9I6GKCapRMFpU29nyNzLAq3DJ7PA6IGfSy/vk/r5HbxxJoomlZ1BZjhDezLn2nPXFqLQunQMeBozA/Eh/v7GZ0tZrhX/LEmq1aqfXtB+4mBX//gDMwCcvBlCLeBsVDXQOzPvCXr7qKTIFdXiQZhlopEt3RFEQb6CoQF5XSHisq3wbLxK8wTa01+0jcknJZGHYYCUCgYEA6vfKnukjge2J7ptmIegu5j+vZmbAcnKBJ1C8xrcb4VyGAFGWotUcUMTa8lKrjYXAhKnj/FeKRu19VhfIjEGEno/7zLoQz9j8QLlQAykIBoP9TTXI380K2Kns8uYdSUNCtJTE8CmqOpHJrKVcUy/lxREfv8SmHFPKD9jTC9t7ZKsCgYEAzQJvLgrp9zrYTmbAxKxXeswiJsBlMo4atlIs+R0xH3xVyWQe8Ng2GqH4BqSvXIqdthTq8YNmid6H42KRvaVVX7k5Bfkx6Qdghi+YWkedSKB1mKteUsEPzRtjcLUAABgu4SFGX/YuMCCZkXpNx3McnVRNmW5dY+qxlhSQ+3EXEXMCgYEAnnu8IytFU+GQY2xVmxEscQkLmZo8u/UXwBjo+2+OUpdBmv1tCS+NBb2BoGi6ZZ6Nl+2vZQj2r5iILYWlM1UNypV7VT87D7Zfjphvq3IFg7+LHoTklG+MnU8gD0W/Aydm2r5thz/THeYvjU+L0mBALoe6TnKpR/oMFFw/HYRQ2jkCgYEAmbLPg+duzYnyjaT/tPO4ijntCLyJokNjx3kIeqPmJkLjVh+YCt0ugv0XpHNnfav23YIFOphXEdoiatmFhncj8KY/GDlhr+F1/mREhrrWMpMKVzFzf/t6Sz3TabZpj6iRzPtTdbJtomtuduEI2xV0SIfhvbw+jCByj6BPqhN5Rf0CgYAVth2Yg6oTblLTXxsEYFifo0AsTxPVwVT+Ow+W+MeAYsRvEBnyE6pLIcUhzj/yOrXSTOsYdmP6YoghJT/mh/rp4kHXFuiqcne5WRhdzhRJFRTwF6gRQMcJsZbB6JcICv6h2AMUCFHUARhJTLHBSxUlvUHT1hDe09aZ6r964qd6NQ=="
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAimaDmFEr8HxMkqUFrd94aGBp2hjpfsOqVWdepqER1kSFv1cwmusOY6wAXhf5GHuKuz9MwjP1LZRr3PmRM/OwD9BFf7t/CKRBn7LFIwCWQ+yN8IWCWcNxSZfAkMRIncA5lAY1r4iqaG++ze7KDLjQD1xI2UaFhUBwKiBoTljC6hfmAYzUPObox+aQq5LUYY98O8DPzZbjDhHc5wq+KqaRfKefcqFkx/6Vy/GGgUjIA8tydN0MOUzZMgu4i8Z4E2H3R98Mt0oEJEFQksZ3NQxqaulO7kl4H/gU97KGEKfjmd2w1opO7FKgihtD2DYcXgtO8iRBhMFz5owc4k2wvKFUCQIDAQAB"
	encryptKey = "3fj+6fTqw4xNl1Bvf91vAg=="

	aliClient, err = alipay.NewAliClient(appId, aliPublicKey, appPrivateKey, "RSA2", false)
	if err != nil {
		fmt.Printf("NewAliClient error:%s\n", err.Error())
		return
	}
	return
}
func TestAlipayTradePrecreate(t *testing.T) {
	req := alipay.TradePrecreateRequestParams{
		OtherRequestParams: alipay.OtherRequestParams{
			NeedEncrypt:  false,
			ReturnUrl:    "",
			NotifyUrl:    "",
			AppAuthToken: "",
		},
		OutTradeNo:  fmt.Sprintf("%d",time.Now().UnixNano()),
		TotalAmount: 100,
		Subject:     "统一收单线下交易预创建",
	}
	aliClient.AddEncryptKey(encryptKey)
	res, err := aliClient.TradePrecreateRequest(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log("返回值：", string(bytes))

}

func TestTradeAppPay(t *testing.T) {
	req := alipay.TradeAppPayRequestParams{
		OutTradeNo:  fmt.Sprintf("%d",time.Now().UnixNano()),
		TotalAmount: "0.01",
		Subject:     "app支付接口2.0",
	}
	res, err := aliClient.TradeAppPayRequest(req)
	if err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log("返回值：", string(bytes))
}

func TestTradePagePayRequest(t *testing.T) {
	req := alipay.TradePagePayRequestParams{
		OtherRequestParams: alipay.OtherRequestParams{
			NeedEncrypt:  false,
			ReturnUrl:    "",
			NotifyUrl:    "",
			AppAuthToken: "",
		},
		OutTradeNo:  fmt.Sprintf("%d",time.Now().UnixNano()),
		TotalAmount: "100",
		Subject:     "统一收单下单并支付页面接口",
		ProductCode: "FAST_INSTANT_TRADE_PAY",
	}
	html, urlRe, err := aliClient.TradePagePayRequest(req)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(urlRe)
	t.Log("POST 返回值：", html)
}

func TestFundTransUniTransfer(t *testing.T) {
	req := alipay.FundTransUniTransferRequestParams{
		OutBizNo:   fmt.Sprintf("%d",time.Now().UnixNano()),
		TransAmount: 0.01,
		ProductCode: "TRANS_ACCOUNT_NO_PWD",
		BizScene:    "DIRECT_TRANSFER",
		OrderTitle:  "单笔转账接口调试",
		PayeeInfo: &alipay.Participant{
			Identity:     "2088102175557825",
			IdentityType: "ALIPAY_USER_ID",
			Name:         "沙箱环境",
		},
		Remark:         "单笔转账接口调试备注",
		BusinessParams: "{\"payer_show_name_use_alias\":\"true\"}",
	}
	res, err := aliClient.FundTransUniTransferRequest(req)
	if err != nil {
		t.Log(err.Error())
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log(string(bytes))
}

func TestSyncVerifySign(t *testing.T) {
	var str string
	// 证书模式
	str = "{\"alipay_trade_precreate_response\":{\n  \"code\":\"10000\",\n  \"msg\":\"Success\",\n  \"out_trade_no\":\"6141161365682511\",\n  \"qr_code\":\"https:\\/\\/qr.alipay.com\\/bax03206ug0kulveltqc80a8\"\n},\n\"alipay_cert_sn\":\"80121e8b64901cf31d529c70dd6cd8c4\",\n\"sign\":\"VrgnnGgRMNApB1QlNJimiOt5ocGn4a4pbXjdoqjHtnYMWPYGX9AS0ELt8YikVAl6LPfsD7hjSyGWGjwaAYJjzH1MH7B2/T3He0kLezuWHsikao2ktCjTrX0tmUfoMUBCxKGGuDHtmasQi4yAoDk+ux7og1J5tL49yWiiwgaJoBE=\"\n}"
	// 非证书模式
	//str = "{\"alipay_trade_precreate_response\":{\"code\":\"10000\",\"msg\":\"Success\",\"out_trade_no\":\"123456666\",\"qr_code\":\"https:\\/\\/qr.alipay.com\\/bax09815z1h5tif3pfax002d\"},\"sign\":\"AcUFuQRN8hq4jwhEvJkNqEcRnBtj0ClWUPZYzn9reTwG6wnA+GHzNBQq2or28rxJwzgMf6QVDyFyS8zh28yPA5Dw6hFESeoDZGSgjwtaNou8No+4Vuh1opHE38lUV878YvH1cSJvYk/356siImSimYd8Sc5mK10kkgiVBeFghBzfBQvTjG8Em9IQsTXLz8WnZLdScdtI7mkHW/Cxa0Or0XOPIBJEzO3f03ip/RDGOjBVEH9EAXCWCv3Y14sFJ09Z9HKhTWHdphpqnq+BEAMfHoipYPX5hPWHS0RWl/vCIUcsbyCQacqftUXvybBFjFJrn1Iffy0unSmmFy7AHsrg1A==\"}"
	var res alipay.TradePrecreateResponseParams
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		t.Log(err)
		return
	}
	bytes, _ := json.Marshal(res)
	t.Log("返回值：", string(bytes))
	c, err := NewAliClient()
	if err != nil {
		return
	}
	var result bool
	result, err = c.SyncVerifySign(str, "alipay.trade.precreate")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(result)
}

func TestAesCBCDecrypt(t *testing.T) {
	//plaintext:="TestAesCBCDecrypt"
	//encryptStr,err:=utils.AesCBCEncrypt(plaintext,[]byte(encryptKey))
	//if err!=nil{
	//	t.Log(err)
	//	return
	//}
	////func AesCBCEncrypt(plaintext string, secretKey []byte) string {
	////ciphertext:="0qsDWRG3FrvPTo6nDKqWk2T6k9GkQIRAskrWSaHzRssaVgqJD5Toc3PZ3yI8M13tzzfMTXUSoL+XHrfnma20bLvWopFU4wwC5az8aAxc33srybWXHTl6CtUuMC/ETJaiW7QhOHtXKeuFtampu4Wsbw=="
	//res,err:=utils.AesCBCDecrypt(encryptStr,[]byte(encryptKey))
	//if err!=nil{
	//	t.Log(err)
	//	return
	//}
	//t.Log(res)
}