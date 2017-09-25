package sdk

import (
    "encoding/json"
    "strings"
    "bytes"
    "io/ioutil"
    "errors"
    "net/url"
    "fmt"
)

/**
 * @author dzh
 * @date 20/09/2017 20:23
 * @since 0.0.1
 */

// Yunpian APIs' names
const (
    FLOW  = "flow"
    SIGN  = "sign"
    SMS   = "sms"
    TPL   = "tpl"
    USER  = "user"
    VOICE = "voice"
)

const (
    V1 = "v1"
    V2 = "v2"
)

type YunpianApi interface {
    Name() string

    Host() string
    SetHost(host string) YunpianApi

    Version() string
    SetVersion(v string) YunpianApi

    Path() string
    SetPath(path string) YunpianApi

    Apikey() string
    SetApikey(apikey string) YunpianApi

    Charset() string
    SetCharset(charset string) YunpianApi

    HttpClnt() YunpianClient
    Init(clnt YunpianClient) YunpianApi

    Url() string
    Post(param map[string]string, h ResultHandler, r *Result) *Result
}

type YunpianApiResult interface {
    Result(rsp []byte, h ResultHandler, r *Result) *Result
}

func (api *YunpianApiOption) Url() string {
    return strings.Join([]string{api.host, api.version, api.name, api.path}, "/")
}

func (api *YunpianApiOption) Post(param map[string]string, h ResultHandler, r *Result) *Result {
    data := url.Values{}
    if param != nil {
        for k, v := range param {
            data.Add(k, v)
        }
    }
    if rsp, err := api.clnt.Post(api.Url(), data.Encode(), nil, api.charset); err == nil {
        defer rsp.Body.Close()
        if body, err := ioutil.ReadAll(rsp.Body); err == nil {
            api.Result(body, h, r)
        } else {
            r.Error(err)
        }
    } else {
        r.Error(err)
    }
    return r
}

func (api *YunpianApiOption) Result(rsp []byte, h ResultHandler, r *Result) *Result {
    data := h.Parse(rsp)
    code := h.Code(data, api.version)
    if code == SUCC {
        h.Succ(code, data, r)
    } else {
        h.Fail(code, data, r)
    }
    return r
}

type ResultHandler interface {
    Parse(rsp []byte) interface{}

    Code(rsp interface{}, version string) int

    Succ(code int, rsp interface{}, r *Result)
    Fail(code int, rsp interface{}, r *Result)
}

type YunpianResultHandler struct {
    Data func(rsp interface{}) interface{}
}

func (h *YunpianResultHandler) Succ(code int, rsp interface{}, r *Result) {
    RspMsgDetail(rsp, r).SetData(h.Data(rsp)).SetCode(code)
}

func (h *YunpianResultHandler) Fail(code int, rsp interface{}, r *Result) {
    RspMsgDetail(rsp, r).SetCode(code)
}

func RspMsgDetail(rsp interface{}, r *Result) *Result {
    switch rsp.(type) {
    case map[string]interface{}:
        if msg, found := rsp.(map[string]interface{})[MSG]; found {
            r.SetMsg(fmt.Sprint(msg))
        }
        if detail, found := rsp.(map[string]interface{})[DETAIL]; found {
            r.SetDetail(fmt.Sprint(detail))
        }
    }
    return r
}

func (h *YunpianResultHandler) Code(rsp interface{}, version string) int {
    if rsp == nil {
        return UNKOWN
    }
    if version == "" {
        version = V2
    }

    switch rsp.(type) {
    case map[string]interface{}:
        if code, found := rsp.(map[string]interface{})[CODE]; found {
            return int(code.(float64))
        } else {
            switch version {
            case V1:
                return UNKOWN
            case V2:
                return SUCC
                //TODO others version
            }
        }
    default:
        return SUCC
    }
    return UNKOWN
}

func (h *YunpianResultHandler) Parse(rsp []byte) interface{} {
    var r interface{} //TODO custom decode struct
    if err := json.NewDecoder(bytes.NewBuffer(rsp)).Decode(&r); err != nil {
        panic("Parse rsp:" + err.Error())
    }
    return r
}

type YunpianApiOption struct {
    name    string
    host    string
    version string
    path    string
    apikey  string
    charset string

    clnt YunpianClient
}

func (api *YunpianApiOption) Name() string {
    return api.name
}

//func (api *YunpianApiOption) SetName(name string) YunpianApi {
//    api.name = name
//    return api
//}

func (api *YunpianApiOption) Host() string {
    return api.host
}

func (api *YunpianApiOption) SetHost(h string) YunpianApi {
    api.host = h
    return api
}

func (api *YunpianApiOption) Version() string {
    return api.version
}

func (api *YunpianApiOption) SetVersion(v string) YunpianApi {
    api.version = v
    return api
}

func (api *YunpianApiOption) Path() string {
    return api.version
}

func (api *YunpianApiOption) SetPath(p string) YunpianApi {
    api.path = p
    return api
}

func (api *YunpianApiOption) Apikey() string {
    return api.apikey
}

func (api *YunpianApiOption) SetApikey(apikey string) YunpianApi {
    api.apikey = apikey
    return api
}

func (api *YunpianApiOption) Charset() string {
    return api.charset
}

func (api *YunpianApiOption) SetCharset(ch string) YunpianApi {
    api.charset = ch
    return api
}

func (api *YunpianApiOption) HttpClnt() YunpianClient {
    return api.clnt
}

func (api *YunpianApiOption) Init(clnt YunpianClient) YunpianApi {
    api.clnt = clnt
    api.SetApikey(clnt.Apikey()).SetCharset(clnt.Conf().Http.Charset).SetVersion(clnt.Conf().Version)
    return api
}

func (api *YunpianApiOption) VerifyParam(param map[string]string, must []string, r *Result) error {
    if param == nil {
        return errors.New("param is nil")
    }

    if _, found := param[APIKEY]; !found {
        param[APIKEY] = api.Apikey()
    }

    for _, m := range must {
        if _, found := param[m]; !found {
            return errors.New("miss param:" + m)
        }
    }
    return nil
}
