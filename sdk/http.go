package sdk

/**
 * @author dzh
 * @date 20/09/2017 21:53
 * @since 0.0.1
 */
import (
    "net/http"
    "time"
    "net"
    "strings"
    "io/ioutil"
    "github.com/yunpian/yunpian-go-sdk/sdk/api"
)

type HttpClnt struct {
    apikey string
    conf   YunpianConf

    http *http.Client
}

func New(apikey string) YunpianClient {
    return &HttpClnt{
        apikey,
        *OnlineConf,
        createHttp(OnlineConf),
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
        clnt.http = http;
    }
    return clnt
}

func (clnt *HttpClnt) Conf() *YunpianConf {
    return &clnt.conf
}

func (clnt *HttpClnt) Apikey() string {
    return clnt.apikey
}

func (clnt *HttpClnt) Flow() (flow api.FlowApi) {
    return clnt.api(api.Flow)
}

func (clnt *HttpClnt) Sign() (sign api.SignApi) {
    return clnt.api(api.Sign)
}

func (clnt *HttpClnt) Sms() (sms api.SmsApi) {
    return clnt.api(api.Sms)
}

func (clnt *HttpClnt) Tpl() (tpl api.TplApi) {
    return clnt.api(api.Tpl)
}

func (clnt *HttpClnt) User() (user api.UserApi) {
    return clnt.api(api.User)
}

func (clnt *HttpClnt) Voice() (voice api.VoiceApi) {
    return clnt.api(api.Voice)
}

func (clnt *HttpClnt) Post(url string, data string, headers map[string]string, charset string) (string, error) {
    req, err := http.NewRequest("POST", url, strings.NewReader(data))
    if err != nil {
        return "", err
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
    if rsp, err := clnt.http.Do(req); err == nil {
        defer rsp.Body.Close()
        bodyBytes, err2 := ioutil.ReadAll(rsp.Body)
        return string(bodyBytes), err2
    } else {
        return "", err
    }
}

func (clnt *HttpClnt) Close() {

}

func (clnt *HttpClnt) api(name string) api.YunpianApi {
    var _api api.YunpianApi

    switch name {
    case api.Sms:
        _api = api.NewSms().SetHost(clnt.conf.SmsHost)
    case api.Voice:
        _api = api.NewVoice().SetHost(clnt.conf.VoiceHost)
    case api.Flow:
        _api = api.NewFlow().SetHost(clnt.conf.FlowHost)
    case api.Sign:
        _api = api.NewSign().SetHost(clnt.conf.SignHost)
    case api.Tpl:
        _api = api.NewTpl().SetHost(clnt.conf.TplHost)
    case api.User:
        _api = api.NewUser().SetHost(clnt.conf.UserHost)
    default:
        panic("Not found api" + name)
    }

    _api.SetApikey(clnt.Apikey()).SetCharset(clnt.conf.Http.Charset).SetVersion(clnt.conf.Version)
    _api.SetHttpClnt(clnt.http)
    return _api
}
