package accountapi

// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
type AccountErrorResponse struct {
	ErrorMessage *string `json:"error_message,omitempty"`
	ErrorCode    *string `json:"error_code,omitempty"`
}

type Account struct {
	Data *AccountData `json:"data,omitempty"`
}
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}
type AccountUserDefinedData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type AccountAttributes struct {
	Country                 *string                  `json:"country,omitempty"`
	BaseCurrency            string                   `json:"base_currency,omitempty"`
	BankID                  string                   `json:"bank_id,omitempty"`
	BankIDCode              string                   `json:"bank_id_code,omitempty"`
	AccountNumber           string                   `json:"account_number,omitempty"`
	Bic                     string                   `json:"bic,omitempty"`
	Iban                    string                   `json:"iban,omitempty"`
	CustomerID              string                   `json:"customer_id,omitempty"`
	Name                    []string                 `json:"name,omitempty"`
	AlternativeNames        []string                 `json:"alternative_names,omitempty"`
	AccountClassification   *string                  `json:"account_classification,omitempty"`
	JointAccount            *bool                    `json:"joint_account,omitempty"`
	AccountMatchingOptOut   *bool                    `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification string                   `json:"secondary_identification,omitempty"`
	Switched                *bool                    `json:"switched,omitempty"`
	Status                  *string                  `json:"status,omitempty"`
	StatusReason            *string                  `json:"status_reason,omitempty"`
	UserDefinedData         []AccountUserDefinedData `json:"user_defined_data,omitempty"`
	ValidationType          *string                  `json:"validation_type,omitempty"`
	ReferenceMask           *string                  `json:"reference_mask,omitempty"`
	AcceptanceQualifier     *string                  `json:"acceptance_qualifier,omitempty"`
}
