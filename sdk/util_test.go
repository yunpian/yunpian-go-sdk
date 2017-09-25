package sdk

import (
    "testing"
)

/**
 * @author dzh
 * @date 25/09/2017 16:28
 * @since 0.0.1
 */

func TestUrlEncodedAndJoin(t *testing.T) {
    r := UrlEncodedAndJoin("a=", "b&", "c ", "http://")
    t.Log(r)
}
