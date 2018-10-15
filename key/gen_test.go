package key

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGen(t *testing.T) {
	type ret struct{ Name, ID string }
	expected := ret{Name: "foo", ID: "bar"}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(expected)
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("error on url.Parse(): %s", err)
	}

	r, err := Gen(u, "foo")
	if err != nil {
		t.Fatalf("error on Gen(): %s", err)
	}
	defer r.Close()

	var got ret
	if err := json.NewDecoder(r).Decode(&got); err != nil {
		t.Fatalf("error on decoder.Decode(): %s", err)
	}

	if got != expected {
		t.Fatalf("Expected %v, but got %v", expected, got)
	}
}
