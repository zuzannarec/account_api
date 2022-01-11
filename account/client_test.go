package account

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClientNoOpts(t *testing.T) {
	_, err := NewClient()
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))
}

func TestCreateClientOpts(t *testing.T) {
	_, err := NewClient(WithReqTimeout(100), WithBaseURL("localhost:8080"), WithLogger(NewDefaultLogger()))
	assert.Nil(t, err, fmt.Sprintf("could not create client %v", err))
}

func TestCreateClientInvalidBaseURL(t *testing.T) {
	invalidURL := "://invalid"
	c, err := NewClient(WithBaseURL(invalidURL))
	assert.Nil(t, c)
	assert.NotNil(t, err)
}
