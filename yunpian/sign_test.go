package yunpian

import (
	"testing"
)

func TestSignGet(t *testing.T) {
	sign := TestClient.Sign()
	input := &SignGetRequest{
		PageNum:  1,
		PageSize: 10,
	}

	resp, err := sign.Get(input)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
