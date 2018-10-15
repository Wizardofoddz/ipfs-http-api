package dag

import (
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// Resolve resolves a ipld reference in IPFS
func Resolve(ipfsURL *url.URL, address string) (string, error) {
	query := url.Values{}
	query.Add("arg", address)

	dagResolveURL := *ipfsURL
	dagResolveURL.Path = "/api/v0/dag/resolve"
	dagResolveURL.RawQuery = query.Encode()

	debug("Resolve %v", dagResolveURL.String())
	res, err := DefaultClient.Get(dagResolveURL.String())
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return "", errors.Errorf("unsuccessful response: %s", res.Status)
	}

	resolveResponse := struct {
		Cid struct {
			Address string `json:"/"`
		}
	}{}

	if err := json.NewDecoder(res.Body).Decode(&resolveResponse); err != nil {
		return "", err
	}

	return resolveResponse.Cid.Address, nil
}
