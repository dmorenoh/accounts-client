package accountsclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountInvalidRequest(t *testing.T) {
	// given a valid account creation request
	cmd := CreateAccountCommand{
		Country:               "GB",
		BaseCurrency:          "GBP",
		Name:                  "David Moreno",
		AccountClassification: "Invalid",
	}

	// when request account creation
	accountClient := NewAccountApiClient()
	res, err := accountClient.createAccount(cmd)

	// then
	assert.Nil(t, res, "Not null")
	assert.NotNil(t, err, "Null")
	assert.Equal(t, err.Code, 400, "expeced invalid request")
}

func TestCreateAccountValidRequest(t *testing.T) {
	// given a valid account creation request
	cmd := CreateAccountCommand{
		Country:               "GB",
		BaseCurrency:          "GBP",
		Name:                  "David Moreno",
		AccountClassification: "Personal",
	}

	// when request account creation
	accountClient := NewAccountApiClient()
	res, err := accountClient.createAccount(cmd)

	// then
	assert.NotNil(t, res, "Not null")
	assert.Nil(t, err, "Null")
	assert.NotNil(t, res.Attributes, "Not null")
	assert.Equal(t, MyOrganizationID, res.OrganisationID, "unexpected OrganisationId")
	assert.NotNil(t, res.ID, "unexpected Id")
	assert.Equal(t, cmd.Name, res.Attributes.Name)
}

func TestFetchExistingAccount(t *testing.T) {

	// given an existing account
	req := CreateAccountCommand{
		Country:               "GB",
		BaseCurrency:          "GBP",
		Name:                  "David Moreno",
		AccountClassification: "Personal",
	}

	accountClient := NewAccountApiClient()
	res, _ := accountClient.createAccount(req)

	// when fetching this account
	res2, err2 := accountClient.fetchAccount(res.ID)
	assert.NotNil(t, res2, "Not null")
	assert.Zero(t, err2, "Null")
	assert.Equal(t, res.ID, res2.ID, "unexpected Id")
}

func TestFetchNonExistigAccount(t *testing.T) {

	// when fetching this account
	accountClient := NewAccountApiClient()
	res, err := accountClient.fetchAccount("1111")

	assert.Nil(t, res, "Not null")
	assert.NotNil(t, err, "Null")
	assert.Equal(t, err.Code, 400, "unexpected Id")
}

func TestListAccounts(t *testing.T) {

	accountClient := NewAccountApiClient()
	pageOpt := PageOptions{
		Number: 0,
		Size:   100,
	}
	res, err := accountClient.list(pageOpt)
	assert.NotNil(t, res, "Not null")
	assert.Nil(t, err, "Null")
}

func TestDeleteExisitingAccount(t *testing.T) {

	// given an existing account
	req := CreateAccountCommand{
		Country:               "GB",
		BaseCurrency:          "GBP",
		Name:                  "David Moreno",
		AccountClassification: "Personal",
	}

	accountClient := NewAccountApiClient()
	res, _ := accountClient.createAccount(req)

	resf, _ := accountClient.fetchAccount(res.ID)
	assert.NotNil(t, resf, "Not null")

	// when request to delete
	cmd := DeleteAccountCommand{
		AccountID: resf.ID,
		Version:   resf.Version,
	}
	errD := accountClient.delete(cmd)
	assert.Nil(t, errD, "Null")
}
