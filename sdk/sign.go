package sdk

/**
 * @author dzh
 * @date 20/09/2017 20:49
 * @since 0.0.1
 */

type SignApi interface {
    YunpianApi
}

type SignApiImpl struct {
    YunpianApiOption
}

func NewSign() SignApi {
    sign := &SignApiImpl{}
    sign.name = SIGN
    return sign
}
