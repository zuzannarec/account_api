package account

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	createAccountPath = "%s/organisation/accounts"
	fetchAccountPath  = "%s/organisation/accounts/%s"
	deleteAccountPath = "%s/organisation/accounts/%s?version=%d"
)

func (c *Client) CreateAccount(ctx context.Context, data *Account) (*Account, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf(createAccountPath, c.baseURL), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/vnd.api+json")

	res := &Account{}
	if err := c.doRequest(req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) FetchAccount(ctx context.Context, accountID string) (*Account, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(fetchAccountPath, c.baseURL, accountID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.api+json")

	res := &Account{}
	if err := c.doRequest(req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteAccount(ctx context.Context, accountID string, version int) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf(deleteAccountPath, c.baseURL, accountID, version), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.api+json")

	if err := c.doRequest(req, nil); err != nil {
		return err
	}
	return nil
}
