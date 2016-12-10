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

func TestClient_GetCheckoutKey_notFound(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/checkout-key/not-real", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"message": "checkout key not found"}`)
	})
	key, apiResp := testClient.GetCheckoutKey(testUsername, testReponame, "not-real")
	if apiResp.Success() {
		t.Errorf("Expected response not to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status not found but got %v", apiResp.Response.StatusCode)
	}
	if apiResp.ErrorResponse.Message != "checkout key not found" {
		t.Errorf("Expected checkout key not found but got %v", apiResp.ErrorResponse.Message)
	}
	if key.Type == "deploy-key" {
		t.Errorf("Expected not to get type deploy-key")
	}
}

func TestClient_GetCheckoutKey(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	testFingerprint := "c9:0b:1c:4f:d5:65:56:b9:ad:88:f9:81:2b:37:74:2f"
	path := fmt.Sprintf("/project/%s/%s/checkout-key/%s", testUsername, testReponame, testFingerprint)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"public_key": "ssh-rsa ...", "fingerprint": "%s", "type": "deploy-key"}`, testFingerprint)
	})
	key, apiResp := testClient.GetCheckoutKey(testUsername, testReponame, testFingerprint)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if key.Type != "deploy-key" {
		t.Errorf("Expected key type to be deploy-key but got %s", key.Type)
	}
	if key.Fingerprint != testFingerprint {
		t.Errorf("Expected key type to be %s but got %s", testFingerprint, key.Fingerprint)
	}
}
