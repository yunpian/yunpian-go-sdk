package sdk

/**
 * @author dzh
 * @date 20/09/2017 20:49
 * @since 0.0.1
 */

type VoiceApi interface {
    YunpianApi
}

type VoiceApiImpl struct {
    YunpianApiOption
}

func NewVoice() VoiceApi {
    voice := &VoiceApiImpl{}
    voice.name = VOICE
    return voice
}
