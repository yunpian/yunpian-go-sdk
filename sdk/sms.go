package sdk

import "fmt"

/**
 * https://www.yunpian.com/api2.0/sms.html
 *
 * @author dzh
 * @date 20/09/2017 20:23
 * @since 0.0.1
 */

type SmsApi interface {
    YunpianApi
    Send(param map[string]string) (r *Result)
    SingleSend(param map[string]string) (r *Result)
    BatchSend(param map[string]string) (r *Result)
    MultiSend(param map[string]string) (r *Result)
    MultiSendV1(param map[string]string) (r *Result)
    PullStatus(param map[string]string) (r *Result)
    PullReply(param map[string]string) (r *Result)
    GetReply(param map[string]string) (r *Result)
    GetBlackWord(param map[string]string) (r *Result)
    GetRecord(param map[string]string) (r *Result)
    Count(param map[string]string) (r *Result)
    TplSend(param map[string]string) (r *Result)
    TplSingleSend(param map[string]string) (r *Result)
    TplBatchSend(param map[string]string) (r *Result)
}

type SmsApiOption struct {
    YunpianApiOption
}

func NewSms() SmsApi {
    sms := &SmsApiOption{}
    sms.name = SMS
    return sms
}

// <h1>智能匹配模板发送 only v1</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// mobile String 是 接收的手机号;发送多个手机号请以逗号分隔，一次不要超过1000个
// 国际短信仅支持单号码发送，国际号码需包含国际地区前缀号码，格式必须是"+"号开头("+"号需要urlencode处理，否则会出现格式错误)，国际号码不以"+"开头将被认为是中国地区的号码
// （针对国际短信，mobile参数会自动格式化到E.164格式，可能会造成传入mobile参数跟后续的状态报告中的号码不一致。E.164格式说明，参见：
// https://en.wikipedia.org/wiki/E.164） 单号码：15205201314
// 多号码：15205201314,15205201315 国际短信：+93701234567
// </p>
// <p>
// text String 是 短信内容 【云片网】您的验证码是1234
// </p>
// <p>
// extend String 否 扩展号。默认不开放，如有需要请联系客服申请 001
// </p>
// <p>
// uid String 否 该条短信在您业务系统内的ID，比如订单号或者短信发送记录的流水号。填写后发送状态返回值内将包含这个ID
// 默认不开放，如有需要请联系客服申请 10001
// </p>
// <p>
// callback_url String 否
// 本条短信状态报告推送地址。短信发送后将向这个地址推送短信发送报告。"后台-系统设置-数据推送与获取”可以做批量设置。如果后台已经设置地址的情况下，单次请求内也包含此参数，将以请求内的推送地址为准。
// </p>
func (sms *SmsApiOption) Send(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, MOBILE, TEXT}
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[RESULT]; found {
                    return data
                }
            }
            return nil
        },
    }
    return sms.SetPath("send.json").Post(param, h, r)
}

// <h1>单条发送</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// mobile String 是
// 接收的手机号；仅支持单号码发送；国际号码需包含国际地区前缀号码，格式必须是"+"号开头("+"号需要urlencode处理，否则会出现格式错误)，国际号码不以"+"开头将被认为是中国地区的号码
// （针对国际短信，mobile参数会自动格式化到E.164格式，可能会造成传入mobile参数跟后续的状态报告中的号码不一致。E.164格式说明，参见：
// https://en.wikipedia.org/wiki/E.164） 国内号码：15205201314
// 国际号码：urlencode("+93701234567");
// </p>
// <p>
// text String 是 短信内容 【云片网】您的验证码是1234
// </p>
// <p>
// extend String 否 扩展号。默认不开放，如有需要请联系客服申请 001
// </p>
// <p>
// uid String 否 该条短信在您业务系统内的ID，比如订单号或者短信发送记录的流水号。填写后发送状态返回值内将包含这个ID
// 默认不开放，如有需要请联系客服申请 10001
// </p>
// <p>
// callback_url String 否
// 本条短信状态报告推送地址。短信发送后将向这个地址推送短信发送报告。"后台-系统设置-数据推送与获取”可以做批量设置。如果后台已经设置地址的情况下，单次请求内也包含此参数，将以请求内的推送地址为准。
// http://your_receive_url_address
// </p>
func (sms *SmsApiOption) SingleSend(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, MOBILE, TEXT}
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V2:
                return rsp
            }
            return nil
        },
    }
    return sms.SetPath("single_send.json").Post(param, h, r)
}

// <h1>批量发送</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// mobile String 是 接收的手机号；发送多个手机号请以逗号分隔，一次不要超过1000个。 单号码：15205201314
// 多号码：15205201314,15205201315
// </p>
// <p>
// text String 是 短信内容 【云片网】您的验证码是1234
// </p>
// <p>
// extend String 否 扩展号。默认不开放，如有需要请联系客服申请 001
// </p>
// <p>
// uid String 否 该条短信在您业务系统内的ID，比如订单号或者短信发送记录的流水号。填写后发送状态返回值内将包含这个ID
// 默认不开放，如有需要请联系客服申请 10001
// </p>
// <p>
// callback_url String 否
// 本条短信状态报告推送地址。短信发送后将向这个地址推送短信发送报告。"后台-系统设置-数据推送与获取”可以做批量设置。如果后台已经设置地址的情况下，单次请求内也包含此参数，将以请求内的推送地址为准。
// http://your_receive_url_address
// </p>
func (sms *SmsApiOption) BatchSend(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, MOBILE, TEXT}
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V2:
                return rsp
            }
            return nil
        },
    }
    return sms.SetPath("batch_send.json").Post(param, h, r)
}

