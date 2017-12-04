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
