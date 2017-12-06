package yunpian

import (
	"testing"
)

func TestSMSPullStatusV1(t *testing.T) {
	sms := TestClient.SMS()
	resp, err := sms.PullStatusV1(10)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestSMSPullReplyV1(t *testing.T) {
	sms := TestClient.SMS()
	resp, err := sms.PullReplyV1(10)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestSMSGetReplyV1(t *testing.T) {
	sms := TestClient.SMS()
	input := &SMSGetReplyV1Request{
		PageNum:   1,
		PageSize:  10,
		StartTime: "2017-12-01 00:00:00",
		EndTime:   "2017-12-05 23:59:59",
	}

	resp, err := sms.GetReplyV1(input)
	if err != nil {
		t.Error(err)
	}

	t.Log(resp)
}

func TestSMSGetBlackWorldV1(t *testing.T) {
	sms := TestClient.SMS()
	resp, err := sms.GetBlackWorldV1("收放高利贷")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestSMSGetRecordV1(t *testing.T) {
	sms := TestClient.SMS()
	input := &SMSGetRecordV1Request{
		PageNum:   1,
		PageSize:  10,
		StartTime: "2017-11-30 12:00:00",
		EndTime:   "2017-12-09 12:00:00",
	}

	resp, err := sms.GetRecordV1(input)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestSMSCountV1(t *testing.T) {
	sms := TestClient.SMS()
	input := &SMSCountV1Request{
		StartTime: "2017-11-30 12:00:00",
		EndTime:   "2017-12-09 12:00:00",
	}

	resp, err := sms.CountV1(input)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
