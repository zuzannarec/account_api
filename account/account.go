package account

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) CreateAccount(ctx context.Context, data Account) (*Account, error) {
	body, _ := json.Marshal(data)
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/organisation/accounts", c.baseURL), bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/vnd.api+json")

	res := Account{}
	if err := c.doRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) FetchAccount(ctx context.Context, accountID string) (*Account, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/organisation/accounts/%s", c.baseURL, accountID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.api+json")

	res := Account{}
	if err := c.doRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteAccount(ctx context.Context, accountID string, version int) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf("%s/organisation/accounts/%s?version=%d", c.baseURL, accountID, version), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.api+json")

	if err := c.doRequest(req); err != nil {
		return err
	}
	return nil
}
