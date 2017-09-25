package sdk

import "testing"

/**
 * @author dzh
 * @date 25/09/2017 16:59
 * @since 0.0.1
 */

func TestTplAdd(t *testing.T) {
    tpl := TestClnt.Tpl()
    param := NewParam(1)
    param[TPL_CONTENT] = "【云片网】您的验证码是1234"
    r := tpl.Add(param)
    t.Log(r)

    //
    tpl.SetVersion(V1)
    r = tpl.Add(param)
    t.Log(r)
}

func TestTplGet(t *testing.T) {
    tpl := TestClnt.Tpl()
    param := NewParam(1)
    r := tpl.Get(param)
    t.Log(r)

    //
    tpl.SetVersion(V1)
    r = tpl.Get(param)
    t.Log(r)
}

func TestTplDel(t *testing.T) {
    tpl := TestClnt.Tpl()
    param := NewParam(1)
    param[TPL_ID] = "1"
    r := tpl.Del(param)
    t.Log(r)

    //
    tpl.SetVersion(V1)
    r = tpl.Del(param)
    t.Log(r)
}

func TestTplGetDefault(t *testing.T) {
    tpl := TestClnt.Tpl()
    param := NewParam(1)
    param[TPL_ID] = "1"
    r := tpl.GetDefault(param)
    t.Log(r)

    //
    tpl.SetVersion(V1)
    r = tpl.GetDefault(param)
    t.Log(r)
}

func TestTplUpdate(t *testing.T) {
    tpl := TestClnt.Tpl()
    param := NewParam(2)
    param[TPL_ID] = "665"
    param[TPL_CONTENT] = "【云片网】您的验证码是#code#"
    r := tpl.Update(param)
    t.Log(r)

    //
    tpl.SetVersion(V1)
    r = tpl.Update(param)
    t.Log(r)
}

func TestTplAddVoiceNotify(t *testing.T) {
    tpl := TestClnt.Tpl()
    param := NewParam(1)
    param[TPL_CONTENT] = "应用#name#在#time#无法响应"
    r := tpl.AddVoiceNotify(param)
    t.Log(r)

}

func TestTplUpdateVoiceNotify(t *testing.T) {
    tpl := TestClnt.Tpl()
    param := NewParam(2)
    param[TPL_ID] = "3373"
    param[TPL_CONTENT] = "应用#name#在#time#无法响应1"
    r := tpl.UpdateVoiceNotify(param)
    t.Log(r)

}
