package circleci

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClient_CheckoutKeys(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/checkout-key", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{"public_key": "ssh-rsa ...", "type": "deploy-key"},
			{"public_key": "BAZ", "type": "deploy-key"}
		]`)
	})
	keys, apiResp := testClient.CheckoutKeys(testUsername, testReponame)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if len(keys) != 2 {
		t.Errorf("Expected to get %d checkout-keys but got %d", 2, len(keys))
	}
	if keys[0].Type != "deploy-key" {
		t.Errorf("Expected to get type deploy-key but got %s", keys[0].Type)
	}
}
