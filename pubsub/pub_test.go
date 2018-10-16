package pubsub

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestPublish(t *testing.T) {
	for name, tt := range map[string]struct {
		f       func(http.ResponseWriter, *http.Request)
		topic   string
		payload string
		wantErr bool
	}{
		"happy path": {
			f: func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(tt.f))
			defer ts.Close()

			u, _ := url.Parse(ts.URL)
			if err := Publish(u, tt.topic, tt.payload); (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
