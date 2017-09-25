package sdk

import "testing"

/**
 * @author dzh
 * @date 25/09/2017 17:12
 * @since 0.0.1
 */

func TestUserGet(t *testing.T) {
    user := TestClnt.User()
    param := NewParam(1)
    r := user.Get(param)
    t.Log(r)

    //
    user.SetVersion(V1)
    r = user.Get(param)
    t.Log(r)
}

func TestUserSet(t *testing.T) {
    user := TestClnt.User()
    param := NewParam(3)
    param[EMERGENCY_CONTACT] = "dzh"
    param[EMERGENCY_MOBILE] = "18616020610"
    param[ALARM_BALANCE] = "10"
    r := user.Set(param)
    t.Log(r)

    //
    user.SetVersion(V1)
    r = user.Set(param)
    t.Log(r)
}
