package dag

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestResolve(t *testing.T) {
	for name, tc := range map[string]struct {
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
		"404": {
			f: func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
		},
		"invalid json response": {
			f: func(w http.ResponseWriter, _ *http.Request) {
				w.Write([]byte(`this is not valid JSON`))
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

			addr, err := Resolve(u, tc.expect)
			if err == nil && !tc.noErr {
				t.Error("expected an error, but got nil")
			}

			if err != nil && tc.noErr {
				t.Errorf("error on Resolve(): %s", err)
			}

			if addr != tc.expect {
				t.Errorf(`expected addr to be %q, but got %q`, tc.expect, addr)
			}
		})
	}
}
