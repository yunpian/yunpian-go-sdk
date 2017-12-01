package yunpian

import (
	"encoding/json"
	"testing"
)

func TestSingleSend(t *testing.T) {
	sms := TestClient.SMS()
	input := &SingleSendRequest{
		Mobile: "13320942172",
		Text:   "您的验证码是1234。如非本人操作，请忽略本短信",
	}

	resp, err := sms.SingleSend(input)
	if err != nil {
		t.Error(err)
	}
	if !resp.IsSuccess() {
		t.Errorf("测试短信发送失败：%s", resp.Message)
	}
}

func TestDecodeSingleSendResponseBody(t *testing.T) {
	body := `{"code":0,"msg":"发送成功","count":1,"fee":0.045,"unit":"RMB","mobile":"13320942172","sid":19473255234}`
	var resp SingleSendResponse
	err := json.Unmarshal([]byte(body), &resp)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTotalFee(t *testing.T) {
	sms := TestClient.SMS()
	input := &GetTotalFeeRequest{
		Date: "2017-12-01",
	}

	resp, err := sms.GetTotalFee(input)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.Count)
}
