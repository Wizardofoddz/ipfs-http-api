package swarm

import (
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// DefaultClient is the default net/http.Client that this package will use when
// making HTTP requests
var DefaultClient = http.DefaultClient

// Peers list peers with open connections
func Peers(ipfsURL *url.URL) (io.ReadCloser, error) {
	query := url.Values{}

	swarmPeersURL := *ipfsURL
	swarmPeersURL.Path = "/api/v0/swarm/peers"
	swarmPeersURL.RawQuery = query.Encode()

	debug("Peers %v", swarmPeersURL.String())
	res, err := DefaultClient.Get(swarmPeersURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "http.Get failed")
	}

	return res.Body, nil
}
