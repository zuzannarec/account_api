package account

import (
	"context"
	"testing"
)

func Test(t *testing.T) {
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

func TestGetAccount(t *testing.T) {
	c, err := NewClient(WithLogger(NewStdOutLogger()))
	if err != nil {
		t.Fatalf("could not create client %v", err)
	}
	ctx := context.Background()
	_, err = c.FetchAccount(ctx, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	if err != nil {
		t.Fatalf("request failed %v", err)
	}
}
