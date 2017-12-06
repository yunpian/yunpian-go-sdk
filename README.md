yunpian-go-sdk
================================
[云片](https://www.yunpian.com/) SDK

## 快速开始
- 使用云片sdk

```go
import ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"

// 发送短信
clnt := ypclnt.New("apikey")
param := ypclnt.NewParam(2)
param[ypclnt.MOBILE] = "18616020610"
param[ypclnt.TEXT] = "【云片网】您的验证码是1234"
r := clnt.Sms().SingleSend(param)

//账户:clnt.User() 签名:clnt.Sign() 模版:clnt.Tpl() 短信:clnt.Sms() 语音:clnt.Voice() 流量:clnt.Flow()
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
    - util.go   工具函数
- 分支说明
    - master最新稳定发布版本
    - develop待发布版本，贡献的代码请pull request到这里:)

## 联系我们
[云片支持 QQ](https://static.meiqia.com/dist/standalone.html?eid=30951&groupid=0d20ab23ab4702939552b3f81978012f&metadata={"name":"github"})

SDK开源QQ群

<img src="doc/sdk_qq.jpeg" width="15%" alt="SDK开源QQ群"/>

## 文档链接
- [api文档](https://www.yunpian.com/api2.0/guide.html)


## 开源SDK列表
- [飞牛巴士](https://github.com/FeiniuBus/yunpian-go-sdk/tree/master/yunpian)
