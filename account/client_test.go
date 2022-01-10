package accountapi

import (
	"testing"
)

func TestCreateClient(t *testing.T) {
	_, err := NewClient()
	if err != nil {
		t.Fatalf("could not create client %v", err)
	}
}
