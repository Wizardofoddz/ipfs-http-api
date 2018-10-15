package dag

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type ErrorReader struct{}

func (e *ErrorReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("I only error")
}

var _ io.Reader = &ErrorReader{}

func TestPut(t *testing.T) {
	for name, tc := range map[string]struct {
		r      io.Reader
		f      func(http.ResponseWriter, *http.Request)
		expect string
		noErr  bool
	}{
		"happy path": {
			expect: "foo-addr",
			noErr:  true,
			f: func(w http.ResponseWriter, _ *http.Request) {
				w.Write([]byte(`{"Cid":{"/":"foo-addr"}}`))
			},
		},
		"invalid json response": {
			f: func(w http.ResponseWriter, _ *http.Request) {
				w.Write([]byte(`this is not valid JSON`))
			},
		},
		"invalid request reader": {
			r: &ErrorReader{},
		},
		"no server": {},
		"404": {
			f: func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			server := "http://notexist.example"
			if tc.f != nil {
				ts := httptest.NewServer(http.HandlerFunc(tc.f))
				defer ts.Close()
				server = ts.URL
			}

			u, err := url.Parse(server)
			if err != nil {
				t.Fatalf("error on url.Parse(): %s", err)
			}

			if tc.r == nil {
				tc.r = bytes.NewBuffer([]byte(`"foo"`))
			}

			addr, err := Put(u, tc.r)
			if err == nil && !tc.noErr {
				t.Error("expected an error, but got nil")
			}

			if err != nil && tc.noErr {
				t.Errorf("error on Put(): %s", err)
			}

			if addr != tc.expect {
				t.Errorf(`expected addr to be %q, but got %q`, tc.expect, addr)
			}
		})
	}

}
