package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mm-ooto/alipay"
	"net/http"
)

var aliClient2 *alipay.AliClient

func init() {
	appId := "2016091200490539"
	appPrivateKey := "MIIEpAIBAAKCAQEAvCqpPIlU/VM78FXnVWoO/iPbiIg/TSdZO2UmKdd+sLtVRqhy+doS08pzw/Qq5XUSL4QwVhoRShoQmGKtAAxI2YshCnk61WdHZogOlmniQlundZmflRLQgNrEnGftcuNO9xMAgwXtke1akgXDHjpkyGhl8M6yJ2cwou9wH1x2VDsW0cFVdgOT7uQD2W81CCGkPeMu5XfxKbBoKRE8OPNMsKAE2Bd6dPStmg8w3E2ZCGo8xVLH+vdSyGX30iYZHFP65WFoQt4ANJU/OgaWrIOcGdskFGTEtGeKEDoPw27LY/VD4I+2KolGalv8qcSZME+3FrmlHL+7I2ZgVKsVUHqT0QIDAQABAoIBAHmuqOSR9tkfU1qXYtMUk/97FsPTQARX1teXELfsOGx3qKzZ0AiNIrG9cWGd64OZUppRxKRZlSazdlnlLfUi/JVZ6JMKVKaedEj04WIZtQyukrt1DgLsONOrJYvzlVU/c9hJfII+eiRtNq3JdiV9I6GKCapRMFpU29nyNzLAq3DJ7PA6IGfSy/vk/r5HbxxJoomlZ1BZjhDezLn2nPXFqLQunQMeBozA/Eh/v7GZ0tZrhX/LEmq1aqfXtB+4mBX//gDMwCcvBlCLeBsVDXQOzPvCXr7qKTIFdXiQZhlopEt3RFEQb6CoQF5XSHisq3wbLxK8wTa01+0jcknJZGHYYCUCgYEA6vfKnukjge2J7ptmIegu5j+vZmbAcnKBJ1C8xrcb4VyGAFGWotUcUMTa8lKrjYXAhKnj/FeKRu19VhfIjEGEno/7zLoQz9j8QLlQAykIBoP9TTXI380K2Kns8uYdSUNCtJTE8CmqOpHJrKVcUy/lxREfv8SmHFPKD9jTC9t7ZKsCgYEAzQJvLgrp9zrYTmbAxKxXeswiJsBlMo4atlIs+R0xH3xVyWQe8Ng2GqH4BqSvXIqdthTq8YNmid6H42KRvaVVX7k5Bfkx6Qdghi+YWkedSKB1mKteUsEPzRtjcLUAABgu4SFGX/YuMCCZkXpNx3McnVRNmW5dY+qxlhSQ+3EXEXMCgYEAnnu8IytFU+GQY2xVmxEscQkLmZo8u/UXwBjo+2+OUpdBmv1tCS+NBb2BoGi6ZZ6Nl+2vZQj2r5iILYWlM1UNypV7VT87D7Zfjphvq3IFg7+LHoTklG+MnU8gD0W/Aydm2r5thz/THeYvjU+L0mBALoe6TnKpR/oMFFw/HYRQ2jkCgYEAmbLPg+duzYnyjaT/tPO4ijntCLyJokNjx3kIeqPmJkLjVh+YCt0ugv0XpHNnfav23YIFOphXEdoiatmFhncj8KY/GDlhr+F1/mREhrrWMpMKVzFzf/t6Sz3TabZpj6iRzPtTdbJtomtuduEI2xV0SIfhvbw+jCByj6BPqhN5Rf0CgYAVth2Yg6oTblLTXxsEYFifo0AsTxPVwVT+Ow+W+MeAYsRvEBnyE6pLIcUhzj/yOrXSTOsYdmP6YoghJT/mh/rp4kHXFuiqcne5WRhdzhRJFRTwF6gRQMcJsZbB6JcICv6h2AMUCFHUARhJTLHBSxUlvUHT1hDe09aZ6r964qd6NQ=="
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAimaDmFEr8HxMkqUFrd94aGBp2hjpfsOqVWdepqER1kSFv1cwmusOY6wAXhf5GHuKuz9MwjP1LZRr3PmRM/OwD9BFf7t/CKRBn7LFIwCWQ+yN8IWCWcNxSZfAkMRIncA5lAY1r4iqaG++ze7KDLjQD1xI2UaFhUBwKiBoTljC6hfmAYzUPObox+aQq5LUYY98O8DPzZbjDhHc5wq+KqaRfKefcqFkx/6Vy/GGgUjIA8tydN0MOUzZMgu4i8Z4E2H3R98Mt0oEJEFQksZ3NQxqaulO7kl4H/gU97KGEKfjmd2w1opO7FKgihtD2DYcXgtO8iRBhMFz5owc4k2wvKFUCQIDAQAB"
	var err error
	aliClient2, err = alipay.NewAliClient(appId, aliPublicKey, appPrivateKey, "RSA2", false)
	if err != nil {
		fmt.Printf("NewAliClient error:%s\n", err.Error())
		return
	}
}

func main() {
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
	router.Run("192.168.10.254:7888")
}
