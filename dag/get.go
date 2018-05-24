package dag

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/pkg/errors"

	"github.com/computes/ipfs-http-api/http"
)

// Get retrieves a dag object from IPFS
func Get(ipfsURL *url.URL, address string) (io.ReadCloser, error) {
	query := url.Values{}
	query.Add("arg", address)

	dagGetURL := *ipfsURL
	dagGetURL.Path = "/api/v0/dag/get"
	dagGetURL.RawQuery = query.Encode()

	debug("Get %v", dagGetURL.String())
	res, err := http.Get(dagGetURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "Get failed")
	}
	return res, nil
}

// GetBytes retrieves a dag object from IPFS and reads the whole buffer
// into memory
func GetBytes(ipfsURL *url.URL, address string) ([]byte, error) {
	reader, err := Get(ipfsURL, address)
	if reader != nil {
		defer reader.Close()
	}
	if err != nil {
		return nil, errors.Wrap(err, "GetBytes failed")
	}

	message := json.RawMessage{}
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&message)
	if err != nil {
		return nil, errors.Wrap(err, "json.Decode failed")
	}

	return []byte(message), nil
}
