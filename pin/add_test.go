package pin

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestAdd(t *testing.T) {
	for name, tc := range map[string]struct {
		f     func(http.ResponseWriter, *http.Request)
		noErr bool
	}{
		"happy path": {
			noErr: true,
			f: func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
		},
		"404": {
			f: func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(tc.f))
			u, _ := url.Parse(ts.URL)

			err := Add(u, "foo-addr")
			if tc.noErr && err != nil {
				t.Errorf("unexpected error on Add(): %s", err)
			}
			if !tc.noErr && err == nil {
				t.Error("expected an error, but got none")
			}
		})

	}
}
