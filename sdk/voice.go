package sdk

import (
    "fmt"
)

/**
 * https://www.yunpian.com/api2.0/voice.html
 *
 * @author dzh
 * @date 20/09/2017 20:49
 * @since 0.0.1
 */

type VoiceApi interface {
    YunpianApi
    Send(param map[string]string) (r *Result)
    TplNotify(param map[string]string) (r *Result)
    PullStatus(param map[string]string) (r *Result)
}

type VoiceApiOption struct {
    YunpianApiOption
}

func NewVoice() VoiceApi {
    voice := &VoiceApiOption{}
    voice.name = VOICE
    return voice
}

// <h1>发语音验证码</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// mobile String 是 接收的手机号、固话（需加区号） 15205201314 01088880000
// </p>
// <p>
// code String 是 验证码，支持4~6位阿拉伯数字 1234
// </p>
// <p>
// encrypt String 否 加密方式 使用加密 tea (不再使用)
// </p>
// <p>
// _sign String 否 签名字段 参考使用加密 393d079e0a00912335adfe46f4a2e10f (不再使用)
// </p>
// <p>
// callback_url String 否 本条语音验证码状态报告推送地址 http://your_receive_url_address
// </p>
// <p>
// display_num String 否 透传号码，为保证全国范围的呼通率，云片会自动选择最佳的线路，透传的主叫号码也会相应变化。
// 如需透传固定号码则需要单独注册报备，为了确保号码真实有效，客服将要求您使用报备的号码拨打一次客服电话
// </p>
func (voice *VoiceApiOption) Send(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, MOBILE, CODE}
    if err := voice.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch voice.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[RESULT]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return voice.SetPath("send.json").Post(param, h, r)
}

// <h1>发送语音通知</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// mobile String 是 接收的手机号、固话（需加区号） 15205201314 01088880000
// </p>
// <p>
// tpl_id Long 是 审核通过的模版ID 1136
// </p>
// <p>
// tpl_value String 是 模版的变量值
// 如模版内容&quot;课程#name#在#time#开始&quot;,那么这里的值为&quot;name=计算机&amp;time=17点&quot;,注意若出现特殊字符(如&#39;=&#39;,&#39;&amp;&#39;)则需要URLEncode内容
// </p>
// <p>
// callback_url String 否 本条语音验证码状态报告推送地址 http://your_receive_url_address
// </p>
func (voice *VoiceApiOption) TplNotify(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, MOBILE, TPL_ID, TPL_VALUE}
    if err := voice.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch voice.version {
            case V2:
                return rsp
            }
            return nil
        },
    }
    return voice.SetPath("tpl_notify.json").Post(param, h, r)
}

// <h1>获取状态报告</h1>
//
// <p>
// 参数名 是否必须 描述 示例
// </p>
// <p>
// apikey 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// page_size 否 每页个数，最大100个，默认20个 20
// </p>
// <p>
// type Integer 否 拉取类型，1-语音验证码 2-语音通知，默认type=1 1
// </p>

func (voice *VoiceApiOption) PullStatus(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY}
    if err := voice.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch voice.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[VOICE_STATUS]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return voice.SetPath("pull_status.json").Post(param, h, r)
}
