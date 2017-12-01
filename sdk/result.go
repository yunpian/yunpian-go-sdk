package sdk

import "fmt"

/**
 * @author dzh
 * @date 21/09/2017 10:08
 * @since 0.0.1
 */

// Code definition
const (
    SUCC   = 0
    UNKOWN = -50
)

type Result struct {
    Code   int         `json:"code,omitempty" `
    Msg    string      `json:"msg,omitempty" `
    Detail string      `json:"detail,omitempty" `
    Data   interface{} `json:"data,omitempty" `
}

func (r *Result) SetCode(code int) *Result {
    r.Code = code
    return r
}

func (r *Result) SetMsg(msg string) *Result {
    r.Msg = msg
    return r
}

func (r *Result) SetDetail(detail string) *Result {
    r.Detail = detail
    return r
}

func (r *Result) SetData(data interface{}) *Result {
    r.Data = data
    return r
}

func (r *Result) String() string {
    return fmt.Sprintf("%d %s %s %v", r.Code, r.Msg, r.Detail, r.Data)
}

func (r *Result) IsSucc() bool {
    return r.Code == SUCC
}

func (r *Result) Error(err error) *Result {
    return r.SetCode(UNKOWN).SetMsg(err.Error())
}
