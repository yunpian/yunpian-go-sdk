package sdk

/**
 * @author dzh
 * @date 20/09/2017 19:54
 * @since 0.0.1
 */

import (
    "net/http"
    "time"
    "github.com/yunpian/yunpian-go-sdk/sdk/api"
)

type YunpianClient interface {
    Flow() api.FlowApi
    Sign() api.SignApi
    Sms() api.SmsApi
    Tpl() api.TplApi
    User() api.UserApi
    Voice() api.VoiceApi

    api(name string) api.YunpianApi

    Apikey() string
    Conf() *YunpianConf

    WithConf(conf *YunpianConf) YunpianClient
    WithHttp(http *http.Client) YunpianClient

    Post(url string, data string, headers map[string]string, charset string) (string, error)

    Close()
}

type YunpianConf struct {
    Version string

    UserHost,
    SignHost,
    TplHost,
    SmsHost,
    VoiceHost,
    FlowHost string

    // http
    Http *HttpConf
}

type HttpConf struct {
    Timeout   time.Duration
    KeepAlive time.Duration

    MaxIdleConns        int
    IdleConnTimeout     time.Duration // second
    TLSHandshakeTimeout time.Duration

    Charset string
}

var OnlineConf = &YunpianConf{
    Version:   "v2",
    UserHost:  "https://sms.yunpian.com",
    SignHost:  "https://sms.yunpian.com",
    TplHost:   "https://sms.yunpian.com",
    SmsHost:   "https://sms.yunpian.com",
    VoiceHost: "https://voice.yunpian.com",
    FlowHost:  "https://flow.yunpian.com",

    Http: &HttpConf{Timeout: 30, KeepAlive: 30, MaxIdleConns: 100, IdleConnTimeout: 30, TLSHandshakeTimeout: 10,
        Charset: "utf-8"},
}

var DefHeaders = map[string]string{
    "Api-Lang":   "go",
    "Connection": "keep-alive",
}
