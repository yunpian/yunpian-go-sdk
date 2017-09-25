package sdk

/**
 * @author dzh
 * @date 20/09/2017 19:54
 * @since 0.0.1
 */

import (
    "net/http"
    "time"
    "net"
    "strings"
)

type YunpianClient interface {
    Flow() FlowApi
    Sign() SignApi
    Sms() SmsApi
    Tpl() TplApi
    User() UserApi
    Voice() VoiceApi

    Apikey() string
    Conf() *YunpianConf

    // WithConf to initialize YunpianClient and inner *http.Client
    WithConf(conf *YunpianConf) YunpianClient
    // WithHttp to initialize inner *http.Client
    WithHttp(http *http.Client) YunpianClient

    HttpClnt() *http.Client
    Post(url string, data string, headers map[string]string, charset string) (*http.Response, error)

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

var DefOnlineConf = &YunpianConf{
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

type HttpClnt struct {
    apikey string
    conf   YunpianConf
    http   *http.Client
}

func New(apikey string) YunpianClient {
    return &HttpClnt{
        apikey,
        *DefOnlineConf,
        createHttp(DefOnlineConf),
    }
}

func createHttp(conf *YunpianConf) *http.Client {
    tr := &http.Transport{
        DialContext: (&net.Dialer{
            Timeout:   conf.Http.Timeout * time.Second,
            KeepAlive: conf.Http.KeepAlive * time.Second,
            DualStack: true,
        }).DialContext,
        MaxIdleConns:        conf.Http.MaxIdleConns,
        IdleConnTimeout:     conf.Http.IdleConnTimeout * time.Second,
        TLSHandshakeTimeout: conf.Http.TLSHandshakeTimeout * time.Second,
    }
    return &http.Client{Transport: tr}
}

func (clnt *HttpClnt) WithConf(conf *YunpianConf) YunpianClient {
    if conf != nil {
        clnt.conf = *conf
    }
    clnt.http = createHttp(&clnt.conf)
    return clnt
}

func (clnt *HttpClnt) WithHttp(http *http.Client) YunpianClient {
    if http != nil {
        clnt.http = http
    }
    return clnt
}

func (clnt *HttpClnt) HttpClnt() *http.Client {
    return clnt.http
}

func (clnt *HttpClnt) Conf() *YunpianConf {
    return &clnt.conf
}

func (clnt *HttpClnt) Apikey() string {
    return clnt.apikey
}

func (clnt *HttpClnt) Flow() FlowApi {
    flow := NewFlow()
    flow.Init(clnt).SetHost(clnt.Conf().FlowHost)
    return flow
    //return clnt.api(Flow).(FlowApi)
}

func (clnt *HttpClnt) Sign() SignApi {
    sign := NewSign()
    sign.Init(clnt).SetHost(clnt.conf.SignHost)
    return sign
}

func (clnt *HttpClnt) Sms() SmsApi {
    sms := NewSms()
    sms.Init(clnt).SetHost(clnt.conf.SmsHost)
    return sms
}

func (clnt *HttpClnt) Tpl() TplApi {
    tpl := NewTpl()
    tpl.Init(clnt).SetHost(clnt.conf.TplHost)
    return tpl
}

func (clnt *HttpClnt) User() UserApi {
    user := NewUser()
    user.Init(clnt).SetHost(clnt.conf.UserHost)
    return user
}

func (clnt *HttpClnt) Voice() VoiceApi {
    voice := NewVoice()
    voice.Init(clnt).SetHost(clnt.conf.VoiceHost)
    return voice
}

func (clnt *HttpClnt) Post(url string, data string, headers map[string]string, charset string) (*http.Response, error) {
    req, err := http.NewRequest("POST", url, strings.NewReader(data))
    if err != nil {
        return nil, err
    }
    if charset == "" {
        charset = clnt.conf.Http.Charset
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset="+charset)
    //add headers
    for key, val := range DefHeaders {
        req.Header.Add(key, val)
    }
    if headers != nil {
        for key, val := range headers {
            req.Header.Set(key, val)
        }
    }
    return clnt.http.Do(req)
}

func (clnt *HttpClnt) Close() {
    //
}
