package pin

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// DefaultClient is the default net/http.Client that this package will use when
// making HTTP requests
var DefaultClient = http.DefaultClient

// Add pins a an IPFS object recursively
func Add(ipfsURL *url.URL, address string) error {
	query := url.Values{}
	query.Add("arg", address)

	pinAddURL := *ipfsURL
	pinAddURL.Path = "/api/v0/pin/add"
	pinAddURL.RawQuery = query.Encode()

	debug("Add %v", pinAddURL.String())
	resp, err := DefaultClient.Get(pinAddURL.String())
	if err != nil {
		return errors.Wrap(err, "http.Get failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return errors.Errorf("unsuccessful response: %s", resp.Status)
	}

	return nil
}
