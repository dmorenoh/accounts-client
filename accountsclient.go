package accountsclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BaseURLV1 : base url v1
const (
	BaseURLV1        = "http://localhost:8080/v1/organisation"
	MyOrganizationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	AccountType      = "accounts"
)

// Client : client type
type Client struct {
	baseURL    string
	HTTPClient *http.Client
}

func NewAccountApiClient() *Client {

	return &Client{
		baseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type successResponse struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"error_message"`
}

func (er *errorResponse) Error() string {
	return er.Message
}

type PageOptions struct {
	Number int `json:"page[number]"`
	Size   int `json:"page[size]"`
}

func (c *Client) createAccount(cmd CreateAccountCommand) (*AccountResource, *errorResponse) {

	reqBody := cmd.toRequest()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/accounts", c.baseURL), reqBody.toBuffer())

	if err != nil {
		return nil, NewErrorResponse(err)
	}

	res := AccountResource{}

	errorResp := c.sendRequest(req, &res)

	if errorResp != nil {
		return nil, errorResp
	}

	return &res, nil
}

func (c *Client) fetchAccount(id string) (*AccountResource, *errorResponse) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts/%s", c.baseURL, id), nil)

	if err != nil {
		return nil, NewErrorResponse(err)
	}

	res := AccountResource{}

	errorResp := c.sendRequest(req, &res)

	if errorResp != nil {
		return nil, errorResp
	}

	return &res, nil
}

func (c *Client) list(pageOptions PageOptions) (*AccountResources, *errorResponse) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/accounts?page[number]=%d&page[size]=%d", c.baseURL, pageOptions.Number, pageOptions.Size), nil)

	if err != nil {
		return nil, NewErrorResponse(err)
	}

	res := AccountResources{}

	errorResp := c.sendRequest(req, &res.Data)

	if errorResp != nil {
		return nil, errorResp
	}

	return &res, nil
}

func (c *Client) delete(r DeleteAccountCommand) *errorResponse {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/accounts/%s?version=%d", c.baseURL, r.AccountID, r.Version), nil)

	if err != nil {
		return NewErrorResponse(err)
	}

	errorResp := c.sendRequest(req, nil)

	if errorResp != nil {
		return errorResp
	}

	return nil
}

func (c *Client) sendRequest(req *http.Request, responseBody interface{}) *errorResponse {

	req.Header.Set("Content-Type", "application/vnd.api+json")

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return NewErrorResponse(err)
	}

	defer res.Body.Close()

	if isHTTPErrorCode(res) {
		return NewHttpErrorResponse(res)
	}

	if responseBody == nil {
		return nil
	}

	fullResponse := successResponse{
		Data: responseBody,
	}
	if err := json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return NewErrorResponse(err)
	}
	return nil
}

func NewErrorResponse(e error) *errorResponse {
	return &errorResponse{
		Message: e.Error(),
	}
}

func NewHttpErrorResponse(res *http.Response) *errorResponse {
	errRes := errorResponse{
		Code:    res.StatusCode,
		Message: BodyAsString(res),
	}
	return &errRes
}

func BodyAsString(r *http.Response) string {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	return bodyString
}

func isHTTPErrorCode(r *http.Response) bool {

	if r.StatusCode < http.StatusOK || r.StatusCode >= http.StatusBadRequest {
		return true
	}
	return false
}
