package sdk

/**
 * @author dzh
 * @date 20/09/2017 20:23
 * @since 0.0.1
 */

type SmsApi interface {
    YunpianApi
}

type SmsApiImpl struct {
    YunpianApiOption
}

func NewSms() SmsApi {
    sms := &SmsApiImpl{}
    sms.name = SMS
    return sms
}
