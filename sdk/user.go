package sdk

import "fmt"

/**
 * https://www.yunpian.com/api2.0/user.html
 *
 * @author dzh
 * @date 20/09/2017 20:50
 * @since 0.0.1
 */

type UserApi interface {
    YunpianApi
    Get(param map[string]string) (r *Result)
    Set(param map[string]string) (r *Result)
}

type UserApiOption struct {
    YunpianApiOption
}

func NewUser() UserApi {
    user := &UserApiOption{}
    user.name = USER
    return user
}

// <h1>查账户信息</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
func (user *UserApiOption) Get(param map[string]string) (r *Result) {
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
    if err := user.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch user.version {
            case V1:
                if data, found := rsp.(map[string]interface{})[USER]; found {
                    return data
                }
            case V2:
                return rsp
            }
            return nil
        },
    }
    return user.SetPath("get.json").Post(param, h, r)
}

// <h1>修改账户信息</h1>
//
// <p>
// 参数名 类型 是否必须 描述 示例
// </p>
// <p>
// apikey String 是 用户唯一标识 9b11127a9701975c734b8aee81ee3526
// </p>
// <p>
// emergency_contact String 否 紧急联系人 zhangshan
// </p>
// <p>
// emergency_mobile String 否 紧急联系人手机号 13012345678
// </p>
// <p>
// alarm_balance Long 否 短信余额提醒阈值。 一天只提示一次 100
// </p>
func (user *UserApiOption) Set(param map[string]string) (r *Result) {
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
    if err := user.VerifyParam(param, must, r); err != nil {
        return r.Error(err)
    }

    h := &YunpianResultHandler{
        Data: func(rsp interface{}) interface{} {
            switch user.version {
            case V2:
                return rsp
            }
            return nil
        },
    }
    return user.SetPath("set.json").Post(param, h, r)
}
