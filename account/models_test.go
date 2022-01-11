package account

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelUnmarshal(t *testing.T) {
	body := []byte(`
	{
		"data": {
			"attributes": {
				"country":"GB",
				"base_currency":"GBP",
				"bank_id":"400300",
				"bank_id_code":"GBDSC",
				"account_number":"41426819",
				"bic":"NWBKGB22",
				"iban":"GB11NWBK40030041426819",
				"customer_id":"999",
				"name":["Tom Cruise"],
				"alternative_names":["Sam Holder"],
				"account_classification":"Personal",
				"joint_account":false,
				"account_matching_opt_out":false,
				"secondary_identification":"A1B2C3D4",
				"switched": false,
				"status": "pending",
				"status_reason": "unspecified",
				"user_defined_data": [
					{
					  "key":"Some account related key",
					  "value":"Some account related value"
					}
				],
				"validation_type":"card",
				"reference_mask":"############",
				"acceptance_qualifier":"same_day"
			},
			"created_on":"2021-12-21T19:40:44.721Z",
			"id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			"modified_on":"2021-12-21T19:40:44.721Z",
			"organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"type":"accounts",
			"version":0
		},
		"links": {
			"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
		}
	}`)
	v := Account{}
	err := json.Unmarshal(body, &v)
	assert.Nil(t, err, fmt.Sprintf("could not unmarshal Account json data %v", err))
}
