package api

/**
 * @author dzh
 * @date 20/09/2017 20:50
 * @since 0.0.1
 */

type UserApi interface {
    YunpianApi
}

type UserApiImpl struct {
    YunpianApiOption
}

func NewUser() UserApi {
    user := &UserApiImpl{}
    user.name = User
    return user
}
