package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	Headers    map[string]string
	HTTPClient *http.Client
}

type RequestOptions struct {
	Method  string
	URL     string
	Body    interface{}
	Headers map[string]string
}

type Response struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}

func NewHttpClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
		},
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *Client) Request(opts RequestOptions) (*Response, error) {
	url := opts.URL
	if c.BaseURL != "" {
		url = c.BaseURL + url
	}

	var bodyReader io.Reader
	if opts.Body != nil {
		jsonBody, err := json.Marshal(opts.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(opts.Method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}
	for key, value := range opts.Headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    resp.Header,
	}, nil
}

func (c *Client) Get(url string, headers map[string]string) (*Response, error) {
	return c.Request(RequestOptions{
		Method:  "GET",
		URL:     url,
		Headers: headers,
	})
}

func (c *Client) Post(url string, body interface{}, headers map[string]string) (*Response, error) {
	return c.Request(RequestOptions{
		Method:  "POST",
		URL:     url,
		Body:    body,
		Headers: headers,
	})
}

func (c *Client) Put(url string, body interface{}, headers map[string]string) (*Response, error) {
	return c.Request(RequestOptions{
		Method:  "PUT",
		URL:     url,
		Body:    body,
		Headers: headers,
	})
}

func (c *Client) Delete(url string, headers map[string]string) (*Response, error) {
	return c.Request(RequestOptions{
		Method:  "DELETE",
		URL:     url,
		Headers: headers,
	})
}
