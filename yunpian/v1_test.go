package yunpian

import (
	"testing"
)

func TestUserGetV1(t *testing.T) {
	user := TestClient.User()
	resp, err := user.GetV1()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestUserSetV1(t *testing.T) {
	user := TestClient.User()
	input := &UserSetRequest{
		AlarmBalance: 15,
	}
	resp, err := user.SetV1(input)
	if err != nil {
		t.Error(err)
	}

	t.Log(resp)
}

func TestTPLGetV1(t *testing.T) {
	tpl := TestClient.TPL()
	resp, err := tpl.GetV1(2079056)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestTPLListV1(t *testing.T) {
	tpl := TestClient.TPL()
	resp, err := tpl.ListV1()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestVoicePullStatusV1(t *testing.T) {
	voice := TestClient.Voice()
	resp, err := voice.PullStatusV1(10)
	if err != nil {
		t.Log(err)
	}
	t.Log(resp)
}
