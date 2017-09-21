package sdk

/**
 * @author dzh
 * @date 21/09/2017 10:08
 * @since 0.0.1
 */

type Result struct {
    code   int
    msg    string
    detail string
    data   interface{}
}

func (r *Result) Code() int {
    return r.code
}

func (r *Result) Msg() string {
    return r.msg
}

func (r *Result) Detail() string {
    return r.detail
}

func (r *Result) Data() interface{} {
    return r.data
}
