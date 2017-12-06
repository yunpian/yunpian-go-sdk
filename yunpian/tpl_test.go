package yunpian

import (
	"testing"
)

func TestTPLGet(t *testing.T) {
	tpl := TestClient.TPL()
	input := &TPLGetRequest{ID: 2077620}
	resp, err := tpl.Get(input)
	if err != nil {
		t.Error(err)
	}

	t.Log(resp.Content)
}

func TestTPLList(t *testing.T) {
	tpl := TestClient.TPL()
	resp, err := tpl.List()
	if err != nil {
		t.Error(err)
	}

	t.Log(len(resp))
}

func TestTPLListDefault(t *testing.T) {
	tpl := TestClient.TPL()
	resp, err := tpl.ListDefault()
	if err != nil {
		t.Error(err)
	}

	t.Log(len(resp))
}
