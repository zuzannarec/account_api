package accountapi

import (
	"encoding/json"
	"testing"
)

func TestModelUnmarshal(t *testing.T) {
	body := []byte(`{"data":{"attributes":{"alternative_names":null,"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","name":["Tom Cruise"]},"created_on":"2021-12-21T19:40:44.721Z","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-12-21T19:40:44.721Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}}`)
	v := Account{}
	err := json.Unmarshal(body, &v)
	if err != nil {
		t.Fatalf("could not unmarshal Account json data %v", err)
	}
}
