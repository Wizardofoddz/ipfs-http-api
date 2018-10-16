package key

import (
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// DefaultClient is the default net/http.Client that this package will use when
// making HTTP requests
var DefaultClient = http.DefaultClient

// Gen will create a new IPFS key
func Gen(ipfsURL *url.URL, name string) (io.ReadCloser, error) {
	query := url.Values{}
	query.Add("arg", name)
	query.Add("type", "ed25519")

	keyGenURL := *ipfsURL
	keyGenURL.Path = "/api/v0/key/gen"
	keyGenURL.RawQuery = query.Encode()

	debug("Get %v", keyGenURL.String())
	res, err := DefaultClient.Get(keyGenURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "http.Get failed")
	}

	if res.StatusCode/100 != 2 {
		res.Body.Close()
		return nil, errors.Errorf("unsuccessful response: %s", res.Status)
	}

	return res.Body, nil
}
