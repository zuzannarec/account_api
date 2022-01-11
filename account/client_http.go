package account

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	reqTimeout time.Duration
	logger     Logger
	baseURL    *url.URL
	netClient  *http.Client
}

type Option func(*Client) error

func WithReqTimeout(timeout time.Duration) func(*Client) error {
	return func(c *Client) error {
		c.reqTimeout = timeout
		return nil
	}
}

func WithBaseURL(u string) func(*Client) error {
	return func(c *Client) error {
		parsed, err := url.Parse(u)
		if err != nil {
			return fmt.Errorf("invalid URL %s", u)
		}
		c.baseURL = parsed
		return nil
	}
}

func WithLogger(l Logger) func(*Client) error {
	return func(c *Client) error {
		c.logger = l
		return nil
	}
}

func NewClient(opts ...Option) (*Client, error) {
	host, _ := url.Parse("http://:8080/v1")
	c := &Client{
		reqTimeout: 100 * time.Millisecond,
		logger:     NewDefaultLogger(),
		baseURL:    host,
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, fmt.Errorf("failed to set option %w", err)
		}
	}
	c.netClient = &http.Client{
		Timeout: c.reqTimeout,
	}
	return c, nil
}

func (client *Client) doRequest(req *http.Request, body *Account) error {
	response, err := client.netClient.Do(req)
	if err != nil {
		client.logger.Debugf("request failed %w", err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= http.StatusBadRequest {
		errResp, err := io.ReadAll(response.Body)
		if err == nil && len(errResp) > 0 {
			client.logger.Debugf("request failed, status code %d, error response %v",
				response.StatusCode, string(errResp))
			return fmt.Errorf("request failed, status code %d, error response %v",
				response.StatusCode, string(errResp))
		}
		client.logger.Debugf("request failed with unknown error, status code %d", response.StatusCode)
		return fmt.Errorf("request failed with unknown error, status code %d", response.StatusCode)
	}

	if body == nil {
		return nil
	}

	if err := json.NewDecoder(response.Body).Decode(body); err != nil {
		client.logger.Debugf("could not decode response %w", err)
		return err
	}

	return nil
}
