package yunpian

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client provides a client to the YunPian API
type Client struct {
	config Config
}

// NewClient returns a new client
func NewClient(config *Config) *Client {
	cfg := DefaultConfig()
	if config != nil {
		cfg.MergeIn(config)
	}

	return &Client{config: *cfg}
}

// request is used to help build up a request
type request struct {
	config *Config
	method string
	url    *url.URL
	params url.Values
	body   io.Reader
	header http.Header
	ctx    context.Context
}

func (c *Client) newRequest(method, endpoint, path string) *request {
	r := &request{
		config: &c.config,
		method: method,
		params: make(map[string][]string),
		header: make(http.Header),
		ctx:    c.config.Context,
	}

	u := &url.URL{
		Host: endpoint,
		Path: path,
	}

	if *c.config.UseSSL {
		u.Scheme = "https"
	} else {
		u.Scheme = "http"
	}

	r.url = u
	return r
}

func (c *Client) doRequest(r *request) (*http.Response, error) {
	req, err := r.toHTTP()
	if err != nil {
		return nil, err
	}

	return c.config.HTTPClient.Do(req)
}

func (r *request) toHTTP() (*http.Request, error) {
	r.url.RawQuery = r.params.Encode()

	req, err := http.NewRequest(r.method, r.url.RequestURI(), r.body)
	if err != nil {
		return nil, err
	}

	req.URL.Host = r.url.Host
	req.URL.Scheme = r.url.Scheme
	req.Host = r.url.Host
	req.Header = r.header

	req.Header.Set("Api-Lang", "go")
	req.Header.Set("Connection", "Keep-Alive")

	if r.ctx != nil {
		return req.WithContext(r.ctx), nil
	}

	return req, nil
}

// encodeFormBody is used to Form encode a request body
func (c *Client) encodeFormBody(obj interface{}) (io.Reader, error) {
	encoder := NewEncoder()
	form := url.Values{}
	err := encoder.Encode(obj, form)
	if err != nil {
		return nil, err
	}

	form.Set("apikey", *c.config.APIKey)
	return strings.NewReader(form.Encode()), nil
}

func (c *Client) handleResponse(resp *http.Response, out interface{}) error {
	if resp.StatusCode >= http.StatusBadRequest {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("请求失败：%s", string(body))
	}

	return c.decodeJSONBody(resp, out)
}

// decodeJSONBody is used to JSON decode a body
func (c *Client) decodeJSONBody(resp *http.Response, out interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return fmt.Errorf("解析响应体失败，原始响应体为：%s", string(body))
	}
	return nil
}
