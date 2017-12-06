package yunpian

import (
	"testing"
)

func TestVoicePullStatus(t *testing.T) {
	voice := TestClient.Voice()
	input := &VoicePullStatusRequest{
		PageSize: 10,
	}
	resp, err := voice.PullStatus(input)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
