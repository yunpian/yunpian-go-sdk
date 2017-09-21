package api

/**
 * @author dzh
 * @date 20/09/2017 20:50
 * @since 0.0.1
 */

type TplApi interface {
    YunpianApi
}

type TplApiImpl struct {
    YunpianApiOption
}

func NewTpl() TplApi {
    tpl := &TplApiImpl{}
    tpl.name = Tpl
    return tpl
}
