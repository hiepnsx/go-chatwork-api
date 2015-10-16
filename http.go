package gochatwork

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var apiVersion = "/v1/"

type apiConnection interface {
	Get(endPoint string, params url.Values, config *config) ([]byte, error)
	Post(endPoint string, params url.Values, config *config) ([]byte, error)
}

// http interface
type httpImp struct {
}

func (h *httpImp) Get(endPoint string, params url.Values, config *config) ([]byte, error) {
	return h.connection("GET", endPoint, params, config)
}

func (h *httpImp) Post(endPoint string, params url.Values, config *config) ([]byte, error) {
	return h.connection("POST", endPoint, params, config)
}

func (h *httpImp) connection(method string, endPoint string, params url.Values, config *config) ([]byte, error) {
	if config == nil || config.token == "" {
		return make([]byte, 0), fmt.Errorf("No auth token")
	}

	req, _ := http.NewRequest(method, config.url+apiVersion+endPoint, nil)
	req.Header.Set("X-ChatWorkToken", config.token)
	req.URL.RawQuery = params.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
