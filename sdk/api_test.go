package sdk

import (
    "testing"
    "strings"
    "io/ioutil"
)

/**
 * @author dzh
 * @date 24/09/2017 21:57
 * @since 0.0.1
 */

func TestUrl(t *testing.T) {
    url := strings.Join([]string{"1", "v2", "send.json"}, "/")
    t.Log(url)
}

var DevConf = &YunpianConf{
    Version:   "v2",
    UserHost:  "https://test-api.yunpian.com",
    SignHost:  "https://test-api.yunpian.com",
    TplHost:   "https://test-api.yunpian.com",
    SmsHost:   "https://test-api.yunpian.com",
    VoiceHost: "https://test-api.yunpian.com",
    FlowHost:  "https://test-api.yunpian.com",

    Http: &HttpConf{Timeout: 30, KeepAlive: 30, MaxIdleConns: 100, IdleConnTimeout: 30, TLSHandshakeTimeout: 10,
        Charset: "utf-8"},
}

const TEST_APIKEY = ""

var TestClnt = New(TEST_APIKEY).WithConf(DevConf)

func TestHttpPost(t *testing.T) {
    rsp, err := TestClnt.Post("https://test-api.yunpian.com/v2/user/get.json", "apikey="+TEST_APIKEY, nil, "")
    if err == nil {
        defer rsp.Body.Close()
        body, _ := ioutil.ReadAll(rsp.Body)
        t.Log(string(body))
    }

}
