package pubsub

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// DefaultClient is the default net/http.Client that this package will use when
// making HTTP requests
var DefaultClient = http.DefaultClient

// Publish will publish the content to a given URL
func Publish(ipfsURL *url.URL, topic, payload string) error {
	query := url.Values{}
	query.Add("arg", topic)
	query.Add("arg", payload)

	pubURL := *ipfsURL
	pubURL.Path = "/api/v0/pubsub/pub"
	pubURL.RawQuery = query.Encode()

	debug("Publish %v", pubURL.String())
	res, err := DefaultClient.Get(pubURL.String())
	if err != nil {
		return errors.Wrap(err, "http.Get failed")
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return errors.Errorf("unsuccessful response: %s", res.Status)
	}

	return nil
}
