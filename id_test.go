package ipfs

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestID(t *testing.T) {
	expected := `"foo"`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expected))
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("error on url.Parse(): %s", err)
	}

	r, err := ID(u)
	if err != nil {
		t.Fatalf("error on ID(): %s", err)
	}
	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("error on ioutil.ReadAll(): %s", err)
	}

	if string(body) != expected {
		t.Fatalf("Expected body == %s, Actual body == %q", expected, body)
	}
}

func TestIDBytes(t *testing.T) {
	expected := `"foo"`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expected))
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("error on url.Parse(): %s", err)
	}

	body, err := IDBytes(u)
	if err != nil {
		t.Fatalf("error on ID(): %s", err)
	}

	if string(body) != expected {
		t.Fatalf("Expected body == %s, Actual body == %q", expected, body)
	}
}
