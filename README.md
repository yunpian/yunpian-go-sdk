yunpian-go-sdk
================================
[云片](https://www.yunpian.com/) SDK

## 快速开始
- 使用云片sdk

```go
import ypclnt "github.com/yunpian/yupian-go-client/sdk"

// 发送短信示例
clnt := ypclnt.New("apikey")
param := ypclnt.NewParam(2)
param[ypclnt.MOBILE] = "18616020***"
param[ypclnt.TEXT] = "【云片网】您的验证码是1234"
r := ypclnt.sms().single_send(param)

//账户:clnt.user() 签名:clnt.sign() 模版:clnt.tpl() 短信:clnt.sms() 语音:clnt.voice() 流量:clnt.flow()
```

## 配置说明
- 自定义sdk配置
```go
clnt.WithConf(&YunpianConf{..})
```
**注**: 参考默认配置sdk.DefOnlineConf
- 自定义httpclient
```go
clnt.WithHttp(customHttp) //*http.Client
```

## 源码说明
- sdk
    - api.go    云片接口公共功能
    - client.go 云片客户端定义、配置
    - fields.go 常量定义
    - flow.go   流量
    - sign.go   签名
    - sms.go    短信
    - tpl.go    模版
    - user.go   用户
    - voice.go  语音
    - util.go   工具函数·

## 联系我们
[云片支持 QQ](https://static.meiqia.com/dist/standalone.html?eid=30951&groupid=0d20ab23ab4702939552b3f81978012f&metadata={"name":"github"})

SDK开源QQ群

<img src="doc/sdk_qq.jpeg" width="15%" alt="SDK开源QQ群"/>

## 文档链接
- [api文档](https://www.yunpian.com/api2.0/guide.html)

