package account

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFetchDeleteAccount(t *testing.T) {
	c, err := NewClient(WithLogger(NewStdOutLogger()))
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))

	body := []byte(`
	{
		"data": {
			"type": "accounts",
			"id": "26bab9a2-b9ec-4ab9-8fdf-e4bad7087f06",
			"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"attributes": 
			{
				"name": ["Milo Jones"],
				"country": "PL",
				"base_currency": "PLN",
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"bic": "NWBKGB22",
				"user_defined_data": [
					{
						"key": "Some account related key",
						"value": "Some account related value"
					}
				],
				"validation_type": "card",
				"reference_mask": "############",
				"acceptance_qualifier": "same_day"
			}
		}
	}`)

	account := Account{}
	err = json.Unmarshal(body, &account)
	assert.Nil(t, err, fmt.Sprintf("could not unmarshal Account json data %v", err))
	ctx := context.Background()
	_, err = c.Create(ctx, &account)
	assert.Nil(t, err, fmt.Sprintf("could not create account %v", err))

	id := "26bab9a2-b9ec-4ab9-8fdf-e4bad7087f06"

	resp, err := c.Fetch(ctx, id)
	assert.Nil(t, err, fmt.Sprintf("get request for account %s failed %v", id, err))
	version := resp.Data.Version

	err = c.Delete(ctx, id, *version)
	assert.Nil(t, err, fmt.Sprintf("get request for account %s failed %v", id, err))
}

func TestCreateAccountTwice(t *testing.T) {
	c, err := NewClient(WithLogger(NewStdOutLogger()))
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))

	body := []byte(`
	{
		"data": {
			"type": "accounts",
			"id": "26bab9a2-b9ec-4ab9-8fdf-e4bad7087f06",
			"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"attributes": {
				"name": ["Milo Jones"],
				"country": "PL",
				"base_currency": "PLN",
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"bic": "NWBKGB22",
				"user_defined_data": [
					{
						"key": "Some account related key",
						"value": "Some account related value"
					}
				],
				"validation_type": "card",
				"reference_mask": "############",
				"acceptance_qualifier": "same_day"
			}
		}
	}`)

	account := Account{}
	err = json.Unmarshal(body, &account)
	assert.Nil(t, err, fmt.Sprintf("could not unmarshal Account json data %v", err))
	ctx := context.Background()
	_, err = c.Create(ctx, &account)
	assert.Nil(t, err, fmt.Sprintf("could not create account %v", err))

	_, err = c.Create(ctx, &account)
	assert.NotNil(t, err)
	expectedErr := "status code 409"
	assert.Containsf(t, err.Error(), expectedErr, "expected error containing %q, got %s", expectedErr, err)

	id := "26bab9a2-b9ec-4ab9-8fdf-e4bad7087f06"

	resp, err := c.Fetch(ctx, id)
	assert.Nil(t, err, fmt.Sprintf("get request for account %s failed %v", id, err))
	version := resp.Data.Version

	err = c.Delete(ctx, id, *version)
	assert.Nil(t, err, fmt.Sprintf("get request for account %s failed %v", id, err))
}

func TestCreateAccountWithoutMandatoryField(t *testing.T) {
	c, err := NewClient(WithLogger(NewStdOutLogger()))
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))

	body := []byte(`
	{
		"data": {
			"type": "accounts",
			"id": "26bab9a2-b9ec-4ab9-8fdf-e4bad7087f06",
			"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"attributes": {
				"country": "PL",
				"base_currency": "PLN",
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"bic": "NWBKGB22",
				"validation_type": "card",
				"reference_mask": "############",
				"acceptance_qualifier": "same_day"
			}
		}
	}`)

	account := Account{}
	err = json.Unmarshal(body, &account)
	assert.Nil(t, err, fmt.Sprintf("could not unmarshal Account json data %v", err))
	ctx := context.Background()
	_, err = c.Create(ctx, &account)
	assert.NotNil(t, err)
	expectedErr := "status code 400"
	assert.Containsf(t, err.Error(), expectedErr, "expected error containing %q, got %s", expectedErr, err)
}

func TestCreateAccountEmptyPayload(t *testing.T) {
	c, err := NewClient(WithLogger(NewStdOutLogger()))
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))

	account := Account{}
	assert.Nil(t, err, fmt.Sprintf("could not unmarshal Account json data %v", err))
	ctx := context.Background()
	_, err = c.Create(ctx, &account)
	assert.NotNil(t, err)
	expectedErr := "status code 500"
	assert.Containsf(t, err.Error(), expectedErr, "expected error containing %q, got %s", expectedErr, err)
}

func TestCreateAccountNilPayload(t *testing.T) {
	c, err := NewClient(WithLogger(NewStdOutLogger()))
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))

	assert.Nil(t, err, fmt.Sprintf("could not unmarshal Account json data %v", err))
	ctx := context.Background()
	_, err = c.Create(ctx, nil)
	assert.NotNil(t, err)
	expectedErr := "status code 500"
	assert.Containsf(t, err.Error(), expectedErr, "expected error containing %q, got %s", expectedErr, err)
}

func TestGetNonExistentAccount(t *testing.T) {
	c, err := NewClient(WithLogger(NewStdOutLogger()))
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))

	ctx := context.Background()
	_, err = c.Fetch(ctx, "dd46b81a-c46d-4df2-954d-017a7302b38f")
	assert.NotNil(t, err)
	expectedErr := "status code 404"
	assert.Containsf(t, err.Error(), expectedErr, "expected error containing %q, got %s", expectedErr, err)
}

func TestGetInvalidAccountID(t *testing.T) {
	c, err := NewClient(WithLogger(NewStdOutLogger()))
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))

	ctx := context.Background()
	_, err = c.Fetch(ctx, "invalidID")
	assert.NotNil(t, err)
	expectedErr := "status code 400"
	assert.Containsf(t, err.Error(), expectedErr, "expected error containing %q, got %s", expectedErr, err)
}

func TestDeleteNonExistentAccount(t *testing.T) {
	c, err := NewClient(WithLogger(NewStdOutLogger()))
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))

	ctx := context.Background()
	err = c.Delete(ctx, "dd46b81a-c46d-4df2-954d-017a7302b38f", 0)
	assert.NotNil(t, err)
	expectedErr := "status code 404"
	assert.Containsf(t, err.Error(), expectedErr, "expected error containing %q, got %s", expectedErr, err)
}

func TestDeleteInvalidAccountID(t *testing.T) {
	c, err := NewClient(WithLogger(NewStdOutLogger()))
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))

	ctx := context.Background()
	err = c.Delete(ctx, "invalidID", 0)
	assert.NotNil(t, err)
	expectedErr := "status code 400"
	assert.Containsf(t, err.Error(), expectedErr, "expected error containing %q, got %s", expectedErr, err)
}
