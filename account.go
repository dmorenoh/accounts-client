package accountsclient

import (
	"bytes"
	"encoding/json"
)

type AccountClassificationType string

const (
	PERSONAL AccountClassificationType = "Personal"
	BUSINESS AccountClassificationType = "Business"
)

type AccountResponse struct {
	Data AccountResource `json:"data"`
}

type CreateAccountRequest struct {
	Data AccountResource `json:"data"`
}

func (req *CreateAccountRequest) toBuffer() *bytes.Buffer {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(req)
	return buf
}

type CreateAccountCommand struct {
	ID                      string
	Country                 string
	BaseCurrency            string
	BankID                  string
	BankIDCode              string
	Bic                     string
	Iban                    string
	CustomerID              string
	Name                    string
	AlternativeNames        string
	AccountClassification   AccountClassificationType
	JointAccount            bool
	AccountMatchingOptOut   bool
	SecondaryIdentification string
	Switched                bool
}

type DeleteAccountCommand struct {
	AccountID string
	Version   int
}

type AccountResources struct {
	Data []AccountResource `json:"data"`
}

type AccountResource struct {
	Type           string     `json:"type"`
	ID             string     `json:"id"`
	OrganisationID string     `json:"organisation_id"`
	Version        int        `json:"version"`
	Attributes     Attributes `json:"attributes"`
}

type Attributes struct {
	Country                 string   `json:"country"`
	BaseCurrency            string   `json:"base_currency"`
	AccountNumber           string   `json:"account_number"`
	BankID                  string   `json:"bank_id"`
	BankIDCode              string   `json:"bank_id_code"`
	Bic                     string   `json:"bic"`
	Iban                    string   `json:"iban"`
	Name                    string   `json:"bank_account_name"`
	CustomerID              string   `json:"customer_id"`
	AlternativeNames        []string `json:"alternative_bank_account_names"`
	AccountClassification   string   `json:"account_classification"`
	JointAccount            bool     `json:"joint_account"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out"`
	SecondaryIdentification string   `json:"secondary_identification"`
	Switched                bool     `json:"switched"`
	Status                  string   `json:"status"`
}

func (req *CreateAccountRequest) toJSON() *bytes.Buffer {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(req)
	return buf
}

func (cmd *CreateAccountCommand) toRequest() CreateAccountRequest {
	newAccountResource := AccountResource{
		OrganisationID: MyOrganizationID,
		Type:           AccountType,
		ID:             cmd.ID,
		Attributes: Attributes{
			Country:                 cmd.Country,
			BaseCurrency:            cmd.BaseCurrency,
			BankID:                  cmd.BankID,
			BankIDCode:              cmd.BankIDCode,
			Bic:                     cmd.Bic,
			Name:                    cmd.Name,
			CustomerID:              cmd.CustomerID,
			AccountClassification:   string(cmd.AccountClassification),
			JointAccount:            cmd.JointAccount,
			AccountMatchingOptOut:   cmd.AccountMatchingOptOut,
			SecondaryIdentification: cmd.SecondaryIdentification,
			Switched:                cmd.Switched,
		},
	}
	return CreateAccountRequest{Data: newAccountResource}
}
