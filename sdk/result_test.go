package sdk

import (
    "testing"
    "fmt"
)

/**
 * @author dzh
 * @date 24/09/2017 15:00
 * @since 0.0.1
 */

func TestResult(t *testing.T) {
    r := new(Result)
    r.SetCode(0).SetMsg("succ").SetDetail("nil")
    t.Log(r.String())
}

type PS struct {
}

func TestP(t *testing.T) {
    var p interface{}
    t.Log(p)
    x := handle(&p)
    t.Log(x)
    t.Log(p)
    //var p PS
    //t.Log(*&p)
    //var p = new(PS)
    //t.Log(&*p)

    //if p == nil {
    //    t.Log("nil")
    //}
    //p = nil
    //t.Log(&*p)
    //handle(p)
    //t.Log(p)
    //if p == nil {
    //    t.Log("nil")
    //}
}

func handle(p *interface{}) interface{} {
    fmt.Println(p)
    if p == nil {
        fmt.Println("handle nil")
    }
    *p = *new(PS)
    return p
}
