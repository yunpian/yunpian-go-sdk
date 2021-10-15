# 云片 SDK for Go

yunpian-go-sdk 是由飞牛巴士实现的[云片](https://www.yunpian.com/)接口客户端SDK. 

## Installing

请使用以下命令进行安装

    go get -u github.com/FeiniuBus/yunpian-go-sdk/yunpian

## Getting Start

* 以发送单条短信为例

```Go
package main

import (
	"log"

	"github.com/FeiniuBus/yunpian-go-sdk/yunpian"
)

func main() {
	cfg := yunpian.DefaultConfig().WithAPIKey("xxxxxxxxxxxxx")
	sms := yunpian.NewClient(cfg).SMS()

	input := &yunpian.SingleSendRequest{
		Text:   "【云片网】您的验证码是1234",
		Mobile: "13800138000",
	}

	resp, err := sms.SingleSend(input)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf(resp.Message)
}
```

## Configuration

* `WithUseSSL(use bool)`: 是否开启HTTPS访问
* `WithHTTPClient(client *http.Client)`: 使用自定义HTTPClient