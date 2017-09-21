package sdk

import (
    "runtime"
    "testing"
)

func _TestHello(t *testing.T) {
    print(runtime.NumCPU())
    if v := Hello(); v != "hello" {
        t.Errorf("Expected 'hello', but got '%s'", v)
    }
}

func Hello() string {
    return "hello1"
}

func TestYunpianClient(t *testing.T) {
    apikey := "123456"
    clnt := New(apikey)
    if clnt.Apikey() != apikey {
        t.Errorf("%s", apikey)
    }

}