// <h1>个性化发送</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// mobile String 是
// 接收的手机号；多个手机号请以逗号分隔，一次不要超过1000个且手机号个数必须与短信内容条数相等；不支持国际号码发送；
// 多号码：15205201314,15205201315
// </p>
// <p>
// text String 是
// 短信内容，多个短信内容请使用UTF-8做urlencode后，使用逗号分隔，一次不要超过1000条且短信内容条数必须与手机号个数相等
// 内容示意：UrlEncode("【云片网】您的验证码是1234", "UTF-8") + "," +
// UrlEncode("【云片网】您的验证码是5678", "UTF-8")
// </p>
// <p>
// extend String 否 扩展号。默认不开放，如有需要请联系客服申请 001
// </p>
// <p>
// uid String 否 该条短信在您业务系统内的ID，比如订单号或者短信发送记录的流水号。填写后发送状态返回值内将包含这个ID
// 默认不开放，如有需要请联系客服申请 10001
// </p>
// <p>
// callback_url String 否
// 本条短信状态报告推送地址。短信发送后将向这个地址推送短信发送报告。"后台-系统设置-数据推送与获取”可以做批量设置。如果后台已经设置地址的情况下，单次请求内也包含此参数，将以请求内的推送地址为准。
// http://your_receive_url_address
// </p>
func (sms *SmsApiOption) MultiSend(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, MOBILE, TEXT}
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V2:
                return rsp
            }
            return nil
        },
    }
    return sms.SetPath("multi_send.json").Post(param, h, r)
}

// /v1/sms/multi_send.json
func (sms *SmsApiOption) MultiSendV1(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, MOBILE, TEXT}
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V1:
                return rsp
            }
            return nil
        },
    }
    return sms.SetPath("multi_send.json").Post(param, h, r)
}

// <h1>获取状态报告</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// page_size Integer 否 每页个数，最大100个，默认20个 20
// </p>
func (sms *SmsApiOption) PullStatus(param map[string]string) (r *Result) {
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
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[SMS_STATUS]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return sms.SetPath("pull_status.json").Post(param, h, r)
}

// <h1>获取回复短信</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// page_size Integer 否 每页个数，最大100个，默认20个 20
// </p>
func (sms *SmsApiOption) PullReply(param map[string]string) (r *Result) {
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
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[SMS_REPLY]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return sms.SetPath("pull_reply.json").Post(param, h, r)
}

// <h1>查回复的短信</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// start_time String 是 短信回复开始时间 2013-08-11 00:00:00
// </p>
// <p>
// end_time String 是 短信回复结束时间 2013-08-12 00:00:00
// </p>
// <p>
// page_num Integer 是 页码，默认值为1 1
// </p>
// <p>
// page_size Integer 是 每页个数，最大100个 20
// </p>
// <p>
// mobile String 否 填写时只查该手机号的回复，不填时查所有的回复 15205201314
// </p>
// <p>
// return_fields 否 返回字段（暂未开放
// </p>
// <p>
// sort_fields 否 排序字段（暂未开放） 默认按提交时间降序
// </p>
func (sms *SmsApiOption) GetReply(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, START_TIME, END_TIME, PAGE_NUM, PAGE_SIZE}
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[SMS_REPLY]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return sms.SetPath("get_reply.json").Post(param, h, r)
}

// <h1>查屏蔽词</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// text String 是 要检查的短信模板或者内容 这是一条测试短信
// </p>
func (sms *SmsApiOption) GetBlackWord(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, TEXT}
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
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
    return sms.SetPath("get_black_word.json").Post(param, h, r)
}

// <h1>查短信发送记录</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// mobile String 否 需要查询的手机号 15205201314
// </p>
// <p>
// start_time String 是 短信发送开始时间 2013-08-11 00:00:00
// </p>
// <p>
// end_time String 是 短信发送结束时间 2013-08-12 00:00:00
// </p>
// <p>
// page_num Integer 否 页码，默认值为1 1
// </p>
// <p>
// page_size Integer 否 每页个数，最大100个 20
// </p>
func (sms *SmsApiOption) GetRecord(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, START_TIME, END_TIME}
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[SMS]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return sms.SetPath("get_record.json").Post(param, h, r)
}

