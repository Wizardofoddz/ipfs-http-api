package dag

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	expected := `"foo"`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expected))
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("error on url.Parse(): %s", err)
	}

	r, err := Get(u, "foo-addr")
	if err != nil {
		t.Fatalf("error on Cat(): %s", err)
	}
	defer r.Close()

	var message json.RawMessage
	if err := json.NewDecoder(r).Decode(&message); err != nil {
		t.Fatalf("error on decoder.Decode(): %s", err)
	}

	if got := string(message); got != expected {
		t.Fatalf("Expected body == %q, Actual body == %q", expected, got)
	}
}

func TestGetBytes(t *testing.T) {
	expected := `"foo"`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expected))
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("error on url.Parse(): %s", err)
	}

	message, err := GetBytes(u, "foo-addr")
	if err != nil {
		t.Fatalf("error on GetBytes(): %s", err)
	}

	if got := string(message); got != expected {
		t.Fatalf("Expected body == %q, Actual body == %q", expected, got)
	}
}

func TestGetInterface(t *testing.T) {
	expected := "foo"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(expected)
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("error on url.Parse(): %s", err)
	}

	var got string
	if err := GetInterface(u, "foo-addr", &got); err != nil {
		t.Fatalf("error on GetInterface(): %s", err)
	}

	if got != expected {
		t.Fatalf("Expected body == %q, Actual body == %q", expected, got)
	}
}
