package ipfs

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// DefaultClient is the default net/http.Client that this package will use when
// making HTTP requests
var DefaultClient = http.DefaultClient

// ID returns a reader of the IPFS node info
func ID(ipfsURL *url.URL) (io.ReadCloser, error) {
	idURL := *ipfsURL
	idURL.Path = "/api/v0/id"

	debug("ID %v", idURL.String())
	res, err := DefaultClient.Get(idURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "http.Get failed")
	}

	if res.StatusCode/100 != 2 {
		res.Body.Close()
		return nil, errors.Errorf("unsuccessful response: %s", res.Status)
	}

	return res.Body, nil
}

// IDBytes returns the IPFS node info as bytes
func IDBytes(ipfsURL *url.URL) ([]byte, error) {
	r, err := ID(ipfsURL)
	if err != nil {
		return nil, errors.Wrap(err, "IDBytes failed")
	}
	defer r.Close()

	var message json.RawMessage
	if err := json.NewDecoder(r).Decode(&message); err != nil {
		return nil, errors.Wrap(err, "json.Decode failed")
	}

	return []byte(message), nil
}
