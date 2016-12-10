package circleci

import (
	"fmt"
	"net/http"
)

// CheckoutKey - CircleCI project checkout key API response
type CheckoutKey struct {
	PublicKey   string `json:"public_key"`
	Type        string `json:"type"`
	Fingerprint string `json:"fingerprint"`
	Preferred   bool   `json:"preferred"`
	Time        string `json:"time"`
}

// CheckoutKeys calls the /project/:username/:project/checkout-key endpoint
// to get a list of all the checkout keys for a project
func (client *Client) CheckoutKeys(username, project string) ([]*CheckoutKey, *APIResponse) {
	keys := []*CheckoutKey{}
	path := fmt.Sprintf("/project/%s/%s/checkout-key", username, project)
	apiResp := client.request(http.MethodGet, path, nil, nil, &keys)
	return keys, apiResp
}