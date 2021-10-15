package yunpian

import (
	"testing"
)

func TestFlowGetPackage(t *testing.T) {
	flow := TestClient.Flow()
	resp, err := flow.GetPackageV1("")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestFlowPullStatus(t *testing.T) {
	flow := TestClient.Flow()
	resp, err := flow.PullStatusV1(10)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