// <h1>统计短信条数</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// start_time String 是 短信发送开始时间 2013-08-11 00:00:00
// </p>
// <p>
// end_time String 是 短信发送结束时间 2013-08-12 00:00:00
// </p>
// <p>
// mobile String 否 需要查询的手机号 15205201314
// </p>
// <p>
// page_num Integer 否 页码，默认值为1 1
// </p>
// <p>
// page_size Integer 否 每页个数，最大100个 20
// </p>
func (sms *SmsApiOption) Count(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, START_TIME, END_TIME}
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            if data, found := rsp.(map[string]interface{})[TOTAL]; found {
                return data
            }
            return 0
        },
    }
    return sms.SetPath("count.json").Post(param, h, r)
}

// <h1>指定模板发送 only v1</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// mobile String 是 接收的手机号 15205201314
// </p>
// <p>
// tpl_id Long 是 模板id 1
// </p>
// <p>
// tpl_value String 是 变量名和变量值对。请先对您的变量名和变量值分别进行urlencode再传递。使用参考：代码示例。
// 注：变量名和变量值都不能为空 模板： 【#company#】您的验证码是#code#。 最终发送结果： 【云片网】您的验证码是1234。
// tplvalue=urlencode("#code#") + "=" + urlencode("1234") + "&amp;" +
// urlencode("#company#") + "=" + urlencode("云片网"); 若您直接发送报文请求则使用下面这种形式
// tplvalue=urlencode(urlencode("#code#") + "=" + urlencode("1234") +
// "&amp;" + urlencode("#company#") + "=" + urlencode("云片网"));
// </p>
// <p>
// extend String 否 扩展号。默认不开放，如有需要请联系客服申请 001
// </p>
// <p>
// uid String 否 用户自定义唯一id。最大长度不超过256的字符串。 默认不开放，如有需要请联系客服申请 10001
// </p>
func (sms *SmsApiOption) TplSend(param map[string]string) (r *Result) {
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
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[RESULT]; found {
                    return data
                }
            }
            return nil
        },
    }
    return sms.SetPath("tpl_send.json").Post(param, h, r)
}

// <h1>指定模板单发 only v2</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// mobile String 是
// 接收的手机号（针对国际短信，mobile参数会自动格式化到E.164格式，可能会造成传入mobile参数跟后续的状态报告中的号码不一致。E.164格式说明，参见：
// https://en.wikipedia.org/wiki/E.164） 15205201314
// </p>
// <p>
// tpl_id Long 是 模板id 1
// </p>
// <p>
// tpl_value String 是 变量名和变量值对。请先对您的变量名和变量值分别进行urlencode再传递。使用参考：代码示例。
// 注：变量名和变量值都不能为空 模板： 【#company#】您的验证码是#code#。 最终发送结果： 【云片网】您的验证码是1234。
// tplvalue=urlencode("#code#") + "=" + urlencode("1234") + "&amp;" +
// urlencode("#company#") + "=" + urlencode("云片网"); 若您直接发送报文请求则使用下面这种形式
// tplvalue=urlencode(urlencode("#code#") + "=" + urlencode("1234") +
// "&amp;" + urlencode("#company#") + "=" + urlencode("云片网"));
// </p>
// <p>
// extend String 否 扩展号。默认不开放，如有需要请联系客服申请 001
// </p>
// <p>
// uid String 否 用户自定义唯一id。最大长度不超过256的字符串。 默认不开放，如有需要请联系客服申请 10001
// </p>
func (sms *SmsApiOption) TplSingleSend(param map[string]string) (r *Result) {
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
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V2:
                return rsp
            }
            return nil
        },
    }
    return sms.SetPath("tpl_single_send.json").Post(param, h, r)
}

// <h1>指定模板群发 only v2</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// mobile String 是
// 接收的手机号（针对国际短信，mobile参数会自动格式化到E.164格式，可能会造成传入mobile参数跟后续的状态报告中的号码不一致。E.164格式说明，参见：
// https://en.wikipedia.org/wiki/E.164） 15205201314
// </p>
// <p>
// tpl_id Long 是 模板id 1
// </p>
// <p>
// tpl_value String 是 变量名和变量值对。请先对您的变量名和变量值分别进行urlencode再传递。使用参考：代码示例。
// 注：变量名和变量值都不能为空 模板： 【#company#】您的验证码是#code#。 最终发送结果： 【云片网】您的验证码是1234。
// tplvalue=urlencode("#code#") + "=" + urlencode("1234") + "&amp;" +
// urlencode("#company#") + "=" + urlencode("云片网"); 若您直接发送报文请求则使用下面这种形式
// tplvalue=urlencode(urlencode("#code#") + "=" + urlencode("1234") +
// "&amp;" + urlencode("#company#") + "=" + urlencode("云片网"));
// </p>
// <p>
// extend String 否 扩展号。默认不开放，如有需要请联系客服申请 001
// </p>
// <p>
// uid String 否 用户自定义唯一id。最大长度不超过256的字符串。 默认不开放，如有需要请联系客服申请 10001
// </p>
func (sms *SmsApiOption) TplBatchSend(param map[string]string) (r *Result) {
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
    if err := sms.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch sms.version {
            case V2:
                return rsp
            }
            return nil
        },
    }
    return sms.SetPath("tpl_batch_send.json").Post(param, h, r)
}
