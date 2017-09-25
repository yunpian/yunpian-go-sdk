package sdk

import "testing"

/**
 * @author dzh
 * @date 25/09/2017 14:03
 * @since 0.0.1
 */

func TestAdd(t *testing.T) {
    sign := TestClnt.Sign()
    param := NewParam(5)
    param[SIGN] = "你好吗"
    param[NOTIFY] = "true"
    param[APPLYVIP] = "false"
    param[ISONLYGLOBAL] = "false"
    param[INDUSTRYTYPE] = "其他"
    r := sign.Add(param)
    t.Log(r)
}

func TestUpdate(t *testing.T) {
    sign := TestClnt.Sign()
    param := NewParam(6)
    param[SIGN] = "你好吗1"
    param[OLD_SIGN] = "你好吗"
    param[NOTIFY] = "true"
    param[APPLYVIP] = "false"
    param[ISONLYGLOBAL] = "false"
    param[INDUSTRYTYPE] = "其他"
    r := sign.Update(param)
    t.Log(r)
}

func TestGet(t *testing.T) {
    sign := TestClnt.Sign()
    param := NewParam(5)
    param[SIGN] = "你好吗"
    param[PAGE_NUM] = "1"
    param[PAGE_SIZE] = "3"
    r := sign.Get(param)
    t.Log(r)
}
