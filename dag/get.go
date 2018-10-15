package dag

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

// Get retrieves a dag object from IPFS
func Get(ipfsURL *url.URL, address string) (io.ReadCloser, error) {
	query := url.Values{}
	query.Add("arg", address)

	dagGetURL := *ipfsURL
	dagGetURL.Path = "/api/v0/dag/get"
	dagGetURL.RawQuery = query.Encode()

	debug("Get %v", dagGetURL.String())
	res, err := DefaultClient.Get(dagGetURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "Get failed")
	}
	if res.StatusCode/100 != 2 {
		res.Body.Close()
		return nil, errors.Errorf("unsuccessful response: %s", res.Status)
	}
	return res.Body, nil
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

// GetInterface retrieves a dag object from IPFS and parses it into
// the provided interface
func GetInterface(ipfsURL *url.URL, address string, t interface{}) error {
	buf, err := GetBytes(ipfsURL, address)
	if err != nil {
		return errors.Wrap(err, "DAG.Getbytes failed")
	}

	return json.Unmarshal(buf, t)
}
