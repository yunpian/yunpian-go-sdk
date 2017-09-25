package sdk

import (
    "testing"
)

/**
 * @author dzh
 * @date 24/09/2017 22:58
 * @since 0.0.1
 */

func TestGetPackage(t *testing.T) {
    flow := TestClnt.Flow()
    r := flow.GetPackage(nil)
    t.Log(r)
    sn := r.Data.([]interface{})[0].(map[string]interface{})[SN].(float64)
    t.Log(int(sn))

    //flow.SetVersion(V1)
    //r = flow.GetPackage()
    //t.Log(r)
}

func TestRecharset(t *testing.T) {
    flow := TestClnt.Flow()
    param := NewParam(2)
    param[MOBILE] = "18616020610"
    param[SN] = "1008601"

    r := flow.Recharge(param)
    t.Log(r)

    //flow.SetVersion(V1)
    //r = flow.Recharge(param)
    //t.Log(r)
}

func TestPullStatus(t *testing.T) {
    flow := TestClnt.Flow()

    r := flow.PullStatus(nil)
    t.Log(r)

    //flow.SetVersion(V1)
    //r = flow.PullStatus()
    //t.Log(r)
}
