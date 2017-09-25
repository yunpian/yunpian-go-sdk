package sdk

import (
    "strings"
    "net/url"
)

/**
 * @author dzh
 * @date 25/09/2017 16:02
 * @since 0.0.1
 */

func NewParam(capacity int) map[string]string {
    if capacity < 1 {
        capacity = 0
    }
    return make(map[string]string, capacity+1)
}

func UrlEncodedAndJoin(req ...string) string {
    r := make([]string, len(req))
    for i, str := range req {
        r[i] = url.QueryEscape(str)
    }
    return strings.Join(r, ",")
}
