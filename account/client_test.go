package account

import (
	"testing"
)

func TestClient(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
		})
	}
}

func TestCreateClient(t *testing.T) {
	_, err := NewClient()
	if err != nil {
		t.Fatalf("could not create client %v", err)
	}
}
