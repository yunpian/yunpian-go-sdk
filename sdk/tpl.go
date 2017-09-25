package sdk

import "fmt"

/**
 * https://www.yunpian.com/api2.0/tpl.html
 *
 * @author dzh
 * @date 20/09/2017 20:50
 * @since 0.0.1
 */

type TplApi interface {
    YunpianApi
    GetDefault(param map[string]string) (r *Result)
    Get(param map[string]string) (r *Result)
    Add(param map[string]string) (r *Result)
    Del(param map[string]string) (r *Result)
    Update(param map[string]string) (r *Result)
    AddVoiceNotify(param map[string]string) (r *Result)
    UpdateVoiceNotify(param map[string]string) (r *Result)
}

type TplApiOption struct {
    YunpianApiOption
}

func NewTpl() TplApi {
    tpl := &TplApiOption{}
    tpl.name = TPL
    return tpl
}

// <h1>取默认模板</h1>
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// tpl_id Long 否 模板id，64位长整形。指定id时返回id对应的默认 模板。未指定时返回所有默认模板 1
// </p>
func (tpl *TplApiOption) GetDefault(param map[string]string) (r *Result) {
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
    if err := tpl.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch tpl.version {
            case V2:
                return rsp
            }
            return nil
        },
    }
    return tpl.SetPath("get_default.json").Post(param, h, r)
}

// <h1>取模板</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// tpl_id Long 否 模板id，64位长整形。指定id时返回id对应的 模板。未指定时返回所有模板 1
// </p>
func (tpl *TplApiOption) Get(param map[string]string) (r *Result) {
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
    if err := tpl.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch tpl.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[TEMPLATE]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return tpl.SetPath("get.json").Post(param, h, r)
}

// <h1>添加模板</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// tpl_content String 是 模板内容，必须以带符号【】的签名开头 【云片网】您的验证码是#code#
// </p>
// <p>
// notify_type Integer 否 审核结果短信通知的方式: 0表示需要通知,默认; 1表示仅审核不通过时通知; 2表示仅审核通过时通知;
// 3表示不需要通知 1
// </p>
// <p>
// lang String 否 国际短信模板所需参数，模板语言:简体中文zh_cn; 英文en; 繁体中文 zh_tw; 韩文ko,日文 ja
// zh_cn
// </p>
func (tpl *TplApiOption) Add(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, TPL_CONTENT}
    if err := tpl.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch tpl.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[TEMPLATE]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return tpl.SetPath("add.json").Post(param, h, r)
}

// <h1>删除模板</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// tpl_id Long 是 模板id，64位长整形 9527
// </p>
func (tpl *TplApiOption) Del(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, TPL_ID}
    if err := tpl.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch tpl.version {
            case V2:
                return rsp
            }
            return nil
        },
    }
    return tpl.SetPath("del.json").Post(param, h, r)
}

// <h1>修改模板</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// tpl_id Long 是 模板id，64位长整形，指定id时返回id对应的模板。未指定时返回所有模板 9527
// </p>
// <p>
// tpl_content String 是
// 模板id，64位长整形。指定id时返回id对应的模板。未指定时返回所有模板模板内容，必须以带符号【】的签名开头 【云片网】您的验证码是#code#
// </p>
func (tpl *TplApiOption) Update(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, TPL_ID, TPL_CONTENT}
    if err := tpl.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch tpl.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[TEMPLATE]; found {
                    return data
                }
            case V2:
                if data, found := rsp.(map[string]interface{})[TEMPLATE]; found {
                    return data
                } else {
                    return rsp
                }
            }
            return nil
        },
    }
    return tpl.SetPath("update.json").Post(param, h, r)
}

// <h1>添加语音通知模版</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// tpl_content String 是 模板内容，必须以带符号【】的签名开头 【云片网】您的验证码是#code#
// </p>
// <p>
// notify_type Integer 否 审核结果短信通知的方式: 0表示需要通知,默认; 1表示仅审核不通过时通知; 2表示仅审核通过时通知;
// 3表示不需要通知 1
// </p>
func (tpl *TplApiOption) AddVoiceNotify(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, TPL_CONTENT}
    if err := tpl.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch tpl.version {
            case V2:
                return rsp
            }
            return nil
        },
    }
    return tpl.SetPath("add_voice_notify.json").Post(param, h, r)
}

// <h1>修改语音通知模版</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// tpl_id Long 是 模板id，64位长整形，指定id时返回id对应的模板。未指定时返回所有模板 9527
// </p>
// <p>
// tpl_content String 是
// 模板id，64位长整形。指定id时返回id对应的模板。未指定时返回所有模板模板内容，必须以带符号【】的签名开头 【云片网】您的验证码是#code#
// </p>
func (tpl *TplApiOption) UpdateVoiceNotify(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, TPL_CONTENT}
    if err := tpl.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch tpl.version {
            case V2:
                if data, found := rsp.(map[string]interface{})[TEMPLATE]; found {
                    return data
                } else {
                    return rsp
                }
            }
            return nil
        },
    }
    return tpl.SetPath("update_voice_notify.json").Post(param, h, r)
}
