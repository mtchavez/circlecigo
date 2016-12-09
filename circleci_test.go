package circleci

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/joho/godotenv"
)

const TestToken = "abcd-1234-abcd-1234"

var (
	testMux      *http.ServeMux
	testServer   *httptest.Server
	testClient   *Client
	testUsername = "mtchavez"
	testReponame = "circlecigo"
	testBuildNum = 1
)

func init() {
	godotenv.Load()
}

func startTestServer() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	testClient = defaultClient
	defaultClient.BaseURL, _ = url.Parse(testServer.URL)
}

func stopTestServer() {
	if testServer != nil {
		testServer.Close()
	}
}

func checkMethod(t *testing.T, r *http.Request, expected string) {
	if r.Method != expected {
		t.Errorf("Expected %v method but got %v", expected, r.Method)
	}
}

func TestNewClient(t *testing.T) {
	client := NewClient(TestToken)
	if client.Token != TestToken {
		t.Errorf("Expected Token to be set to %s but got %s", TestToken, client.Token)
	}
}

func TestClient_EnvVars_badJSON(t *testing.T) {
	startTestServer()
	defer stopTestServer()
	path := fmt.Sprintf("/project/%s/%s/envvar", testUsername, testReponame)
	testMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"name": "FOO", "value": "BAR"}, {"name": "BAZ", "value": "FIZZ"`)
	})
	envVars, apiResp := testClient.EnvVars(testUsername, testReponame)
	if apiResp.Success() {
		t.Errorf("Expected response to not be successful")
	}
	if apiResp.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", apiResp.Response.StatusCode)
	}
	expectedError := "unexpected end of JSON input"
	if apiResp.Error.Error() != expectedError {
		t.Errorf("Expected error message of %s but got %s", expectedError, apiResp.Error.Error())
	}
	if len(envVars) != 0 {
		t.Errorf("Expected to get %d envvars but got %d", 0, len(envVars))
	}
}
