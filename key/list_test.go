package key

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestList(t *testing.T) {
	expected := `{"Keys":[]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(expected))
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("error on url.Parse(): %s", err)
	}

	r, err := List(u)
	if err != nil {
		t.Fatalf("error on List(): %s", err)
	}
	defer r.Close()

	var message json.RawMessage
	if err := json.NewDecoder(r).Decode(&message); err != nil {
		t.Fatalf("error on decoder.Decode(): %s", err)
	}

	if got := string(message); got != expected {
		t.Fatalf("Expected %s, but got %s", expected, got)
	}
}
