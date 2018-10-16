package ipfs

import (
	"io"
	"net/url"

	"github.com/pkg/errors"
)

// Cat returns a reader for the data in IPFS located at address
func Cat(ipfsURL *url.URL, address string) (io.ReadCloser, error) {
	query := url.Values{}
	query.Add("arg", address)

	catURL := *ipfsURL
	catURL.Path = "/api/v0/cat"
	catURL.RawQuery = query.Encode()

	debug("Cat %v", catURL.String())
	res, err := DefaultClient.Get(catURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "http.Get failed")
	}

	if res.StatusCode/100 != 2 {
		res.Body.Close()
		return nil, errors.Errorf("unsuccessful response: %s", res.Status)
	}

	return res.Body, nil
}
