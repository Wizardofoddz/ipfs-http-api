package key

import (
	"io"
	"net/url"

	"github.com/pkg/errors"
)

// List will return a list of existing keys
func List(ipfsURL *url.URL) (io.ReadCloser, error) {
	query := url.Values{}

	keyListURL := *ipfsURL
	keyListURL.Path = "/api/v0/key/list"
	keyListURL.RawQuery = query.Encode()

	debug("Get %v", keyListURL.String())
	res, err := DefaultClient.Get(keyListURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "http.Get failed")
	}

	if res.StatusCode/100 != 2 {
		res.Body.Close()
		return nil, errors.Errorf("unsuccessful response: %s", res.Status)
	}

	return res.Body, nil
}
