package sdk

import (
    "fmt"
)

/**
 * https://www.yunpian.com/api2.0/flow.html
 *
 * @author dzh
 * @date 20/09/2017 20:49
 * @since 0.0.1
 */

type FlowApi interface {
    YunpianApi
    GetPackage(param map[string]string) *Result
    Recharge(param map[string]string) *Result
    PullStatus(param map[string]string) *Result
}

type FlowApiOption struct {
    YunpianApiOption
}

func NewFlow() FlowApi {
    flow := &FlowApiOption{}
    flow.name = FLOW
    return flow
}

// <h1>查询流量包</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// carrier String 否 运营商ID 传入该参数则获取指定运营商的流量包， 否则获取所有运营商的流量包 移动：10086 联通：10010
// 电信：10000
// </p>
func (flow *FlowApiOption) GetPackage(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = make(map[string]string, 1)
    }
    must := []string{APIKEY}
    if err := flow.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            //fmt.Println(rsp)
            switch flow.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[FLOW_PACKAGE]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return flow.SetPath("get_package.json").Post(param, h, r)
}

// <h1>充值流量</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
//</p>
// <p>
//mobile String 是 接收的手机号（仅支持大陆号码） 15205201314
// </p>
// <p>
// sn String 是 流量包的唯一ID 点击查看 1008601
// </p>
// <p>
// callback_url String 否 本条流量充值的状态报告推送地址 http: //your_receive_url_address
// </p>
// <p>
// encrypt String 否 加密方式 使用加密 tea (不再使用)
// </p>
// <p>
// _sign String 否 签名字段 参考使用加密 393d079e0a00912335adfe46f4a2e10f (不再使用)
// </p>
func (flow *FlowApiOption) Recharge(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = NewParam(0)
    }
    must := []string{APIKEY, MOBILE, SN}
    if err := flow.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            //fmt.Println(rsp)
            switch flow.version {
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
    return flow.SetPath("recharge.json").Post(param, h, r)
}

// <h1>获取状态报告</h1>
// <p>
//参数名 是否必须 描述 示例
//</p>
// <p>
// apikey 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// page_size 否 每页个数，最大100个，默认20个 20
// </p>
func (flow *FlowApiOption) PullStatus(param map[string]string) (r *Result) {
    r = new(Result)
    defer func() {
        if e := recover(); e != nil {
            r.Error(fmt.Errorf("%v", e))
        }
    }()
    if param == nil {
        param = make(map[string]string, 1)
    }
    must := []string{APIKEY}
    if err := flow.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            //fmt.Println(rsp)
            switch flow.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[FLOW_STATUS]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return flow.SetPath("pull_status.json").Post(param, h, r)
}
