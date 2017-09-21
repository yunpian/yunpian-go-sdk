package api

import (
    "net/http"
)

/**
 * @author dzh
 * @date 20/09/2017 20:23
 * @since 0.0.1
 */

// Rest APIs' names
const (
    Flow  = "flow"
    Sign  = "sign"
    Sms   = "sms"
    Tpl   = "tpl"
    User  = "user"
    Voice = "voice"
)

type YunpianApi interface {
    Name() string
    //SetName(name string) YunpianApi

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

    HttpClnt() *http.Client
    SetHttpClnt(clnt *http.Client) YunpianApi
}

type YunpianApiOption struct {
    name    string
    host    string
    version string
    path    string
    apikey  string
    charset string

    httpClnt *http.Client
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

func (api *YunpianApiOption) HttpClnt() *http.Client {
    return api.httpClnt
}

func (api *YunpianApiOption) SetHttpClnt(clnt *http.Client) YunpianApi {
    api.httpClnt = clnt
    return api
}
