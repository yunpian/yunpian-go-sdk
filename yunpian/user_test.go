package yunpian

import (
	"testing"
)

func TestUserGet(t *testing.T) {
	user := TestClient.User()
	resp, err := user.Get()
	if err != nil {
		t.Error(err)
	}

	t.Log(resp)
}

func TestUserSet(t *testing.T) {
	user := TestClient.User()
	input := &UserSetRequest{
		AlarmBalance: 15,
	}
	resp, err := user.Set(input)
	if err != nil {
		t.Error(err)
	}
	if resp.AlarmBalance != 15 {
		t.Error(resp)
	}
}
