package account

import (
	"fmt"
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
		if timeout > 1000 {
			return fmt.Errorf("invalid timeout value %s", timeout)
		}
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
	host, _ := url.Parse("http://:8080")
	c := &Client{
		reqTimeout: 10,
		logger:     NewDefaultLogger(),
		baseURL:    host,
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return &Client{}, fmt.Errorf("failed to set option %w", err)
		}
	}
	c.netClient = &http.Client{
		Timeout: c.reqTimeout,
	}
	return c, nil
}
