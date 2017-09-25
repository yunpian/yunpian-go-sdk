package sdk

import "testing"

/**
 * @author dzh
 * @date 25/09/2017 15:10
 * @since 0.0.1
 */

func TestSingleSend(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(2)
    param[MOBILE] = "18616020610"
    param[TEXT] = "【云片网】您的验证码是1234"
    r := sms.SingleSend(param)
    t.Log(r)
}

func TestBatchSend(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(2)
    param[MOBILE] = "18616020610"
    param[TEXT] = "【云片网】您的验证码是1234"
    r := sms.BatchSend(param)
    t.Log(r)
}

func TestMultiSend(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(2)
    param[MOBILE] = "18616020610,18616020611"
    param[TEXT] = UrlEncodedAndJoin("【云片网】您的验证码是1234", "【哈哈哈】您的验证码是1123")
    r := sms.MultiSend(param)
    t.Log(r)
}

func TestTplSingleSend(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(3)
    param[MOBILE] = "18616020610"
    param[TPL_ID] = "1"
    param[TPL_VALUE] = "#company#=云片网"
    r := sms.TplSingleSend(param)
    t.Log(r)
}

func TestTplBatchSend(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(3)
    param[MOBILE] = "18616020610"
    param[TPL_ID] = "1"
    param[TPL_VALUE] = "#company#=云片网"
    r := sms.TplBatchSend(param)
    t.Log(r)
}

func TestSmsPullStatus(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(1)
    param[PAGE_SIZE] = "20"
    r := sms.PullStatus(param)
    t.Log(r)

    sms.SetVersion(V1)
    r = sms.PullStatus(param)
    t.Log(r)
}

func TestPullReply(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(1)
    param[PAGE_SIZE] = "20"
    r := sms.PullReply(param)
    t.Log(r)

    sms.SetVersion(V1)
    r = sms.PullReply(param)
    t.Log(r)
}

func TestGetRecord(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(4)
    param[START_TIME] = "2013-08-11 00:00:00"
    param[END_TIME] = "2017-9-25 00:00:00"
    param[PAGE_NUM] = "1"
    param[PAGE_SIZE] = "20"
    r := sms.GetRecord(param)
    t.Log(r)

    sms.SetVersion(V1)
    r = sms.GetRecord(param)
    t.Log(r)
}

func TestGetBlackWord(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(4)
    param[TEXT] = "高利贷,发票"
    r := sms.GetBlackWord(param)
    t.Log(r)

    sms.SetVersion(V1)
    r = sms.GetBlackWord(param)
    t.Log(r)
}

func TestSend(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(2)
    param[MOBILE] = "18616020610"
    param[TEXT] = "【云片网】您的验证码是1234"

    sms.SetVersion(V1)
    r := sms.Send(param)
    t.Log(r)
}

func TestGetReply(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(4)
    param[START_TIME] = "2013-08-11 00:00:00"
    param[END_TIME] = "2017-9-25 00:00:00"
    param[PAGE_NUM] = "1"
    param[PAGE_SIZE] = "20"

    r := sms.GetReply(param)
    t.Log(r)

    sms.SetVersion(V1)
    r = sms.GetReply(param)
    t.Log(r)
}

func TestTplSend(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(3)
    param[MOBILE] = "18616020610"
    param[TPL_ID] = "1"
    param[TPL_VALUE] = "#company#=云片网"

    sms.SetVersion(V1)
    r := sms.TplSend(param)
    t.Log(r)
}

func TestCount(t *testing.T) {
    sms := TestClnt.Sms()
    param := NewParam(4)
    param[START_TIME] = "2013-08-11 00:00:00"
    param[END_TIME] = "2017-9-25 00:00:00"
    param[PAGE_NUM] = "1"
    param[PAGE_SIZE] = "20"

    r := sms.Count(param)
    t.Log(r)

    sms.SetVersion(V1)
    r = sms.GetReply(param)
    t.Log(r)
}
