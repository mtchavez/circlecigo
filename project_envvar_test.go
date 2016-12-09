package circleci

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClient_EnvVars_unauthorized(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/envvar", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message": "You must log in first"}`)
	})
	_, apiResp := testClient.EnvVars(testUsername, testReponame)
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful without token")
	}
	if apiResp.Response.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected unauthorized code but got %v", apiResp.Response.StatusCode)
	}
}

func TestClient_EnvVars(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/envvar", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"name": "FOO", "value": "BAR"}, {"name": "BAZ", "value": "FIZZ"}]`)
	})
	envVars, apiResp := testClient.EnvVars(testUsername, testReponame)
	if !apiResp.Success() {
		t.Errorf("Expected response to be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	if len(envVars) != 2 {
		t.Errorf("Expected to get %d envvars but got %d", 2, len(envVars))
	}
	if envVars[0].Name != "FOO" && envVars[0].Value != "BAR" {
		t.Errorf("Expected to get FOO=BAR but got %s=%s", envVars[0].Name, envVars[0].Value)
	}
}
