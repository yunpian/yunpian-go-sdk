package sdk

import "testing"

/**
 * @author dzh
 * @date 20/09/2017 22:10
 * @since 0.0.1
 */

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

const APIKEY = "2daab1114c69c9c41d1172b0ad8c392d"

func TestHttpPost(t *testing.T) {
    clnt := New(APIKEY)
    clnt.WithConf(DevConf)
    rsp, err := clnt.Post("https://test-api.yunpian.com/v2/user/get.json", "apikey="+APIKEY, nil, "")
    t.Log(rsp, err)
}
