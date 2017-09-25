package sdk

import (
    "testing"
)

/**
 * @author dzh
 * @date 25/09/2017 17:27
 * @since 0.0.1
 */

func TestVoiceSend(t *testing.T) {
    voice := TestClnt.Voice()
    param := NewParam(2)
    param[MOBILE] = "18616020610"
    param[CODE] = "123412"
    r := voice.Send(param)
    t.Log(r)

    voice.SetVersion(V1)
    r = voice.Send(param)
    t.Log(r)
}

func TestVoicePullStatus(t *testing.T) {
    voice := TestClnt.Voice()
    param := NewParam(1)
    r := voice.PullStatus(param)
    t.Log(r)

    voice.SetVersion(V1)
    r = voice.PullStatus(param)
    t.Log(r)
}

func TestVoiceTplNotify(t *testing.T) {
    voice := TestClnt.Voice()
    param := NewParam(3)
    param[MOBILE] = "18616020610"
    param[TPL_ID] = "3373"
    param[TPL_VALUE] = "name=dzh&time=7"
    r := voice.TplNotify(param)
    t.Log(r)

    voice.SetVersion(V1)
    r = voice.TplNotify(param)
    t.Log(r)
}
