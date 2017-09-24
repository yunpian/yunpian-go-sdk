package sdk

import (
    "testing"
    "encoding/json"
    "bytes"
)

/**
 * @author dzh
 * @date 24/09/2017 00:45
 * @since 0.0.1
 */

func TestJsonStr(t *testing.T) {
    str := `{"a":{"b":1}}`
    f := decode(bytes.NewBufferString(str)).(map[string]interface{})
    t.Log(f["a"])

    //str = `[1,2,3]`
    //f = decode(bytes.NewBufferString(str))
    //t.Log(f)
}

func decode(buf *bytes.Buffer) interface{} {
    switch ch, _, _ := bytes.NewReader(buf.Bytes()).ReadRune(); ch {
    case '{':
        var f interface{}
        println(ch)
        json.NewDecoder(buf).Decode(&f)
        return f
    case '[':
        var f interface{}
        println(ch)
        json.NewDecoder(buf).Decode(&f)
        return f
    default:
        panic("invalid json " + string(ch))
    }

}

func TestList(t *testing.T) {
    listJson := []int{1, 2, 3}
    buf := new(bytes.Buffer)
    json.NewEncoder(buf).Encode(listJson)
    t.Log(buf)

    var l interface{}
    json.NewDecoder(buf).Decode(&l)
    switch l.(type) {
    case []interface{}:
        t.Log(l.([]interface{})[0])
    }
    t.Log(l)
}

func TestMap(t *testing.T) {
    mapJson := map[string]int{}
    mapJson["a"] = 1
    //j := string(json.Marshal(str))
    //println(j)
    buf := new(bytes.Buffer)
    json.NewEncoder(buf).Encode(mapJson)
    t.Log(buf.String())

    var f map[string]int
    //json.Unmarshal(buf.Bytes(), &f)
    json.NewDecoder(buf).Decode(&f)
    t.Log(f)

    //var f interface{}
    ////json.Unmarshal(buf.Bytes(), &f)
    //json.NewDecoder(buf).Decode(&f)
    //m := f.(map[string]interface{})
    //for k, v := range m {
    //    switch vv := v.(type) {
    //    case string:
    //        fmt.Println(k, "is string", vv)
    //    case int:
    //        fmt.Println(k, "is int", vv)
    //    case []interface{}:
    //        fmt.Println(k, "is an array:")
    //        for i, u := range vv {
    //            fmt.Println(i, u)
    //        }
    //    default:
    //        fmt.Println(k, "is of a type I don't know how to handle")
    //    }
    //}
}
